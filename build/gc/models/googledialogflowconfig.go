package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GoogledialogflowconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GoogledialogflowconfigDud struct { 
    


    

}

// Googledialogflowconfig
type Googledialogflowconfig struct { 
    // IntegrationId - The integration identifier with which the assistant will fetch transcriptions and knowledge suggestions.
    IntegrationId string `json:"integrationId"`


    // ConversationProfiles - The conversation profiles for which the assistant will fetch transcription and knowledge suggestions.
    ConversationProfiles []Conversationprofile `json:"conversationProfiles"`

}

// String returns a JSON representation of the model
func (o *Googledialogflowconfig) String() string {
    
     o.ConversationProfiles = []Conversationprofile{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Googledialogflowconfig) MarshalJSON() ([]byte, error) {
    type Alias Googledialogflowconfig

    if GoogledialogflowconfigMarshalled {
        return []byte("{}"), nil
    }
    GoogledialogflowconfigMarshalled = true

    return json.Marshal(&struct {
        
        IntegrationId string `json:"integrationId"`
        
        ConversationProfiles []Conversationprofile `json:"conversationProfiles"`
        *Alias
    }{

        


        
        ConversationProfiles: []Conversationprofile{{}},
        

        Alias: (*Alias)(u),
    })
}

