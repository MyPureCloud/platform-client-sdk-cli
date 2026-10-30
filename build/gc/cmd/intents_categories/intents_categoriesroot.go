package intents_categories

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/intents_categories_similar"
)

func init() {
	intents_categoriesCmd.AddCommand(intents_categories_similar.Cmdintents_categories_similar())
	intents_categoriesCmd.Short = utils.GenerateCustomDescription(intents_categoriesCmd.Short, intents_categories_similar.Description, )
	intents_categoriesCmd.Long = intents_categoriesCmd.Short
}
