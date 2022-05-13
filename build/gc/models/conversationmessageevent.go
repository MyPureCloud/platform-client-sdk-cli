package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationmessageeventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationmessageeventDud struct { 
    


    


    


    

}

// Conversationmessageevent - Message event element.  Examples include: system messages, typing indicators, cobrowse offerings.
type Conversationmessageevent struct { 
    // EventType - Type of this event element
    EventType string `json:"eventType"`


    // CoBrowse - CoBrowse event.
    CoBrowse Conversationeventcobrowse `json:"coBrowse"`


    // Typing - Typing event.
    Typing Conversationeventtyping `json:"typing"`


    // Presence - Presence event.
    Presence Conversationeventpresence `json:"presence"`

}

// String returns a JSON representation of the model
func (o *Conversationmessageevent) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationmessageevent) MarshalJSON() ([]byte, error) {
    type Alias Conversationmessageevent

    if ConversationmessageeventMarshalled {
        return []byte("{}"), nil
    }
    ConversationmessageeventMarshalled = true

    return json.Marshal(&struct {
        
        EventType string `json:"eventType"`
        
        CoBrowse Conversationeventcobrowse `json:"coBrowse"`
        
        Typing Conversationeventtyping `json:"typing"`
        
        Presence Conversationeventpresence `json:"presence"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

