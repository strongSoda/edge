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

package routes

import (
	"github.com/netfoundry/ziti-fabric/controller/models"
	"reflect"
	"testing"

	"github.com/netfoundry/ziti-edge/controller/model"
)

func TestServiceApiCreate_ToModelService(t *testing.T) {
	type fields struct {
		Name               *string
		TerminatorStrategy *string
		RoleAttributes     []string
		Tags               map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    *model.Service
		wantErr bool
	}{
		{name: "test all fields", fields: fields{
			Name:               strPtr("bar"),
			TerminatorStrategy: strPtr("default"),
			RoleAttributes:     []string{"id1", "id2"},
			Tags:               map[string]interface{}{"hello": 1, "thing": "hi"},
		}, want: &model.Service{
			BaseEntity: models.BaseEntity{
				Tags: map[string]interface{}{"hello": 1, "thing": "hi"},
			},
			Name:               "bar",
			TerminatorStrategy: "default",
			RoleAttributes:     []string{"id1", "id2"},
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiService := &ServiceApiCreate{
				Name:               tt.fields.Name,
				TerminatorStrategy: tt.fields.TerminatorStrategy,
				RoleAttributes:     tt.fields.RoleAttributes,
				Tags:               tt.fields.Tags,
			}
			got := apiService.ToModel()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToModelService() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServiceApiUpdate_ToModelService(t *testing.T) {
	type fields struct {
		Name               *string
		TerminatorStrategy *string
		Tags               map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    *model.Service
		wantErr bool
	}{
		{name: "test all fields", fields: fields{
			Name:               strPtr("bar"),
			TerminatorStrategy: strPtr("foobar"),
			Tags:               map[string]interface{}{"hello": 1, "thing": "hi"},
		}, want: &model.Service{
			BaseEntity: models.BaseEntity{
				Tags: map[string]interface{}{"hello": 1, "thing": "hi"},
			},
			Name:               "bar",
			TerminatorStrategy: "foobar",
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiService := &ServiceApiUpdate{
				Name:               tt.fields.Name,
				TerminatorStrategy: tt.fields.TerminatorStrategy,
				Tags:               tt.fields.Tags,
			}
			got := apiService.ToModel("")
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToModelService() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func strPtr(val string) *string {
	return &val
}