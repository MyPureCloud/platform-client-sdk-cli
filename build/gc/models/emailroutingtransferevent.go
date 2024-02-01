package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmailroutingtransfereventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmailroutingtransfereventDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Emailroutingtransferevent
type Emailroutingtransferevent struct { 
    // EventId - A unique (V4 UUID) eventId for this event
    EventId string `json:"eventId"`


    // EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EventDateTime time.Time `json:"eventDateTime"`


    // ConversationId - A unique Id (V4 UUID) identifying this conversation
    ConversationId string `json:"conversationId"`


    // TransferType - Indicates the desired type of transfer.
    TransferType string `json:"transferType"`


    // CommandId - The id (V4 UUID) used by the external platform to refer to the transfer in subsequent *Transfer events.
    CommandId string `json:"commandId"`


    // InitiatingCommunicationId - Indicates the desired type of transfer.
    InitiatingCommunicationId string `json:"initiatingCommunicationId"`


    // TargetCommunicationId - The id (V4 UUID) of the communication that is being transferred away from. In many cases this will be the same as the `initiatingCommunicationId`.
    TargetCommunicationId string `json:"targetCommunicationId"`


    // ObjectCommunicationId - The id (V4 UUID) of the communication that is being transferred.
    ObjectCommunicationId string `json:"objectCommunicationId"`


    // DestinationQueueId - The id (V4 UUID) of the desired destination queue that the object communication should be transferred to.
    DestinationQueueId string `json:"destinationQueueId"`


    // LanguageId - The unique identifier (V4 UUID) for the language that should be used to determine the destination for the conversation.
    LanguageId string `json:"languageId"`


    // SkillIds - The unique identifiers (V4 UUID) for the skills that should be used to determine the destination for the conversation.
    SkillIds []string `json:"skillIds"`

}

// String returns a JSON representation of the model
func (o *Emailroutingtransferevent) String() string {
    
    
    
    
    
    
    
    
    
    
     o.SkillIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emailroutingtransferevent) MarshalJSON() ([]byte, error) {
    type Alias Emailroutingtransferevent

    if EmailroutingtransfereventMarshalled {
        return []byte("{}"), nil
    }
    EmailroutingtransfereventMarshalled = true

    return json.Marshal(&struct {
        
        EventId string `json:"eventId"`
        
        EventDateTime time.Time `json:"eventDateTime"`
        
        ConversationId string `json:"conversationId"`
        
        TransferType string `json:"transferType"`
        
        CommandId string `json:"commandId"`
        
        InitiatingCommunicationId string `json:"initiatingCommunicationId"`
        
        TargetCommunicationId string `json:"targetCommunicationId"`
        
        ObjectCommunicationId string `json:"objectCommunicationId"`
        
        DestinationQueueId string `json:"destinationQueueId"`
        
        LanguageId string `json:"languageId"`
        
        SkillIds []string `json:"skillIds"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        
        SkillIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

