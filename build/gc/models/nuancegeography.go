package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NuancegeographyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NuancegeographyDud struct { 
    


    

}

// Nuancegeography - Model for a Nuance bot geography
type Nuancegeography struct { 
    // Id - The geography ID
    Id string `json:"id"`


    // Name - The geography name
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Nuancegeography) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nuancegeography) MarshalJSON() ([]byte, error) {
    type Alias Nuancegeography

    if NuancegeographyMarshalled {
        return []byte("{}"), nil
    }
    NuancegeographyMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

