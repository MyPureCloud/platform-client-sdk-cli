package learning_scheduleslots

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_scheduleslots_query"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_scheduleslots_jobs"
)

func init() {
	learning_scheduleslotsCmd.AddCommand(learning_scheduleslots_query.Cmdlearning_scheduleslots_query())
	learning_scheduleslotsCmd.AddCommand(learning_scheduleslots_jobs.Cmdlearning_scheduleslots_jobs())
	learning_scheduleslotsCmd.Short = utils.GenerateCustomDescription(learning_scheduleslotsCmd.Short, learning_scheduleslots_query.Description, learning_scheduleslots_jobs.Description, )
	learning_scheduleslotsCmd.Long = learning_scheduleslotsCmd.Short
}
