## gc authorization roles get

Get a single organization role.

### Synopsis

Get a single organization role.

```
gc authorization roles get [roleId] [flags]
```

### Options

```
      --expand strings   Which fields, if any, to expand. unusedPermissions returns the permissions not used for the role Valid values: unusedPermissions
  -h, --help             help for get
      --userCount true   Fetch the count of users who have this role granted in at least one division. Setting this value or defaulting to true can lead to slower load times or timeouts for role queries with large member counts. Valid values: true, false Valid values: true, false
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


