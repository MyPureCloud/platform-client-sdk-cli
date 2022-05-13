package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UnreadstatusMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UnreadstatusDud struct { 
    

}

// Unreadstatus
type Unreadstatus struct { 
    // Unread - Sets if the alert is read or unread.
    Unread bool `json:"unread"`

}

// String returns a JSON representation of the model
func (o *Unreadstatus) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Unreadstatus) MarshalJSON() ([]byte, error) {
    type Alias Unreadstatus

    if UnreadstatusMarshalled {
        return []byte("{}"), nil
    }
    UnreadstatusMarshalled = true

    return json.Marshal(&struct {
        
        Unread bool `json:"unread"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

