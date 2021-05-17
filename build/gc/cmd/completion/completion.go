package completion

import (
	"github.com/spf13/cobra"
	"os"
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long: `To load completions:

Bash:

  $ source <(gc completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ gc completion bash > /etc/bash_completion.d/gc
  # macOS:

  # Firstly, install bash-completion if it is not already installed:
  $ brew install bash-completion
  
  # Configure the bash-completion package in your ~/.bash_profile or ~/.profile by adding the following:
  if [ -f $(brew --prefix)/etc/bash_completion ]; then
	  . $(brew --prefix)/etc/bash_completion
  fi
  
  # Finally, run the following command:
  $ gc completion bash > /usr/local/etc/bash_completion.d/gc

Zsh:

  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ gc completion zsh > "${fpath[1]}/_gc"

  # You will need to start a new shell for this setup to take effect.

fish:

  $ gc completion fish | source

  # To load completions for each session, execute once:
  $ gc completion fish > ~/.config/fish/completions/gc.fish

PowerShell:

  PS> gc completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> gc completion powershell > gc.ps1
  # and source this file from your PowerShell profile.
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
	},
}

func Cmdcompletion() *cobra.Command {
	return completionCmd
}