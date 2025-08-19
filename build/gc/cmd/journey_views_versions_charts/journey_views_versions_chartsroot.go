package journey_views_versions_charts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_views_versions_charts_versions"
)

func init() {
	journey_views_versions_chartsCmd.AddCommand(journey_views_versions_charts_versions.Cmdjourney_views_versions_charts_versions())
	journey_views_versions_chartsCmd.Short = utils.GenerateCustomDescription(journey_views_versions_chartsCmd.Short, journey_views_versions_charts_versions.Description, )
	journey_views_versions_chartsCmd.Long = journey_views_versions_chartsCmd.Short
}
