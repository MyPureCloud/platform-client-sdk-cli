package quality_forms_evaluations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_forms_evaluations_aiscoring"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_forms_evaluations_versions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_forms_evaluations_bulk"
)

func init() {
	quality_forms_evaluationsCmd.AddCommand(quality_forms_evaluations_aiscoring.Cmdquality_forms_evaluations_aiscoring())
	quality_forms_evaluationsCmd.AddCommand(quality_forms_evaluations_versions.Cmdquality_forms_evaluations_versions())
	quality_forms_evaluationsCmd.AddCommand(quality_forms_evaluations_bulk.Cmdquality_forms_evaluations_bulk())
	quality_forms_evaluationsCmd.Short = utils.GenerateCustomDescription(quality_forms_evaluationsCmd.Short, quality_forms_evaluations_aiscoring.Description, quality_forms_evaluations_versions.Description, quality_forms_evaluations_bulk.Description, )
	quality_forms_evaluationsCmd.Long = quality_forms_evaluationsCmd.Short
}
