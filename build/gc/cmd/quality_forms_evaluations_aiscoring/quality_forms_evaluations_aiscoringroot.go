package quality_forms_evaluations_aiscoring

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_forms_evaluations_aiscoring_settings"
)

func init() {
	quality_forms_evaluations_aiscoringCmd.AddCommand(quality_forms_evaluations_aiscoring_settings.Cmdquality_forms_evaluations_aiscoring_settings())
	quality_forms_evaluations_aiscoringCmd.Short = utils.GenerateCustomDescription(quality_forms_evaluations_aiscoringCmd.Short, quality_forms_evaluations_aiscoring_settings.Description, )
	quality_forms_evaluations_aiscoringCmd.Long = quality_forms_evaluations_aiscoringCmd.Short
}
