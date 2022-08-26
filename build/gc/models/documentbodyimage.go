package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentbodyimageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentbodyimageDud struct { 
    

}

// Documentbodyimage
type Documentbodyimage struct { 
    // Url - The URL for the image.
    Url string `json:"url"`

}

// String returns a JSON representation of the model
func (o *Documentbodyimage) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentbodyimage) MarshalJSON() ([]byte, error) {
    type Alias Documentbodyimage

    if DocumentbodyimageMarshalled {
        return []byte("{}"), nil
    }
    DocumentbodyimageMarshalled = true

    return json.Marshal(&struct {
        
        Url string `json:"url"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

