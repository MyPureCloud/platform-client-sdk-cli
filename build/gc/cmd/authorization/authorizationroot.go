package authorization

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_divisions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_subjects"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_roles"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_divisionspermitted"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_settings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_permissions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_products"
)

func init() {
	authorizationCmd.AddCommand(authorization_divisions.Cmdauthorization_divisions())
	authorizationCmd.AddCommand(authorization_subjects.Cmdauthorization_subjects())
	authorizationCmd.AddCommand(authorization_roles.Cmdauthorization_roles())
	authorizationCmd.AddCommand(authorization_divisionspermitted.Cmdauthorization_divisionspermitted())
	authorizationCmd.AddCommand(authorization_settings.Cmdauthorization_settings())
	authorizationCmd.AddCommand(authorization_permissions.Cmdauthorization_permissions())
	authorizationCmd.AddCommand(authorization_products.Cmdauthorization_products())
	authorizationCmd.Short = utils.GenerateCustomDescription(authorizationCmd.Short, authorization_divisions.Description, authorization_subjects.Description, authorization_roles.Description, authorization_divisionspermitted.Description, authorization_settings.Description, authorization_permissions.Description, authorization_products.Description, )
	authorizationCmd.Long = authorizationCmd.Short
}
