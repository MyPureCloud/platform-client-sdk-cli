
package notifications

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/topics"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/channels"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/subscriptions"
)

func init() {
	notificationsCmd.AddCommand(topics.Cmdtopics())
	notificationsCmd.AddCommand(channels.Cmdchannels())
	notificationsCmd.AddCommand(subscriptions.Cmdsubscriptions())
}
