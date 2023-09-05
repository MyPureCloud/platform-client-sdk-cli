package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecipientrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecipientrequestDud struct { 
    

}

// Recipientrequest
type Recipientrequest struct { 
    // Flow - An automate flow object which defines the set of actions to be taken, when a message is received by this provisioned phone number.
    Flow Recipientflow `json:"flow"`

}

// String returns a JSON representation of the model
func (o *Recipientrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recipientrequest) MarshalJSON() ([]byte, error) {
    type Alias Recipientrequest

    if RecipientrequestMarshalled {
        return []byte("{}"), nil
    }
    RecipientrequestMarshalled = true

    return json.Marshal(&struct {
        
        Flow Recipientflow `json:"flow"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

