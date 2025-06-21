package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagingconferwithuserrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagingconferwithuserrequestDud struct { 
    

}

// Messagingconferwithuserrequest
type Messagingconferwithuserrequest struct { 
    // TargetUserId - The user ID of the target.
    TargetUserId string `json:"targetUserId"`

}

// String returns a JSON representation of the model
func (o *Messagingconferwithuserrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagingconferwithuserrequest) MarshalJSON() ([]byte, error) {
    type Alias Messagingconferwithuserrequest

    if MessagingconferwithuserrequestMarshalled {
        return []byte("{}"), nil
    }
    MessagingconferwithuserrequestMarshalled = true

    return json.Marshal(&struct {
        
        TargetUserId string `json:"targetUserId"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

