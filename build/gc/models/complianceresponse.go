package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ComplianceresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ComplianceresponseDud struct { 
    

}

// Complianceresponse
type Complianceresponse struct { 
    // Message - Message response
    Message string `json:"message"`

}

// String returns a JSON representation of the model
func (o *Complianceresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Complianceresponse) MarshalJSON() ([]byte, error) {
    type Alias Complianceresponse

    if ComplianceresponseMarshalled {
        return []byte("{}"), nil
    }
    ComplianceresponseMarshalled = true

    return json.Marshal(&struct {
        
        Message string `json:"message"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

