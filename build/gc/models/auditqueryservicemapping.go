package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuditqueryservicemappingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuditqueryservicemappingDud struct { 
    

}

// Auditqueryservicemapping
type Auditqueryservicemapping struct { 
    // Services - List of Services
    Services []Auditqueryservice `json:"services"`

}

// String returns a JSON representation of the model
func (o *Auditqueryservicemapping) String() string {
     o.Services = []Auditqueryservice{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Auditqueryservicemapping) MarshalJSON() ([]byte, error) {
    type Alias Auditqueryservicemapping

    if AuditqueryservicemappingMarshalled {
        return []byte("{}"), nil
    }
    AuditqueryservicemappingMarshalled = true

    return json.Marshal(&struct {
        
        Services []Auditqueryservice `json:"services"`
        *Alias
    }{

        
        Services: []Auditqueryservice{{}},
        

        Alias: (*Alias)(u),
    })
}

