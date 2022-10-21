package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgesearchclientapplicationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgesearchclientapplicationDud struct { 
    


    


    


    

}

// Knowledgesearchclientapplication
type Knowledgesearchclientapplication struct { 
    // VarType - Application type.
    VarType string `json:"type"`


    // Deployment - Application details when type is MessengerKnowledgeApp or SupportCenter.
    Deployment Addressableentityref `json:"deployment"`


    // BotFlow - Application details when type is BotFlow.
    BotFlow Addressableentityref `json:"botFlow"`


    // Assistant - Application details when type is Assistant.
    Assistant Addressableentityref `json:"assistant"`

}

// String returns a JSON representation of the model
func (o *Knowledgesearchclientapplication) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgesearchclientapplication) MarshalJSON() ([]byte, error) {
    type Alias Knowledgesearchclientapplication

    if KnowledgesearchclientapplicationMarshalled {
        return []byte("{}"), nil
    }
    KnowledgesearchclientapplicationMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Deployment Addressableentityref `json:"deployment"`
        
        BotFlow Addressableentityref `json:"botFlow"`
        
        Assistant Addressableentityref `json:"assistant"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

