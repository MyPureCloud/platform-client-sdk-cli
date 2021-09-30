package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Scimv2memberreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Scimv2memberreferenceDud struct { 
    VarType string `json:"type"`


    


    Ref string `json:"$ref"`

}

// Scimv2memberreference - Defines a reference to SCIM group members.
type Scimv2memberreference struct { 
    


    // Value - The ID of the group member. Can be \"userId\" or \"groupId\".
    Value string `json:"value"`


    

}

// String returns a JSON representation of the model
func (o *Scimv2memberreference) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scimv2memberreference) MarshalJSON() ([]byte, error) {
    type Alias Scimv2memberreference

    if Scimv2memberreferenceMarshalled {
        return []byte("{}"), nil
    }
    Scimv2memberreferenceMarshalled = true

    return json.Marshal(&struct { 
        
        
        Value string `json:"value"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

