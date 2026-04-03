## gc workforcemanagement shifttrading trades match jobs create

Matches a shift trade. This route can only be called by the receiving user

### Synopsis

Matches a shift trade. This route can only be called by the receiving user

```
gc workforcemanagement shifttrading trades match jobs create [tradeId] [flags]
```

### Options

```
  -d, --directory string    Directory path with files containing request bodies
  -f, --file string         File name containing the JSON body
      --forceAsync string   Force the result of this operation to be sent asynchronously via notification. For testing/app development purposes Valid values: true, false
  -h, --help                help for create
  -b, --printrequestbody    Print the request body format of the API.
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

* [gc workforcemanagement shifttrading trades match jobs](gc_workforcemanagement_shifttrading_trades_match_jobs.html)	 - /api/v2/workforcemanagement/shifttrading/trades/{tradeId}/match/jobs


