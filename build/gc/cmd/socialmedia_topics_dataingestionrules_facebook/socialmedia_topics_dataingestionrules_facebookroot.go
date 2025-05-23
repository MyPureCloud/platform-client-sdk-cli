package socialmedia_topics_dataingestionrules_facebook

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/socialmedia_topics_dataingestionrules_facebook_versions"
)

func init() {
	socialmedia_topics_dataingestionrules_facebookCmd.AddCommand(socialmedia_topics_dataingestionrules_facebook_versions.Cmdsocialmedia_topics_dataingestionrules_facebook_versions())
	socialmedia_topics_dataingestionrules_facebookCmd.Short = utils.GenerateCustomDescription(socialmedia_topics_dataingestionrules_facebookCmd.Short, socialmedia_topics_dataingestionrules_facebook_versions.Description, )
	socialmedia_topics_dataingestionrules_facebookCmd.Long = socialmedia_topics_dataingestionrules_facebookCmd.Short
}
