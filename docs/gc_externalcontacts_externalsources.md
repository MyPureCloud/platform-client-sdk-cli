## gc externalcontacts externalsources

/api/v2/externalcontacts/externalsources

### Synopsis

/api/v2/externalcontacts/externalsources

### Options

```
  -h, --help   help for externalsources
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

* [gc externalcontacts](gc_externalcontacts.html)	 - /api/v2/externalcontacts
* [gc externalcontacts externalsources create](gc_externalcontacts_externalsources_create.html)	 - Create an External Source
* [gc externalcontacts externalsources delete](gc_externalcontacts_externalsources_delete.html)	 - Delete an External Source. WARNING: Any records that reference this External Source will not be automatically cleaned up. Those records will still be editable, but their External IDs may not be fully viewable.
* [gc externalcontacts externalsources get](gc_externalcontacts_externalsources_get.html)	 - Fetch an External Source
* [gc externalcontacts externalsources list](gc_externalcontacts_externalsources_list.html)	 - Fetch a list of External Sources
* [gc externalcontacts externalsources update](gc_externalcontacts_externalsources_update.html)	 - Update an External Source


