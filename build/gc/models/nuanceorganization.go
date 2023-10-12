package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NuanceorganizationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NuanceorganizationDud struct { 
    


    

}

// Nuanceorganization - Model for a Nuance bot organization
type Nuanceorganization struct { 
    // Id - The organization ID
    Id string `json:"id"`


    // Name - The organization name
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Nuanceorganization) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nuanceorganization) MarshalJSON() ([]byte, error) {
    type Alias Nuanceorganization

    if NuanceorganizationMarshalled {
        return []byte("{}"), nil
    }
    NuanceorganizationMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

