package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SigneddataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SigneddataDud struct { 
    

}

// Signeddata
type Signeddata struct { 
    // Jwt
    Jwt string `json:"jwt"`

}

// String returns a JSON representation of the model
func (o *Signeddata) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Signeddata) MarshalJSON() ([]byte, error) {
    type Alias Signeddata

    if SigneddataMarshalled {
        return []byte("{}"), nil
    }
    SigneddataMarshalled = true

    return json.Marshal(&struct {
        
        Jwt string `json:"jwt"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

