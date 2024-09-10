## gc telephony providers edges logicalinterfaces list

Get edge logical interfaces.

### Synopsis

Get edge logical interfaces.

```
gc telephony providers edges logicalinterfaces list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --edgeIds string           Comma separated list of Edge Id`s - REQUIRED
      --expand strings           Field to expand in the response Valid values: externalTrunkBaseAssignments, phoneTrunkBaseAssignments
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

* [gc telephony providers edges logicalinterfaces](gc_telephony_providers_edges_logicalinterfaces.html)	 - /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces /api/v2/telephony/providers/edges/logicalinterfaces


