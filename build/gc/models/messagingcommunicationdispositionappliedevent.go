package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagingcommunicationdispositionappliedeventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagingcommunicationdispositionappliedeventDud struct { 
    


    


    


    


    


    


    

}

// Messagingcommunicationdispositionappliedevent
type Messagingcommunicationdispositionappliedevent struct { 
    // EventId - A unique (V4 UUID) eventId for this event
    EventId string `json:"eventId"`


    // EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EventDateTime time.Time `json:"eventDateTime"`


    // ConversationId - A unique Id (V4 UUID) identifying this conversation
    ConversationId string `json:"conversationId"`


    // CommunicationId - A unique Id (V4 UUID) identifying this communication
    CommunicationId string `json:"communicationId"`


    // Code - The wrapup-code (V4 UUID) used to disposition this interaction. If this value is not provided the disposition is considered skipped.
    Code string `json:"code"`


    // Notes - Text entered by the agent to describe the interaction or disposition. Ignored if the disposition is considered skipped.
    Notes string `json:"notes"`


    // Tags - The list of tags selected by the agent to describe the interaction or disposition. Ignored if the disposition is considered skipped.
    Tags []string `json:"tags"`

}

// String returns a JSON representation of the model
func (o *Messagingcommunicationdispositionappliedevent) String() string {
    
    
    
    
    
    
     o.Tags = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagingcommunicationdispositionappliedevent) MarshalJSON() ([]byte, error) {
    type Alias Messagingcommunicationdispositionappliedevent

    if MessagingcommunicationdispositionappliedeventMarshalled {
        return []byte("{}"), nil
    }
    MessagingcommunicationdispositionappliedeventMarshalled = true

    return json.Marshal(&struct {
        
        EventId string `json:"eventId"`
        
        EventDateTime time.Time `json:"eventDateTime"`
        
        ConversationId string `json:"conversationId"`
        
        CommunicationId string `json:"communicationId"`
        
        Code string `json:"code"`
        
        Notes string `json:"notes"`
        
        Tags []string `json:"tags"`
        *Alias
    }{

        


        


        


        


        


        


        
        Tags: []string{""},
        

        Alias: (*Alias)(u),
    })
}

