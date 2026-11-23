package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GenerateentryrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GenerateentryrequestDud struct { 
    

}

// Generateentryrequest
type Generateentryrequest struct { 
    // RfcName - The name of the RFC (Reason For Contact).
    RfcName string `json:"rfcName"`

}

// String returns a JSON representation of the model
func (o *Generateentryrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Generateentryrequest) MarshalJSON() ([]byte, error) {
    type Alias Generateentryrequest

    if GenerateentryrequestMarshalled {
        return []byte("{}"), nil
    }
    GenerateentryrequestMarshalled = true

    return json.Marshal(&struct {
        
        RfcName string `json:"rfcName"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

