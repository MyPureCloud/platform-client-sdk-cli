
package queues

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/user_queues"
)

func init() {
	queuesCmd.AddCommand(user_queues.Cmduser_queues())
}
