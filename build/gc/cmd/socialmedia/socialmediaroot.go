package socialmedia

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/socialmedia_topics"
)

func init() {
	socialmediaCmd.AddCommand(socialmedia_topics.Cmdsocialmedia_topics())
	socialmediaCmd.Short = utils.GenerateCustomDescription(socialmediaCmd.Short, socialmedia_topics.Description, )
	socialmediaCmd.Long = socialmediaCmd.Short
}
