package casemanagement_cases_query

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("casemanagement_cases_query", "SWAGGER_OVERRIDE_/api/v2/casemanagement/cases/query")
	casemanagement_cases_queryCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("casemanagement_cases_query"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(casemanagement_cases_queryCmd)
}

func Cmdcasemanagement_cases_query() *cobra.Command {
	return casemanagement_cases_queryCmd
}
