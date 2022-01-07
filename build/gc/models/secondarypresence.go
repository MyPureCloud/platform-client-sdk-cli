package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SecondarypresenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SecondarypresenceDud struct { 
    


    SelfUri string `json:"selfUri"`

}

// Secondarypresence
type Secondarypresence struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *Secondarypresence) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Secondarypresence) MarshalJSON() ([]byte, error) {
    type Alias Secondarypresence

    if SecondarypresenceMarshalled {
        return []byte("{}"), nil
    }
    SecondarypresenceMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

