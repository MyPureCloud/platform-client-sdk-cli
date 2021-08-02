package knowledge

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_documentuploads"
)

func init() {
	knowledgeCmd.AddCommand(knowledge_knowledgebases.Cmdknowledge_knowledgebases())
	knowledgeCmd.AddCommand(knowledge_documentuploads.Cmdknowledge_documentuploads())
	knowledgeCmd.Short = utils.GenerateCustomDescription(knowledgeCmd.Short, knowledge_knowledgebases.Description, knowledge_documentuploads.Description, )
	knowledgeCmd.Long = knowledgeCmd.Short
}
