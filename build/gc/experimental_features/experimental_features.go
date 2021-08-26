package experimental_features

import (
	"fmt"

	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/config"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/models"
)

var allCommands = [...]models.ExperimentalFeature{models.AlternativeFormats}

func IsValidCommand(command string) bool {
	for _, c := range allCommands {
		if command == c.String() {
			return true
		}
	}
	// Temporary clause to maintain dummy_command as valid without listing it in ListAllFeatures
	if command == "dummy_command" {
		return true
	}

	return false
}

func ListAllFeatures(profileName string) {
	if len(allCommands) == 0 {
		fmt.Println("No experimental features have been deployed yet.")
		return
	}

	for _, c := range allCommands {
		fmt.Printf("%v (enabled: %v) - %v\n", c.String(), config.GetExperimentalFeature(profileName, c.String()), c.Description())
	}
}