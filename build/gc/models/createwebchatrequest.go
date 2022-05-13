package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatewebchatrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatewebchatrequestDud struct { 
    


    


    


    


    


    


    

}

// Createwebchatrequest
type Createwebchatrequest struct { 
    // QueueId - The ID of the queue to use for routing the chat conversation.
    QueueId string `json:"queueId"`


    // Provider - The name of the provider that is sourcing the web chat.
    Provider string `json:"provider"`


    // SkillIds - The list of skill ID's to use for routing.
    SkillIds []string `json:"skillIds"`


    // LanguageId - The ID of the langauge to use for routing.
    LanguageId string `json:"languageId"`


    // Priority - The priority to assign to the conversation for routing.
    Priority int `json:"priority"`


    // Attributes - The list of attributes to associate with the customer participant.
    Attributes map[string]string `json:"attributes"`


    // CustomerName - The name of the customer participating in the web chat.
    CustomerName string `json:"customerName"`

}

// String returns a JSON representation of the model
func (o *Createwebchatrequest) String() string {
    
    
     o.SkillIds = []string{""} 
    
    
     o.Attributes = map[string]string{"": ""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createwebchatrequest) MarshalJSON() ([]byte, error) {
    type Alias Createwebchatrequest

    if CreatewebchatrequestMarshalled {
        return []byte("{}"), nil
    }
    CreatewebchatrequestMarshalled = true

    return json.Marshal(&struct {
        
        QueueId string `json:"queueId"`
        
        Provider string `json:"provider"`
        
        SkillIds []string `json:"skillIds"`
        
        LanguageId string `json:"languageId"`
        
        Priority int `json:"priority"`
        
        Attributes map[string]string `json:"attributes"`
        
        CustomerName string `json:"customerName"`
        *Alias
    }{

        


        


        
        SkillIds: []string{""},
        


        


        


        
        Attributes: map[string]string{"": ""},
        


        

        Alias: (*Alias)(u),
    })
}

