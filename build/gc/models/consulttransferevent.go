package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConsulttransfereventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConsulttransfereventDud struct { 
    


    


    


    


    


    

}

// Consulttransferevent
type Consulttransferevent struct { 
    // EventId - A unique (V4 UUID) eventId for this event
    EventId string `json:"eventId"`


    // EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EventDateTime time.Time `json:"eventDateTime"`


    // ConversationId - A unique Id (V4 UUID) identifying this conversation
    ConversationId string `json:"conversationId"`


    // InitiatingCommunicationId - The id (V4 UUID) of the communication representing the participant that is initiating the transfer.
    InitiatingCommunicationId string `json:"initiatingCommunicationId"`


    // DestinationCommunicationId - The id (V4 UUID) of the communication that is being transferred to.
    DestinationCommunicationId string `json:"destinationCommunicationId"`


    // ObjectCommunicationId - The id (V4 UUID) of the communication that is being transferred.
    ObjectCommunicationId string `json:"objectCommunicationId"`

}

// String returns a JSON representation of the model
func (o *Consulttransferevent) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Consulttransferevent) MarshalJSON() ([]byte, error) {
    type Alias Consulttransferevent

    if ConsulttransfereventMarshalled {
        return []byte("{}"), nil
    }
    ConsulttransfereventMarshalled = true

    return json.Marshal(&struct {
        
        EventId string `json:"eventId"`
        
        EventDateTime time.Time `json:"eventDateTime"`
        
        ConversationId string `json:"conversationId"`
        
        InitiatingCommunicationId string `json:"initiatingCommunicationId"`
        
        DestinationCommunicationId string `json:"destinationCommunicationId"`
        
        ObjectCommunicationId string `json:"objectCommunicationId"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

