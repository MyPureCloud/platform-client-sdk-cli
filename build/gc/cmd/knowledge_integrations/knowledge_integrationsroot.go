package knowledge_integrations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_integrations_options"
)

func init() {
	knowledge_integrationsCmd.AddCommand(knowledge_integrations_options.Cmdknowledge_integrations_options())
	knowledge_integrationsCmd.Short = utils.GenerateCustomDescription(knowledge_integrationsCmd.Short, knowledge_integrations_options.Description, )
	knowledge_integrationsCmd.Long = knowledge_integrationsCmd.Short
}
