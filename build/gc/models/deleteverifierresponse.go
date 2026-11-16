package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DeleteverifierresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DeleteverifierresponseDud struct { 
    

}

// Deleteverifierresponse
type Deleteverifierresponse struct { 
    // DefaultVerifier - The verifier that is now the default, or null if no default verifier remains.
    DefaultVerifier Verifier `json:"defaultVerifier"`

}

// String returns a JSON representation of the model
func (o *Deleteverifierresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Deleteverifierresponse) MarshalJSON() ([]byte, error) {
    type Alias Deleteverifierresponse

    if DeleteverifierresponseMarshalled {
        return []byte("{}"), nil
    }
    DeleteverifierresponseMarshalled = true

    return json.Marshal(&struct {
        
        DefaultVerifier Verifier `json:"defaultVerifier"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

