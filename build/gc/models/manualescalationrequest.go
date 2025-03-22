package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ManualescalationrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ManualescalationrequestDud struct { 
    


    

}

// Manualescalationrequest
type Manualescalationrequest struct { 
    // SocialMediaNormalizedMessageId - The social media normalized message ID to be escalated.
    SocialMediaNormalizedMessageId string `json:"socialMediaNormalizedMessageId"`


    // EscalationTarget - The escalation target to which the specified social media normalized message should be routed.
    EscalationTarget Escalationtarget `json:"escalationTarget"`

}

// String returns a JSON representation of the model
func (o *Manualescalationrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Manualescalationrequest) MarshalJSON() ([]byte, error) {
    type Alias Manualescalationrequest

    if ManualescalationrequestMarshalled {
        return []byte("{}"), nil
    }
    ManualescalationrequestMarshalled = true

    return json.Marshal(&struct {
        
        SocialMediaNormalizedMessageId string `json:"socialMediaNormalizedMessageId"`
        
        EscalationTarget Escalationtarget `json:"escalationTarget"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

