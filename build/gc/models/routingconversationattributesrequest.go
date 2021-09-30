package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RoutingconversationattributesrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RoutingconversationattributesrequestDud struct { 
    


    


    

}

// Routingconversationattributesrequest
type Routingconversationattributesrequest struct { 
    // Priority - Priority to be updated on in-queue conversation. Range:[-25000000, 25000000]
    Priority int `json:"priority"`


    // SkillIds - Skills to be updated on in-queue conversation.
    SkillIds []string `json:"skillIds"`


    // LanguageId - Language required on the in-queue conversation.
    LanguageId string `json:"languageId"`

}

// String returns a JSON representation of the model
func (o *Routingconversationattributesrequest) String() string {
    
    
    
    
    
    
     o.SkillIds = []string{""} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Routingconversationattributesrequest) MarshalJSON() ([]byte, error) {
    type Alias Routingconversationattributesrequest

    if RoutingconversationattributesrequestMarshalled {
        return []byte("{}"), nil
    }
    RoutingconversationattributesrequestMarshalled = true

    return json.Marshal(&struct { 
        Priority int `json:"priority"`
        
        SkillIds []string `json:"skillIds"`
        
        LanguageId string `json:"languageId"`
        
        *Alias
    }{
        

        

        

        
        SkillIds: []string{""},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

