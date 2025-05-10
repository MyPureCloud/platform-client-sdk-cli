## gc businessrules decisiontables search list

Search for decision tables.

### Synopsis

Search for decision tables.

```
gc businessrules decisiontables search list [flags]
```

### Options

```
      --after string             The cursor that points to the end of the set of entities that has been returned.
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --before string            The cursor that points to the start of the set of entities that has been returned.
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --name string              Search for decision tables with a name that contains the given search string. Search is case insensitive and will match any table that contains this string in any part of the name. Cannot be combined with schema search. Search results will not be paginated if used.
      --pageSize string          Number of entities to return. Maximum of 200.
      --schemaId string          Search for decision tables that use the schema with this ID. Cannot be combined with name search. Search results will not be paginated if used.
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

* [gc businessrules decisiontables search](gc_businessrules_decisiontables_search.html)	 - /api/v2/businessrules/decisiontables/search


