## gc analytics agents status counts create

Count agents by different groupings

### Synopsis

Count agents by different groupings

```
gc analytics agents status counts create [flags]
```

### Options

```
  -d, --directory string   Directory path with files containing request bodies
  -f, --file string        File name containing the JSON body
      --groupBy strings    Include to choose which groupings to count by and return. If not included it will return only counts grouped by segmentType Valid values: segmentType, presence, routingStatus, isOutOfOffice
  -h, --help               help for create
  -b, --printrequestbody   Print the request body format of the API.
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

* [gc analytics agents status counts](gc_analytics_agents_status_counts.html)	 - /api/v2/analytics/agents/status/counts


