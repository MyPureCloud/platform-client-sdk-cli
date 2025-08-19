package documentationfile_gkn

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/documentationfile_gkn_search"
)

func init() {
	documentationfile_gknCmd.AddCommand(documentationfile_gkn_search.Cmddocumentationfile_gkn_search())
	documentationfile_gknCmd.Short = utils.GenerateCustomDescription(documentationfile_gknCmd.Short, documentationfile_gkn_search.Description, )
	documentationfile_gknCmd.Long = documentationfile_gknCmd.Short
}
