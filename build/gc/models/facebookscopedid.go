package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FacebookscopedidMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FacebookscopedidDud struct { 
    

}

// Facebookscopedid - Scoped ID for a Facebook user interacting with a page or app
type Facebookscopedid struct { 
    // ScopedId - The unique page/app-specific scopedId for the user
    ScopedId string `json:"scopedId"`

}

// String returns a JSON representation of the model
func (o *Facebookscopedid) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Facebookscopedid) MarshalJSON() ([]byte, error) {
    type Alias Facebookscopedid

    if FacebookscopedidMarshalled {
        return []byte("{}"), nil
    }
    FacebookscopedidMarshalled = true

    return json.Marshal(&struct {
        
        ScopedId string `json:"scopedId"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

