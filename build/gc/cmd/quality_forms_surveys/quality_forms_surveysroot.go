package quality_forms_surveys

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_forms_surveys_versions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_forms_surveys_bulk"
)

func init() {
	quality_forms_surveysCmd.AddCommand(quality_forms_surveys_versions.Cmdquality_forms_surveys_versions())
	quality_forms_surveysCmd.AddCommand(quality_forms_surveys_bulk.Cmdquality_forms_surveys_bulk())
	quality_forms_surveysCmd.Short = utils.GenerateCustomDescription(quality_forms_surveysCmd.Short, quality_forms_surveys_versions.Description, quality_forms_surveys_bulk.Description, )
	quality_forms_surveysCmd.Long = quality_forms_surveysCmd.Short
}
