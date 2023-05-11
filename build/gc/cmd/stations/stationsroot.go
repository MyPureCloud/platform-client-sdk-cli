package stations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/stations_associateduser"
)

func init() {
	stationsCmd.AddCommand(stations_associateduser.Cmdstations_associateduser())
	stationsCmd.Short = utils.GenerateCustomDescription(stationsCmd.Short, stations_associateduser.Description, )
	stationsCmd.Long = stationsCmd.Short
}
