package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchuserMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchuserDud struct { 
    


    


    

}

// Patchuser
type Patchuser struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // PreferredName - Preferred full name of agent
    PreferredName string `json:"preferredName"`


    // AcdAutoAnswer - The value that denotes if acdAutoAnswer is set on the user
    AcdAutoAnswer bool `json:"acdAutoAnswer"`

}

// String returns a JSON representation of the model
func (o *Patchuser) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchuser) MarshalJSON() ([]byte, error) {
    type Alias Patchuser

    if PatchuserMarshalled {
        return []byte("{}"), nil
    }
    PatchuserMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        PreferredName string `json:"preferredName"`
        
        AcdAutoAnswer bool `json:"acdAutoAnswer"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

