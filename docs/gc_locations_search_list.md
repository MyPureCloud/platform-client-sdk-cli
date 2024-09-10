## gc locations search list

Search locations using the q64 value returned from a previous search

### Synopsis

Search locations using the q64 value returned from a previous search

```
gc locations search list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --expand strings           Provides more details about a specified resource Valid values: images, addressVerificationDetails
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --q64 string               q64 - REQUIRED
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

* [gc locations search](gc_locations_search.html)	 - /api/v2/locations/search


