package coaching_scheduleslots

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/coaching_scheduleslots_jobs"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/coaching_scheduleslots_query"
)

func init() {
	coaching_scheduleslotsCmd.AddCommand(coaching_scheduleslots_jobs.Cmdcoaching_scheduleslots_jobs())
	coaching_scheduleslotsCmd.AddCommand(coaching_scheduleslots_query.Cmdcoaching_scheduleslots_query())
	coaching_scheduleslotsCmd.Short = utils.GenerateCustomDescription(coaching_scheduleslotsCmd.Short, coaching_scheduleslots_jobs.Description, coaching_scheduleslots_query.Description, )
	coaching_scheduleslotsCmd.Long = coaching_scheduleslotsCmd.Short
}
