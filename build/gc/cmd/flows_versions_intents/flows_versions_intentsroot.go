package flows_versions_intents

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_versions_intents_health"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_versions_intents_utterances"
)

func init() {
	flows_versions_intentsCmd.AddCommand(flows_versions_intents_health.Cmdflows_versions_intents_health())
	flows_versions_intentsCmd.AddCommand(flows_versions_intents_utterances.Cmdflows_versions_intents_utterances())
	flows_versions_intentsCmd.Short = utils.GenerateCustomDescription(flows_versions_intentsCmd.Short, flows_versions_intents_health.Description, flows_versions_intents_utterances.Description, )
	flows_versions_intentsCmd.Long = flows_versions_intentsCmd.Short
}
