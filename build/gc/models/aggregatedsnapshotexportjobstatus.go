package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AggregatedsnapshotexportjobstatusMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AggregatedsnapshotexportjobstatusDud struct { 
    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Aggregatedsnapshotexportjobstatus
type Aggregatedsnapshotexportjobstatus struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Status - The status of the export job
    Status string `json:"status"`


    // DownloadUrl - The download URL for the completed export. Populated when status is Complete
    DownloadUrl string `json:"downloadUrl"`


    // VarError - Error details if the export failed. Populated when status is Error
    VarError Csvexporterrordetails `json:"error"`


    

}

// String returns a JSON representation of the model
func (o *Aggregatedsnapshotexportjobstatus) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Aggregatedsnapshotexportjobstatus) MarshalJSON() ([]byte, error) {
    type Alias Aggregatedsnapshotexportjobstatus

    if AggregatedsnapshotexportjobstatusMarshalled {
        return []byte("{}"), nil
    }
    AggregatedsnapshotexportjobstatusMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Status string `json:"status"`
        
        DownloadUrl string `json:"downloadUrl"`
        
        VarError Csvexporterrordetails `json:"error"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

