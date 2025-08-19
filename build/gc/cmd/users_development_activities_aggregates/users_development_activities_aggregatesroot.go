package users_development_activities_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_development_activities_aggregates_query"
)

func init() {
	users_development_activities_aggregatesCmd.AddCommand(users_development_activities_aggregates_query.Cmdusers_development_activities_aggregates_query())
	users_development_activities_aggregatesCmd.Short = utils.GenerateCustomDescription(users_development_activities_aggregatesCmd.Short, users_development_activities_aggregates_query.Description, )
	users_development_activities_aggregatesCmd.Long = users_development_activities_aggregatesCmd.Short
}
