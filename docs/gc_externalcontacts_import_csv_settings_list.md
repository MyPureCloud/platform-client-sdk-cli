## gc externalcontacts import csv settings list

Retrieve all settings for organization filtered by externalSettingsId if provided

### Synopsis

Retrieve all settings for organization filtered by externalSettingsId if provided

```
gc externalcontacts import csv settings list [flags]
```

### Options

```
      --after string                The cursor that points to the end of the set of entities that has been returned.
  -a, --autopaginate                Automatically paginate through the results stripping page information
      --externalSettingsId string   External Settings Id to filter the list.
      --filtercondition string      Filter list command output based on a given condition or regular expression
  -h, --help                        help for list
      --pageSize string             Number of entities to return. Maximum of 200.
  -s, --stream                      Paginate and stream data as it is being processed leaving page information intact
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

* [gc externalcontacts import csv settings](gc_externalcontacts_import_csv_settings.html)	 - /api/v2/externalcontacts/import/csv/settings

