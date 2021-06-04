package geolocations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/geolocations_settings"
)

func init() {
	geolocationsCmd.AddCommand(geolocations_settings.Cmdgeolocations_settings())
	geolocationsCmd.Short = utils.GenerateCustomDescription(geolocationsCmd.Short, geolocations_settings.Description, )
	geolocationsCmd.Long = geolocationsCmd.Short
}
