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

package persistence

const (
	EntityTypeApiSessions               = "apiSessions"
	EntityTypeAppwans                   = "appwans"
	EntityTypeCas                       = "cas"
	EntityTypeClusters                  = "clusters"
	EntityTypeConfigs                   = "configs"
	EntityTypeConfigTypes               = "configTypes"
	EntityTypeEdgeRouters               = "edgeRouters"
	EntityTypeEdgeRouterPolicies        = "edgeRouterPolicies"
	EntityTypeEventLogs                 = "eventLogs"
	EntityTypeGeoRegions                = "geoRegions"
	EntityTypeIdentities                = "identities"
	EntityTypeIdentityTypes             = "identityTypes"
	EntityTypeServices                  = "services"
	EntityTypeServicePolicies           = "servicePolicies"
	EntityTypeServiceEdgeRouterPolicies = "serviceEdgeRouterPolicies"
	EntityTypeSessions                  = "sessions"
	EntityTypeSessionCerts              = "sessionCerts"
	EntityTypeEnrollments               = "enrollments"
	EntityTypeAuthenticators            = "authenticators"
	EdgeBucket                          = "edge"

	FieldName           = "name"
	FieldSemantic       = "semantic"
	FieldRoleAttributes = "roleAttributes"

	FieldEdgeRouterRoles = "edgeRouterRoles"
	FieldIdentityRoles   = "identityRoles"
	FieldServiceRoles    = "serviceRoles"

	SemanticAllOf = "AllOf"
	SemanticAnyOf = "AnyOf"
)

var validSemantics = []string{SemanticAllOf, SemanticAnyOf}

func toStringStringMap(m map[string]interface{}) map[string]string {
	result := map[string]string{}
	for k, v := range m {
		result[k] = v.(string)
	}
	return result
}

func toStringInterfaceMap(m map[string]string) map[string]interface{} {
	result := map[string]interface{}{}
	for k, v := range m {
		result[k] = v
	}
	return result
}