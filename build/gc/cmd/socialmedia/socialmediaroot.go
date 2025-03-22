package socialmedia

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/socialmedia_analytics"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/socialmedia_topics"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/socialmedia_twitter"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/socialmedia_escalationrules"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/socialmedia_escalations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/socialmedia_messages"
)

func init() {
	socialmediaCmd.AddCommand(socialmedia_analytics.Cmdsocialmedia_analytics())
	socialmediaCmd.AddCommand(socialmedia_topics.Cmdsocialmedia_topics())
	socialmediaCmd.AddCommand(socialmedia_twitter.Cmdsocialmedia_twitter())
	socialmediaCmd.AddCommand(socialmedia_escalationrules.Cmdsocialmedia_escalationrules())
	socialmediaCmd.AddCommand(socialmedia_escalations.Cmdsocialmedia_escalations())
	socialmediaCmd.AddCommand(socialmedia_messages.Cmdsocialmedia_messages())
	socialmediaCmd.Short = utils.GenerateCustomDescription(socialmediaCmd.Short, socialmedia_analytics.Description, socialmedia_topics.Description, socialmedia_twitter.Description, socialmedia_escalationrules.Description, socialmedia_escalations.Description, socialmedia_messages.Description, )
	socialmediaCmd.Long = socialmediaCmd.Short
}
