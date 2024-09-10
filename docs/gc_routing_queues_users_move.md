## gc routing queues users move

DEPRECATED: use POST /routing/queues/{queueId}/members.  Bulk add or delete up to 100 queue members.

### Synopsis

DEPRECATED: use POST /routing/queues/{queueId}/members.  Bulk add or delete up to 100 queue members.

```
gc routing queues users move [queueId] [flags]
```

### Options

```
      --delete string      True to delete queue members Valid values: true, false
  -d, --directory string   Directory path with files containing request bodies
  -f, --file string        File name containing the JSON body
  -h, --help               help for move
  -b, --printrequestbody   Print the request body format of the API.
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

* [gc routing queues users](gc_routing_queues_users.html)	 - /api/v2/routing/queues/{queueId}/users


