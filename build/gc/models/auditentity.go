package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuditentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuditentityDud struct { 
    


    


    


    

}

// Auditentity
type Auditentity struct { 
    // VarType - The type of the entity the action of this AuditEntity targeted.
    VarType string `json:"type"`


    // Id - The id of the entity the action of this AuditEntity targeted.
    Id string `json:"id"`


    // Name - The name of the entity the action of this AuditEntity targeted.
    Name string `json:"name"`


    // SelfUri - The selfUri for this entity.
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Auditentity) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Auditentity) MarshalJSON() ([]byte, error) {
    type Alias Auditentity

    if AuditentityMarshalled {
        return []byte("{}"), nil
    }
    AuditentityMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        SelfUri string `json:"selfUri"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

