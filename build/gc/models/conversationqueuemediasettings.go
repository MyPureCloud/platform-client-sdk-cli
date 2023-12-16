package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationqueuemediasettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationqueuemediasettingsDud struct { 
    


    


    


    

}

// Conversationqueuemediasettings
type Conversationqueuemediasettings struct { 
    // AlertingTimeoutSeconds - Specifies how long the agent has to answer an interaction before being marked as not responding.
    AlertingTimeoutSeconds int `json:"alertingTimeoutSeconds"`


    // AutoAnswerAlertToneSeconds - Specifies the duration of the alerting sound to be played for auto answered interactions.
    AutoAnswerAlertToneSeconds float64 `json:"autoAnswerAlertToneSeconds"`


    // ManualAnswerAlertToneSeconds - Specifies the duration of the alerting sound to be played for manually answered interactions.
    ManualAnswerAlertToneSeconds float64 `json:"manualAnswerAlertToneSeconds"`


    // EnableAutoAnswer - Flag to indicate if auto answer is enabled for the given media type or media subtype.
    EnableAutoAnswer bool `json:"enableAutoAnswer"`

}

// String returns a JSON representation of the model
func (o *Conversationqueuemediasettings) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationqueuemediasettings) MarshalJSON() ([]byte, error) {
    type Alias Conversationqueuemediasettings

    if ConversationqueuemediasettingsMarshalled {
        return []byte("{}"), nil
    }
    ConversationqueuemediasettingsMarshalled = true

    return json.Marshal(&struct {
        
        AlertingTimeoutSeconds int `json:"alertingTimeoutSeconds"`
        
        AutoAnswerAlertToneSeconds float64 `json:"autoAnswerAlertToneSeconds"`
        
        ManualAnswerAlertToneSeconds float64 `json:"manualAnswerAlertToneSeconds"`
        
        EnableAutoAnswer bool `json:"enableAutoAnswer"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

