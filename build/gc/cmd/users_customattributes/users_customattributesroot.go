package users_customattributes

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_customattributes_schemas"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_customattributes_bulk"
)

func init() {
	users_customattributesCmd.AddCommand(users_customattributes_schemas.Cmdusers_customattributes_schemas())
	users_customattributesCmd.AddCommand(users_customattributes_bulk.Cmdusers_customattributes_bulk())
	users_customattributesCmd.Short = utils.GenerateCustomDescription(users_customattributesCmd.Short, users_customattributes_schemas.Description, users_customattributes_bulk.Description, )
	users_customattributesCmd.Long = users_customattributesCmd.Short
}
