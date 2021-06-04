package fax_documents

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/fax_documents_content"
)

func init() {
	fax_documentsCmd.AddCommand(fax_documents_content.Cmdfax_documents_content())
	fax_documentsCmd.Short = utils.GenerateCustomDescription(fax_documentsCmd.Short, fax_documents_content.Description, )
	fax_documentsCmd.Long = fax_documentsCmd.Short
}
