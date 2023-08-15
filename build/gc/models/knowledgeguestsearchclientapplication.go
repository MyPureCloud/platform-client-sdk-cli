package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeguestsearchclientapplicationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeguestsearchclientapplicationDud struct { 
    


    


    


    

}

// Knowledgeguestsearchclientapplication
type Knowledgeguestsearchclientapplication struct { 
    // VarType - Application type.
    VarType string `json:"type"`


    // Deployment - Application details when type is MessengerKnowledgeApp or SupportCenter.
    Deployment Entityreference `json:"deployment"`


    // BotFlow - Application details when type is BotFlow.
    BotFlow Entityreference `json:"botFlow"`


    // Assistant - Application details when type is Assistant.
    Assistant Entityreference `json:"assistant"`

}

// String returns a JSON representation of the model
func (o *Knowledgeguestsearchclientapplication) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeguestsearchclientapplication) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeguestsearchclientapplication

    if KnowledgeguestsearchclientapplicationMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeguestsearchclientapplicationMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Deployment Entityreference `json:"deployment"`
        
        BotFlow Entityreference `json:"botFlow"`
        
        Assistant Entityreference `json:"assistant"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

