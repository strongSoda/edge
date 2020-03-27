// +build apitests

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

package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/netfoundry/ziti-foundation/util/stringz"
	"net/http"
	"reflect"
	"strings"

	"github.com/netfoundry/ziti-edge/controller/apierror"

	"github.com/Jeffail/gabs"
	"github.com/michaelquigley/pfxlog"
)

func (ctx *TestContext) setJsonValue(container *gabs.Container, value interface{}, path ...string) {
	_, err := container.Set(value, path...)
	ctx.req.NoError(err)
}

func (ctx *TestContext) setValue(container *gabs.Container, value interface{}, fields []string, field string) {
	ctx.setValueWithPath(container, value, fields, field, field)
}

func (ctx *TestContext) setValueWithPath(container *gabs.Container, value interface{}, fields []string, field string, path ...string) {
	if len(fields) == 0 || stringz.Contains(fields, field) {
		_, err := container.Set(value, path...)
		ctx.req.NoError(err)
	}
}

func (ctx *TestContext) parseJson(body []byte) *gabs.Container {
	result, err := gabs.ParseJSON(body)
	ctx.req.NoError(err)
	return result
}

func (ctx *TestContext) getEntityId(body []byte) string {
	result := ctx.parseJson(body)
	path := result.S("data", "id")
	ctx.req.NotNil(path)
	return fmt.Sprintf("%v", path.Data())
}

func (ctx *TestContext) pathEquals(container *gabs.Container, val interface{}, path []string) {
	pathValue := container.Search(path...)
	if val == nil || (reflect.TypeOf(val).Kind() == reflect.Map && reflect.ValueOf(val).IsNil()) {
		ctx.req.True(pathValue == nil || pathValue.Data() == nil)
	} else {
		ctx.req.Equal(val, pathValue.Data())
	}
}

func (ctx *TestContext) requireString(container *gabs.Container, path ...string) string {
	pathValue := container.Search(path...)
	ctx.req.NotNil(pathValue)
	result, ok := pathValue.Data().(string)
	ctx.req.True(ok, "%+v must be a string", path)
	return result
}

func (ctx *TestContext) pathEqualsStringSlice(container *gabs.Container, val interface{}, path []string) {
	pathValue := container.Search(path...)
	if val == nil || reflect.ValueOf(val).IsNil() {
		if pathValue != nil {
			ctx.req.Nil(pathValue.Data())
		}
	} else {
		slice := ctx.toStringSlice(pathValue)
		ctx.req.Equal(val, slice)
	}
}

func (ctx *TestContext) requirePath(container *gabs.Container, searchPath ...string) *gabs.Container {
	if len(searchPath) == 1 {
		searchPath = path(searchPath[0])
	}
	elem := container.S(searchPath...)
	ctx.req.NotNil(elem)
	return elem
}

func (ctx *TestContext) requireChildWith(container *gabs.Container, attribute string, value interface{}) *gabs.Container {
	child := ctx.childWith(container, attribute, value)
	ctx.req.NotNil(child, "no child found with %v = %v", attribute, value)
	return child
}

func (ctx *TestContext) requireNoChildWith(container *gabs.Container, attribute string, value interface{}) *gabs.Container {
	child := ctx.childWith(container, attribute, value)
	ctx.req.Nil(child, "child found with %v = %v", attribute, value)
	return child
}

func (ctx *TestContext) childWith(container *gabs.Container, attribute string, value interface{}) *gabs.Container {
	children, err := container.Children()
	ctx.req.NoError(err)
	for _, child := range children {
		attr := child.S(path(attribute)...)
		if attr == nil {
			continue
		}
		if reflect.DeepEqual(attr.Data(), value) {
			return child
		}
		pfxlog.Logger().Infof("for attr %v, %v did not match %v", attribute, value, attr.Data())
	}

	return nil
}

func (ctx *TestContext) toStringSlice(container *gabs.Container) []string {
	var result []string
	if container != nil {
		if container.Data() == nil {
			return nil
		}
		children, err := container.Children()
		ctx.req.NoError(err)
		for _, child := range children {
			val, ok := child.Data().(string)
			ctx.req.True(ok, "expected child to be string value")
			result = append(result, val)
		}
	}
	return result
}

func (ctx *TestContext) requireFieldError(httpStatus int, body []byte, errorCode string, field string) *gabs.Container {
	ctx.req.Equal(http.StatusBadRequest, httpStatus)
	parsed := ctx.parseJson(body)
	ctx.pathEquals(parsed, errorCode, path("error.code"))
	ctx.pathEquals(parsed, field, path("error.cause.field"))
	return parsed
}

func (ctx *TestContext) requireNotFoundError(httpStatus int, body []byte) *gabs.Container {
	ctx.req.Equal(http.StatusNotFound, httpStatus)
	parsed := ctx.parseJson(body)
	ctx.pathEquals(parsed, apierror.NotFoundCode, path("error.code"))
	ctx.pathEquals(parsed, "The resource requested was not found or is no longer available", path("error.message"))
	return parsed
}

func (ctx *TestContext) requireUnauthorizedError(httpStatus int, body []byte) *gabs.Container {
	ctx.req.Equal(http.StatusUnauthorized, httpStatus)
	parsed := ctx.parseJson(body)
	ctx.pathEquals(parsed, apierror.UnauthorizedCode, path("error.code"))
	ctx.pathEquals(parsed, "The request could not be completed. The session is not authorized or the credentials are invalid", path("error.message"))
	return parsed
}

func (ctx *TestContext) logJson(data []byte) {
	if ctx.enabledJsonLogging {
		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, data, "", "    "); err == nil {
			if _, err := fmt.Printf("Result:\n%s\n", prettyJSON.String()); err != nil {
				panic(err)
			}
		} else {
			if _, err := fmt.Printf("Result:\n%s\n", data); err != nil {
				panic(err)
			}
		}
	}
}

func path(path ...string) []string {
	if len(path) == 1 {
		return strings.Split(path[0], ".")
	}
	return path
}

func s(vals ...string) []string {
	return vals
}