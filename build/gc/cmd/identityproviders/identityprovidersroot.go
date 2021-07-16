package identityproviders

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/identityproviders_adfs"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/identityproviders_cic"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/identityproviders_gsuite"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/identityproviders_generic"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/identityproviders_identitynow"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/identityproviders_okta"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/identityproviders_onelogin"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/identityproviders_ping"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/identityproviders_purecloud"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/identityproviders_pureengage"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/identityproviders_salesforce"
)

func init() {
	identityprovidersCmd.AddCommand(identityproviders_adfs.Cmdidentityproviders_adfs())
	identityprovidersCmd.AddCommand(identityproviders_cic.Cmdidentityproviders_cic())
	identityprovidersCmd.AddCommand(identityproviders_gsuite.Cmdidentityproviders_gsuite())
	identityprovidersCmd.AddCommand(identityproviders_generic.Cmdidentityproviders_generic())
	identityprovidersCmd.AddCommand(identityproviders_identitynow.Cmdidentityproviders_identitynow())
	identityprovidersCmd.AddCommand(identityproviders_okta.Cmdidentityproviders_okta())
	identityprovidersCmd.AddCommand(identityproviders_onelogin.Cmdidentityproviders_onelogin())
	identityprovidersCmd.AddCommand(identityproviders_ping.Cmdidentityproviders_ping())
	identityprovidersCmd.AddCommand(identityproviders_purecloud.Cmdidentityproviders_purecloud())
	identityprovidersCmd.AddCommand(identityproviders_pureengage.Cmdidentityproviders_pureengage())
	identityprovidersCmd.AddCommand(identityproviders_salesforce.Cmdidentityproviders_salesforce())
	identityprovidersCmd.Short = utils.GenerateCustomDescription(identityprovidersCmd.Short, identityproviders_adfs.Description, identityproviders_cic.Description, identityproviders_gsuite.Description, identityproviders_generic.Description, identityproviders_identitynow.Description, identityproviders_okta.Description, identityproviders_onelogin.Description, identityproviders_ping.Description, identityproviders_purecloud.Description, identityproviders_pureengage.Description, identityproviders_salesforce.Description, )
	identityprovidersCmd.Long = identityprovidersCmd.Short
}
