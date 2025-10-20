package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IdentityresolutionqueueconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IdentityresolutionqueueconfigDud struct { 
    

}

// Identityresolutionqueueconfig
type Identityresolutionqueueconfig struct { 
    // CallOnBehalfOfQueue
    CallOnBehalfOfQueue Outboundqueueidentityresolutionconfig `json:"callOnBehalfOfQueue"`

}

// String returns a JSON representation of the model
func (o *Identityresolutionqueueconfig) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Identityresolutionqueueconfig) MarshalJSON() ([]byte, error) {
    type Alias Identityresolutionqueueconfig

    if IdentityresolutionqueueconfigMarshalled {
        return []byte("{}"), nil
    }
    IdentityresolutionqueueconfigMarshalled = true

    return json.Marshal(&struct {
        
        CallOnBehalfOfQueue Outboundqueueidentityresolutionconfig `json:"callOnBehalfOfQueue"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

