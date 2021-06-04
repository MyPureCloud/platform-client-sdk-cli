package quality_forms_evaluations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_forms_evaluations_versions"
)

func init() {
	quality_forms_evaluationsCmd.AddCommand(quality_forms_evaluations_versions.Cmdquality_forms_evaluations_versions())
	quality_forms_evaluationsCmd.Short = utils.GenerateCustomDescription(quality_forms_evaluationsCmd.Short, quality_forms_evaluations_versions.Description, )
	quality_forms_evaluationsCmd.Long = quality_forms_evaluationsCmd.Short
}
