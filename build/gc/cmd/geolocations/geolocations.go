package geolocations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("geolocations", "SWAGGER_OVERRIDE_/api/v2/geolocations")
	geolocationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("geolocations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(geolocationsCmd)
}

func Cmdgeolocations() *cobra.Command {
	return geolocationsCmd
}
