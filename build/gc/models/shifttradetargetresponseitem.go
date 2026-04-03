package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShifttradetargetresponseitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShifttradetargetresponseitemDud struct { 
    

}

// Shifttradetargetresponseitem
type Shifttradetargetresponseitem struct { 
    // User - The user to whom the shift trade request was sent in a direct trade
    User Userreference `json:"user"`

}

// String returns a JSON representation of the model
func (o *Shifttradetargetresponseitem) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shifttradetargetresponseitem) MarshalJSON() ([]byte, error) {
    type Alias Shifttradetargetresponseitem

    if ShifttradetargetresponseitemMarshalled {
        return []byte("{}"), nil
    }
    ShifttradetargetresponseitemMarshalled = true

    return json.Marshal(&struct {
        
        User Userreference `json:"user"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

