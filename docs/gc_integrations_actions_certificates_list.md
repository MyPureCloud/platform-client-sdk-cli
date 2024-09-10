## gc integrations actions certificates list

Retrieves the available mTLS client certificates in use. This endpoint will return inconsistent results while a certificate rotation is in progress.

### Synopsis

Retrieves the available mTLS client certificates in use. This endpoint will return inconsistent results while a certificate rotation is in progress.

```
gc integrations actions certificates list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --status string            Indicates the validity of the certificate in question. Valid values: Current, Upcoming
  -s, --stream                   Paginate and stream data as it is being processed leaving page information intact
      --varType string           Indicates the type of the certificate. Valid values: Client
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

* [gc integrations actions certificates](gc_integrations_actions_certificates.html)	 - /api/v2/integrations/actions/certificates


