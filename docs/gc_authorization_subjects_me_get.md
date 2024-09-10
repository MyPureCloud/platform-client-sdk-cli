## gc authorization subjects me get

Returns a listing of roles and permissions for the currently authenticated user.

### Synopsis

Returns a listing of roles and permissions for the currently authenticated user.

```
gc authorization subjects me get [flags]
```

### Options

```
  -h, --help                       help for get
      --includeDuplicates string   Include multiple entries with the same role and division but different subjects Valid values: true, false Valid values: true, false
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

* [gc authorization subjects me](gc_authorization_subjects_me.html)	 - /api/v2/authorization/subjects/me


