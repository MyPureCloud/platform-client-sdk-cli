package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebdeploymentheadlessmodeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebdeploymentheadlessmodeDud struct { 
    

}

// Webdeploymentheadlessmode
type Webdeploymentheadlessmode struct { 
    // Enabled - Whether or not Headless Mode is enabled
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Webdeploymentheadlessmode) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webdeploymentheadlessmode) MarshalJSON() ([]byte, error) {
    type Alias Webdeploymentheadlessmode

    if WebdeploymentheadlessmodeMarshalled {
        return []byte("{}"), nil
    }
    WebdeploymentheadlessmodeMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

