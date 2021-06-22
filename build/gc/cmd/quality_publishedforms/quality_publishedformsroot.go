package quality_publishedforms

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_publishedforms_surveys"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_publishedforms_evaluations"
)

func init() {
	quality_publishedformsCmd.AddCommand(quality_publishedforms_surveys.Cmdquality_publishedforms_surveys())
	quality_publishedformsCmd.AddCommand(quality_publishedforms_evaluations.Cmdquality_publishedforms_evaluations())
	quality_publishedformsCmd.Short = utils.GenerateCustomDescription(quality_publishedformsCmd.Short, quality_publishedforms_surveys.Description, quality_publishedforms_evaluations.Description, )
	quality_publishedformsCmd.Long = quality_publishedformsCmd.Short
}
