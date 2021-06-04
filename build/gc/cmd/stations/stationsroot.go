package stations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/stations_associateduser"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/stations_settings"
)

func init() {
	stationsCmd.AddCommand(stations_associateduser.Cmdstations_associateduser())
	stationsCmd.AddCommand(stations_settings.Cmdstations_settings())
	stationsCmd.Short = utils.GenerateCustomDescription(stationsCmd.Short, stations_associateduser.Description, stations_settings.Description, )
	stationsCmd.Long = stationsCmd.Short
}
