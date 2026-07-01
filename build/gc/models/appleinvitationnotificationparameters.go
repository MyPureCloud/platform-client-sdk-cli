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
    // ReferenceId - Provides context for the notification message, such as an order number or case ID.
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

