package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GenerateentryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GenerateentryresponseDud struct { 
    


    

}

// Generateentryresponse
type Generateentryresponse struct { 
    // SnapshotId - The snapshot ID for the generation job.
    SnapshotId string `json:"snapshotId"`


    // RfcName - The name of the RFC (Reason For Contact).
    RfcName string `json:"rfcName"`

}

// String returns a JSON representation of the model
func (o *Generateentryresponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Generateentryresponse) MarshalJSON() ([]byte, error) {
    type Alias Generateentryresponse

    if GenerateentryresponseMarshalled {
        return []byte("{}"), nil
    }
    GenerateentryresponseMarshalled = true

    return json.Marshal(&struct {
        
        SnapshotId string `json:"snapshotId"`
        
        RfcName string `json:"rfcName"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

