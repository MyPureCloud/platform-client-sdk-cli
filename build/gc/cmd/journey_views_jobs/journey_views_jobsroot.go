package journey_views_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_views_jobs_me"
)

func init() {
	journey_views_jobsCmd.AddCommand(journey_views_jobs_me.Cmdjourney_views_jobs_me())
	journey_views_jobsCmd.Short = utils.GenerateCustomDescription(journey_views_jobsCmd.Short, journey_views_jobs_me.Description, )
	journey_views_jobsCmd.Long = journey_views_jobsCmd.Short
}
