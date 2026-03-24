package license_infer

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/license_infer_permissions"
)

func init() {
	license_inferCmd.AddCommand(license_infer_permissions.Cmdlicense_infer_permissions())
	license_inferCmd.Short = utils.GenerateCustomDescription(license_inferCmd.Short, license_infer_permissions.Description, )
	license_inferCmd.Long = license_inferCmd.Short
}
