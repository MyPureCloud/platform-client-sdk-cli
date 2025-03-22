package socialmedia_analytics_messages

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/socialmedia_analytics_messages_jobs"
)

func init() {
	socialmedia_analytics_messagesCmd.AddCommand(socialmedia_analytics_messages_jobs.Cmdsocialmedia_analytics_messages_jobs())
	socialmedia_analytics_messagesCmd.Short = utils.GenerateCustomDescription(socialmedia_analytics_messagesCmd.Short, socialmedia_analytics_messages_jobs.Description, )
	socialmedia_analytics_messagesCmd.Long = socialmedia_analytics_messagesCmd.Short
}
