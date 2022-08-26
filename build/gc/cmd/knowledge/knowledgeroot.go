package knowledge

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_documentuploads"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_guest"
)

func init() {
	knowledgeCmd.AddCommand(knowledge_knowledgebases.Cmdknowledge_knowledgebases())
	knowledgeCmd.AddCommand(knowledge_documentuploads.Cmdknowledge_documentuploads())
	knowledgeCmd.AddCommand(knowledge_guest.Cmdknowledge_guest())
	knowledgeCmd.Short = utils.GenerateCustomDescription(knowledgeCmd.Short, knowledge_knowledgebases.Description, knowledge_documentuploads.Description, knowledge_guest.Description, )
	knowledgeCmd.Long = knowledgeCmd.Short
}
