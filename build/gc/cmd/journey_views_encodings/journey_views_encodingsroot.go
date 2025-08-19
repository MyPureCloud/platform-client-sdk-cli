package journey_views_encodings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_views_encodings_validate"
)

func init() {
	journey_views_encodingsCmd.AddCommand(journey_views_encodings_validate.Cmdjourney_views_encodings_validate())
	journey_views_encodingsCmd.Short = utils.GenerateCustomDescription(journey_views_encodingsCmd.Short, journey_views_encodings_validate.Description, )
	journey_views_encodingsCmd.Long = journey_views_encodingsCmd.Short
}
