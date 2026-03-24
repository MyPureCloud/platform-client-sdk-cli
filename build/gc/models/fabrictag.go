package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FabrictagMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FabrictagDud struct { 
    

}

// Fabrictag
type Fabrictag struct { 
    // Name - The name of the tag
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Fabrictag) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Fabrictag) MarshalJSON() ([]byte, error) {
    type Alias Fabrictag

    if FabrictagMarshalled {
        return []byte("{}"), nil
    }
    FabrictagMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

