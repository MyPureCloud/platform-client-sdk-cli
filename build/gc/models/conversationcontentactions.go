package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontentactionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontentactionsDud struct { 
    


    


    

}

// Conversationcontentactions - User actions available on the content. All actions are optional and all actions are executed simultaneously.
type Conversationcontentactions struct { 
    // Url - A URL of a web page to direct the user to.
    Url string `json:"url"`


    // UrlTarget - The target window in which to open the URL. If empty will open a blank page or tab.
    UrlTarget string `json:"urlTarget"`


    // Textback - Text to be sent back in reply when the item is selected.
    Textback string `json:"textback"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentactions) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentactions) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentactions

    if ConversationcontentactionsMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontentactionsMarshalled = true

    return json.Marshal(&struct { 
        Url string `json:"url"`
        
        UrlTarget string `json:"urlTarget"`
        
        Textback string `json:"textback"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

