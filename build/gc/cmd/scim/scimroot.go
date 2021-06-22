package scim

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/scim_v2"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/scim_groups"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/scim_resourcetypes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/scim_serviceproviderconfig"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/scim_schemas"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/scim_users"
)

func init() {
	scimCmd.AddCommand(scim_v2.Cmdscim_v2())
	scimCmd.AddCommand(scim_groups.Cmdscim_groups())
	scimCmd.AddCommand(scim_resourcetypes.Cmdscim_resourcetypes())
	scimCmd.AddCommand(scim_serviceproviderconfig.Cmdscim_serviceproviderconfig())
	scimCmd.AddCommand(scim_schemas.Cmdscim_schemas())
	scimCmd.AddCommand(scim_users.Cmdscim_users())
	scimCmd.Short = utils.GenerateCustomDescription(scimCmd.Short, scim_v2.Description, scim_groups.Description, scim_resourcetypes.Description, scim_serviceproviderconfig.Description, scim_schemas.Description, scim_users.Description, )
	scimCmd.Long = scimCmd.Short
}
