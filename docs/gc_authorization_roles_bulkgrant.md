## gc authorization roles bulkgrant

Bulk-grant subjects and divisions with an organization role.

### Synopsis

Bulk-grant subjects and divisions with an organization role.

```
gc authorization roles bulkgrant [roleId] [flags]
```

### Options

```
  -d, --directory string     Directory path with files containing request bodies
  -f, --file string          File name containing the JSON body
  -h, --help                 help for bulkgrant
  -b, --printrequestbody     Print the request body format of the API.
      --subjectType string   what the type of the subjects are (PC_GROUP, PC_USER or PC_OAUTH_CLIENT)
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

* [gc authorization roles](gc_authorization_roles.html)	 - /api/v2/authorization/roles


