package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContentmoderationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContentmoderationDud struct { 
    


    

}

// Contentmoderation
type Contentmoderation struct { 
    // Flag - The Content Moderation Flag of the message.
    Flag string `json:"flag"`


    // Categories - The Content Moderation Categories of the message.
    Categories []string `json:"categories"`

}

// String returns a JSON representation of the model
func (o *Contentmoderation) String() string {
    
     o.Categories = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contentmoderation) MarshalJSON() ([]byte, error) {
    type Alias Contentmoderation

    if ContentmoderationMarshalled {
        return []byte("{}"), nil
    }
    ContentmoderationMarshalled = true

    return json.Marshal(&struct {
        
        Flag string `json:"flag"`
        
        Categories []string `json:"categories"`
        *Alias
    }{

        


        
        Categories: []string{""},
        

        Alias: (*Alias)(u),
    })
}

