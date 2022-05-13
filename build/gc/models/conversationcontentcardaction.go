package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontentcardactionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontentcardactionDud struct { 
    


    


    


    

}

// Conversationcontentcardaction - CardAction Object
type Conversationcontentcardaction struct { 
    // VarType - Describes the type of action.
    VarType string `json:"type"`


    // Text - The response text from the button click.
    Text string `json:"text"`


    // Payload - Text to be returned as the payload from a ButtonResponse when a button is clicked.
    Payload string `json:"payload"`


    // Url - A URL of a web page to direct the user to.
    Url string `json:"url"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentcardaction) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentcardaction) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentcardaction

    if ConversationcontentcardactionMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontentcardactionMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Text string `json:"text"`
        
        Payload string `json:"payload"`
        
        Url string `json:"url"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

