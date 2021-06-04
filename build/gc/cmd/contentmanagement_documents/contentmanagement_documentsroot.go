package contentmanagement_documents

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/contentmanagement_documents_audits"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/contentmanagement_documents_content"
)

func init() {
	contentmanagement_documentsCmd.AddCommand(contentmanagement_documents_audits.Cmdcontentmanagement_documents_audits())
	contentmanagement_documentsCmd.AddCommand(contentmanagement_documents_content.Cmdcontentmanagement_documents_content())
	contentmanagement_documentsCmd.Short = utils.GenerateCustomDescription(contentmanagement_documentsCmd.Short, contentmanagement_documents_audits.Description, contentmanagement_documents_content.Description, )
	contentmanagement_documentsCmd.Long = contentmanagement_documentsCmd.Short
}
