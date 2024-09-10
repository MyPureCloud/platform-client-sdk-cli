## gc stations list

Get the list of available stations.

### Synopsis

Get the list of available stations.

```
gc stations list [flags]
```

### Options

```
  -a, --autopaginate              Automatically paginate through the results stripping page information
      --filtercondition string    Filter list command output based on a given condition or regular expression
  -h, --help                      help for list
      --id string                 Comma separated list of stationIds
      --lineAppearanceId string   lineAppearanceId
      --name string               Name
      --pageNumber int            Page number (default 1)
      --pageSize int              Page size (default 25)
      --sortBy string             Sort by
  -s, --stream                    Paginate and stream data as it is being processed leaving page information intact
      --userSelectable string     True for stations that the user can select otherwise false
      --webRtcUserId string       Filter for the webRtc station of the webRtcUserId
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

* [gc stations](gc_stations.html)	 - /api/v2/stations


