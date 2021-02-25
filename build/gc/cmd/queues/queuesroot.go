
package queues

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_queues"
)

func init() {
	queuesCmd.AddCommand(users_queues.Cmdusers_queues())
}
