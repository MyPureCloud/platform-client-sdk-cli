package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuditqueryserviceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuditqueryserviceDud struct { 
    


    

}

// Auditqueryservice
type Auditqueryservice struct { 
    // Name - Name of the Service
    Name string `json:"name"`


    // Entities - List of Entities
    Entities []Auditqueryentity `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Auditqueryservice) String() string {
    
     o.Entities = []Auditqueryentity{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Auditqueryservice) MarshalJSON() ([]byte, error) {
    type Alias Auditqueryservice

    if AuditqueryserviceMarshalled {
        return []byte("{}"), nil
    }
    AuditqueryserviceMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Entities []Auditqueryentity `json:"entities"`
        *Alias
    }{

        


        
        Entities: []Auditqueryentity{{}},
        

        Alias: (*Alias)(u),
    })
}

