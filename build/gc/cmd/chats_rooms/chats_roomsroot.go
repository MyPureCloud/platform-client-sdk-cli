package chats_rooms

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/chats_rooms_messages"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/chats_rooms_participants"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/chats_rooms_pinnedmessages"
)

func init() {
	chats_roomsCmd.AddCommand(chats_rooms_messages.Cmdchats_rooms_messages())
	chats_roomsCmd.AddCommand(chats_rooms_participants.Cmdchats_rooms_participants())
	chats_roomsCmd.AddCommand(chats_rooms_pinnedmessages.Cmdchats_rooms_pinnedmessages())
	chats_roomsCmd.Short = utils.GenerateCustomDescription(chats_roomsCmd.Short, chats_rooms_messages.Description, chats_rooms_participants.Description, chats_rooms_pinnedmessages.Description, )
	chats_roomsCmd.Long = chats_roomsCmd.Short
}
