## gc completion

Generate completion script

### Synopsis

To load completions:

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


```
gc completion [bash|zsh|fish|powershell]
```

### Options

```
  -h, --help   help for completion
```

### Options inherited from parent commands

```
      --accesstoken string    accessToken override
      --clientid string       clientId override
      --clientsecret string   clientSecret override
      --environment string    environment override. E.g. mypurecloud.com.au or ap-southeast-2
  -i, --indicateprogress      Trace progress indicators to stderr
      --inputformat string    Data input format. Supported formats: YAML, JSON
      --outputformat string   Data output format. Supported formats: YAML, JSON
  -p, --profile string        Name of the profile to use for configuring the cli (default "DEFAULT")
      --transform string      Provide a Go template file for transforming output data
      --transformstr string   Provide a Go template string for transforming output data
```

### SEE ALSO

* [gc](gc.html)	 - gc is a CLI for interacting with Genesys Cloud


