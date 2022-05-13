package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NludetectionrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NludetectionrequestDud struct { 
    


    

}

// Nludetectionrequest
type Nludetectionrequest struct { 
    // Input - The input subject to NLU detection.
    Input Nludetectioninput `json:"input"`


    // Context - The context for the input to NLU detection.
    Context Nludetectioncontext `json:"context"`

}

// String returns a JSON representation of the model
func (o *Nludetectionrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nludetectionrequest) MarshalJSON() ([]byte, error) {
    type Alias Nludetectionrequest

    if NludetectionrequestMarshalled {
        return []byte("{}"), nil
    }
    NludetectionrequestMarshalled = true

    return json.Marshal(&struct {
        
        Input Nludetectioninput `json:"input"`
        
        Context Nludetectioncontext `json:"context"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

