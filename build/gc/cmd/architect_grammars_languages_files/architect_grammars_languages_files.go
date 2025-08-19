package architect_grammars_languages_files

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("architect_grammars_languages_files", "SWAGGER_OVERRIDE_/api/v2/architect/grammars/{grammarId}/languages/{languageCode}/files")
	architect_grammars_languages_filesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("architect_grammars_languages_files"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(architect_grammars_languages_filesCmd)
}

func Cmdarchitect_grammars_languages_files() *cobra.Command {
	return architect_grammars_languages_filesCmd
}
