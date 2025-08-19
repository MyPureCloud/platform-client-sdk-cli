package users_routinglanguages

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_routinglanguages_bulk"
)

func init() {
	users_routinglanguagesCmd.AddCommand(users_routinglanguages_bulk.Cmdusers_routinglanguages_bulk())
	users_routinglanguagesCmd.Short = utils.GenerateCustomDescription(users_routinglanguagesCmd.Short, users_routinglanguages_bulk.Description, )
	users_routinglanguagesCmd.Long = users_routinglanguagesCmd.Short
}
