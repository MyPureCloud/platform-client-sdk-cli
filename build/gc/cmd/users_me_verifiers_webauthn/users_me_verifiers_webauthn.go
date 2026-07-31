package users_me_verifiers_webauthn

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("users_me_verifiers_webauthn", "SWAGGER_OVERRIDE_/api/v2/users/me/verifiers/webauthn")
	users_me_verifiers_webauthnCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("users_me_verifiers_webauthn"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(users_me_verifiers_webauthnCmd)
}

func Cmdusers_me_verifiers_webauthn() *cobra.Command {
	return users_me_verifiers_webauthnCmd
}
