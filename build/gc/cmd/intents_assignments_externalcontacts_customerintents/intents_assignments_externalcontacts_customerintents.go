package intents_assignments_externalcontacts_customerintents

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("intents_assignments_externalcontacts_customerintents", "SWAGGER_OVERRIDE_/api/v2/intents/assignments/externalcontacts/{externalContactId}/customerintents")
	intents_assignments_externalcontacts_customerintentsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("intents_assignments_externalcontacts_customerintents"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(intents_assignments_externalcontacts_customerintentsCmd)
}

func Cmdintents_assignments_externalcontacts_customerintents() *cobra.Command {
	return intents_assignments_externalcontacts_customerintentsCmd
}
