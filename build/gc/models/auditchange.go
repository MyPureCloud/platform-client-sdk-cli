package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuditchangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuditchangeDud struct { 
    


    


    


    

}

// Auditchange
type Auditchange struct { 
    // Property
    Property string `json:"property"`


    // Entity
    Entity Auditentityreference `json:"entity"`


    // OldValues
    OldValues []string `json:"oldValues"`


    // NewValues
    NewValues []string `json:"newValues"`

}

// String returns a JSON representation of the model
func (o *Auditchange) String() string {
    
    
    
    
    
    
    
    
    
    
     o.OldValues = []string{""} 
    
    
    
     o.NewValues = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Auditchange) MarshalJSON() ([]byte, error) {
    type Alias Auditchange

    if AuditchangeMarshalled {
        return []byte("{}"), nil
    }
    AuditchangeMarshalled = true

    return json.Marshal(&struct { 
        Property string `json:"property"`
        
        Entity Auditentityreference `json:"entity"`
        
        OldValues []string `json:"oldValues"`
        
        NewValues []string `json:"newValues"`
        
        *Alias
    }{
        

        

        

        

        

        
        OldValues: []string{""},
        

        

        
        NewValues: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}

