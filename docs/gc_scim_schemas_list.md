## gc scim schemas list

Get a list of SCIM schemas

### Synopsis

Get a list of SCIM schemas

```
gc scim schemas list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --filter string            Filtered results are invalid and return 403 Unauthorized.
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
  -s, --stream                   Paginate and stream data as it is being processed leaving page information intact
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

* [gc scim schemas](gc_scim_schemas.html)	 - /api/v2/scim/schemas

