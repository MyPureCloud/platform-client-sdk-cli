## gc quality calibrations list

Get the list of calibrations

### Synopsis

Get the list of calibrations

```
gc quality calibrations list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --calibratorId string      user id of calibrator - REQUIRED
      --conversationId string    conversation id
      --endTime string           end of the calibration query. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
      --expand strings           variable name requested by expand list
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --nextPage string          next page token
      --pageNumber string        The page number requested (default "1")
      --pageSize string          The total page size requested (default "25")
      --previousPage string      Previous page token
      --sortBy string            variable name requested to sort by
      --startTime string         Beginning of the calibration query. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
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

* [gc quality calibrations](gc_quality_calibrations.html)	 - /api/v2/quality/calibrations


