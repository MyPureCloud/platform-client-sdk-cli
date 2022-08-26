package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumenttextMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumenttextDud struct { 
    


    


    

}

// Documenttext
type Documenttext struct { 
    // Text - Text.
    Text string `json:"text"`


    // Marks - The unique list of marks (whether it is bold and/or underlined etc.) for the text.
    Marks []string `json:"marks"`


    // Hyperlink - The URL of the page that the hyperlink goes to.
    Hyperlink string `json:"hyperlink"`

}

// String returns a JSON representation of the model
func (o *Documenttext) String() string {
    
     o.Marks = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documenttext) MarshalJSON() ([]byte, error) {
    type Alias Documenttext

    if DocumenttextMarshalled {
        return []byte("{}"), nil
    }
    DocumenttextMarshalled = true

    return json.Marshal(&struct {
        
        Text string `json:"text"`
        
        Marks []string `json:"marks"`
        
        Hyperlink string `json:"hyperlink"`
        *Alias
    }{

        


        
        Marks: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

