package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MarkdownMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MarkdownDud struct { 
    

}

// Markdown
type Markdown struct { 
    // Enabled - whether or not markdown is enabled
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Markdown) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Markdown) MarshalJSON() ([]byte, error) {
    type Alias Markdown

    if MarkdownMarshalled {
        return []byte("{}"), nil
    }
    MarkdownMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

