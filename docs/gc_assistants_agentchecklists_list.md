## gc assistants agentchecklists list

Get the list of agent checklists

### Synopsis

Get the list of agent checklists

```
gc assistants agentchecklists list [flags]
```

### Options

```
      --after string             The cursor that points to the end of the set of entities that has been returned.
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --before string            The cursor that points to the start of the set of entities that has been returned.
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --language string          The agent checklist language filter applied to the listing.
      --namePrefix string        The agent checklist name prefix filter applied to the listing.
      --pageSize string          The page size for the listing. The max that will be returned is 100.
      --sortBy string            The field to sort by for the listing. Valid values: dateModified, language, name
      --sortOrder string         The sort order for the listing Valid values: asc, desc
  -s, --stream                   Paginate and stream data as it is being processed leaving page information intact
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

* [gc assistants agentchecklists](gc_assistants_agentchecklists.html)	 - /api/v2/assistants/agentchecklists


