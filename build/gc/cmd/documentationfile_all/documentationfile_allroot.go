package documentationfile_all

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/documentationfile_all_search"
)

func init() {
	documentationfile_allCmd.AddCommand(documentationfile_all_search.Cmddocumentationfile_all_search())
	documentationfile_allCmd.Short = utils.GenerateCustomDescription(documentationfile_allCmd.Short, documentationfile_all_search.Description, )
	documentationfile_allCmd.Long = documentationfile_allCmd.Short
}
