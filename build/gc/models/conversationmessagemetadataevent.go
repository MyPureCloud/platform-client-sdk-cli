package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationmessagemetadataeventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationmessagemetadataeventDud struct { 
    


    

}

// Conversationmessagemetadataevent - Metadata information about a message event.
type Conversationmessagemetadataevent struct { 
    // EventType - Type of this event element
    EventType string `json:"eventType"`


    // SubType - Event subtype
    SubType string `json:"subType"`

}

// String returns a JSON representation of the model
func (o *Conversationmessagemetadataevent) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationmessagemetadataevent) MarshalJSON() ([]byte, error) {
    type Alias Conversationmessagemetadataevent

    if ConversationmessagemetadataeventMarshalled {
        return []byte("{}"), nil
    }
    ConversationmessagemetadataeventMarshalled = true

    return json.Marshal(&struct {
        
        EventType string `json:"eventType"`
        
        SubType string `json:"subType"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

