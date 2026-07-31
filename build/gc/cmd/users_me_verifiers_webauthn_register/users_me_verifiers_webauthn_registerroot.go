package users_me_verifiers_webauthn_register

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_me_verifiers_webauthn_register_options"
)

func init() {
	users_me_verifiers_webauthn_registerCmd.AddCommand(users_me_verifiers_webauthn_register_options.Cmdusers_me_verifiers_webauthn_register_options())
	users_me_verifiers_webauthn_registerCmd.Short = utils.GenerateCustomDescription(users_me_verifiers_webauthn_registerCmd.Short, users_me_verifiers_webauthn_register_options.Description, )
	users_me_verifiers_webauthn_registerCmd.Long = users_me_verifiers_webauthn_registerCmd.Short
}
