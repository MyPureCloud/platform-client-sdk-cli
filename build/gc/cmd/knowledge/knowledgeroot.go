package knowledge

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases"
)

func init() {
	knowledgeCmd.AddCommand(knowledge_knowledgebases.Cmdknowledge_knowledgebases())
	knowledgeCmd.Short = utils.GenerateCustomDescription(knowledgeCmd.Short, knowledge_knowledgebases.Description, )
	knowledgeCmd.Long = knowledgeCmd.Short
}
