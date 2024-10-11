package journey_views_versions_jobs_results

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_views_versions_jobs_results_charts"
)

func init() {
	journey_views_versions_jobs_resultsCmd.AddCommand(journey_views_versions_jobs_results_charts.Cmdjourney_views_versions_jobs_results_charts())
	journey_views_versions_jobs_resultsCmd.Short = utils.GenerateCustomDescription(journey_views_versions_jobs_resultsCmd.Short, journey_views_versions_jobs_results_charts.Description, )
	journey_views_versions_jobs_resultsCmd.Long = journey_views_versions_jobs_resultsCmd.Short
}
