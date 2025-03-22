package socialmedia_analytics

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/socialmedia_analytics_aggregates"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/socialmedia_analytics_messages"
)

func init() {
	socialmedia_analyticsCmd.AddCommand(socialmedia_analytics_aggregates.Cmdsocialmedia_analytics_aggregates())
	socialmedia_analyticsCmd.AddCommand(socialmedia_analytics_messages.Cmdsocialmedia_analytics_messages())
	socialmedia_analyticsCmd.Short = utils.GenerateCustomDescription(socialmedia_analyticsCmd.Short, socialmedia_analytics_aggregates.Description, socialmedia_analytics_messages.Description, )
	socialmedia_analyticsCmd.Long = socialmedia_analyticsCmd.Short
}
