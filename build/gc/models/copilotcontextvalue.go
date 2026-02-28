package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CopilotcontextvalueMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CopilotcontextvalueDud struct { 
    


    


    


    

}

// Copilotcontextvalue
type Copilotcontextvalue struct { 
    // Name - Name of the context.
    Name string `json:"name"`


    // VarType - Type of the context.
    VarType string `json:"type"`


    // ParticipantDataProperties - Participant data properties.
    ParticipantDataProperties Participantdataproperties `json:"participantDataProperties"`


    // ConversationAttributeProperties - Conversation attribute properties.
    ConversationAttributeProperties Conversationattributeproperties `json:"conversationAttributeProperties"`

}

// String returns a JSON representation of the model
func (o *Copilotcontextvalue) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Copilotcontextvalue) MarshalJSON() ([]byte, error) {
    type Alias Copilotcontextvalue

    if CopilotcontextvalueMarshalled {
        return []byte("{}"), nil
    }
    CopilotcontextvalueMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        ParticipantDataProperties Participantdataproperties `json:"participantDataProperties"`
        
        ConversationAttributeProperties Conversationattributeproperties `json:"conversationAttributeProperties"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

