package voicemail_groups

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("voicemail_groups", "SWAGGER_OVERRIDE_/api/v2/voicemail/groups")
	voicemail_groupsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("voicemail_groups"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(voicemail_groupsCmd)
}

func Cmdvoicemail_groups() *cobra.Command {
	return voicemail_groupsCmd
}
