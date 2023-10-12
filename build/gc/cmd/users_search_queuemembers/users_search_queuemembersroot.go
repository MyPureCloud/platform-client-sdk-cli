package users_search_queuemembers

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_search_queuemembers_manage"
)

func init() {
	users_search_queuemembersCmd.AddCommand(users_search_queuemembers_manage.Cmdusers_search_queuemembers_manage())
	users_search_queuemembersCmd.Short = utils.GenerateCustomDescription(users_search_queuemembersCmd.Short, users_search_queuemembers_manage.Description, )
	users_search_queuemembersCmd.Long = users_search_queuemembersCmd.Short
}
