package casemanagement

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("casemanagement", "SWAGGER_OVERRIDE_/api/v2/casemanagement")
	casemanagementCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("casemanagement"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(casemanagementCmd)
}

func Cmdcasemanagement() *cobra.Command {
	return casemanagementCmd
}
