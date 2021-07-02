package documentationfile

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/documentationfile_search"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/documentationfile_gkn"
)

func init() {
	documentationfileCmd.AddCommand(documentationfile_search.Cmddocumentationfile_search())
	documentationfileCmd.AddCommand(documentationfile_gkn.Cmddocumentationfile_gkn())
	documentationfileCmd.Short = utils.GenerateCustomDescription(documentationfileCmd.Short, documentationfile_search.Description, documentationfile_gkn.Description, )
	documentationfileCmd.Long = documentationfileCmd.Short
}
