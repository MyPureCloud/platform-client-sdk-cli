## gc presence definitions list

Get a list of Presence Definitions

### Synopsis

Get a list of Presence Definitions

```
gc presence definitions list [flags]
```

### Options

```
      --deactivated string   Deactivated query can be TRUE or FALSE
      --divisionId strings   One or more division IDs. If nothing is provided, the definitions associated withthe list of divisions that the user has access to will be returned.
  -h, --help                 help for list
      --localeCode string    The locale code to fetch for the presence definition. Use ALL to fetch everything. Valid values: ALL, he, fr, en_US, da, de, it, cs, es, fi, ar, ja, ko, nl, no, pl, pt_BR, pt_PT, ru, sv, th, tr, uk, zh_CN, zh_TW
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

* [gc presence definitions](gc_presence_definitions.html)	 - /api/v2/presence/definitions


