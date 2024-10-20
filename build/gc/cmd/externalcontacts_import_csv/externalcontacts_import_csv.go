package externalcontacts_import_csv

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("externalcontacts_import_csv", "SWAGGER_OVERRIDE_/api/v2/externalcontacts/import/csv")
	externalcontacts_import_csvCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("externalcontacts_import_csv"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(externalcontacts_import_csvCmd)
}

func Cmdexternalcontacts_import_csv() *cobra.Command {
	return externalcontacts_import_csvCmd
}
