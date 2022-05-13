package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuditqueryentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuditqueryentityDud struct { 
    


    

}

// Auditqueryentity
type Auditqueryentity struct { 
    // Name - Name of the Entity
    Name string `json:"name"`


    // Actions - List of Actions
    Actions []string `json:"actions"`

}

// String returns a JSON representation of the model
func (o *Auditqueryentity) String() string {
    
     o.Actions = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Auditqueryentity) MarshalJSON() ([]byte, error) {
    type Alias Auditqueryentity

    if AuditqueryentityMarshalled {
        return []byte("{}"), nil
    }
    AuditqueryentityMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Actions []string `json:"actions"`
        *Alias
    }{

        


        
        Actions: []string{""},
        

        Alias: (*Alias)(u),
    })
}

