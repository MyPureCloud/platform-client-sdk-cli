package uploads_publicassets

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("uploads_publicassets", "SWAGGER_OVERRIDE_/api/v2/uploads/publicassets")
	uploads_publicassetsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("uploads_publicassets"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(uploads_publicassetsCmd)
}

func Cmduploads_publicassets() *cobra.Command {
	return uploads_publicassetsCmd
}
