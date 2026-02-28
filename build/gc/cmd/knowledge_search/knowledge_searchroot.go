package knowledge_search

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_search_preview"
)

func init() {
	knowledge_searchCmd.AddCommand(knowledge_search_preview.Cmdknowledge_search_preview())
	knowledge_searchCmd.Short = utils.GenerateCustomDescription(knowledge_searchCmd.Short, knowledge_search_preview.Description, )
	knowledge_searchCmd.Long = knowledge_searchCmd.Short
}
