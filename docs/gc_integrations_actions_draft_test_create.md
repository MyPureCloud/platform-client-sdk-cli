## gc integrations actions draft test create

Test the execution of a draft. Responses will show execution steps broken out with intermediate results to help in debugging.

### Synopsis

Test the execution of a draft. Responses will show execution steps broken out with intermediate results to help in debugging.

```
gc integrations actions draft test create [actionId] [flags]
```

### Options

```
  -d, --directory string   Directory path with files containing request bodies
  -f, --file string        File name containing the JSON body
      --flatten string     Indicates the response should be reformatted, based on Architect`s flattening format. Valid values: true, false
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

* [gc integrations actions draft test](gc_integrations_actions_draft_test.html)	 - /api/v2/integrations/actions/{actionId}/draft/test


