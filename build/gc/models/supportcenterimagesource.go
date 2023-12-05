package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SupportcenterimagesourceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SupportcenterimagesourceDud struct { 
    

}

// Supportcenterimagesource
type Supportcenterimagesource struct { 
    // DefaultUrl - Default URL for image
    DefaultUrl string `json:"defaultUrl"`

}

// String returns a JSON representation of the model
func (o *Supportcenterimagesource) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Supportcenterimagesource) MarshalJSON() ([]byte, error) {
    type Alias Supportcenterimagesource

    if SupportcenterimagesourceMarshalled {
        return []byte("{}"), nil
    }
    SupportcenterimagesourceMarshalled = true

    return json.Marshal(&struct {
        
        DefaultUrl string `json:"defaultUrl"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

