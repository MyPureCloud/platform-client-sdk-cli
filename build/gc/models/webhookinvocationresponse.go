package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebhookinvocationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebhookinvocationresponseDud struct { 
    

}

// Webhookinvocationresponse
type Webhookinvocationresponse struct { 
    // InvocationId - The id of the Webhook Invocation
    InvocationId string `json:"invocationId"`

}

// String returns a JSON representation of the model
func (o *Webhookinvocationresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webhookinvocationresponse) MarshalJSON() ([]byte, error) {
    type Alias Webhookinvocationresponse

    if WebhookinvocationresponseMarshalled {
        return []byte("{}"), nil
    }
    WebhookinvocationresponseMarshalled = true

    return json.Marshal(&struct {
        
        InvocationId string `json:"invocationId"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

