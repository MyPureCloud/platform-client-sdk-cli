## gc telephony providers edges trunkbasesettings list

Get Trunk Base Settings listing

### Synopsis

Get Trunk Base Settings listing

```
gc telephony providers edges trunkbasesettings list [flags]
```

### Options

```
  -a, --autopaginate              Automatically paginate through the results stripping page information
      --expand strings            Fields to expand in the response, comma-separated Valid values: properties
      --filtercondition string    Filter list command output based on a given condition or regular expression
  -h, --help                      help for list
      --ignoreHidden string       Set this to true to not receive trunk properties that are meant to be hidden or for internal system usage only. Valid values: true, false
      --managed string            Filter by managed Valid values: true, false
      --name string               Name of the TrunkBase to filter by
      --pageNumber int            Page number (default 1)
      --pageSize int              Page size (default 25)
      --recordingEnabled string   Filter trunks by recording enabled Valid values: true, false
      --sortBy string             Value by which to sort
      --sortOrder string          Sort order
  -s, --stream                    Paginate and stream data as it is being processed leaving page information intact
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

* [gc telephony providers edges trunkbasesettings](gc_telephony_providers_edges_trunkbasesettings.html)	 - /api/v2/telephony/providers/edges/trunkbasesettings


