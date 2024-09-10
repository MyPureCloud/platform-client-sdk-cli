## gc messaging settings default

/api/v2/messaging/settings/default

### Synopsis

/api/v2/messaging/settings/default

### Options

```
  -h, --help   help for default
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

* [gc messaging settings](gc_messaging_settings.html)	 - /api/v2/messaging/settings
* [gc messaging settings default delete](gc_messaging_settings_default_delete.html)	 - Delete the organization`s default setting, a global default will be applied to integrations without settings
* [gc messaging settings default get](gc_messaging_settings_default_get.html)	 - Get the organization`s default settings that will be used as the default when creating an integration.
* [gc messaging settings default update](gc_messaging_settings_default_update.html)	 - Set the organization`s default settings that may be applied to an integration when it is created.


