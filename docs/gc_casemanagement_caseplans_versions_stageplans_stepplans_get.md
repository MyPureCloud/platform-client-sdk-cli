## gc casemanagement caseplans versions stageplans stepplans get

Get a Stepplan.

### Synopsis

Get a Stepplan.

```
gc casemanagement caseplans versions stageplans stepplans get [caseplanId] [versionId] [stageplanId] [stepplanId] [flags]
```

### Options

```
      --expands strings   Which fields to expand. Valid values: stageplan, caseplan, worktype
  -h, --help              help for get
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

* [gc casemanagement caseplans versions stageplans stepplans](gc_casemanagement_caseplans_versions_stageplans_stepplans.html)	 - /api/v2/casemanagement/caseplans/{caseplanId}/versions/{versionId}/stageplans/{stageplanId}/stepplans


