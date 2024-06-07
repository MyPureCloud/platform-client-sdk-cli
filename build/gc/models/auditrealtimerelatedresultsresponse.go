package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuditrealtimerelatedresultsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuditrealtimerelatedresultsresponseDud struct { 
    

}

// Auditrealtimerelatedresultsresponse
type Auditrealtimerelatedresultsresponse struct { 
    // Entities
    Entities []Auditlogmessage `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Auditrealtimerelatedresultsresponse) String() string {
     o.Entities = []Auditlogmessage{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Auditrealtimerelatedresultsresponse) MarshalJSON() ([]byte, error) {
    type Alias Auditrealtimerelatedresultsresponse

    if AuditrealtimerelatedresultsresponseMarshalled {
        return []byte("{}"), nil
    }
    AuditrealtimerelatedresultsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Auditlogmessage `json:"entities"`
        *Alias
    }{

        
        Entities: []Auditlogmessage{{}},
        

        Alias: (*Alias)(u),
    })
}

