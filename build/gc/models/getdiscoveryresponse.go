package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GetdiscoveryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GetdiscoveryresponseDud struct { 
    


    


    


    

}

// Getdiscoveryresponse
type Getdiscoveryresponse struct { 
    // SnapshotId - The snapshot ID for the discovery job.
    SnapshotId string `json:"snapshotId"`


    // SnapshotStatus - The current status of the discovery snapshot.
    SnapshotStatus string `json:"snapshotStatus"`


    // Items - List of RFC metadata items (only present when status is READY).
    Items []Rfcmetadata `json:"items"`


    // ErrorCode - Error code (only present when status is FAILED).
    ErrorCode string `json:"errorCode"`

}

// String returns a JSON representation of the model
func (o *Getdiscoveryresponse) String() string {
    
    
     o.Items = []Rfcmetadata{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Getdiscoveryresponse) MarshalJSON() ([]byte, error) {
    type Alias Getdiscoveryresponse

    if GetdiscoveryresponseMarshalled {
        return []byte("{}"), nil
    }
    GetdiscoveryresponseMarshalled = true

    return json.Marshal(&struct {
        
        SnapshotId string `json:"snapshotId"`
        
        SnapshotStatus string `json:"snapshotStatus"`
        
        Items []Rfcmetadata `json:"items"`
        
        ErrorCode string `json:"errorCode"`
        *Alias
    }{

        


        


        
        Items: []Rfcmetadata{{}},
        


        

        Alias: (*Alias)(u),
    })
}

