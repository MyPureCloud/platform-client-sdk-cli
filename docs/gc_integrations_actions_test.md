## gc integrations actions test

/api/v2/integrations/actions/{actionId}/test

### Synopsis

/api/v2/integrations/actions/{actionId}/test

### Options

```
  -h, --help   help for test
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

* [gc integrations actions](gc_integrations_actions.html)	 - /api/v2/integrations/actions
* [gc integrations actions test create](gc_integrations_actions_test_create.html)	 - Test the execution of an action. Responses will show execution steps broken out with intermediate results to help in debugging.

