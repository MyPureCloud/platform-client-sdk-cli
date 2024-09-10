## gc externalcontacts organizations get

Fetch an external organization

### Synopsis

Fetch an external organization

```
gc externalcontacts organizations get [externalOrganizationId] [flags]
```

### Options

```
      --expand strings           which fields, if any, to expand (externalDataSources) Valid values: externalDataSources
  -h, --help                     help for get
      --includeTrustors string   (true or false) whether or not to include trustor information embedded in the externalOrganization Valid values: true, false
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

* [gc externalcontacts organizations](gc_externalcontacts_organizations.html)	 - /api/v2/externalcontacts/organizations


