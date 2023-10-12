package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UtilizationtagMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UtilizationtagDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Utilizationtag
type Utilizationtag struct { 
    


    // Name
    Name string `json:"name"`


    

}

// String returns a JSON representation of the model
func (o *Utilizationtag) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Utilizationtag) MarshalJSON() ([]byte, error) {
    type Alias Utilizationtag

    if UtilizationtagMarshalled {
        return []byte("{}"), nil
    }
    UtilizationtagMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

