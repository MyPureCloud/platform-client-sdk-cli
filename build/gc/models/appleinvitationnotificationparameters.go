package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AppleinvitationnotificationparametersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AppleinvitationnotificationparametersDud struct { 
    

}

// Appleinvitationnotificationparameters - Notification parameters for Apple Invitation
type Appleinvitationnotificationparameters struct { 
    // ReferenceId - An opaque, caller-supplied string that provides business context for the notification message (e.g., an order number or case ID). Max: 1000 characters. Must not be empty and must not contain '?' or apostrophe (') characters.
    ReferenceId string `json:"referenceId"`

}

// String returns a JSON representation of the model
func (o *Appleinvitationnotificationparameters) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Appleinvitationnotificationparameters) MarshalJSON() ([]byte, error) {
    type Alias Appleinvitationnotificationparameters

    if AppleinvitationnotificationparametersMarshalled {
        return []byte("{}"), nil
    }
    AppleinvitationnotificationparametersMarshalled = true

    return json.Marshal(&struct {
        
        ReferenceId string `json:"referenceId"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

