package conversations_customattributes

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_customattributes_bulk"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_customattributes_search"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_customattributes_schemas"
)

func init() {
	conversations_customattributesCmd.AddCommand(conversations_customattributes_bulk.Cmdconversations_customattributes_bulk())
	conversations_customattributesCmd.AddCommand(conversations_customattributes_search.Cmdconversations_customattributes_search())
	conversations_customattributesCmd.AddCommand(conversations_customattributes_schemas.Cmdconversations_customattributes_schemas())
	conversations_customattributesCmd.Short = utils.GenerateCustomDescription(conversations_customattributesCmd.Short, conversations_customattributes_bulk.Description, conversations_customattributes_search.Description, conversations_customattributes_schemas.Description, )
	conversations_customattributesCmd.Long = conversations_customattributesCmd.Short
}
