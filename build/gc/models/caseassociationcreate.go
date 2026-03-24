package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CaseassociationcreateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CaseassociationcreateDud struct { 
    


    

}

// Caseassociationcreate
type Caseassociationcreate struct { 
    // WorkitemId - The ID of the workitem to associate with the case.
    WorkitemId string `json:"workitemId"`


    // ConversationId - The ID of the conversation to associate with the case.
    ConversationId string `json:"conversationId"`

}

// String returns a JSON representation of the model
func (o *Caseassociationcreate) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Caseassociationcreate) MarshalJSON() ([]byte, error) {
    type Alias Caseassociationcreate

    if CaseassociationcreateMarshalled {
        return []byte("{}"), nil
    }
    CaseassociationcreateMarshalled = true

    return json.Marshal(&struct {
        
        WorkitemId string `json:"workitemId"`
        
        ConversationId string `json:"conversationId"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

