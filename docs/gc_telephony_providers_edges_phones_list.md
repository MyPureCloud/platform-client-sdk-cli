## gc telephony providers edges phones list

Get a list of Phone Instances. A maximum of 10,000 results is returned when filtering the results or sorting by a field other than the ID. Sorting by only the ID has no result limit. Each filter supports a wildcard, *, as a value to search for partial values.

### Synopsis

Get a list of Phone Instances. A maximum of 10,000 results is returned when filtering the results or sorting by a field other than the ID. Sorting by only the ID has no result limit. Each filter supports a wildcard, *, as a value to search for partial values.

```
gc telephony providers edges phones list [flags]
```

### Options

```
  -a, --autopaginate                              Automatically paginate through the results stripping page information
      --expand strings                            Fields to expand in the response, comma-separated Valid values: properties, site, status, status.primaryEdgesStatus, status.secondaryEdgesStatus, phoneBaseSettings, lines
      --fields strings                            Fields and properties to get, comma-separated Valid values: webRtcUser, properties.*, lines.loggedInUser, lines.defaultForUser
      --filtercondition string                    Filter list command output based on a given condition or regular expression
  -h, --help                                      help for list
      --linesDefaultForUserId string              Filter by lines.defaultForUser.id
      --linesId string                            Filter by lines.id
      --linesLoggedInUserId string                Filter by lines.loggedInUser.id
      --linesName string                          Filter by lines.name
      --name string                               Name of the Phone to filter by, comma-separated
      --pageNumber string                         Page number (default "1")
      --pageSize string                           Page size (default "25")
      --phoneBaseSettingsId string                Filter by phoneBaseSettings.id
      --phoneHardwareId string                    Filter by phone_hardwareId
      --secondaryStatusOperationalStatus string   The secondary status to filter by
      --siteId string                             Filter by site.id
      --sortBy string                             The field to sort by Valid values: id, name, status.operationalStatus, secondaryStatus.operationalStatus
      --sortOrder string                          Sort order
      --statusOperationalStatus string            The primary status to filter by
  -s, --stream                                    Paginate and stream data as it is being processed leaving page information intact
      --webRtcUserId string                       Filter by webRtcUser.id
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

* [gc telephony providers edges phones](gc_telephony_providers_edges_phones.html)	 - /api/v2/telephony/providers/edges/phones


