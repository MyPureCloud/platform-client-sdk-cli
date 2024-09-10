## gc authorization roles list

Retrieve a list of all roles defined for the organization

### Synopsis

Retrieve a list of all roles defined for the organization

```
gc authorization roles list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --defaultRoleId strings    
      --expand strings           variable name requested by expand list
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --id strings               id
      --name string              
      --nextPage string          next page token
      --pageNumber int           The page number requested (default 1)
      --pageSize int             The total page size requested (default 25)
      --permission strings       
      --previousPage string      Previous page token
      --sortBy string            variable name requested to sort by
  -s, --stream                   Paginate and stream data as it is being processed leaving page information intact
      --userCount string          Valid values: true, false
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


