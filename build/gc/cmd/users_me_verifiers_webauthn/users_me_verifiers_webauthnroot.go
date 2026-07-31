package users_me_verifiers_webauthn

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_me_verifiers_webauthn_register"
)

func init() {
	users_me_verifiers_webauthnCmd.AddCommand(users_me_verifiers_webauthn_register.Cmdusers_me_verifiers_webauthn_register())
	users_me_verifiers_webauthnCmd.Short = utils.GenerateCustomDescription(users_me_verifiers_webauthnCmd.Short, users_me_verifiers_webauthn_register.Description, )
	users_me_verifiers_webauthnCmd.Long = users_me_verifiers_webauthnCmd.Short
}
