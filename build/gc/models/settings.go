package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SettingsDud struct { 
    


    

}

// Settings
type Settings struct { 
    // CommunicationBasedACW - Communication Based ACW
    CommunicationBasedACW bool `json:"communicationBasedACW"`


    // IncludeNonAgentConversationSummary - Display communication summary
    IncludeNonAgentConversationSummary bool `json:"includeNonAgentConversationSummary"`

}

// String returns a JSON representation of the model
func (o *Settings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Settings) MarshalJSON() ([]byte, error) {
    type Alias Settings

    if SettingsMarshalled {
        return []byte("{}"), nil
    }
    SettingsMarshalled = true

    return json.Marshal(&struct {
        
        CommunicationBasedACW bool `json:"communicationBasedACW"`
        
        IncludeNonAgentConversationSummary bool `json:"includeNonAgentConversationSummary"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

