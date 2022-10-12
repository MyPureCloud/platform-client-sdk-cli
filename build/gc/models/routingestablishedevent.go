package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RoutingestablishedeventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RoutingestablishedeventDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Routingestablishedevent
type Routingestablishedevent struct { 
    // EventId - A unique (V4 UUID) eventId for this event
    EventId string `json:"eventId"`


    // EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EventDateTime time.Time `json:"eventDateTime"`


    // ConversationId - A unique Id (V4 UUID) identifying this conversation
    ConversationId string `json:"conversationId"`


    // CommunicationId - A unique Id (V4 UUID) identifying this communication
    CommunicationId string `json:"communicationId"`


    // PhoneNumber - Identifies the phone number used to reach this queue if it is different from the information that would be accessed by queueId.
    PhoneNumber string `json:"phoneNumber"`


    // QueueId - The id (V4 UUID) of the queue that is routing this conversation.
    QueueId string `json:"queueId"`


    // SkillIds - The unique identifiers (V4 UUID) for the skills that should be used to determine the destination for the conversation.
    SkillIds []string `json:"skillIds"`


    // LanguageId - The unique identifier (V4 UUID) for the language that should be used to determine the destination for the conversation.
    LanguageId string `json:"languageId"`


    // InitialConfiguration - Metadata about this communication.
    InitialConfiguration Initialconfiguration `json:"initialConfiguration"`


    // SourceConfiguration - Metadata about the source of this communication's interaction.
    SourceConfiguration Sourceconfiguration `json:"sourceConfiguration"`

}

// String returns a JSON representation of the model
func (o *Routingestablishedevent) String() string {
    
    
    
    
    
    
     o.SkillIds = []string{""} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Routingestablishedevent) MarshalJSON() ([]byte, error) {
    type Alias Routingestablishedevent

    if RoutingestablishedeventMarshalled {
        return []byte("{}"), nil
    }
    RoutingestablishedeventMarshalled = true

    return json.Marshal(&struct {
        
        EventId string `json:"eventId"`
        
        EventDateTime time.Time `json:"eventDateTime"`
        
        ConversationId string `json:"conversationId"`
        
        CommunicationId string `json:"communicationId"`
        
        PhoneNumber string `json:"phoneNumber"`
        
        QueueId string `json:"queueId"`
        
        SkillIds []string `json:"skillIds"`
        
        LanguageId string `json:"languageId"`
        
        InitialConfiguration Initialconfiguration `json:"initialConfiguration"`
        
        SourceConfiguration Sourceconfiguration `json:"sourceConfiguration"`
        *Alias
    }{

        


        


        


        


        


        


        
        SkillIds: []string{""},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

