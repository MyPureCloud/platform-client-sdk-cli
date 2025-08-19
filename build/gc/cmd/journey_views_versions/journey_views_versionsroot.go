package journey_views_versions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_views_versions_charts"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_views_versions_jobs"
)

func init() {
	journey_views_versionsCmd.AddCommand(journey_views_versions_charts.Cmdjourney_views_versions_charts())
	journey_views_versionsCmd.AddCommand(journey_views_versions_jobs.Cmdjourney_views_versions_jobs())
	journey_views_versionsCmd.Short = utils.GenerateCustomDescription(journey_views_versionsCmd.Short, journey_views_versions_charts.Description, journey_views_versions_jobs.Description, )
	journey_views_versionsCmd.Long = journey_views_versionsCmd.Short
}
