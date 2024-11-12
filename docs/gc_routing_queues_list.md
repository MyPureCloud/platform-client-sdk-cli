## gc routing queues list

Get list of queues.

### Synopsis

Get list of queues.

```
gc routing queues list [flags]
```

### Options

```
  -a, --autopaginate                     Automatically paginate through the results stripping page information
      --cannedResponseLibraryId string   Include only queues explicitly associated with the specified canned response library ID
      --divisionId strings               Include only queues in the specified division ID(s)
      --filtercondition string           Filter list command output based on a given condition or regular expression
      --hasPeer string                   Include only queues with a peer ID Valid values: true, false
  -h, --help                             help for list
      --id strings                       Include only queues with the specified ID(s)
      --name string                      Include only queues with the given name (leading and trailing asterisks allowed)
      --pageNumber string                Page number (default "1")
      --pageSize string                  Page size (default "25")
      --peerId strings                   Include only queues with the specified peer ID(s)
      --sortOrder string                 Note: results are sorted by name. Valid values: asc, desc
  -s, --stream                           Paginate and stream data as it is being processed leaving page information intact
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

* [gc routing queues](gc_routing_queues.html)	 - /api/v2/routing/queues


