package users_development_activities

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_development_activities_aggregates"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_development_activities_me"
)

func init() {
	users_development_activitiesCmd.AddCommand(users_development_activities_aggregates.Cmdusers_development_activities_aggregates())
	users_development_activitiesCmd.AddCommand(users_development_activities_me.Cmdusers_development_activities_me())
	users_development_activitiesCmd.Short = utils.GenerateCustomDescription(users_development_activitiesCmd.Short, users_development_activities_aggregates.Description, users_development_activities_me.Description, )
	users_development_activitiesCmd.Long = users_development_activitiesCmd.Short
}
