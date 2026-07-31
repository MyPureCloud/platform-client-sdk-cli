package quality_publishedforms_surveys

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_publishedforms_surveys_divisionviews"
)

func init() {
	quality_publishedforms_surveysCmd.AddCommand(quality_publishedforms_surveys_divisionviews.Cmdquality_publishedforms_surveys_divisionviews())
	quality_publishedforms_surveysCmd.Short = utils.GenerateCustomDescription(quality_publishedforms_surveysCmd.Short, quality_publishedforms_surveys_divisionviews.Description, )
	quality_publishedforms_surveysCmd.Long = quality_publishedforms_surveysCmd.Short
}
