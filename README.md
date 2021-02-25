---
title: Platform API CLI SDK
ispreview: true
---

This CLI focuses mainly on exposing the following operations:

- **get** - Get a defined object by ID
- **list** - Get a list of all the Genesys Cloud objects of a particular type for an organization
- **create** - Create a Genesys Cloud object via data passed in via stdin or via file.
- **update** - Update a Genesys Cloud object via data passed in via stdin or via file.
- **delete** - Delete a Genesys Cloud object via data passed in via stdin or via file.

**Note:** Most operations exposed by the `list` command and some operations exposed by the `get` command support the `pageSize` parameter. It's important to set this to an appropriate value if a large number of resources are being listed to reduce the load on the API and reduce wait times.

At the time of writing the following objects are supported by the CLI:
- campaigns
- divisions
- edges
- externalcontacts
- groups
- locations
- notifications
- phones
- queues
- responsemgt
- roles
- sites
- skills
- stations
- usage
- users

In addition, a few objects provide additional capabilities, including:
- edges - Rebooting Edges
- phones - Rebooting Phones
- queues - Get estimated wait time
- queues - Get users assigned to queue
- notifications - Manage and listen to notifications from the event service
- usage - Retrieve Usage Reports

# Installation

The [Developer Center](https://developer.mypurecloud.com/api/rest/command-line-interface/) contains installation instructions for supported operating systems.

# Setup and Configuration
All access to Genesys Cloud is provided via OAuth client credentials.  Security is enforced via the OAuth client credentials.

To setup and configure your gc CLI run `gc profiles new` command and answer the questions.  If everything works correctly you should have a file created in your home directory called .gc/config.toml.  If for some reason you have problems with the `gc profiles new` command, you can manually create the config file by following the steps below.

-  Create .gc directory in your home directory
-  Create a config.toml file in this directory.

```
[DEFAULT]
environment="mypurecloud.com" 
client_credentials="OAUTH CLIENT CREDENTIAL GRANT"
client_secret="OAUTH CLIENT SECRET"

[hollywoo]
environment="mypurecloud.com"
client_credentials="OAUTH CLIENT CREDENTIAL GRANT"
client_secret="OAUTH CLIENT SECRET"
```

**Note:** You can setup up multiple profiles.  The default profile is what will be used by the CLI by default.  You can use a different profile by passing in a `-p=profile_name` flag on the CLI.

# Using the CLI
The CLI follows standard POSIX command name and command flag parameter styles.  To see all of the available objects you can issue a `gc` command.  To see all the sub-commands under a particular entity (eg. users) type `gc <<subcommand>>`.  For example to see all of the users in the org you can type `gc users list --autopaginate`.

# Application logging
The CLI logs Information, Warning and Fatal data to a log file, the location of which depends on the operating system:

### Linux
`/tmp/GenesysCloud`

### Windows
`%TEMP%\GenesysCloud`

### macOS
`~/Library/Logs/GenesysCloud`

# Tracing progress information
Passing the flag `-i` or `--indicateprogress` to any command will result in progress information traced to stderr and written to the application log file.  For example to see progress information for a list operation and ignore API output, use `gc users list --autopaginate -i > /dev/null`.

# Pagnation
As of version `3.0.0` the default behaviour will be to *not* automatically paginate any paginatable resources. To automatically paginate, you must pass the `--autopaginate` or `-a` flag. In addition, there is a new `--stream` or `-s` flag for paginatable resources. This will paginate through the results and print them one page at a time leaving the page information intact.

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
gc users delete 
```

# Additional Tools
Since this is a CLI, the output from the tool can be passed to other command tools and scripts.  Two of the most common helpful tools are:

[jq](https://stedolan.github.io/jq/) - A JSON query, filter and transformation tools.  This tool can do just about anything with JSON.  For instance to retrieve all of the members of a specific queue, lookup their CLI and names and transform the output into CSV you issue the following CSV you can issue the following `gc` and `jq` commands:

```
gc queues list --autopaginate | jq -c '.[] | select( .name | contains("Chat2"))' | jq -r '. | .id' | xargs -I{} gc queues users {} | jq -r '.[] | [.id,.name] | @csv'> output.csv
```

This command:
- retrieves a list of all the queues
- filters out the queue named Chat2 
- parses out its CLI
- uses xargs to pipe the CLI over to the `gc queue users` command
- parses out the CLI and name field for each record
- formats results into a csv format 
- redirects the output to a file called output.csv

[yq](https://github.com/mikefarah/yq) - A YAML conversion and manipulation tool.  You could use `yq` to convert yaml to and from JSON.  For example, to convert the output of the `gc users list` to yaml via `gc users list --autopaginate | yq r - -P`

## Versioning

The SDK's version is incremented according to the [Semantic Versioning Specification](https://semver.org/). The decision to increment version numbers is determined by [diffing the Platform API's swagger](https://github.com/purecloudlabs/platform-client-sdk-common/blob/master/modules/swaggerDiff.js) for automated builds, and optionally forcing a version bump when a build is triggered manually (e.g. releasing a bugfix).


## Support

This package is intended to be forwards compatible with v2 of Genesys Cloud's Platform API. While the general policy for the API is not to introduce breaking changes, there are certain additions and changes to the API that cause breaking changes for the SDK, often due to the way the API is expressed in its swagger definition. Because of this, the SDK can have a major version bump while the API remains at major version 2. While the CLI SDK is intended to be forward compatible, patches will only be released to the latest version. For these reasons, it is strongly recommended that all applications using this CLI SDK are kept up to date and use the latest version of the CLI SDK.

For any issues, questions, or suggestions for the CLI SDK, visit the [Genesys Cloud Developer Forum](https://developer.mypurecloud.com/forum/).
