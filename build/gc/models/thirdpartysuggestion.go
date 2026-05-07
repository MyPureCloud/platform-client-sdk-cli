package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ThirdpartysuggestionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ThirdpartysuggestionDud struct { 
    

}

// Thirdpartysuggestion
type Thirdpartysuggestion struct { 
    // Text - The third party suggestion text.
    Text string `json:"text"`

}

// String returns a JSON representation of the model
func (o *Thirdpartysuggestion) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Thirdpartysuggestion) MarshalJSON() ([]byte, error) {
    type Alias Thirdpartysuggestion

    if ThirdpartysuggestionMarshalled {
        return []byte("{}"), nil
    }
    ThirdpartysuggestionMarshalled = true

    return json.Marshal(&struct {
        
        Text string `json:"text"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

