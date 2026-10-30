package search

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/search_suggest"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/search_settings"
)

func init() {
	searchCmd.AddCommand(search_suggest.Cmdsearch_suggest())
	searchCmd.AddCommand(search_settings.Cmdsearch_settings())
	searchCmd.Short = utils.GenerateCustomDescription(searchCmd.Short, search_suggest.Description, search_settings.Description, )
	searchCmd.Long = searchCmd.Short
}
