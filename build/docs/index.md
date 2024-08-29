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

To setup and configure your gc CLI run `gc profiles new` command and answer the questions. If everything works correctly you should have a file created in your home directory called .gc/config.toml.  If for some reason you have problems with the `gc profiles new` command, you can manually create the config file by following the steps below.

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

# Preview API's

Preview API's are included in the CLI. These resources are subject to both breaking and non-breaking changes at any time without notice. This includes, but is not limited to, changing resource names, paths, contracts, documentation, and removing resources entirely. For a full list of the preview API's see [here](https://developer.genesys.cloud/platform/preview-apis)

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
As of version `3.0.0` the default behaviour will be to *not* automatically paginate any paginatable resources. 
To automatically paginate, you must pass the `--autopaginate` or `-a` flag. 

Alternatively, you can set auto-pagination in your configuration file.

To enable auto-pagination, run the following command:

```
gc autopagination enable
```

To disable auto-pagination:

```
gc autopagination disable
```

To check if auto-pagination in config is enabled or disabled:

```
gc autopagination status
```
In addition, there is a new `--stream` or `-s` flag for paginatable resources. This will paginate through the results and print them one page at a time leaving the page information intact.

# Proxy Configuration

To add a proxy configuration for the CLI , you can pass file parameter with proxy configuration
in it.

The `ProxyConfiguration` object has 3 properties that determine the URL for proxying.
Port - Port of the Proxy server
Host - Host Ip or DNS of the proxy server
Protocol - Protocol required to connect to the Proxy (http or https)

The 'ProxyConfiguration' has another section which is an optional section. If the proxy requires authentication to connect to
'userName' and 'password' needs to be mentioned under the 'ProxyConfiguration'. This section can be removed from the JSON file if no authentication is required.

If you have different proxyParams for authorise APIs and other APIs, you can provide a new key in pathParams section with name 'login' for authorise sdk calls.

JSON configuration file:
```
{
  "host": "hostname",
  "protocol": "http",
  "port": "8888",
  "userName": "username",
  "password": "password",
  "pathParams": {
    "login": "loginpath"
  }
}
```
Command to Enable Proxy:
```
gc proxy --file=proxy.json
```
To disable this Proxy Configuration, run the following command.
```
gc proxy disable
```
# Gateway Configuration

To add a Gateway configuration for the CLI , you can pass file parameter with Gateway configuration
in it.

The `GateWayConfiguration` object has 3 properties that determine the URL for a Gateway.
Port - Port of the Gateway
Host - Host Ip or DNS of the Gateway
Protocol - Protocol required to connect to the Gateway (http or https)

The 'GateWayConfiguration' has another section which is an optional section. If the Gateway requires authentication to connect to
'userName' and 'password' needs to be mentioned under the 'GateWayConfiguration'. This section can be removed from the JSON file if no authentication is required.

If you have different pathParams for authorise APIs and other APIs, you can provide a new key in pathParams section with name 'login' for authorise sdk calls and 'api' for other api calls.

JSON configuration file:
```
{
  "host": "hostname",
  "protocol": "http",
  "port": "8888",
  "userName": "username",
  "password": "password",
  "pathParams": {
    "login": "loginpath",
    "api":"apipath"
  }
}
```
Command to Enable Gateway:
```
gc gateway --file=gateway.json
```
To disable this Gateway Configuration, run the following command.
```
gc gateway disable
```

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

# OAuth Implicit Grant

When creating a profile with the `gc profiles new` command, you may choose between using a Client Credential Grant, an Implicit Grant or a PKCE Grant. During the Implicit Login process, the user is redirected to their browser where they can authenticate themselves by logging into their Genesys Cloud org. This method can be preferable as it provides context about the user to which the access token was distributed. Unlike the Client Credential Grant, which is a general grant.

To perform an implicit login, the OAuth Client under **`Admin > Integrations > OAuth`** needs to be configured properly by setting the **Grant Types** property to "Token Implicit Grant (Browser)", and by adding the permissible redirect URIs (e.g. `http://localhost:8080`).

Optionally, and for added security, you can choose to open an HTTPS connection for this procedure. In this case, a self-signed certificate is generated locally, and you will need to select **Advanced > Proceed to 127.0.0.1** in the browser window as this certificate is not recognized by an official Certificate Authority. For secure HTTP connections, be sure to prepend the Client's redirect URI with "https://" instead of "http://" in your Genesys Cloud org.

For more information about the Implicit Grant login process, check out [this article](https://developer.genesys.cloud/api/rest/authorization/use-implicit-grant) on our Developer Center.

# OAuth PKCE Grant

When creating a profile with the `gc profiles new` command, you may choose between using a Client Credential Grant, an Implicit Grant or a PKCE Grant. During the PKCE Login process, the user is redirected to their browser where they can authenticate themselves by logging into their Genesys Cloud org. This method can be preferable as it provides context about the user to which the access token was distributed. Unlike the Client Credential Grant, which is a general grant.

To perform a PKCE login, the OAuth Client under **`Admin > Integrations > OAuth`** needs to be configured properly by setting the **Grant Types** property to "Code Authorization", and by adding the permissible redirect URIs (e.g. `http://localhost:8080`).

Optionally, and for added security, you can choose to open an HTTPS connection for this procedure. In this case, a self-signed certificate is generated locally, and you will need to select **Advanced > Proceed to 127.0.0.1** in the browser window as this certificate is not recognized by an official Certificate Authority. For secure HTTP connections, be sure to prepend the Client's redirect URI with "https://" instead of "http://" in your Genesys Cloud org.

For more information about the PKCE Grant login process, check out [this article](https://developer.genesys.cloud/api/rest/authorization/use-pkce) on our Developer Center.

# Experimental Features

Experimental `CLI` features were released in version 18.0.0 of the `Genesys Cloud CLI`. Experimental features allow us to release ideas around the `CLI` early and give customers an opportunity to give feedback on the features before the work on them is finalized. Experimental features are tied to the `CLI` binary, so when we release an experimental feature or promote an experimental feature, you will need to download the `CLI` binary that matches the release or promotion of that feature.

**Note:** Experimental `CLI` features are subject to breaking and non-breaking changes at any time.


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

The `Alternative Formats` feature allows you to specify the input and output format of the `CLI`. Alternative formats are provided to the `Genesys Cloud CLI` by passing the preferred input format to the `--inputformat` flag or by passing the preferred output format to the `--outputformat` flag.   The currently supported formats are: `YAML` and `JSON`.

**Note:** The default format for the `CLI` is `JSON`.

To use the `Alternative Formats` feature, pass a supported value (e.g `YAML` or `JSON`) to the `--inputformat` flag or to the `--outputformat` flag.

To input `YAML` data:

```
--inputformat=yaml
```
To output `YAML` data:

```
--outputformat=yaml
```

**Note:** The `--inputformat` flag and the `--outputformat` flag are not case-sensitive. Writing `--inputformat=yaml` is the same as writing `--inputformat=YAML`.

### Input Format

If you want to input a `YAML` file in the `CLI`, pass the file path to the `--file` flag and include the `--inputformat=yaml` command when making a request.

**Example Request:**  

```
gc analytics conversations details query create --file=./query.yaml --inputformat=yaml
```

### Output Format

If you want to format your output data as `YAML`, include the `--outputformat=yaml` command when making a request.

**Example Request:**  

```
gc users get f3dc94ca-acec-4ee4-a07e-ca7503ddbd62 --outputformat=yaml
```

### Setting Input and Output Formats In Config

Additionally, the desired input and output formats can be pinned in the configuration file to avoid providing the above flags in every API call.

To set the output format to `YAML` or `JSON`, run the following command:

```
gc alternativeformats setoutput [format]
```

And similarly, to set the input format:

```
gc alternativeformats setinput [format]
```

The following commands can be used to check which data formats are set in the configuration file.

To get the input format: 

```
gc alternativeformats getinput
```

To get the output format:

```
gc alternativeformats getoutput
```

**Note:** Where no data format is specified, the `CLI` will default to `JSON`.

## Transform Data

The `Transform Data` feature uses `Go` templates for transforming output data. Templates are provided to the `Genesys Cloud CLI` as either a file containing the template string and passed to the `--transform` flag or as a "raw" template string and input directly in the terminal and passed to the `--transformstr` flag. If you would like to read more about `Go` templates, see Go's template package [text/template](https://pkg.go.dev/text/template) for more info.

**Note:**  You may also use [sprig template functions](http://masterminds.github.io/sprig/) with the Transform Data feature.

To input a local `Go` template file:

```
--transform=./tmpl.gotmpl
```

To input a `Go` template file stored in a remote GitHub repository, any of the following approaches are acceptable:
```
--transform=https://github.com/foobar/path/to/file.gotmpl
--transform=github.com/foobar/path/to/file.gotmpl
--transform=https://raw.githubusercontent.com/foobar/path/to/file.gotmpl
--transform=raw.githubusercontent.com/foobar/path/to/file.gotmpl
```

To input a `Go` template string:

```
--transformstr='your_go_template_string'
```

### Template File

If you want to input a `Go` template file in the `CLI`, pass the template file path to the `--transform` flag when making a request.

**Example Request:**

```
gc users get f3dc94ca-acec-4ee4-a07e-ca7503ddbd62 --transform=./tmpl.gotmpl
```
### Template String

If you want to input a `Go` template string in the `CLI`, pass the "raw" template string to the `--transformstr` flag surrounded by single quotes when making a request.

**Example Request:**

```
gc users get f3dc94ca-acec-4ee4-a07e-ca7503ddbd62 --transformstr='your_go_template_string'
```
### Transforming Output Data To CSV Format

**Note:** To write a `Go` template, you must know the structure of the data in advance.

This is an example `Go` template file for transforming output data to `CSV` format.

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

So to get a list of users and transform the output to `CSV` format, run the following command:

```
gc users list -a --transform=./tmpl.gotmpl
```


## List Filtering

The `filtercondition` flag allows you to filter the output of a list command to include only objects that match a given condition. For example, if you wish to list only users with the name Foo Bar, you could run the following command:

```bash
gc users list --autopaginate --filtercondition="name==Foo Bar"
```

To filter based on a nested attribute, separate each key with a full stop. 

```bash
gc users list --autopaginate --filtercondition="division.name==Home"
```

Other examples:

* Integers: 
  * `gc outbound dnclists list --autopaginate --filtercondition="version==5"`
  * `gc outbound dnclists list --autopaginate --filtercondition="version<5"`
* Booleans:
  * `gc routing queues list --filtercondition="autoAnswerOnly==true"`
* String/Array contains:
  * `gc users list --autopaginate --filtercondition="name contains Foo"`
  * `gc outbound contactlists list --autopaginate --filtercondition="columnNames contains Cell"`
* Regex (the expression to the right of the match operator is passed directly to the Match function is the regexp package. See the documentation for this function [here](https://pkg.go.dev/regexp#Match)):
  * `gc users list --autopaginate --filtercondition="name match ^Foo (.{3})$"`

A few things to note: 
* AND/OR operators are not yet supported. 
* When no matching records are found, the string value "null" is returned. 
* This feature is available as of version 49.2.0

## Additional Resources

1. Experimental CLI feature: Alternative Formats: [Introducing Alternative Formats - blog post](https://developer.genesys.cloud/blog/2021-08-31-new-experimental-cli-feature-alternative-formats/).
2. Experimental CLI feature: Transform Data: [Introducing Transform Data - blog post](https://developer.genesys.cloud/blog/2021-10-01-experimental-feature-transform-data/)
3. Go's template package documentation: [text/template](https://pkg.go.dev/text/template).
4. The sprig library: [sprig template functions](http://masterminds.github.io/sprig/).

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

To create or update multiple objects, you can pass a directory path to the `--directory` flag or to the `-d` short flag containing JSON files, where each file is a request body for either creating or updating data.

For example, creating multiple user objects:

`./users-directory`

```
user-1.json // each file is a request body for creating a user
user-2.json
user-3.json
```

To create multiple users, run the following command:

```
gc users create -d ./users-directory
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
