package casemanagement_cases_terminate

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("casemanagement_cases_terminate", "SWAGGER_OVERRIDE_/api/v2/casemanagement/cases/{caseId}/terminate")
	casemanagement_cases_terminateCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("casemanagement_cases_terminate"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(casemanagement_cases_terminateCmd)
}

func Cmdcasemanagement_cases_terminate() *cobra.Command {
	return casemanagement_cases_terminateCmd
}
