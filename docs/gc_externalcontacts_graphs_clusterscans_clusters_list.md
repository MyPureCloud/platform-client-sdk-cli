## gc externalcontacts graphs clusterscans clusters list

Returns a list of clusters found by a scan

### Synopsis

Returns a list of clusters found by a scan

```
gc externalcontacts graphs clusterscans clusters list [scanId] [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --cursor string            Cursor to continue scanning
      --divisionIds strings      which divisions to filter results to, up to 50 (defaults to all divisions use has access to)
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --limit string             Max number of records to return (must be between 1 and 100) (default "20")
      --mergeInfoStatus string   which merge statuses to filter results to Valid values: AutoQueued, AutoSucceeded, AutoFailed, ManualQueued, ManualSucceeded, ManualFailed, NotMerged
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

* [gc externalcontacts graphs clusterscans clusters](gc_externalcontacts_graphs_clusterscans_clusters.html)	 - /api/v2/externalcontacts/graphs/clusterscans/{scanId}/clusters


