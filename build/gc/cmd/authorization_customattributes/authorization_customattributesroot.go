package authorization_customattributes

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_customattributes_itemtypes"
)

func init() {
	authorization_customattributesCmd.AddCommand(authorization_customattributes_itemtypes.Cmdauthorization_customattributes_itemtypes())
	authorization_customattributesCmd.Short = utils.GenerateCustomDescription(authorization_customattributesCmd.Short, authorization_customattributes_itemtypes.Description, )
	authorization_customattributesCmd.Long = authorization_customattributesCmd.Short
}
