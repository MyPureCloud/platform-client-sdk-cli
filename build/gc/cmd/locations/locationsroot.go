package locations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/locations_search"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/locations_sublocations"
)

func init() {
	locationsCmd.AddCommand(locations_search.Cmdlocations_search())
	locationsCmd.AddCommand(locations_sublocations.Cmdlocations_sublocations())
	locationsCmd.Short = utils.GenerateCustomDescription(locationsCmd.Short, locations_search.Description, locations_sublocations.Description, )
	locationsCmd.Long = locationsCmd.Short
}
