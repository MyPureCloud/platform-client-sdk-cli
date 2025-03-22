package socialmedia_twitter_historical

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/socialmedia_twitter_historical_tweets"
)

func init() {
	socialmedia_twitter_historicalCmd.AddCommand(socialmedia_twitter_historical_tweets.Cmdsocialmedia_twitter_historical_tweets())
	socialmedia_twitter_historicalCmd.Short = utils.GenerateCustomDescription(socialmedia_twitter_historicalCmd.Short, socialmedia_twitter_historical_tweets.Description, )
	socialmedia_twitter_historicalCmd.Long = socialmedia_twitter_historicalCmd.Short
}
