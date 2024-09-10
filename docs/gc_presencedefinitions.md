## gc presencedefinitions

/api/v2/presencedefinitions

### Synopsis

/api/v2/presencedefinitions

### Options

```
  -h, --help   help for presencedefinitions
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

* [gc](gc.html)	 - gc is a CLI for interacting with Genesys Cloud
* [gc presencedefinitions create](gc_presencedefinitions_create.html)	 - Create a Presence Definition. Apps should migrate to use POST /api/v2/presence/definitions instead
* [gc presencedefinitions delete](gc_presencedefinitions_delete.html)	 - Delete a Presence Definition. Apps should migrate to use DELETE /api/v2/presence/definitions/{definitionId} instead
* [gc presencedefinitions get](gc_presencedefinitions_get.html)	 - Get a Presence Definition. Apps should migrate to use GET /api/v2/presence/definitions/{definitionId} instead
* [gc presencedefinitions list](gc_presencedefinitions_list.html)	 - Get an Organization`s list of Presence Definitions. Apps should migrate to use GET /api/v2/presence/definitions instead
* [gc presencedefinitions update](gc_presencedefinitions_update.html)	 - Update a Presence Definition. Apps should migrate to use PUT /api/v2/presence/definitions/{definitionId} instead)


