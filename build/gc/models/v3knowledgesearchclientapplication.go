package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3knowledgesearchclientapplicationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3knowledgesearchclientapplicationDud struct { 
    


    


    

}

// V3knowledgesearchclientapplication
type V3knowledgesearchclientapplication struct { 
    // VarType - Application type.
    VarType string `json:"type"`


    // BotFlow - Application details when type is BotFlow.
    BotFlow Addressableentityref `json:"botFlow"`


    // Assistant - Application details when type is Assistant.
    Assistant Addressableentityref `json:"assistant"`

}

// String returns a JSON representation of the model
func (o *V3knowledgesearchclientapplication) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3knowledgesearchclientapplication) MarshalJSON() ([]byte, error) {
    type Alias V3knowledgesearchclientapplication

    if V3knowledgesearchclientapplicationMarshalled {
        return []byte("{}"), nil
    }
    V3knowledgesearchclientapplicationMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        BotFlow Addressableentityref `json:"botFlow"`
        
        Assistant Addressableentityref `json:"assistant"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

