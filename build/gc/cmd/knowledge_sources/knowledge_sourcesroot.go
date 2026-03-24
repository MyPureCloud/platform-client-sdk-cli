package knowledge_sources

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_sources_synchronizations"
)

func init() {
	knowledge_sourcesCmd.AddCommand(knowledge_sources_synchronizations.Cmdknowledge_sources_synchronizations())
	knowledge_sourcesCmd.Short = utils.GenerateCustomDescription(knowledge_sourcesCmd.Short, knowledge_sources_synchronizations.Description, )
	knowledge_sourcesCmd.Long = knowledge_sourcesCmd.Short
}
