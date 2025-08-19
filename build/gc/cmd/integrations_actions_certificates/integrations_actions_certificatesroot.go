package integrations_actions_certificates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_actions_certificates_truststore"
)

func init() {
	integrations_actions_certificatesCmd.AddCommand(integrations_actions_certificates_truststore.Cmdintegrations_actions_certificates_truststore())
	integrations_actions_certificatesCmd.Short = utils.GenerateCustomDescription(integrations_actions_certificatesCmd.Short, integrations_actions_certificates_truststore.Description, )
	integrations_actions_certificatesCmd.Long = integrations_actions_certificatesCmd.Short
}
