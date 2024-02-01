package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmailprogresstransfereventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmailprogresstransfereventDud struct { 
    


    


    


    


    


    

}

// Emailprogresstransferevent
type Emailprogresstransferevent struct { 
    // EventId - A unique (V4 UUID) eventId for this event
    EventId string `json:"eventId"`


    // EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EventDateTime time.Time `json:"eventDateTime"`


    // ConversationId - A unique Id (V4 UUID) identifying this conversation
    ConversationId string `json:"conversationId"`


    // CommandId - The id (V4 UUID) used to identify the transfer already started by the external platform.
    CommandId string `json:"commandId"`


    // ObjectCommunicationId - The id (V4 UUID) of the communication that is being transferred.
    ObjectCommunicationId string `json:"objectCommunicationId"`


    // DestinationCommunicationId - The id (V4 UUID) of the communication that is being transferred to.
    DestinationCommunicationId string `json:"destinationCommunicationId"`

}

// String returns a JSON representation of the model
func (o *Emailprogresstransferevent) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emailprogresstransferevent) MarshalJSON() ([]byte, error) {
    type Alias Emailprogresstransferevent

    if EmailprogresstransfereventMarshalled {
        return []byte("{}"), nil
    }
    EmailprogresstransfereventMarshalled = true

    return json.Marshal(&struct {
        
        EventId string `json:"eventId"`
        
        EventDateTime time.Time `json:"eventDateTime"`
        
        ConversationId string `json:"conversationId"`
        
        CommandId string `json:"commandId"`
        
        ObjectCommunicationId string `json:"objectCommunicationId"`
        
        DestinationCommunicationId string `json:"destinationCommunicationId"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

