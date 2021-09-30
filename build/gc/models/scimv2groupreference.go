package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Scimv2groupreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Scimv2groupreferenceDud struct { 
    VarType string `json:"type"`


    


    Ref string `json:"$ref"`

}

// Scimv2groupreference - Defines a reference to SCIM groups.
type Scimv2groupreference struct { 
    


    // Value - The ID of the group member. Can be \"userId\" or \"groupId\".
    Value string `json:"value"`


    

}

// String returns a JSON representation of the model
func (o *Scimv2groupreference) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scimv2groupreference) MarshalJSON() ([]byte, error) {
    type Alias Scimv2groupreference

    if Scimv2groupreferenceMarshalled {
        return []byte("{}"), nil
    }
    Scimv2groupreferenceMarshalled = true

    return json.Marshal(&struct { 
        
        
        Value string `json:"value"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

