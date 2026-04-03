package users_station

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_station_defaultstation"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_station_associatedstation"
)

func init() {
	users_stationCmd.AddCommand(users_station_defaultstation.Cmdusers_station_defaultstation())
	users_stationCmd.AddCommand(users_station_associatedstation.Cmdusers_station_associatedstation())
	users_stationCmd.Short = utils.GenerateCustomDescription(users_stationCmd.Short, users_station_defaultstation.Description, users_station_associatedstation.Description, )
	users_stationCmd.Long = users_stationCmd.Short
}
