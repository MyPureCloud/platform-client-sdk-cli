package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpensocialmediamessagecontentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpensocialmediamessagecontentDud struct { 
    


    


    


    

}

// Opensocialmediamessagecontent - Message content element.
type Opensocialmediamessagecontent struct { 
    // ContentType - Type of this content element.
    ContentType string `json:"contentType"`


    // Attachment - Attachment content.
    Attachment Conversationcontentattachment `json:"attachment"`


    // Text - A content type containing text.
    Text Conversationcontenttext `json:"text"`


    // Reaction - A set of reactions to a message.
    Reaction Conversationcontentreaction `json:"reaction"`

}

// String returns a JSON representation of the model
func (o *Opensocialmediamessagecontent) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Opensocialmediamessagecontent) MarshalJSON() ([]byte, error) {
    type Alias Opensocialmediamessagecontent

    if OpensocialmediamessagecontentMarshalled {
        return []byte("{}"), nil
    }
    OpensocialmediamessagecontentMarshalled = true

    return json.Marshal(&struct {
        
        ContentType string `json:"contentType"`
        
        Attachment Conversationcontentattachment `json:"attachment"`
        
        Text Conversationcontenttext `json:"text"`
        
        Reaction Conversationcontentreaction `json:"reaction"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

