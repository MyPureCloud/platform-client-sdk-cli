package certificate

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/certificate_details"
)

func init() {
	certificateCmd.AddCommand(certificate_details.Cmdcertificate_details())
	certificateCmd.Short = utils.GenerateCustomDescription(certificateCmd.Short, certificate_details.Description, )
	certificateCmd.Long = certificateCmd.Short
}
