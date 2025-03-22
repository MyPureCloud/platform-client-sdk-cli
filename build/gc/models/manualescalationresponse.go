package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ManualescalationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ManualescalationresponseDud struct { 
    


    


    


    

}

// Manualescalationresponse
type Manualescalationresponse struct { 
    // SocialMediaNormalizedMessageId - The Id of the message that got escalated.
    SocialMediaNormalizedMessageId string `json:"socialMediaNormalizedMessageId"`


    // ConversationNormalizedMessageId - The ID of the message in the conversation.
    ConversationNormalizedMessageId string `json:"conversationNormalizedMessageId"`


    // EscalationTarget - The target integration configuration used for an social media message.
    EscalationTarget Escalationtarget `json:"escalationTarget"`


    // EscalationStatus - Escalation Status of the message.
    EscalationStatus string `json:"escalationStatus"`

}

// String returns a JSON representation of the model
func (o *Manualescalationresponse) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Manualescalationresponse) MarshalJSON() ([]byte, error) {
    type Alias Manualescalationresponse

    if ManualescalationresponseMarshalled {
        return []byte("{}"), nil
    }
    ManualescalationresponseMarshalled = true

    return json.Marshal(&struct {
        
        SocialMediaNormalizedMessageId string `json:"socialMediaNormalizedMessageId"`
        
        ConversationNormalizedMessageId string `json:"conversationNormalizedMessageId"`
        
        EscalationTarget Escalationtarget `json:"escalationTarget"`
        
        EscalationStatus string `json:"escalationStatus"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

