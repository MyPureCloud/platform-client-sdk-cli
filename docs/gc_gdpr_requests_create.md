## gc gdpr requests create

Submit a new GDPR request

### Synopsis

Submit a new GDPR request

```
gc gdpr requests create [flags]
```

### Options

```
      --deleteConfirmed string   Confirm delete Valid values: true, false
  -d, --directory string         Directory path with files containing request bodies
  -f, --file string              File name containing the JSON body
  -h, --help                     help for create
  -b, --printrequestbody         Print the request body format of the API.
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

* [gc gdpr requests](gc_gdpr_requests.html)	 - /api/v2/gdpr/requests

