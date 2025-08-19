package locations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/locations_sublocations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/locations_search"
)

func init() {
	locationsCmd.AddCommand(locations_sublocations.Cmdlocations_sublocations())
	locationsCmd.AddCommand(locations_search.Cmdlocations_search())
	locationsCmd.Short = utils.GenerateCustomDescription(locationsCmd.Short, locations_sublocations.Description, locations_search.Description, )
	locationsCmd.Long = locationsCmd.Short
}
