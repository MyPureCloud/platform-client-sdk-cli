package authorization_subjects_divisions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("authorization_subjects_divisions", "SWAGGER_OVERRIDE_/api/v2/authorization/subjects/{subjectId}/divisions")
	authorization_subjects_divisionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("authorization_subjects_divisions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(authorization_subjects_divisionsCmd)
}

func Cmdauthorization_subjects_divisions() *cobra.Command {
	return authorization_subjects_divisionsCmd
}
