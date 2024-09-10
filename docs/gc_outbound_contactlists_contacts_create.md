## gc outbound contactlists contacts create

Add contacts to a contact list.

### Synopsis

Add contacts to a contact list.

```
gc outbound contactlists contacts create [contactListId] [flags]
```

### Options

```
      --clearSystemData string   Clear system data. True means the system columns (attempts, callable status, etc) stored on the contact will be cleared if the contact already exists; false means they won`t. Valid values: true, false
  -d, --directory string         Directory path with files containing request bodies
      --doNotQueue priority      Do not queue. True means that updated contacts will not have their positions in the queue altered, so contacts that have already been dialed will not be redialed. For new contacts, this parameter has no effect; False means that updated contacts will be re-queued, according to the priority parameter. Valid values: true, false
  -f, --file string              File name containing the JSON body
  -h, --help                     help for create
  -b, --printrequestbody         Print the request body format of the API.
      --priority string          Contact priority. True means the contact(s) will be dialed next; false means the contact will go to the end of the contact queue. Valid values: true, false
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

* [gc outbound contactlists contacts](gc_outbound_contactlists_contacts.html)	 - /api/v2/outbound/contactlists/{contactListId}/contacts


