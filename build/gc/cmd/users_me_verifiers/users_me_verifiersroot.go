package users_me_verifiers

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_me_verifiers_totp"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_me_verifiers_webauthn"
)

func init() {
	users_me_verifiersCmd.AddCommand(users_me_verifiers_totp.Cmdusers_me_verifiers_totp())
	users_me_verifiersCmd.AddCommand(users_me_verifiers_webauthn.Cmdusers_me_verifiers_webauthn())
	users_me_verifiersCmd.Short = utils.GenerateCustomDescription(users_me_verifiersCmd.Short, users_me_verifiers_totp.Description, users_me_verifiers_webauthn.Description, )
	users_me_verifiersCmd.Long = users_me_verifiersCmd.Short
}
