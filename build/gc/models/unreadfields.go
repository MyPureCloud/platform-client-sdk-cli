package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UnreadfieldsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UnreadfieldsDud struct { 
    

}

// Unreadfields
type Unreadfields struct { 
    // State - The new unread state of the alert
    State bool `json:"state"`

}

// String returns a JSON representation of the model
func (o *Unreadfields) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Unreadfields) MarshalJSON() ([]byte, error) {
    type Alias Unreadfields

    if UnreadfieldsMarshalled {
        return []byte("{}"), nil
    }
    UnreadfieldsMarshalled = true

    return json.Marshal(&struct {
        
        State bool `json:"state"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

