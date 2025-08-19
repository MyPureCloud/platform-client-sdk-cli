package uploads

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("uploads", "SWAGGER_OVERRIDE_/api/v2/uploads")
	uploadsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("uploads"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(uploadsCmd)
}

func Cmduploads() *cobra.Command {
	return uploadsCmd
}
