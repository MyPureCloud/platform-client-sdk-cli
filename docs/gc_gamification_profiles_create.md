## gc gamification profiles create

Create a new custom performance profile

### Synopsis

Create a new custom performance profile

```
gc gamification profiles create [flags]
```

### Options

```
      --copyMetrics string   Flag to copy metrics. If set to false, there will be no metrics associated with the new profile. If set to true or is absent (the default behavior), all metrics from the default profile will be copied over into the new profile. Valid values: true, false
  -d, --directory string     Directory path with files containing request bodies
  -f, --file string          File name containing the JSON body
  -h, --help                 help for create
  -b, --printrequestbody     Print the request body format of the API.
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

* [gc gamification profiles](gc_gamification_profiles.html)	 - /api/v2/gamification/profiles


