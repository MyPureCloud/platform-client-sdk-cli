package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShifttradetargetrequestitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShifttradetargetrequestitemDud struct { 
    

}

// Shifttradetargetrequestitem
type Shifttradetargetrequestitem struct { 
    // UserId - The user to whom the shift trade request was sent in a direct trade
    UserId string `json:"userId"`

}

// String returns a JSON representation of the model
func (o *Shifttradetargetrequestitem) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shifttradetargetrequestitem) MarshalJSON() ([]byte, error) {
    type Alias Shifttradetargetrequestitem

    if ShifttradetargetrequestitemMarshalled {
        return []byte("{}"), nil
    }
    ShifttradetargetrequestitemMarshalled = true

    return json.Marshal(&struct {
        
        UserId string `json:"userId"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

