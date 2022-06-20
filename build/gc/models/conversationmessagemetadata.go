package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationmessagemetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationmessagemetadataDud struct { 
    


    


    

}

// Conversationmessagemetadata - Metadata information about a message.
type Conversationmessagemetadata struct { 
    // VarType - Message type.
    VarType string `json:"type"`


    // Events - List of events metadata
    Events []Conversationmessagemetadataevent `json:"events"`


    // Content - List of message content
    Content []Conversationmessagemetadatacontent `json:"content"`

}

// String returns a JSON representation of the model
func (o *Conversationmessagemetadata) String() string {
    
     o.Events = []Conversationmessagemetadataevent{{}} 
     o.Content = []Conversationmessagemetadatacontent{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationmessagemetadata) MarshalJSON() ([]byte, error) {
    type Alias Conversationmessagemetadata

    if ConversationmessagemetadataMarshalled {
        return []byte("{}"), nil
    }
    ConversationmessagemetadataMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Events []Conversationmessagemetadataevent `json:"events"`
        
        Content []Conversationmessagemetadatacontent `json:"content"`
        *Alias
    }{

        


        
        Events: []Conversationmessagemetadataevent{{}},
        


        
        Content: []Conversationmessagemetadatacontent{{}},
        

        Alias: (*Alias)(u),
    })
}

