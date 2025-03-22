package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SocialmediamessageescalationinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SocialmediamessageescalationinfoDud struct { 
    

}

// Socialmediamessageescalationinfo
type Socialmediamessageescalationinfo struct { 
    // EscalationStatus
    EscalationStatus string `json:"escalationStatus"`

}

// String returns a JSON representation of the model
func (o *Socialmediamessageescalationinfo) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Socialmediamessageescalationinfo) MarshalJSON() ([]byte, error) {
    type Alias Socialmediamessageescalationinfo

    if SocialmediamessageescalationinfoMarshalled {
        return []byte("{}"), nil
    }
    SocialmediamessageescalationinfoMarshalled = true

    return json.Marshal(&struct {
        
        EscalationStatus string `json:"escalationStatus"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

