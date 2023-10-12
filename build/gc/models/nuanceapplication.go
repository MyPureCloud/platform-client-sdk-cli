package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NuanceapplicationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NuanceapplicationDud struct { 
    


    


    

}

// Nuanceapplication - Model for a Nuance bot application
type Nuanceapplication struct { 
    // Id - The application ID
    Id string `json:"id"`


    // Tag - The application Tag
    Tag string `json:"tag"`


    // Name - The application name
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Nuanceapplication) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nuanceapplication) MarshalJSON() ([]byte, error) {
    type Alias Nuanceapplication

    if NuanceapplicationMarshalled {
        return []byte("{}"), nil
    }
    NuanceapplicationMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Tag string `json:"tag"`
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

