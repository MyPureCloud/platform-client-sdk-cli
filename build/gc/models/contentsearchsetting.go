package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContentsearchsettingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContentsearchsettingDud struct { 
    

}

// Contentsearchsetting
type Contentsearchsetting struct { 
    // ContentSearchEnabled - If true, enables full-text search of interaction transcripts
    ContentSearchEnabled bool `json:"contentSearchEnabled"`

}

// String returns a JSON representation of the model
func (o *Contentsearchsetting) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contentsearchsetting) MarshalJSON() ([]byte, error) {
    type Alias Contentsearchsetting

    if ContentsearchsettingMarshalled {
        return []byte("{}"), nil
    }
    ContentsearchsettingMarshalled = true

    return json.Marshal(&struct {
        
        ContentSearchEnabled bool `json:"contentSearchEnabled"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

