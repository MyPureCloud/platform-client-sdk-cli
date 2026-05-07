package users_stations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_stations_me"
)

func init() {
	users_stationsCmd.AddCommand(users_stations_me.Cmdusers_stations_me())
	users_stationsCmd.Short = utils.GenerateCustomDescription(users_stationsCmd.Short, users_stations_me.Description, )
	users_stationsCmd.Long = users_stationsCmd.Short
}
