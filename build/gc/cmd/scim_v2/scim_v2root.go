package scim_v2

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/scim_v2_resourcetypes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/scim_v2_schemas"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/scim_v2_groups"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/scim_v2_serviceproviderconfig"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/scim_v2_users"
)

func init() {
	scim_v2Cmd.AddCommand(scim_v2_resourcetypes.Cmdscim_v2_resourcetypes())
	scim_v2Cmd.AddCommand(scim_v2_schemas.Cmdscim_v2_schemas())
	scim_v2Cmd.AddCommand(scim_v2_groups.Cmdscim_v2_groups())
	scim_v2Cmd.AddCommand(scim_v2_serviceproviderconfig.Cmdscim_v2_serviceproviderconfig())
	scim_v2Cmd.AddCommand(scim_v2_users.Cmdscim_v2_users())
	scim_v2Cmd.Short = utils.GenerateCustomDescription(scim_v2Cmd.Short, scim_v2_resourcetypes.Description, scim_v2_schemas.Description, scim_v2_groups.Description, scim_v2_serviceproviderconfig.Description, scim_v2_users.Description, )
	scim_v2Cmd.Long = scim_v2Cmd.Short
}
