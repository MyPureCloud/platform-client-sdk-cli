package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NamedentitydefinitionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NamedentitydefinitionDud struct { 
    


    

}

// Namedentitydefinition
type Namedentitydefinition struct { 
    // Name - The name of the entity.
    Name string `json:"name"`


    // VarType - The name of the entity type.
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Namedentitydefinition) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Namedentitydefinition) MarshalJSON() ([]byte, error) {
    type Alias Namedentitydefinition

    if NamedentitydefinitionMarshalled {
        return []byte("{}"), nil
    }
    NamedentitydefinitionMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

