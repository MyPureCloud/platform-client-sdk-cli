package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IntentreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IntentreferenceDud struct { 
    Id string `json:"id"`


    

}

// Intentreference
type Intentreference struct { 
    


    // Name - The name of the intent.
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Intentreference) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Intentreference) MarshalJSON() ([]byte, error) {
    type Alias Intentreference

    if IntentreferenceMarshalled {
        return []byte("{}"), nil
    }
    IntentreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

