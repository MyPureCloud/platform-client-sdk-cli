## gc routing queues users

/api/v2/routing/queues/{queueId}/users

### Synopsis

/api/v2/routing/queues/{queueId}/users

### Options

```
  -h, --help   help for users
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
* [gc routing queues users activate](gc_routing_queues_users_activate.html)	 - DEPRECATED: use PATCH /routing/queues/{queueId}/members.  Join or unjoin a set of users for a queue.
* [gc routing queues users delete](gc_routing_queues_users_delete.html)	 - DEPRECATED: use DELETE /routing/queues/{queueId}/members/{memberId}.  Delete queue member.
* [gc routing queues users list](gc_routing_queues_users_list.html)	 - DEPRECATED: use GET /routing/queues/{queueId}/members.  Get the members of this queue.
* [gc routing queues users move](gc_routing_queues_users_move.html)	 - DEPRECATED: use POST /routing/queues/{queueId}/members.  Bulk add or delete up to 100 queue members.
* [gc routing queues users update](gc_routing_queues_users_update.html)	 - DEPRECATED: use PATCH /routing/queues/{queueId}/members/{memberId}.  Update the ring number OR joined status for a User in a Queue.


