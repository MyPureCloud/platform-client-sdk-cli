## gc organizations limits namespaces limits counts list

Get estimated limit counts for a namespace and limit name. This is not a source of truth for limit values but a record of estimates to facilitate limit threshold tracking.

### Synopsis

Get estimated limit counts for a namespace and limit name. This is not a source of truth for limit values but a record of estimates to facilitate limit threshold tracking.

```
gc organizations limits namespaces limits counts list [namespaceName] [limitName] [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --cursor string            Cursor provided when retrieving the last page
      --entityId string          entity id of the count
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
  -s, --stream                   Paginate and stream data as it is being processed leaving page information intact
      --userId string            userid of the count
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

* [gc organizations limits namespaces limits counts](gc_organizations_limits_namespaces_limits_counts.html)	 - /api/v2/organizations/limits/namespaces/{namespaceName}/limits/{limitName}/counts


