package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExternalsourceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExternalsourceDud struct { 
    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Externalsource
type Externalsource struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Name - The name of the external source. Must be unique. Max: 200 characters. Leading and trailing whitespace stripped.
    Name string `json:"name"`


    // Active
    Active bool `json:"active"`


    // LinkConfiguration
    LinkConfiguration Linkconfiguration `json:"linkConfiguration"`


    

}

// String returns a JSON representation of the model
func (o *Externalsource) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Externalsource) MarshalJSON() ([]byte, error) {
    type Alias Externalsource

    if ExternalsourceMarshalled {
        return []byte("{}"), nil
    }
    ExternalsourceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Active bool `json:"active"`
        
        LinkConfiguration Linkconfiguration `json:"linkConfiguration"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

