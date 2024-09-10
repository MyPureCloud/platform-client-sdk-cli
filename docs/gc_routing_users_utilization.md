## gc routing users utilization

/api/v2/routing/users/{userId}/utilization

### Synopsis

/api/v2/routing/users/{userId}/utilization

### Options

```
  -h, --help   help for utilization
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

* [gc routing users](gc_routing_users.html)	 - /api/v2/routing/users
* [gc routing users utilization delete](gc_routing_users_utilization_delete.html)	 - Delete the user`s max utilization settings and revert to the organization-wide default.
* [gc routing users utilization get](gc_routing_users_utilization_get.html)	 - Get the user`s max utilization settings.  If not configured, the organization-wide default is returned.
* [gc routing users utilization update](gc_routing_users_utilization_update.html)	 - Update the user`s max utilization settings.  Include only those media types requiring custom configuration.


