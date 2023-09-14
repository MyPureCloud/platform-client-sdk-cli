package flows_versions_intents_utterances

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_versions_intents_utterances_health"
)

func init() {
	flows_versions_intents_utterancesCmd.AddCommand(flows_versions_intents_utterances_health.Cmdflows_versions_intents_utterances_health())
	flows_versions_intents_utterancesCmd.Short = utils.GenerateCustomDescription(flows_versions_intents_utterancesCmd.Short, flows_versions_intents_utterances_health.Description, )
	flows_versions_intents_utterancesCmd.Long = flows_versions_intents_utterancesCmd.Short
}
