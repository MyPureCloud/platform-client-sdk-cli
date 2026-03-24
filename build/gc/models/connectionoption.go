package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConnectionoptionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConnectionoptionDud struct { 
    


    

}

// Connectionoption
type Connectionoption struct { 
    // Id - The id of the connection option.
    Id string `json:"id"`


    // Name - The name of the connection option.
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Connectionoption) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Connectionoption) MarshalJSON() ([]byte, error) {
    type Alias Connectionoption

    if ConnectionoptionMarshalled {
        return []byte("{}"), nil
    }
    ConnectionoptionMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

