## gc analytics dataextraction downloads metadata list

Get metadata on files available for extraction

### Synopsis

Get metadata on files available for extraction

```
gc analytics dataextraction downloads metadata list [flags]
```

### Options

```
      --after string             The cursor that points to the end of the set of entities that has been returned.
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --before string            The cursor that points to the start of the set of entities that has been returned.
      --dataSchema string        Data schema like conversations
      --dateEnd string           End DateTime filter. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
      --dateStart string         Start DateTime filter. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --pageSize string          Number of entities to return. Maximum of 200.
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

* [gc analytics dataextraction downloads metadata](gc_analytics_dataextraction_downloads_metadata.html)	 - /api/v2/analytics/dataextraction/downloads/metadata


