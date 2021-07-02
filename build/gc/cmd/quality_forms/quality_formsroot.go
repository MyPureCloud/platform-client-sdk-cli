package quality_forms

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_forms_versions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_forms_surveys"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_forms_evaluations"
)

func init() {
	quality_formsCmd.AddCommand(quality_forms_versions.Cmdquality_forms_versions())
	quality_formsCmd.AddCommand(quality_forms_surveys.Cmdquality_forms_surveys())
	quality_formsCmd.AddCommand(quality_forms_evaluations.Cmdquality_forms_evaluations())
	quality_formsCmd.Short = utils.GenerateCustomDescription(quality_formsCmd.Short, quality_forms_versions.Description, quality_forms_surveys.Description, quality_forms_evaluations.Description, )
	quality_formsCmd.Long = quality_formsCmd.Short
}
