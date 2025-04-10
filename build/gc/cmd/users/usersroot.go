package users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_verifiers"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_agentui"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_roles"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_callforwarding"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_superiors"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_directreports"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_favorites"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_adjacents"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_profile"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_geolocations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_greetings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_outofoffice"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_presences"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_chats"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_externalid"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_me"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_password"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_profileskills"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_queues"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_invite"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_routinglanguages"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_routingskills"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_routingstatus"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_skillgroups"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_state"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_station"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_trustors"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_bulk"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_search"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_development"
)

func init() {
	usersCmd.AddCommand(users_verifiers.Cmdusers_verifiers())
	usersCmd.AddCommand(users_agentui.Cmdusers_agentui())
	usersCmd.AddCommand(users_roles.Cmdusers_roles())
	usersCmd.AddCommand(users_callforwarding.Cmdusers_callforwarding())
	usersCmd.AddCommand(users_superiors.Cmdusers_superiors())
	usersCmd.AddCommand(users_directreports.Cmdusers_directreports())
	usersCmd.AddCommand(users_favorites.Cmdusers_favorites())
	usersCmd.AddCommand(users_adjacents.Cmdusers_adjacents())
	usersCmd.AddCommand(users_profile.Cmdusers_profile())
	usersCmd.AddCommand(users_geolocations.Cmdusers_geolocations())
	usersCmd.AddCommand(users_greetings.Cmdusers_greetings())
	usersCmd.AddCommand(users_outofoffice.Cmdusers_outofoffice())
	usersCmd.AddCommand(users_presences.Cmdusers_presences())
	usersCmd.AddCommand(users_chats.Cmdusers_chats())
	usersCmd.AddCommand(users_externalid.Cmdusers_externalid())
	usersCmd.AddCommand(users_me.Cmdusers_me())
	usersCmd.AddCommand(users_password.Cmdusers_password())
	usersCmd.AddCommand(users_profileskills.Cmdusers_profileskills())
	usersCmd.AddCommand(users_queues.Cmdusers_queues())
	usersCmd.AddCommand(users_invite.Cmdusers_invite())
	usersCmd.AddCommand(users_routinglanguages.Cmdusers_routinglanguages())
	usersCmd.AddCommand(users_routingskills.Cmdusers_routingskills())
	usersCmd.AddCommand(users_routingstatus.Cmdusers_routingstatus())
	usersCmd.AddCommand(users_skillgroups.Cmdusers_skillgroups())
	usersCmd.AddCommand(users_state.Cmdusers_state())
	usersCmd.AddCommand(users_station.Cmdusers_station())
	usersCmd.AddCommand(users_trustors.Cmdusers_trustors())
	usersCmd.AddCommand(users_bulk.Cmdusers_bulk())
	usersCmd.AddCommand(users_search.Cmdusers_search())
	usersCmd.AddCommand(users_development.Cmdusers_development())
	usersCmd.Short = utils.GenerateCustomDescription(usersCmd.Short, users_verifiers.Description, users_agentui.Description, users_roles.Description, users_callforwarding.Description, users_superiors.Description, users_directreports.Description, users_favorites.Description, users_adjacents.Description, users_profile.Description, users_geolocations.Description, users_greetings.Description, users_outofoffice.Description, users_presences.Description, users_chats.Description, users_externalid.Description, users_me.Description, users_password.Description, users_profileskills.Description, users_queues.Description, users_invite.Description, users_routinglanguages.Description, users_routingskills.Description, users_routingstatus.Description, users_skillgroups.Description, users_state.Description, users_station.Description, users_trustors.Description, users_bulk.Description, users_search.Description, users_development.Description, )
	usersCmd.Long = usersCmd.Short
}
