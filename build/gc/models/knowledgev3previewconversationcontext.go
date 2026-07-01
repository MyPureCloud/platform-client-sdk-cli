package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Knowledgev3previewconversationcontextMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Knowledgev3previewconversationcontextDud struct { 
    

}

// Knowledgev3previewconversationcontext
type Knowledgev3previewconversationcontext struct { 
    // MediaType - The media type to simulate for the preview.
    MediaType string `json:"mediaType"`

}

// String returns a JSON representation of the model
func (o *Knowledgev3previewconversationcontext) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgev3previewconversationcontext) MarshalJSON() ([]byte, error) {
    type Alias Knowledgev3previewconversationcontext

    if Knowledgev3previewconversationcontextMarshalled {
        return []byte("{}"), nil
    }
    Knowledgev3previewconversationcontextMarshalled = true

    return json.Marshal(&struct {
        
        MediaType string `json:"mediaType"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

