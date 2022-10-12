package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PhonetransfereventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PhonetransfereventDud struct { 
    


    


    


    


    


    


    


    


    

}

// Phonetransferevent
type Phonetransferevent struct { 
    // EventId - A unique (V4 UUID) eventId for this event
    EventId string `json:"eventId"`


    // EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EventDateTime time.Time `json:"eventDateTime"`


    // ConversationId - A unique Id (V4 UUID) identifying this conversation
    ConversationId string `json:"conversationId"`


    // TransferType - Indicates the desired type of transfer.
    TransferType string `json:"transferType"`


    // CommandId - The id (V4 UUID) used by the external platform to refer to the transfer in subsequent Transfer events.
    CommandId string `json:"commandId"`


    // InitiatingCommunicationId - The id (V4 UUID) of the communication representing the participant that is initiating the transfer.
    InitiatingCommunicationId string `json:"initiatingCommunicationId"`


    // TargetCommunicationId - The id (V4 UUID) of the communication that is being transferred away from. In many cases this will be the same as the `initiatingCommunicationId`.
    TargetCommunicationId string `json:"targetCommunicationId"`


    // ObjectCommunicationId - The id (V4 UUID) of the communication that is being transferred.
    ObjectCommunicationId string `json:"objectCommunicationId"`


    // DestinationPhoneNumber - The desired destination phone number that the object communication should be transferred to.
    DestinationPhoneNumber string `json:"destinationPhoneNumber"`

}

// String returns a JSON representation of the model
func (o *Phonetransferevent) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Phonetransferevent) MarshalJSON() ([]byte, error) {
    type Alias Phonetransferevent

    if PhonetransfereventMarshalled {
        return []byte("{}"), nil
    }
    PhonetransfereventMarshalled = true

    return json.Marshal(&struct {
        
        EventId string `json:"eventId"`
        
        EventDateTime time.Time `json:"eventDateTime"`
        
        ConversationId string `json:"conversationId"`
        
        TransferType string `json:"transferType"`
        
        CommandId string `json:"commandId"`
        
        InitiatingCommunicationId string `json:"initiatingCommunicationId"`
        
        TargetCommunicationId string `json:"targetCommunicationId"`
        
        ObjectCommunicationId string `json:"objectCommunicationId"`
        
        DestinationPhoneNumber string `json:"destinationPhoneNumber"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

