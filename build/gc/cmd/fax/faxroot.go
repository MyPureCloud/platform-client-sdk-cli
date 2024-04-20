package fax

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/fax_settings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/fax_documents"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/fax_summary"
)

func init() {
	faxCmd.AddCommand(fax_settings.Cmdfax_settings())
	faxCmd.AddCommand(fax_documents.Cmdfax_documents())
	faxCmd.AddCommand(fax_summary.Cmdfax_summary())
	faxCmd.Short = utils.GenerateCustomDescription(faxCmd.Short, fax_settings.Description, fax_documents.Description, fax_summary.Description, )
	faxCmd.Long = faxCmd.Short
}
