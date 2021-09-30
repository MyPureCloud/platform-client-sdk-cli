---
title: Platform API CLI SDK
---

This CLI focuses mainly on exposing the following operations:

- **get** - Get a defined object by ID
- **list** - Get a list of all the Genesys Cloud objects of a particular type for an organization
- **create** - Create a Genesys Cloud object via data passed in via stdin or via file.
- **update** - Update a Genesys Cloud object via data passed in via stdin or via file.
- **delete** - Delete a Genesys Cloud object via data passed in via stdin or via file.

**Note:** Most operations exposed by the `list` command and some operations exposed by the `get` command support the `pageSize` parameter. It's important to set this to an appropriate value if a large number of resources are being listed to reduce the load on the API and reduce wait times.

# Installation

The [Developer Center](https://developer.mypurecloud.com/api/rest/command-line-interface/) contains installation instructions for supported operating systems.

# Setup and Configuration

All access to Genesys Cloud is provided via OAuth client credentials.  Security is enforced via the OAuth client credentials.

If you are using the CLI in a pipeline or don't wish to set up a config file, see [environment variables](#environment-variables) and [parameter overrides](#parameter-overrides).  

The "environment" value can be provided as the API base path such as `mypurecloud.com` or AWS region.

To setup and configure your gc CLI run `gc profiles new` command and answer the questions.  If everything works correctly you should have a file created in your home directory called .gc/config.toml.  If for some reason you have problems with the `gc profiles new` command, you can manually create the config file by following the steps below.

-  Create .gc directory in your home directory
-  Create a config.toml file in this directory.

```
[DEFAULT]
environment="mypurecloud.com" 
client_credentials="OAUTH CLIENT CREDENTIAL GRANT"
client_secret="OAUTH CLIENT SECRET"

[test_pro_1]
environment="mypurecloud.com"
client_credentials="OAUTH CLIENT CREDENTIAL GRANT"
client_secret="OAUTH CLIENT SECRET"

[test_pro_2]
environment="mypurecloud.com"
access_token="OAUTH ACCESS TOKEN"
```

**Note:** You can setup up multiple profiles.  The default profile is what will be used by the CLI by default.  You can use a different profile by passing in a `-p=profile_name` flag on the CLI.

## Environment variables

The following environment variables can be used as overrides or alternatives to their corresponding config file values

```
GENESYSCLOUD_REGION
GENESYSCLOUD_OAUTHCLIENT_ID
GENESYSCLOUD_OAUTHCLIENT_SECRET
GENESYSCLOUD_ACCESS_TOKEN
```

`GENESYSCLOUD_REGION` can refer to the API base path or AWS region

## Parameter overrides

The following flags can be used as overrides or alternatives to their corresponding config file values

```
--environment
--clientid
--clientsecret
--accesstoken
```

# Using the CLI
The CLI follows standard POSIX command name and command flag parameter styles.  To see all of the available objects you can issue a `gc` command.  To see all the sub-commands under a particular entity (eg. users) type `gc <<subcommand>>`.  For example to see all of the users in the org you can type `gc users list --autopaginate`.

# Application logging

By default, the CLI does not log information to a file. To enable logging, use the following command:

```
gc logging enable
```  

And to disable:

```
gc logging disable
```

The default location of the log file depends on the operating system:

### Linux
`/tmp/GenesysCloud/gc.log`

### Windows
`%TEMP%\GenesysCloud\gc.txt`

### macOS
`~/Library/Logs/GenesysCloud/gc.log`

To set a custom path, use the following command:

```
gc logging path /path/to/log.log
```

Logging is configured on a per-profile basis so the above commands will only configure logging for the default profile.

# Tracing progress information
Passing the flag `-i` or `--indicateprogress` to any command will result in progress information traced to stderr and written to the application log file.  For example, to see progress information for a list operation and ignore API output, use `gc users list --autopaginate -i > /dev/null`.

# Notifications
Create a channel:
```
gc notifications channels create
```
List all channels created using the current OAuth token:
```
gc notifications channels list
```
Listen to notifications from a channel:
```
gc notifications channels listen [CHANNEL_ID]
```
Using the `--noheartbeat` flag to filter out the heartbeat from the event stream:
```
gc notifications channels listen [CHANNEL_ID] --nohearbeat
```

# Pagination
As of version `3.0.0` the default behaviour will be to *not* automatically paginate any paginatable resources. To automatically paginate, you must pass the `--autopaginate` or `-a` flag. In addition, there is a new `--stream` or `-s` flag for paginatable resources. This will paginate through the results and print them one page at a time leaving the page information intact.

# Autocompletion

See below for instructions on how to enable autocompletion for supported shells:

### Bash (macOs)

Firstly, install `bash-completion` if it is not already installed:

```
brew install bash-completion
```

Configure the `bash-completion` package in your `~/.bash_profile` or `~/.profile` by adding the following:

```
if [ -f $(brew --prefix)/etc/bash_completion ]; then
    . $(brew --prefix)/etc/bash_completion
fi
```

Finally, run the following command:

```
gc completion bash > /usr/local/etc/bash_completion.d/gc
```

### Bash (Linux)

```
gc completion bash > /etc/bash_completion.d/gc
```

### Zsh

If shell completion is not already enabled in your environment, you will need to enable it. You can execute the following once:

```
echo "autoload -U compinit; compinit" >> ~/.zshrc
```

```
gc completion zsh > "${fpath[1]}/_gc"
```

You will need to start a new shell for this setup to take effect.

### Fish

```
gc completion fish > ~/.config/fish/completions/gc.fish
```

### Powershell

```
gc completion powershell > gc.ps1
```

And source this file from your PowerShell profile.  

To see further examples of enabling autocompletion, run the following command:

```
gc completion --help
```

# Operation data

Running any API operation command followed by '--help' will print out the following information:

* Operation. e.g. `POST /api/v2/users`
* Permissions (where applicable)
* Documentation. e.g. https://developer.genesys.cloud/api/rest/v2/users/#post-api-v2-users

# Experimental Features

Experimental CLI features were released in version 18.0.0 of the Genesys Cloud CLI. Experimental features allow us to release ideas around the CLI early and give customers an opportunity to give feedback on the features before the work on them is finalized. Experimental features are tied to the CLI binary, so when we release an experimental feature or promote an experimental feature, you will need to download the CLI binary that matches the release or promotion of that feature.

**Note:** Experimental CLI features are subject to breaking and non-breaking changes at any time.


To see the list of experimental features, run the following command:

```
gc experimental list
```

To enable an experimental feature:

```
gc experimental enable [feature_name]
```

To disable an experimental feature:

```
gc experimental disable [feature_name]
```

**Note:** By default, experimental features are turned off. To use an experimental feature you must explicitly enable the feature.

## Alternative Formats

To use the Alternative Formats feature, first enable the feature with:

```
gc experimental enable alternative_formats
```

Once enabled, we can pass a supported value (e.g `YAML` or `JSON`) to the `--inputformat` flag or to the `--outputformat` flag.

To input `YAML` data:

```
--inputformat=yaml
```
To output `YAML` data:

```
--outputformat=yaml
```

**Note:** The `--inputformat` and `--outputformat` flags are not case sensitive. Writing `--inputformat=yaml` is the same as writing `--inputformat=YAML`.

### Input Format

If you want to input a `YAML` file when making a request, simply include the `--inputformat=yaml` command and pass the file path to the `--file` flag.

**Example Request:**  

```
gc analytics conversations details query create --file=./query.yaml --inputformat=yaml
```

### Output Format

If you want to format your response as `YAML`, simply include the `--outputformat=yaml` command.

**Example Request:**  

```
gc users get f3dc94ca-acec-4ee4-a07e-ca7503ddbd62 --outputformat=yaml
```

<br>
Additionally, the desired input and output formats can be pinned in the configuration file to avoid providing the above flags in every API call.

To set the output format to YAML or JSON, run:

```
gc alternativeformats setoutput [format]
```

And similarly, to set the input format:

```
gc alternativeformats setinput [format]
```

The following commands can be used to check which data formats are set in the configurations:

```
gc alternativeformats getinput
```

```
gc alternativeformats getoutput
```

**Note:** Where no data format is specified, the default is always JSON.

## Transform Data

**Note:** The Transform Data feature uses `Go` templates for transforming output data. Checkout Go's template package [text/template](https://pkg.go.dev/text/template) for more info.

To use the Transform Data feature, first enable the feature with:

```
gc experimental enable transform_data
```
Once enabled, we can pass either a `Go` template file to the `--transform` flag or a `Go` template string to the `--transformstr` flag.

To input a `Go` template file:

```
--transform=./tmpl.gotmpl
```

To input a `Go` template string:

```
--transformstr='your_go_template_string'
```

### Template File

If you want to input a `Go` template file to transform output data, simply pass the template file path to the `--transform` flag when making a request.

**Example Request:**

```
gc users get f3dc94ca-acec-4ee4-a07e-ca7503ddbd62 --transform=./tmpl.gotmpl
```
### Template String

If you want to input a `Go` template string to transform output data, simply pass your template string to the `--transformstr` flag surrounded by single quotes.

**Example Request:**

```
gc users get f3dc94ca-acec-4ee4-a07e-ca7503ddbd62 --transformstr='your_go_template_string'
```
### Transforming Output Data To CSV Format

An example `Go` template file for transforming output data to `CSV` format.

**Note:** To write a `Go` template, you must know the structure of the data in advance.

`tmpl.gotmpl`

```
{{- range . -}}
id: {{.id}},name: {{.name}},division: 
    {{- range $key, $val := .division -}}
        {{$key}}: {{$val}},
    {{- end -}}
    chat:
    {{- range $key, $val := .chat -}}
        {{$key}}: {{$val}},
    {{- end -}}
    email: {{.email}},primaryContactInfo:
    {{- range .primaryContactInfo -}}
        {{- range $key, $val := . -}}
            {{$key}}: {{$val}},
        {{- end -}}
    {{- end -}}
    addresses:
    {{- range .addresses -}}
        {{- range $key, $val := . -}}
            {{$key}}: {{$val}},
        {{- end -}}
    {{- end -}}
    state: {{.state}},username: {{.username}},version: {{.version}},acdAutoAnswer: {{.acdAutoAnswer}},selfUri: {{.selfUri}},
{{- end -}}
```

To get a list of all users and transform the output to `CSV` format by using a `Go` template, run the following command:

```
gc users list -a --transform=./tmpl.gotmpl
```

## Additional Resources

1. Checkout the [Alternative Formats](https://developer.genesys.cloud/blog/2021-08-31-new-experimental-cli-feature-alternative-formats/) blog post for further reading.
2. See Go's [template package](https://pkg.go.dev/text/template) documentation for more info on Go templates.

# Examples

To create users you can pass in JSON via stdin or via the -f file path command.  So for example, to create the user in [gc/examples-json/create-user.json](https://github.com/MyPureCloud/platform-client-sdk-cli/blob/main/build/gc/example-json/create-user.json) you would issue the following commands:

```
cat examples-json/create-user.json | gc users create
```

or

```
gc users create -f examples-json/create-user.json
```

To perform an update:

```
cat examples-json/create-user.json | gc users update 
```

or

```
gc users create -f examples-json/create-user.json 
```

To perform a delete:

```
gc users delete [userId]
```

# Additional Tools
Since this is a CLI, the output from the tool can be passed to other command tools and scripts.  Two of the most common helpful tools are:

[jq](https://stedolan.github.io/jq/) - A JSON query, filter and transformation tools.  This tool can do just about anything with JSON.  For instance to retrieve all of the members of a specific queue, lookup their CLI and names and transform the output into CSV you issue the following CSV you can issue the following `gc` and `jq` commands:

```
gc routing queues list --autopaginate | jq -c '.[] | select( .name | contains("Chat2"))' | jq -r '. | .id' | xargs -I{} gc routing queues members {} | jq -r '.[] | [.id,.name] | @csv'> output.csv
```

This command:
- retrieves a list of all the queues
- filters out the queue named Chat2 
- parses out its CLI
- uses xargs to pipe the CLI over to the `gc routing queues members` command
- parses out the CLI and name field for each record
- formats results into a csv format 
- redirects the output to a file called output.csv

[yq](https://github.com/mikefarah/yq) - A YAML conversion and manipulation tool.  You could use `yq` to convert yaml to and from JSON.  For example, to convert the output of the `gc users list` to yaml via `gc users list --autopaginate | yq r - -P`

## Versioning

The SDK's version is incremented according to the [Semantic Versioning Specification](https://semver.org/). The decision to increment version numbers is determined by [diffing the Platform API's swagger](https://github.com/purecloudlabs/platform-client-sdk-common/blob/master/modules/swaggerDiff.js) for automated builds, and optionally forcing a version bump when a build is triggered manually (e.g. releasing a bugfix).


## Support

This package is intended to be forwards compatible with v2 of Genesys Cloud's Platform API. While the general policy for the API is not to introduce breaking changes, there are certain additions and changes to the API that cause breaking changes for the SDK, often due to the way the API is expressed in its swagger definition. Because of this, the SDK can have a major version bump while the API remains at major version 2. While the CLI SDK is intended to be forward compatible, patches will only be released to the latest version. For these reasons, it is strongly recommended that all applications using this CLI SDK are kept up to date and use the latest version of the CLI SDK.

For any issues, questions, or suggestions for the CLI SDK, visit the [Genesys Cloud Developer Forum](https://developer.mypurecloud.com/forum/).
