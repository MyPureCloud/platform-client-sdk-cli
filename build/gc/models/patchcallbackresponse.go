package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchcallbackresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchcallbackresponseDud struct { 
    


    

}

// Patchcallbackresponse
type Patchcallbackresponse struct { 
    // Conversation - The conversation associated with the callback
    Conversation Domainentityref `json:"conversation"`


    // CallbackIdentifiers - The list of communication identifiers for the callback participants
    CallbackIdentifiers []Callbackidentifier `json:"callbackIdentifiers"`

}

// String returns a JSON representation of the model
func (o *Patchcallbackresponse) String() string {
    
     o.CallbackIdentifiers = []Callbackidentifier{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchcallbackresponse) MarshalJSON() ([]byte, error) {
    type Alias Patchcallbackresponse

    if PatchcallbackresponseMarshalled {
        return []byte("{}"), nil
    }
    PatchcallbackresponseMarshalled = true

    return json.Marshal(&struct {
        
        Conversation Domainentityref `json:"conversation"`
        
        CallbackIdentifiers []Callbackidentifier `json:"callbackIdentifiers"`
        *Alias
    }{

        


        
        CallbackIdentifiers: []Callbackidentifier{{}},
        

        Alias: (*Alias)(u),
    })
}

