## gc notifications channels subscriptions

/api/v2/notifications/channels/{channelId}/subscriptions

### Synopsis

/api/v2/notifications/channels/{channelId}/subscriptions

### Options

```
  -h, --help   help for subscriptions
```

### Options inherited from parent commands

```
      --accesstoken string    accessToken override
      --clientid string       clientId override
      --clientsecret string   clientSecret override
      --environment string    environment override. E.g. mypurecloud.com.au or ap-southeast-2
  -i, --indicateprogress      Trace progress indicators to stderr
      --inputformat string    Data input format. Supported formats: YAML, JSON
      --noheartbeat           Filters out the heartbeat from the event stream when using the notification service
      --outputformat string   Data output format. Supported formats: YAML, JSON
  -p, --profile string        Name of the profile to use for configuring the cli (default "DEFAULT")
      --transform string      Provide a Go template file for transforming output data
      --transformstr string   Provide a Go template string for transforming output data
```

### SEE ALSO

* [gc notifications channels](gc_notifications_channels.html)	 - /api/v2/notifications/channels
* [gc notifications channels subscriptions delete](gc_notifications_channels_subscriptions_delete.html)	 - Remove all subscriptions
* [gc notifications channels subscriptions list](gc_notifications_channels_subscriptions_list.html)	 - The list of all subscriptions for this channel
* [gc notifications channels subscriptions subscribe](gc_notifications_channels_subscriptions_subscribe.html)	 - Add a list of subscriptions to the existing list of subscriptions
* [gc notifications channels subscriptions update](gc_notifications_channels_subscriptions_update.html)	 - Replace the current list of subscriptions with a new list.


