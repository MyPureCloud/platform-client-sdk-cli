package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuditentityreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuditentityreferenceDud struct { 
    


    


    


    


    

}

// Auditentityreference
type Auditentityreference struct { 
    // Id
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // VarType
    VarType string `json:"type"`


    // Action
    Action string `json:"action"`

}

// String returns a JSON representation of the model
func (o *Auditentityreference) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Auditentityreference) MarshalJSON() ([]byte, error) {
    type Alias Auditentityreference

    if AuditentityreferenceMarshalled {
        return []byte("{}"), nil
    }
    AuditentityreferenceMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        SelfUri string `json:"selfUri"`
        
        VarType string `json:"type"`
        
        Action string `json:"action"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

