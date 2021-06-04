package search

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/search_suggest"
)

func init() {
	searchCmd.AddCommand(search_suggest.Cmdsearch_suggest())
	searchCmd.Short = utils.GenerateCustomDescription(searchCmd.Short, search_suggest.Description, )
	searchCmd.Long = searchCmd.Short
}
