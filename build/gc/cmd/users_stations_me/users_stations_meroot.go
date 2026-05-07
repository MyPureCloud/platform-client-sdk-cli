package users_stations_me

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_stations_me_associatedstation"
)

func init() {
	users_stations_meCmd.AddCommand(users_stations_me_associatedstation.Cmdusers_stations_me_associatedstation())
	users_stations_meCmd.Short = utils.GenerateCustomDescription(users_stations_meCmd.Short, users_stations_me_associatedstation.Description, )
	users_stations_meCmd.Long = users_stations_meCmd.Short
}
