package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UnreadmetricMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UnreadmetricDud struct { 
    

}

// Unreadmetric
type Unreadmetric struct { 
    // Count - The count of unread alerts for a specific rule type.
    Count int `json:"count"`

}

// String returns a JSON representation of the model
func (o *Unreadmetric) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Unreadmetric) MarshalJSON() ([]byte, error) {
    type Alias Unreadmetric

    if UnreadmetricMarshalled {
        return []byte("{}"), nil
    }
    UnreadmetricMarshalled = true

    return json.Marshal(&struct {
        
        Count int `json:"count"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

