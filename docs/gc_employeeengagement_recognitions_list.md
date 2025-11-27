## gc employeeengagement recognitions list

Gets sent recognitions

### Synopsis

Gets sent recognitions

```
gc employeeengagement recognitions list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --dateEnd string           The end date of the search range. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
      --dateStart string         The start date of the search range. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
      --direction string         The direction of the recognitions. Valid values: sent, received
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --pageNumber string        Page number (default "1")
      --pageSize string          Page size (default "100")
      --recipient string         The ID of the recipient (when direction is sent).
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

* [gc employeeengagement recognitions](gc_employeeengagement_recognitions.html)	 - /api/v2/employeeengagement/recognitions


