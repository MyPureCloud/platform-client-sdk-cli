## gc orgauthorization trustees default create

Create a new organization authorization trust with Customer Care. This is required to grant your regional Customer Care organization access to your organization.

### Synopsis

Create a new organization authorization trust with Customer Care. This is required to grant your regional Customer Care organization access to your organization.

```
gc orgauthorization trustees default create [flags]
```

### Options

```
      --assignDefaultRole string   Assign Admin role to default pairing with Customer Care Valid values: true, false
      --autoExpire string          Automatically expire pairing after 30 days Valid values: true, false
  -h, --help                       help for create
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

* [gc orgauthorization trustees default](gc_orgauthorization_trustees_default.html)	 - /api/v2/orgauthorization/trustees/default


