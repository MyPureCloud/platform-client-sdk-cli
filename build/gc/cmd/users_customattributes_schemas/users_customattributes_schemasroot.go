package users_customattributes_schemas

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_customattributes_schemas_coretypes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_customattributes_schemas_limits"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_customattributes_schemas_versions"
)

func init() {
	users_customattributes_schemasCmd.AddCommand(users_customattributes_schemas_coretypes.Cmdusers_customattributes_schemas_coretypes())
	users_customattributes_schemasCmd.AddCommand(users_customattributes_schemas_limits.Cmdusers_customattributes_schemas_limits())
	users_customattributes_schemasCmd.AddCommand(users_customattributes_schemas_versions.Cmdusers_customattributes_schemas_versions())
	users_customattributes_schemasCmd.Short = utils.GenerateCustomDescription(users_customattributes_schemasCmd.Short, users_customattributes_schemas_coretypes.Description, users_customattributes_schemas_limits.Description, users_customattributes_schemas_versions.Description, )
	users_customattributes_schemasCmd.Long = users_customattributes_schemasCmd.Short
}
