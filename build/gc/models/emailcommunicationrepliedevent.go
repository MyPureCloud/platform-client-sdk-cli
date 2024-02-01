package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmailcommunicationrepliedeventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmailcommunicationrepliedeventDud struct { 
    


    


    


    

}

// Emailcommunicationrepliedevent
type Emailcommunicationrepliedevent struct { 
    // EventId - A unique (V4 UUID) eventId for this event
    EventId string `json:"eventId"`


    // EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EventDateTime time.Time `json:"eventDateTime"`


    // ConversationId - A unique Id (V4 UUID) identifying this conversation
    ConversationId string `json:"conversationId"`


    // CommunicationId - A unique Id (V4 UUID) identifying this communication.
    CommunicationId string `json:"communicationId"`

}

// String returns a JSON representation of the model
func (o *Emailcommunicationrepliedevent) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emailcommunicationrepliedevent) MarshalJSON() ([]byte, error) {
    type Alias Emailcommunicationrepliedevent

    if EmailcommunicationrepliedeventMarshalled {
        return []byte("{}"), nil
    }
    EmailcommunicationrepliedeventMarshalled = true

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

