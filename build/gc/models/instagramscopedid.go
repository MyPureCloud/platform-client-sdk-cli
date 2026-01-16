package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InstagramscopedidMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InstagramscopedidDud struct { 
    

}

// Instagramscopedid - Scoped ID for an Instagram user interacting with a page or app
type Instagramscopedid struct { 
    // ScopedId - The unique page/app-specific scopedId for the user. Max: 255 characters. Leading and trailing whitespace stripped.
    ScopedId string `json:"scopedId"`

}

// String returns a JSON representation of the model
func (o *Instagramscopedid) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Instagramscopedid) MarshalJSON() ([]byte, error) {
    type Alias Instagramscopedid

    if InstagramscopedidMarshalled {
        return []byte("{}"), nil
    }
    InstagramscopedidMarshalled = true

    return json.Marshal(&struct {
        
        ScopedId string `json:"scopedId"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

