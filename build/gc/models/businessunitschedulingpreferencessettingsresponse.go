package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BusinessunitschedulingpreferencessettingsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BusinessunitschedulingpreferencessettingsresponseDud struct { 
    

}

// Businessunitschedulingpreferencessettingsresponse
type Businessunitschedulingpreferencessettingsresponse struct { 
    // Enabled - Indicates whether scheduling preferences are enabled for the business unit
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Businessunitschedulingpreferencessettingsresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Businessunitschedulingpreferencessettingsresponse) MarshalJSON() ([]byte, error) {
    type Alias Businessunitschedulingpreferencessettingsresponse

    if BusinessunitschedulingpreferencessettingsresponseMarshalled {
        return []byte("{}"), nil
    }
    BusinessunitschedulingpreferencessettingsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

