package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationmessagecontentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationmessagecontentDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Conversationmessagecontent - Message content element.
type Conversationmessagecontent struct { 
    // ContentType - Type of this content element. If contentType = \"Attachment\" only one item is allowed.
    ContentType string `json:"contentType"`


    // Location - Location content.
    Location Conversationcontentlocation `json:"location"`


    // Story - Ephemeral story content.
    Story Conversationcontentstory `json:"story"`


    // Attachment - Attachment content.
    Attachment Conversationcontentattachment `json:"attachment"`


    // QuickReply - Quick reply content.
    QuickReply Conversationcontentquickreply `json:"quickReply"`


    // Template - Template notification content.
    Template Conversationcontentnotificationtemplate `json:"template"`


    // ButtonResponse - Button response content.
    ButtonResponse Conversationcontentbuttonresponse `json:"buttonResponse"`


    // Generic - Generic Template Object (Deprecated).
    Generic Conversationcontentgeneric `json:"generic"`


    // Card - Card (Generic Template) Object
    Card Conversationcontentcard `json:"card"`


    // Carousel - Carousel (Multiple Generic Template) Object
    Carousel Conversationcontentcarousel `json:"carousel"`

}

// String returns a JSON representation of the model
func (o *Conversationmessagecontent) String() string {
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationmessagecontent) MarshalJSON() ([]byte, error) {
    type Alias Conversationmessagecontent

    if ConversationmessagecontentMarshalled {
        return []byte("{}"), nil
    }
    ConversationmessagecontentMarshalled = true

    return json.Marshal(&struct {
        
        ContentType string `json:"contentType"`
        
        Location Conversationcontentlocation `json:"location"`
        
        Story Conversationcontentstory `json:"story"`
        
        Attachment Conversationcontentattachment `json:"attachment"`
        
        QuickReply Conversationcontentquickreply `json:"quickReply"`
        
        Template Conversationcontentnotificationtemplate `json:"template"`
        
        ButtonResponse Conversationcontentbuttonresponse `json:"buttonResponse"`
        
        Generic Conversationcontentgeneric `json:"generic"`
        
        Card Conversationcontentcard `json:"card"`
        
        Carousel Conversationcontentcarousel `json:"carousel"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

