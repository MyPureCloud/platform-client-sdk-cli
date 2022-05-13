package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UrlresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UrlresponseDud struct { 
    

}

// Urlresponse
type Urlresponse struct { 
    // Url
    Url string `json:"url"`

}

// String returns a JSON representation of the model
func (o *Urlresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Urlresponse) MarshalJSON() ([]byte, error) {
    type Alias Urlresponse

    if UrlresponseMarshalled {
        return []byte("{}"), nil
    }
    UrlresponseMarshalled = true

    return json.Marshal(&struct {
        
        Url string `json:"url"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

