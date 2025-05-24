package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontentrequiredcontactfieldMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontentrequiredcontactfieldDud struct { 
    

}

// Conversationcontentrequiredcontactfield - Contact fields a merchant requires to complete a payment request.
type Conversationcontentrequiredcontactfield struct { 
    // ContactField - The name of the contact field
    ContactField string `json:"contactField"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentrequiredcontactfield) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentrequiredcontactfield) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentrequiredcontactfield

    if ConversationcontentrequiredcontactfieldMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontentrequiredcontactfieldMarshalled = true

    return json.Marshal(&struct {
        
        ContactField string `json:"contactField"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

