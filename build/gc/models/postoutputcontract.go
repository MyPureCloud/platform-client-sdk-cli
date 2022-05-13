package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PostoutputcontractMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PostoutputcontractDud struct { 
    

}

// Postoutputcontract - The schemas defining all of the expected responses/outputs.
type Postoutputcontract struct { 
    // SuccessSchema - JSON schema that defines the transformed, successful result that will be sent back to the caller.
    SuccessSchema Jsonschemadocument `json:"successSchema"`

}

// String returns a JSON representation of the model
func (o *Postoutputcontract) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Postoutputcontract) MarshalJSON() ([]byte, error) {
    type Alias Postoutputcontract

    if PostoutputcontractMarshalled {
        return []byte("{}"), nil
    }
    PostoutputcontractMarshalled = true

    return json.Marshal(&struct {
        
        SuccessSchema Jsonschemadocument `json:"successSchema"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

