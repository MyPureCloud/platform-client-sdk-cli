package stations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/stations_settings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/stations_associateduser"
)

func init() {
	stationsCmd.AddCommand(stations_settings.Cmdstations_settings())
	stationsCmd.AddCommand(stations_associateduser.Cmdstations_associateduser())
	stationsCmd.Short = utils.GenerateCustomDescription(stationsCmd.Short, stations_settings.Description, stations_associateduser.Description, )
	stationsCmd.Long = stationsCmd.Short
}
