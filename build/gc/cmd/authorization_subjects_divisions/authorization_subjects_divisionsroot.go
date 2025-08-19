package authorization_subjects_divisions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_subjects_divisions_roles"
)

func init() {
	authorization_subjects_divisionsCmd.AddCommand(authorization_subjects_divisions_roles.Cmdauthorization_subjects_divisions_roles())
	authorization_subjects_divisionsCmd.Short = utils.GenerateCustomDescription(authorization_subjects_divisionsCmd.Short, authorization_subjects_divisions_roles.Description, )
	authorization_subjects_divisionsCmd.Long = authorization_subjects_divisionsCmd.Short
}
