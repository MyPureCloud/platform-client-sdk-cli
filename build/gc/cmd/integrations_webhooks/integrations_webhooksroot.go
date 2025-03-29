package integrations_webhooks

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_webhooks_events"
)

func init() {
	integrations_webhooksCmd.AddCommand(integrations_webhooks_events.Cmdintegrations_webhooks_events())
	integrations_webhooksCmd.Short = utils.GenerateCustomDescription(integrations_webhooksCmd.Short, integrations_webhooks_events.Description, )
	integrations_webhooksCmd.Long = integrations_webhooksCmd.Short
}
