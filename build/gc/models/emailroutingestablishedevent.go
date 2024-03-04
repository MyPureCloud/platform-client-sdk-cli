package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmailroutingestablishedeventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmailroutingestablishedeventDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Emailroutingestablishedevent
type Emailroutingestablishedevent struct { 
    // EventId - A unique (V4 UUID) eventId for this event
    EventId string `json:"eventId"`


    // EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EventDateTime time.Time `json:"eventDateTime"`


    // ConversationId - A unique Id (V4 UUID) identifying this conversation
    ConversationId string `json:"conversationId"`


    // CommunicationId - A unique Id (V4 UUID) identifying this communication.
    CommunicationId string `json:"communicationId"`


    // QueueId - The id of the queue that is routing this conversation.
    QueueId string `json:"queueId"`


    // SkillIds - The unique identifiers for the skills that should be used to determine the destination for the conversation.
    SkillIds []string `json:"skillIds"`


    // LanguageId - The unique identifier for the language that should be used to determine the destination for the conversation.
    LanguageId string `json:"languageId"`


    // Label - An optional label that categorizes the conversation. Max-utilization settings can be configured at a per-label level.
    Label string `json:"label"`


    // InitialConfiguration - Metadata about this communication.
    InitialConfiguration Emailinitialconfiguration `json:"initialConfiguration"`


    // SourceConfiguration - Metadata about the source of this communication's interaction.
    SourceConfiguration Sourceconfiguration `json:"sourceConfiguration"`

}

// String returns a JSON representation of the model
func (o *Emailroutingestablishedevent) String() string {
    
    
    
    
    
     o.SkillIds = []string{""} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emailroutingestablishedevent) MarshalJSON() ([]byte, error) {
    type Alias Emailroutingestablishedevent

    if EmailroutingestablishedeventMarshalled {
        return []byte("{}"), nil
    }
    EmailroutingestablishedeventMarshalled = true

    return json.Marshal(&struct {
        
        EventId string `json:"eventId"`
        
        EventDateTime time.Time `json:"eventDateTime"`
        
        ConversationId string `json:"conversationId"`
        
        CommunicationId string `json:"communicationId"`
        
        QueueId string `json:"queueId"`
        
        SkillIds []string `json:"skillIds"`
        
        LanguageId string `json:"languageId"`
        
        Label string `json:"label"`
        
        InitialConfiguration Emailinitialconfiguration `json:"initialConfiguration"`
        
        SourceConfiguration Sourceconfiguration `json:"sourceConfiguration"`
        *Alias
    }{

        


        


        


        


        


        
        SkillIds: []string{""},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

