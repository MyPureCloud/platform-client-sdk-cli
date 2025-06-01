package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InputmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InputmessageDud struct { 
    


    


    

}

// Inputmessage
type Inputmessage struct { 
    // VarType - Message type.
    VarType string `json:"type"`


    // Text - The text of the message being sent to the bot.
    Text string `json:"text"`


    // Content - A list of content elements in the message.
    Content []Conversationmessagecontent `json:"content"`

}

// String returns a JSON representation of the model
func (o *Inputmessage) String() string {
    
    
     o.Content = []Conversationmessagecontent{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Inputmessage) MarshalJSON() ([]byte, error) {
    type Alias Inputmessage

    if InputmessageMarshalled {
        return []byte("{}"), nil
    }
    InputmessageMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Text string `json:"text"`
        
        Content []Conversationmessagecontent `json:"content"`
        *Alias
    }{

        


        


        
        Content: []Conversationmessagecontent{{}},
        

        Alias: (*Alias)(u),
    })
}

