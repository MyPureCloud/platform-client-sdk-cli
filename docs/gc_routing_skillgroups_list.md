## gc routing skillgroups list

Get skill group listing

### Synopsis

Get skill group listing

```
gc routing skillgroups list [flags]
```

### Options

```
      --after string             The cursor that points to the next item
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --before string            The cursor that points to the previous item
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --name string              Return only skill group names whose names start with this value (case-insensitive matching)
      --pageSize string          Page size (default "25")
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

* [gc routing skillgroups](gc_routing_skillgroups.html)	 - /api/v2/routing/skillgroups


