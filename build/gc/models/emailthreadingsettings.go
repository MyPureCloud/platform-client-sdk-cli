package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmailthreadingsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmailthreadingsettingsDud struct { 
    


    

}

// Emailthreadingsettings
type Emailthreadingsettings struct { 
    // StartNewConversationOnSubjectChange - This setting controls whether a new conversation is started if the subject of an inbound email is different from the subject of the current conversation. RE: and FWD: prefixes in any language are ignored.
    StartNewConversationOnSubjectChange bool `json:"startNewConversationOnSubjectChange"`


    // TimeoutInMinutes - In minutes, how long an email conversation should keep being threaded after being disconnected.
    TimeoutInMinutes int `json:"timeoutInMinutes"`

}

// String returns a JSON representation of the model
func (o *Emailthreadingsettings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emailthreadingsettings) MarshalJSON() ([]byte, error) {
    type Alias Emailthreadingsettings

    if EmailthreadingsettingsMarshalled {
        return []byte("{}"), nil
    }
    EmailthreadingsettingsMarshalled = true

    return json.Marshal(&struct {
        
        StartNewConversationOnSubjectChange bool `json:"startNewConversationOnSubjectChange"`
        
        TimeoutInMinutes int `json:"timeoutInMinutes"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

