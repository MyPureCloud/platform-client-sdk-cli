package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MandatorypostcallactioninputMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MandatorypostcallactioninputDud struct { 
    


    

}

// Mandatorypostcallactioninput
type Mandatorypostcallactioninput struct { 
    // Destination - The destination phone number or phone id to send the flow to after completion.  If null or blank will not update.
    Destination string `json:"destination"`


    // InvocationData - The invocation data to send the destination after completion.  If null or blank will not update.  Invocation data is only supported when the destination is an Inbound Call flow or Voice Survey flow.
    InvocationData string `json:"invocationData"`

}

// String returns a JSON representation of the model
func (o *Mandatorypostcallactioninput) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mandatorypostcallactioninput) MarshalJSON() ([]byte, error) {
    type Alias Mandatorypostcallactioninput

    if MandatorypostcallactioninputMarshalled {
        return []byte("{}"), nil
    }
    MandatorypostcallactioninputMarshalled = true

    return json.Marshal(&struct {
        
        Destination string `json:"destination"`
        
        InvocationData string `json:"invocationData"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

