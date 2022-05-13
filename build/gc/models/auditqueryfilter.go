package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuditqueryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuditqueryfilterDud struct { 
    


    

}

// Auditqueryfilter
type Auditqueryfilter struct { 
    // Property - Name of the property to filter.
    Property string `json:"property"`


    // Value - Value of the property to filter.
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Auditqueryfilter) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Auditqueryfilter) MarshalJSON() ([]byte, error) {
    type Alias Auditqueryfilter

    if AuditqueryfilterMarshalled {
        return []byte("{}"), nil
    }
    AuditqueryfilterMarshalled = true

    return json.Marshal(&struct {
        
        Property string `json:"property"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

