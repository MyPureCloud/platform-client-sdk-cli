package architect_ivrs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_ivrs_divisionviews"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_ivrs_identityresolution"
)

func init() {
	architect_ivrsCmd.AddCommand(architect_ivrs_divisionviews.Cmdarchitect_ivrs_divisionviews())
	architect_ivrsCmd.AddCommand(architect_ivrs_identityresolution.Cmdarchitect_ivrs_identityresolution())
	architect_ivrsCmd.Short = utils.GenerateCustomDescription(architect_ivrsCmd.Short, architect_ivrs_divisionviews.Description, architect_ivrs_identityresolution.Description, )
	architect_ivrsCmd.Long = architect_ivrsCmd.Short
}
