package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CaseassociationqueryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CaseassociationqueryDud struct { 
    


    


    


    

}

// Caseassociationquery
type Caseassociationquery struct { 
    // PageSize - The number of entities to return in the response.
    PageSize int `json:"pageSize"`


    // After - The cursor that points to the end of the set of entities that has been returned.
    After string `json:"after"`


    // WorkitemId - The Workitem ID to query by.
    WorkitemId string `json:"workitemId"`


    // ConversationId - The conversation ID to query by.
    ConversationId string `json:"conversationId"`

}

// String returns a JSON representation of the model
func (o *Caseassociationquery) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Caseassociationquery) MarshalJSON() ([]byte, error) {
    type Alias Caseassociationquery

    if CaseassociationqueryMarshalled {
        return []byte("{}"), nil
    }
    CaseassociationqueryMarshalled = true

    return json.Marshal(&struct {
        
        PageSize int `json:"pageSize"`
        
        After string `json:"after"`
        
        WorkitemId string `json:"workitemId"`
        
        ConversationId string `json:"conversationId"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

