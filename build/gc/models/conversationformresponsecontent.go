package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationformresponsecontentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationformresponsecontentDud struct { 
    


    

}

// Conversationformresponsecontent - Message content element for form responses
type Conversationformresponsecontent struct { 
    // ContentType - Type of this content element.
    ContentType string `json:"contentType"`


    // ButtonResponse - Button response content.
    ButtonResponse Conversationcontentbuttonresponse `json:"buttonResponse"`

}

// String returns a JSON representation of the model
func (o *Conversationformresponsecontent) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationformresponsecontent) MarshalJSON() ([]byte, error) {
    type Alias Conversationformresponsecontent

    if ConversationformresponsecontentMarshalled {
        return []byte("{}"), nil
    }
    ConversationformresponsecontentMarshalled = true

    return json.Marshal(&struct {
        
        ContentType string `json:"contentType"`
        
        ButtonResponse Conversationcontentbuttonresponse `json:"buttonResponse"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

