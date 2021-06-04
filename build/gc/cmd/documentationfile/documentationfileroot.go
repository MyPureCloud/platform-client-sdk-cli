package documentationfile

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/documentationfile_gkn"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/documentationfile_search"
)

func init() {
	documentationfileCmd.AddCommand(documentationfile_gkn.Cmddocumentationfile_gkn())
	documentationfileCmd.AddCommand(documentationfile_search.Cmddocumentationfile_search())
	documentationfileCmd.Short = utils.GenerateCustomDescription(documentationfileCmd.Short, documentationfile_gkn.Description, documentationfile_search.Description, )
	documentationfileCmd.Long = documentationfileCmd.Short
}
