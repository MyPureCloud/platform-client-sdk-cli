## gc routing queues members

/api/v2/routing/queues/{queueId}/members

### Synopsis

/api/v2/routing/queues/{queueId}/members

### Options

```
  -h, --help   help for members
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

* [gc routing queues](gc_routing_queues.html)	 - /api/v2/routing/queues
* [gc routing queues members activate](gc_routing_queues_members_activate.html)	 - Join or unjoin a set of users for a queue
* [gc routing queues members delete](gc_routing_queues_members_delete.html)	 - Delete a queue member.
* [gc routing queues members list](gc_routing_queues_members_list.html)	 - Get the members of this queue.
* [gc routing queues members move](gc_routing_queues_members_move.html)	 - Bulk add or delete up to 100 queue members
* [gc routing queues members update](gc_routing_queues_members_update.html)	 - Update the ring number OR joined status for a queue member.


