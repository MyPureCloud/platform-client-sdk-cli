package knowledge

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_documentuploads"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_guest"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_integrations"
)

func init() {
	knowledgeCmd.AddCommand(knowledge_knowledgebases.Cmdknowledge_knowledgebases())
	knowledgeCmd.AddCommand(knowledge_documentuploads.Cmdknowledge_documentuploads())
	knowledgeCmd.AddCommand(knowledge_guest.Cmdknowledge_guest())
	knowledgeCmd.AddCommand(knowledge_integrations.Cmdknowledge_integrations())
	knowledgeCmd.Short = utils.GenerateCustomDescription(knowledgeCmd.Short, knowledge_knowledgebases.Description, knowledge_documentuploads.Description, knowledge_guest.Description, knowledge_integrations.Description, )
	knowledgeCmd.Long = knowledgeCmd.Short
}
