package license

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/license_users"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/license_organization"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/license_toggles"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/license_definitions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/license_infer"
)

func init() {
	licenseCmd.AddCommand(license_users.Cmdlicense_users())
	licenseCmd.AddCommand(license_organization.Cmdlicense_organization())
	licenseCmd.AddCommand(license_toggles.Cmdlicense_toggles())
	licenseCmd.AddCommand(license_definitions.Cmdlicense_definitions())
	licenseCmd.AddCommand(license_infer.Cmdlicense_infer())
	licenseCmd.Short = utils.GenerateCustomDescription(licenseCmd.Short, license_users.Description, license_organization.Description, license_toggles.Description, license_definitions.Description, license_infer.Description, )
	licenseCmd.Long = licenseCmd.Short
}
