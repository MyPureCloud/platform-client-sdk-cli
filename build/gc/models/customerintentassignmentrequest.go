package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CustomerintentassignmentrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CustomerintentassignmentrequestDud struct { 
    


    


    


    

}

// Customerintentassignmentrequest
type Customerintentassignmentrequest struct { 
    // SourceId - Id of the source of assignment
    SourceId string `json:"sourceId"`


    // SessionId - Id of session assignment occurred in
    SessionId string `json:"sessionId"`


    // ConversationId - Id of conversation assignment occurred in
    ConversationId string `json:"conversationId"`


    // SourceType - Type of source of assignment
    SourceType string `json:"sourceType"`

}

// String returns a JSON representation of the model
func (o *Customerintentassignmentrequest) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Customerintentassignmentrequest) MarshalJSON() ([]byte, error) {
    type Alias Customerintentassignmentrequest

    if CustomerintentassignmentrequestMarshalled {
        return []byte("{}"), nil
    }
    CustomerintentassignmentrequestMarshalled = true

    return json.Marshal(&struct {
        
        SourceId string `json:"sourceId"`
        
        SessionId string `json:"sessionId"`
        
        ConversationId string `json:"conversationId"`
        
        SourceType string `json:"sourceType"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

