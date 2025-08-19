package journey_views_data

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_views_data_details"
)

func init() {
	journey_views_dataCmd.AddCommand(journey_views_data_details.Cmdjourney_views_data_details())
	journey_views_dataCmd.Short = utils.GenerateCustomDescription(journey_views_dataCmd.Short, journey_views_data_details.Description, )
	journey_views_dataCmd.Long = journey_views_dataCmd.Short
}
