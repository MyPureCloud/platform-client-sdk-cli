## gc authorization subjects divisions roles create

Make a grant of a role in a division

### Synopsis

Make a grant of a role in a division

```
gc authorization subjects divisions roles create [subjectId] [divisionId] [roleId] [flags]
```

### Options

```
  -h, --help                 help for create
      --subjectType string   what the type of the subject is: PC_GROUP, PC_USER or PC_OAUTH_CLIENT (note: for cross-org authorization, please use the Organization Authorization endpoints)
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

* [gc authorization subjects divisions roles](gc_authorization_subjects_divisions_roles.html)	 - /api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles


