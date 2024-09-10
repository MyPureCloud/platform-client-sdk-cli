## gc integrations actions execute

/api/v2/integrations/actions/{actionId}/execute

### Synopsis

/api/v2/integrations/actions/{actionId}/execute

### Options

```
  -h, --help   help for execute
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
* [gc integrations actions execute create](gc_integrations_actions_execute_create.html)	 - Execute Action and return response from 3rd party.  Responses will follow the schemas defined on the Action for success and error.


