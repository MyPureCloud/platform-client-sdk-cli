package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpenmessagecontentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpenmessagecontentDud struct { 
    


    

}

// Openmessagecontent - Message content element.
type Openmessagecontent struct { 
    // ContentType - Type of this content element. If contentType = \"Attachment\" only one item is allowed.
    ContentType string `json:"contentType"`


    // Attachment - Attachment content.
    Attachment Contentattachment `json:"attachment"`

}

// String returns a JSON representation of the model
func (o *Openmessagecontent) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Openmessagecontent) MarshalJSON() ([]byte, error) {
    type Alias Openmessagecontent

    if OpenmessagecontentMarshalled {
        return []byte("{}"), nil
    }
    OpenmessagecontentMarshalled = true

    return json.Marshal(&struct {
        
        ContentType string `json:"contentType"`
        
        Attachment Contentattachment `json:"attachment"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

