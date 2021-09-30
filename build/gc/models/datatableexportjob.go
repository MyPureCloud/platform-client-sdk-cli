package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DatatableexportjobMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DatatableexportjobDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Datatableexportjob - State information for an export job of rows from a datatable
type Datatableexportjob struct { 
    


    // Name
    Name string `json:"name"`


    // Owner - The PureCloud user who started the export job
    Owner Addressableentityref `json:"owner"`


    // Status - The status of the export job
    Status string `json:"status"`


    // DateCreated - The timestamp of when the export began. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateCompleted - The timestamp of when the export stopped (either successfully or unsuccessfully). Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCompleted time.Time `json:"dateCompleted"`


    // DownloadURI - The URL of the location at which the caller can download the export file, when available
    DownloadURI string `json:"downloadURI"`


    // ErrorInformation - Any error information, or null of the processing is not in an error state
    ErrorInformation Errorbody `json:"errorInformation"`


    // CountRecordsProcessed - The current count of the number of records processed
    CountRecordsProcessed int `json:"countRecordsProcessed"`


    

}

// String returns a JSON representation of the model
func (o *Datatableexportjob) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Datatableexportjob) MarshalJSON() ([]byte, error) {
    type Alias Datatableexportjob

    if DatatableexportjobMarshalled {
        return []byte("{}"), nil
    }
    DatatableexportjobMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Owner Addressableentityref `json:"owner"`
        
        Status string `json:"status"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateCompleted time.Time `json:"dateCompleted"`
        
        DownloadURI string `json:"downloadURI"`
        
        ErrorInformation Errorbody `json:"errorInformation"`
        
        CountRecordsProcessed int `json:"countRecordsProcessed"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

