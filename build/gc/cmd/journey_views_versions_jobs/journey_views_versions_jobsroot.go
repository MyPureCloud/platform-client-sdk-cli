package journey_views_versions_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_views_versions_jobs_latest"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_views_versions_jobs_results"
)

func init() {
	journey_views_versions_jobsCmd.AddCommand(journey_views_versions_jobs_latest.Cmdjourney_views_versions_jobs_latest())
	journey_views_versions_jobsCmd.AddCommand(journey_views_versions_jobs_results.Cmdjourney_views_versions_jobs_results())
	journey_views_versions_jobsCmd.Short = utils.GenerateCustomDescription(journey_views_versions_jobsCmd.Short, journey_views_versions_jobs_latest.Description, journey_views_versions_jobs_results.Description, )
	journey_views_versions_jobsCmd.Long = journey_views_versions_jobsCmd.Short
}
