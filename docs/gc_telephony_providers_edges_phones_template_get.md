## gc telephony providers edges phones template get

Get a Phone instance template based on a Phone Base Settings object. This object can then be modified and saved as a new Phone instance

### Synopsis

Get a Phone instance template based on a Phone Base Settings object. This object can then be modified and saved as a new Phone instance

```
gc telephony providers edges phones template get [flags]
```

### Options

```
  -h, --help                         help for get
      --phoneBaseSettingsId string   The id of a Phone Base Settings object upon which to base this Phone - REQUIRED
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

* [gc telephony providers edges phones template](gc_telephony_providers_edges_phones_template.html)	 - /api/v2/telephony/providers/edges/phones/template


