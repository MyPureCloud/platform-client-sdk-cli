package knowledge_knowledgebases

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_languages"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_search"
)

func init() {
	knowledge_knowledgebasesCmd.AddCommand(knowledge_knowledgebases_languages.Cmdknowledge_knowledgebases_languages())
	knowledge_knowledgebasesCmd.AddCommand(knowledge_knowledgebases_search.Cmdknowledge_knowledgebases_search())
	knowledge_knowledgebasesCmd.Short = utils.GenerateCustomDescription(knowledge_knowledgebasesCmd.Short, knowledge_knowledgebases_languages.Description, knowledge_knowledgebases_search.Description, )
	knowledge_knowledgebasesCmd.Long = knowledge_knowledgebasesCmd.Short
}
