package users_development

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_development_activities"
)

func init() {
	users_developmentCmd.AddCommand(users_development_activities.Cmdusers_development_activities())
	users_developmentCmd.Short = utils.GenerateCustomDescription(users_developmentCmd.Short, users_development_activities.Description, )
	users_developmentCmd.Long = users_developmentCmd.Short
}
