package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserreferenceDud struct { 
    


    SelfUri string `json:"selfUri"`

}

// Userreference
type Userreference struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *Userreference) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userreference) MarshalJSON() ([]byte, error) {
    type Alias Userreference

    if UserreferenceMarshalled {
        return []byte("{}"), nil
    }
    UserreferenceMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

