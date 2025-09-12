package quality_publishedforms_evaluations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_publishedforms_evaluations_divisionviews"
)

func init() {
	quality_publishedforms_evaluationsCmd.AddCommand(quality_publishedforms_evaluations_divisionviews.Cmdquality_publishedforms_evaluations_divisionviews())
	quality_publishedforms_evaluationsCmd.Short = utils.GenerateCustomDescription(quality_publishedforms_evaluationsCmd.Short, quality_publishedforms_evaluations_divisionviews.Description, )
	quality_publishedforms_evaluationsCmd.Long = quality_publishedforms_evaluationsCmd.Short
}
