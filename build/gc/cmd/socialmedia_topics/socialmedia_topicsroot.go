package socialmedia_topics

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/socialmedia_topics_dataingestionrules"
)

func init() {
	socialmedia_topicsCmd.AddCommand(socialmedia_topics_dataingestionrules.Cmdsocialmedia_topics_dataingestionrules())
	socialmedia_topicsCmd.Short = utils.GenerateCustomDescription(socialmedia_topicsCmd.Short, socialmedia_topics_dataingestionrules.Description, )
	socialmedia_topicsCmd.Long = socialmedia_topicsCmd.Short
}
