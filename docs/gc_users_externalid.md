## gc users externalid

/api/v2/users/{userId}/externalid /api/v2/users/{userId}/externalid/{authorityName} /api/v2/users/externalid/{authorityName}

### Synopsis

/api/v2/users/{userId}/externalid /api/v2/users/{userId}/externalid/{authorityName} /api/v2/users/externalid/{authorityName}

### Options

```
  -h, --help   help for externalid
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

* [gc users](gc_users.html)	 - /api/v2/users
* [gc users externalid create](gc_users_externalid_create.html)	 - Create mapping between external identifier and user. Limit 100 per entity.
* [gc users externalid delete](gc_users_externalid_delete.html)	 - Delete the external identifier for user.
* [gc users externalid getexternalidforuserauthority](gc_users_externalid_getexternalidforuserauthority.html)	 - Get the external identifier of user for an authority.
* [gc users externalid getexternalidsforuser](gc_users_externalid_getexternalidsforuser.html)	 - Get the external identifiers for a user.
* [gc users externalid getuserbyexternalid](gc_users_externalid_getuserbyexternalid.html)	 - Get the user associated with external identifier.


