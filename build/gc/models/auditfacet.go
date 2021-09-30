package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuditfacetMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuditfacetDud struct { 
    


    

}

// Auditfacet
type Auditfacet struct { 
    // Name - The name of the field on which to facet.
    Name string `json:"name"`


    // VarType - The type of the facet, DATE or STRING.
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Auditfacet) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Auditfacet) MarshalJSON() ([]byte, error) {
    type Alias Auditfacet

    if AuditfacetMarshalled {
        return []byte("{}"), nil
    }
    AuditfacetMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

