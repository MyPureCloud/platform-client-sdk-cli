package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReplymessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReplymessageDud struct { 
    


    


    

}

// Replymessage
type Replymessage struct { 
    // VarType - Message type.
    VarType string `json:"type"`


    // Text - Message text.
    Text string `json:"text"`


    // Content - A list of content elements in the message
    Content []Conversationmessagecontent `json:"content"`

}

// String returns a JSON representation of the model
func (o *Replymessage) String() string {
    
    
     o.Content = []Conversationmessagecontent{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Replymessage) MarshalJSON() ([]byte, error) {
    type Alias Replymessage

    if ReplymessageMarshalled {
        return []byte("{}"), nil
    }
    ReplymessageMarshalled = true

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

