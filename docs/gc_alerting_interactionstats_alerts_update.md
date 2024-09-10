## gc alerting interactionstats alerts update

Update an interaction stats alert read status

### Synopsis

Update an interaction stats alert read status

```
gc alerting interactionstats alerts update [alertId] [flags]
```

### Options

```
  -d, --directory string   Directory path with files containing request bodies
      --expand strings     Which fields, if any, to expand Valid values: notificationUsers
  -f, --file string        File name containing the JSON body
  -h, --help               help for update
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

* [gc alerting interactionstats alerts](gc_alerting_interactionstats_alerts.html)	 - /api/v2/alerting/interactionstats/alerts


