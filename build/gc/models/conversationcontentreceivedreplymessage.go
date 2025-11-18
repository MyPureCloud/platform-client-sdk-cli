package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontentreceivedreplymessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontentreceivedreplymessageDud struct { 
    


    


    


    


    

}

// Conversationcontentreceivedreplymessage - ReceivedReplyMessage content object.
type Conversationcontentreceivedreplymessage struct { 
    // Header - Text to show in the header.
    Header string `json:"header"`


    // Title - Text to show in the title.
    Title string `json:"title"`


    // Subtitle - Text to show in the subtitle.
    Subtitle string `json:"subtitle"`


    // ButtonLabel - Text to show on the button label.
    ButtonLabel string `json:"buttonLabel"`


    // ImageUrl - URL of an image.
    ImageUrl string `json:"imageUrl"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentreceivedreplymessage) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentreceivedreplymessage) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentreceivedreplymessage

    if ConversationcontentreceivedreplymessageMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontentreceivedreplymessageMarshalled = true

    return json.Marshal(&struct {
        
        Header string `json:"header"`
        
        Title string `json:"title"`
        
        Subtitle string `json:"subtitle"`
        
        ButtonLabel string `json:"buttonLabel"`
        
        ImageUrl string `json:"imageUrl"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

