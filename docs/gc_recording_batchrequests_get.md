## gc recording batchrequests get

Get the status and results for a batch request job, only the user that submitted the job may retrieve results. Each result may contain either a URL to a recording or an error; additionally, a recording could be associated with multiple results.

### Synopsis

Get the status and results for a batch request job, only the user that submitted the job may retrieve results. Each result may contain either a URL to a recording or an error; additionally, a recording could be associated with multiple results.

```
gc recording batchrequests get [jobId] [flags]
```

### Options

```
  -h, --help   help for get
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

* [gc recording batchrequests](gc_recording_batchrequests.html)	 - /api/v2/recording/batchrequests


