package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactimportfieldMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactimportfieldDud struct { 
    


    

}

// Contactimportfield
type Contactimportfield struct { 
    // Name - Field name
    Name string `json:"name"`


    // Included - Should we import this field
    Included bool `json:"included"`

}

// String returns a JSON representation of the model
func (o *Contactimportfield) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactimportfield) MarshalJSON() ([]byte, error) {
    type Alias Contactimportfield

    if ContactimportfieldMarshalled {
        return []byte("{}"), nil
    }
    ContactimportfieldMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Included bool `json:"included"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

