## gc flows datatables import csv jobs

/api/v2/flows/datatables/{datatableId}/import/csv/jobs

### Synopsis

/api/v2/flows/datatables/{datatableId}/import/csv/jobs

### Options

```
  -h, --help   help for jobs
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

* [gc flows datatables import csv](gc_flows_datatables_import_csv.html)	 - /api/v2/flows/datatables/{datatableId}/import/csv
* [gc flows datatables import csv jobs create](gc_flows_datatables_import_csv_jobs_create.html)	 - Begin an import process for importing rows from a CSV file into a datatable. CSV file is uploaded by performing a PUT request against the URL in the returned `uploadURI` field. Headers for the PUT request must contain all headers contained in the returned `uploadHeaders` field.


