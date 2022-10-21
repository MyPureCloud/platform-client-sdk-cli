package externalcontacts_contacts_journey

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("externalcontacts_contacts_journey", "SWAGGER_OVERRIDE_/api/v2/externalcontacts/contacts/{contactId}/journey")
	externalcontacts_contacts_journeyCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("externalcontacts_contacts_journey"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(externalcontacts_contacts_journeyCmd)
}

func Cmdexternalcontacts_contacts_journey() *cobra.Command {
	return externalcontacts_contacts_journeyCmd
}
