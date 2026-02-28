package knowledge

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_integrations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_documentuploads"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_settings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_search"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_guest"
)

func init() {
	knowledgeCmd.AddCommand(knowledge_knowledgebases.Cmdknowledge_knowledgebases())
	knowledgeCmd.AddCommand(knowledge_integrations.Cmdknowledge_integrations())
	knowledgeCmd.AddCommand(knowledge_documentuploads.Cmdknowledge_documentuploads())
	knowledgeCmd.AddCommand(knowledge_settings.Cmdknowledge_settings())
	knowledgeCmd.AddCommand(knowledge_search.Cmdknowledge_search())
	knowledgeCmd.AddCommand(knowledge_guest.Cmdknowledge_guest())
	knowledgeCmd.Short = utils.GenerateCustomDescription(knowledgeCmd.Short, knowledge_knowledgebases.Description, knowledge_integrations.Description, knowledge_documentuploads.Description, knowledge_settings.Description, knowledge_search.Description, knowledge_guest.Description, )
	knowledgeCmd.Long = knowledgeCmd.Short
}
