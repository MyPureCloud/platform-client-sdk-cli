package authorization_divisions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_divisions_home"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_divisions_limit"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_divisions_objects"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_divisions_restore"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_divisions_grants"
)

func init() {
	authorization_divisionsCmd.AddCommand(authorization_divisions_home.Cmdauthorization_divisions_home())
	authorization_divisionsCmd.AddCommand(authorization_divisions_limit.Cmdauthorization_divisions_limit())
	authorization_divisionsCmd.AddCommand(authorization_divisions_objects.Cmdauthorization_divisions_objects())
	authorization_divisionsCmd.AddCommand(authorization_divisions_restore.Cmdauthorization_divisions_restore())
	authorization_divisionsCmd.AddCommand(authorization_divisions_grants.Cmdauthorization_divisions_grants())
	authorization_divisionsCmd.Short = utils.GenerateCustomDescription(authorization_divisionsCmd.Short, authorization_divisions_home.Description, authorization_divisions_limit.Description, authorization_divisions_objects.Description, authorization_divisions_restore.Description, authorization_divisions_grants.Description, )
	authorization_divisionsCmd.Long = authorization_divisionsCmd.Short
}
