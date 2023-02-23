package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TransferinitiatorMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TransferinitiatorDud struct { 
    

}

// Transferinitiator
type Transferinitiator struct { 
    // UserId - The id of the user who initiated the command if it was initiated by a user.
    UserId string `json:"userId"`

}

// String returns a JSON representation of the model
func (o *Transferinitiator) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transferinitiator) MarshalJSON() ([]byte, error) {
    type Alias Transferinitiator

    if TransferinitiatorMarshalled {
        return []byte("{}"), nil
    }
    TransferinitiatorMarshalled = true

    return json.Marshal(&struct {
        
        UserId string `json:"userId"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

