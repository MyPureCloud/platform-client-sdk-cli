package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GetgenerationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GetgenerationresponseDud struct { 
    


    


    


    


    

}

// Getgenerationresponse
type Getgenerationresponse struct { 
    // SnapshotId - The snapshot ID for the generation job.
    SnapshotId string `json:"snapshotId"`


    // RfcName - The name of the RFC (Reason For Contact).
    RfcName string `json:"rfcName"`


    // GenerationStatus - The current status of the generation job.
    GenerationStatus string `json:"generationStatus"`


    // GuideDraft - The generated guide draft (only present when status is READY).
    GuideDraft string `json:"guideDraft"`


    // ErrorCode - Error code (only present when status is FAILED).
    ErrorCode string `json:"errorCode"`

}

// String returns a JSON representation of the model
func (o *Getgenerationresponse) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Getgenerationresponse) MarshalJSON() ([]byte, error) {
    type Alias Getgenerationresponse

    if GetgenerationresponseMarshalled {
        return []byte("{}"), nil
    }
    GetgenerationresponseMarshalled = true

    return json.Marshal(&struct {
        
        SnapshotId string `json:"snapshotId"`
        
        RfcName string `json:"rfcName"`
        
        GenerationStatus string `json:"generationStatus"`
        
        GuideDraft string `json:"guideDraft"`
        
        ErrorCode string `json:"errorCode"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

