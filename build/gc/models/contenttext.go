package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContenttextMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContenttextDud struct { 
    


    

}

// Contenttext - Message content element containing text only.
type Contenttext struct { 
    // VarType - Type of text content.
    VarType string `json:"type"`


    // Body - Text to be shown for this content element.
    Body string `json:"body"`

}

// String returns a JSON representation of the model
func (o *Contenttext) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contenttext) MarshalJSON() ([]byte, error) {
    type Alias Contenttext

    if ContenttextMarshalled {
        return []byte("{}"), nil
    }
    ContenttextMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Body string `json:"body"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

