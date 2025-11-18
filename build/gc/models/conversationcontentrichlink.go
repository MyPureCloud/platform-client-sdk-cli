package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontentrichlinkMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontentrichlinkDud struct { 
    


    


    


    


    

}

// Conversationcontentrichlink - A Rich Link attachment
type Conversationcontentrichlink struct { 
    // Header - Header for the Rich Link.
    Header Conversationcontentrichlinkheader `json:"header"`


    // Footer - Footer text for the Rich Link.
    Footer string `json:"footer"`


    // Text - Text for the body of the Rich Link.
    Text string `json:"text"`


    // UrlLabel - Label for the URL of the Rich Link.
    UrlLabel string `json:"urlLabel"`


    // Url - Url for the Rich Link.
    Url string `json:"url"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentrichlink) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentrichlink) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentrichlink

    if ConversationcontentrichlinkMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontentrichlinkMarshalled = true

    return json.Marshal(&struct {
        
        Header Conversationcontentrichlinkheader `json:"header"`
        
        Footer string `json:"footer"`
        
        Text string `json:"text"`
        
        UrlLabel string `json:"urlLabel"`
        
        Url string `json:"url"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

