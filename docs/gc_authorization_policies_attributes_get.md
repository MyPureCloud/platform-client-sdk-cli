## gc authorization policies attributes get

Get the list of attributes used to evaluate an access control policy with the specified policy ID

### Synopsis

Get the list of attributes used to evaluate an access control policy with the specified policy ID

```
gc authorization policies attributes get [policyId] [flags]
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

* [gc authorization policies attributes](gc_authorization_policies_attributes.html)	 - /api/v2/authorization/policies/{policyId}/attributes


