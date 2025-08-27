## gc assistants queues

/api/v2/assistants/{assistantId}/queues /api/v2/assistants/queues

### Synopsis

/api/v2/assistants/{assistantId}/queues /api/v2/assistants/queues

### Options

```
  -h, --help   help for queues
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

* [gc assistants](gc_assistants.html)	 - /api/v2/assistants
* [gc assistants queues createAssistantQueue](gc_assistants_queues_createAssistantQueue.html)	 - Create a queue assistant association.
* [gc assistants queues delete](gc_assistants_queues_delete.html)	 - Disassociate the queues from an assistant for the given assistant ID and queue IDs.
* [gc assistants queues get](gc_assistants_queues_get.html)	 - Get queue Information for an assistant.
* [gc assistants queues getAssistantQueues](gc_assistants_queues_getAssistantQueues.html)	 - Get all the queues associated with an assistant.
* [gc assistants queues list](gc_assistants_queues_list.html)	 - Get all queues assigned to any assistant.
* [gc assistants queues removeQueueFromAssistant](gc_assistants_queues_removeQueueFromAssistant.html)	 - Disassociate a queue from an assistant.
* [gc assistants queues update](gc_assistants_queues_update.html)	 - Update Queues for an Assistant.
* [gc assistants queues users](gc_assistants_queues_users.html)	 - /api/v2/assistants/{assistantId}/queues/{queueId}/users


