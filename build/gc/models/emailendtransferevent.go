package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmailendtransfereventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmailendtransfereventDud struct { 
    


    


    


    


    


    

}

// Emailendtransferevent
type Emailendtransferevent struct { 
    // EventId - A unique (V4 UUID) eventId for this event
    EventId string `json:"eventId"`


    // EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EventDateTime time.Time `json:"eventDateTime"`


    // ConversationId - A unique Id (V4 UUID) identifying this conversation
    ConversationId string `json:"conversationId"`


    // CommandId - The id (V4 UUID) used to identify the transfer already started by the external platform.
    CommandId string `json:"commandId"`


    // FinalState - Indicates whether the transfer completed successfully, was cancelled, or failed for some reason.
    FinalState string `json:"finalState"`


    // ObjectCommunicationId - The id (V4 UUID) of the communication that was being transferred.
    ObjectCommunicationId string `json:"objectCommunicationId"`

}

// String returns a JSON representation of the model
func (o *Emailendtransferevent) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emailendtransferevent) MarshalJSON() ([]byte, error) {
    type Alias Emailendtransferevent

    if EmailendtransfereventMarshalled {
        return []byte("{}"), nil
    }
    EmailendtransfereventMarshalled = true

    return json.Marshal(&struct {
        
        EventId string `json:"eventId"`
        
        EventDateTime time.Time `json:"eventDateTime"`
        
        ConversationId string `json:"conversationId"`
        
        CommandId string `json:"commandId"`
        
        FinalState string `json:"finalState"`
        
        ObjectCommunicationId string `json:"objectCommunicationId"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

