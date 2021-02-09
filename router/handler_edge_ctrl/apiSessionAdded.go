/*
	Copyright NetFoundry, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package handler_edge_ctrl

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/edge/controller/env"
	"github.com/openziti/edge/controller/sync_strats"
	"github.com/openziti/edge/pb/edge_ctrl_pb"
	"github.com/openziti/edge/router/internal/fabric"
	"github.com/openziti/foundation/channel2"
	"sync"
	"time"
)

type apiSessionAddedHandler struct {
	control     channel2.Channel
	sm          fabric.StateManager
	syncTracker *apiSessionSyncTracker
	reqChan     chan *edge_ctrl_pb.ApiSessionAdded
	syncReady   chan []*edge_ctrl_pb.ApiSession
	syncFail    chan error
}

func NewApiSessionAddedHandler(sm fabric.StateManager, control channel2.Channel) *apiSessionAddedHandler {
	handler := &apiSessionAddedHandler{
		control:   control,
		sm:        sm,
		reqChan:   make(chan *edge_ctrl_pb.ApiSessionAdded, 100),
		syncReady: make(chan []*edge_ctrl_pb.ApiSession, 0),
		syncFail:  make(chan error, 0),
	}

	go handler.startRecieveSync()
	go handler.startSyncApplier()
	go handler.startSyncFail()

	return handler
}

func (h *apiSessionAddedHandler) ContentType() int32 {
	return env.ApiSessionAddedType
}

func (h *apiSessionAddedHandler) HandleReceive(msg *channel2.Message, ch channel2.Channel) {
	go func() {
		req := &edge_ctrl_pb.ApiSessionAdded{}
		if err := proto.Unmarshal(msg.Body, req); err == nil {
			if req.IsFullState {
				h.reqChan <- req
			} else {
				for _, session := range req.ApiSessions {
					h.sm.AddApiSession(session)
				}
			}
		} else {
			pfxlog.Logger().Panic("could not convert message as api session added")
		}
	}()
}

func (h *apiSessionAddedHandler) startSyncApplier() {
	for apiSessions := range h.syncReady {
		for _, apiSession := range apiSessions {
			h.sm.AddApiSession(apiSession)
		}
		h.sm.RemoveMissingApiSessions(apiSessions)

		pfxlog.Logger().Infof("finished sychronizing api sessions [count: %d]", len(apiSessions))
	}
}

func (h *apiSessionAddedHandler) startSyncFail() {

	for err := range h.syncFail {
		h.syncTracker.Stop()
		h.syncTracker = nil

		resync := &edge_ctrl_pb.RequestClientReSync{
			Reason: fmt.Sprintf("error during sync: %v", err),
		}

		resyncProto, _ := proto.Marshal(resync)

		resyncMsg := channel2.NewMessage(env.RequestClientReSyncType, resyncProto)
		_ = h.control.Send(resyncMsg)
	}
}

func (h *apiSessionAddedHandler) legacySync(req *edge_ctrl_pb.ApiSessionAdded) {
	pfxlog.Logger().Warn("using legacy sync logic some connections may be dropped")
	for _, apiSession := range req.ApiSessions {
		h.sm.AddApiSession(apiSession)
	}

	h.sm.RemoveMissingApiSessions(req.ApiSessions)
}

func (h *apiSessionAddedHandler) startRecieveSync() {
	for req := range h.reqChan {
		switch req.SyncStrategy {
		case string(sync_strats.RouterSyncStrategyInstant):
			h.instantSync(req)
		case "":
			pfxlog.Logger().Warn("syncStrategy is not specifieid, old controller?")
			h.legacySync(req)
		default:
			pfxlog.Logger().Warnf("syncStrategy [%s] is not supported", req.SyncStrategy)
			h.legacySync(req)
		}
	}
}

func (h *apiSessionAddedHandler) instantSync(req *edge_ctrl_pb.ApiSessionAdded) {
	logger := pfxlog.Logger().WithField("strategy", req.SyncStrategy)

	state := &sync_strats.InstantSyncState{}

	if req.SyncState == "" {
		logger.Panic("syncState is empty, cannot continue")
	}

	if err := json.Unmarshal([]byte(req.SyncState), state); err != nil {
		logger.Panicf("could not parse sync state [%s], error: %v", req.SyncState, err)
	}

	//if no id or the sync id is newer, reset
	if h.syncTracker == nil || h.syncTracker.syncId == "" || h.syncTracker.syncId < state.Id {
		logger.Warnf("new syncId [%s], resetting", state.Id)

		if h.syncTracker != nil {
			h.syncTracker.Stop()
		}

		h.syncTracker = newApiSessionSyncTracker(state.Id)

		h.syncTracker.StartDeadline(h.syncReady, h.syncFail, 20*time.Second)
	}

	//ignore older syncs
	if h.syncTracker.syncId > state.Id {
		logger.Warnf("older syncId [%s], ignoring", state.Id)
		return
	}

	h.syncTracker.Add(&apiSessionAddedWithState{
		InstantSyncState: state,
		ApiSessionAdded:  req,
	})

}

type apiSessionSyncTracker struct {
	syncId           string
	syncLastRecieved bool
	reqsWithState    map[int]*apiSessionAddedWithState
	hasLast          bool
	lastSeq          int
	stop             chan struct{}
	deadline         sync.Once
	isDone           bool
}

func newApiSessionSyncTracker(id string) *apiSessionSyncTracker {
	return &apiSessionSyncTracker{
		syncId:        id,
		reqsWithState: map[int]*apiSessionAddedWithState{},
		stop:          make(chan struct{}, 0),
	}
}

func (tracker *apiSessionSyncTracker) Add(reqWithState *apiSessionAddedWithState) {
	tracker.reqsWithState[reqWithState.Sequence] = reqWithState

	if reqWithState.IsLast {
		tracker.hasLast = true
		tracker.lastSeq = reqWithState.Sequence
	}
}

func (tracker *apiSessionSyncTracker) Stop() {
	if tracker != nil {
		tracker.stop <- struct{}{}
	}
}

func (tracker *apiSessionSyncTracker) StartDeadline(syncReady chan []*edge_ctrl_pb.ApiSession, syncFail chan error, timeout time.Duration) {
	tracker.deadline.Do(func() {
		go func() {
			ticker := time.NewTicker(1 * time.Second)
			select {
			case <-tracker.stop:
				tracker.reqsWithState = nil
				syncFail <- nil
				return
			case <-ticker.C:
				if tracker.HasAll() {
					syncReady <- tracker.all()
					return
				}
			case <-time.After(timeout):
				tracker.reqsWithState = nil
				syncFail <- errors.New("timeout, did not receive all updates in time")
				return
			}
		}()
	})
}

func (tracker *apiSessionSyncTracker) HasAll() bool {
	if !tracker.hasLast {
		return false
	}

	hasAll := true

	for i := 0; i <= tracker.lastSeq; i++ {
		if req, ok := tracker.reqsWithState[i]; !ok && req == nil {
			hasAll = false
			break
		}
	}

	return hasAll
}

func (tracker *apiSessionSyncTracker) all() []*edge_ctrl_pb.ApiSession {
	var result []*edge_ctrl_pb.ApiSession
	for i := 0; i <= tracker.lastSeq; i++ {
		if req, ok := tracker.reqsWithState[i]; ok {
			for _, apiSession := range req.ApiSessions {
				result = append(result, apiSession)
			}
		} else {
			pfxlog.Logger().WithField("strategy", sync_strats.RouterSyncStrategyInstant).Error("all failed to have all update sequences")
		}
	}

	return result
}

type apiSessionAddedWithState struct {
	*sync_strats.InstantSyncState
	*edge_ctrl_pb.ApiSessionAdded
}

type apiSessionSyncResult struct {
	apiSessions []*apiSessionAddedWithState
}
