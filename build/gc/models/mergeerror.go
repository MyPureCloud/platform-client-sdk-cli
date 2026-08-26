package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MergeerrorMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MergeerrorDud struct { 
    Message string `json:"message"`


    Code string `json:"code"`

}

// Mergeerror
type Mergeerror struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Mergeerror) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mergeerror) MarshalJSON() ([]byte, error) {
    type Alias Mergeerror

    if MergeerrorMarshalled {
        return []byte("{}"), nil
    }
    MergeerrorMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

