package conversations_participants_attributes

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_participants_attributes_search"
)

func init() {
	conversations_participants_attributesCmd.AddCommand(conversations_participants_attributes_search.Cmdconversations_participants_attributes_search())
	conversations_participants_attributesCmd.Short = utils.GenerateCustomDescription(conversations_participants_attributesCmd.Short, conversations_participants_attributes_search.Description, )
	conversations_participants_attributesCmd.Long = conversations_participants_attributesCmd.Short
}
