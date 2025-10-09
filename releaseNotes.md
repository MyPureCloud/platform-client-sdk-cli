Platform API version: 9648


## Release Notes
When creating a profile with the gc profiles new command, for an Implicit Grant or a PKCE Grant, the user is redirected to their browser where they can authenticate themselves by logging into their Genesys Cloud org.
By default, the linux version of the CLI will use the "xdg-open" command to open the browser.
In order to provide flexibility on Windows Subsystem for Linux (WSL), where "xdg-open" may not be installed, you can make use of the GENESYSCLOUD_BROWSER environment variable to change the browser command. Set the GENESYSCLOUD_BROWSER environment variable with the desired command before running the CLI.

# Major Changes (0 changes)


# Minor Changes (0 changes)


# Point Changes (0 changes)
