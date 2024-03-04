package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagingcommunicationansweredeventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagingcommunicationansweredeventDud struct { 
    


    


    


    

}

// Messagingcommunicationansweredevent
type Messagingcommunicationansweredevent struct { 
    // EventId - A unique (V4 UUID) eventId for this event
    EventId string `json:"eventId"`


    // EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EventDateTime time.Time `json:"eventDateTime"`


    // ConversationId - A unique Id (V4 UUID) identifying this conversation
    ConversationId string `json:"conversationId"`


    // CommunicationId - A unique Id (V4 UUID) identifying this communication
    CommunicationId string `json:"communicationId"`

}

// String returns a JSON representation of the model
func (o *Messagingcommunicationansweredevent) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagingcommunicationansweredevent) MarshalJSON() ([]byte, error) {
    type Alias Messagingcommunicationansweredevent

    if MessagingcommunicationansweredeventMarshalled {
        return []byte("{}"), nil
    }
    MessagingcommunicationansweredeventMarshalled = true

    return json.Marshal(&struct {
        
        EventId string `json:"eventId"`
        
        EventDateTime time.Time `json:"eventDateTime"`
        
        ConversationId string `json:"conversationId"`
        
        CommunicationId string `json:"communicationId"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

