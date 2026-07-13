package externalcontacts_notes

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("externalcontacts_notes", "SWAGGER_OVERRIDE_/api/v2/externalcontacts/notes")
	externalcontacts_notesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("externalcontacts_notes"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(externalcontacts_notesCmd)
}

func Cmdexternalcontacts_notes() *cobra.Command {
	return externalcontacts_notesCmd
}
