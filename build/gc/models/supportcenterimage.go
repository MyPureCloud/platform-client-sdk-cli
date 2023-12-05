package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SupportcenterimageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SupportcenterimageDud struct { 
    

}

// Supportcenterimage
type Supportcenterimage struct { 
    // Source - Source URLs for image
    Source Supportcenterimagesource `json:"source"`

}

// String returns a JSON representation of the model
func (o *Supportcenterimage) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Supportcenterimage) MarshalJSON() ([]byte, error) {
    type Alias Supportcenterimage

    if SupportcenterimageMarshalled {
        return []byte("{}"), nil
    }
    SupportcenterimageMarshalled = true

    return json.Marshal(&struct {
        
        Source Supportcenterimagesource `json:"source"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

