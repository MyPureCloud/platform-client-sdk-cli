package quality_surveys

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_surveys_scoring"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_surveys_scorable"
)

func init() {
	quality_surveysCmd.AddCommand(quality_surveys_scoring.Cmdquality_surveys_scoring())
	quality_surveysCmd.AddCommand(quality_surveys_scorable.Cmdquality_surveys_scorable())
	quality_surveysCmd.Short = utils.GenerateCustomDescription(quality_surveysCmd.Short, quality_surveys_scoring.Description, quality_surveys_scorable.Description, )
	quality_surveysCmd.Long = quality_surveysCmd.Short
}
