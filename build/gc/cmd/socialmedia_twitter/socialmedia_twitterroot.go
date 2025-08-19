package socialmedia_twitter

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/socialmedia_twitter_historical"
)

func init() {
	socialmedia_twitterCmd.AddCommand(socialmedia_twitter_historical.Cmdsocialmedia_twitter_historical())
	socialmedia_twitterCmd.Short = utils.GenerateCustomDescription(socialmedia_twitterCmd.Short, socialmedia_twitter_historical.Description, )
	socialmedia_twitterCmd.Long = socialmedia_twitterCmd.Short
}
