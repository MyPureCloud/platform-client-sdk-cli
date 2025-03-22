package socialmedia_analytics_messages_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/socialmedia_analytics_messages_jobs_results"
)

func init() {
	socialmedia_analytics_messages_jobsCmd.AddCommand(socialmedia_analytics_messages_jobs_results.Cmdsocialmedia_analytics_messages_jobs_results())
	socialmedia_analytics_messages_jobsCmd.Short = utils.GenerateCustomDescription(socialmedia_analytics_messages_jobsCmd.Short, socialmedia_analytics_messages_jobs_results.Description, )
	socialmedia_analytics_messages_jobsCmd.Long = socialmedia_analytics_messages_jobsCmd.Short
}
