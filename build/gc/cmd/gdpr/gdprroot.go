package gdpr

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gdpr_requests"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gdpr_subjects"
)

func init() {
	gdprCmd.AddCommand(gdpr_requests.Cmdgdpr_requests())
	gdprCmd.AddCommand(gdpr_subjects.Cmdgdpr_subjects())
	gdprCmd.Short = utils.GenerateCustomDescription(gdprCmd.Short, gdpr_requests.Description, gdpr_subjects.Description, )
	gdprCmd.Long = gdprCmd.Short
}
