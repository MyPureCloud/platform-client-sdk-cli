package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentbodyvideoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentbodyvideoDud struct { 
    

}

// Documentbodyvideo
type Documentbodyvideo struct { 
    // Url - The URL for the video.
    Url string `json:"url"`

}

// String returns a JSON representation of the model
func (o *Documentbodyvideo) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentbodyvideo) MarshalJSON() ([]byte, error) {
    type Alias Documentbodyvideo

    if DocumentbodyvideoMarshalled {
        return []byte("{}"), nil
    }
    DocumentbodyvideoMarshalled = true

    return json.Marshal(&struct {
        
        Url string `json:"url"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

