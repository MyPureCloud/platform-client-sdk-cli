## gc alerting interactionstats rules list

Get an interaction stats rule list

### Synopsis

Get an interaction stats rule list

```
gc alerting interactionstats rules list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --expand strings           Which fields, if any, to expand Valid values: notificationUsers
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

* [gc alerting interactionstats rules](gc_alerting_interactionstats_rules.html)	 - /api/v2/alerting/interactionstats/rules

