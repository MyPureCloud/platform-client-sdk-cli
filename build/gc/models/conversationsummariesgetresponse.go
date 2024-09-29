package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationsummariesgetresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationsummariesgetresponseDud struct { 
    


    


    

}

// Conversationsummariesgetresponse - Defines the summaries of a conversation.
type Conversationsummariesgetresponse struct { 
    // Conversation - The conversation object.
    Conversation Addressableentityref `json:"conversation"`


    // Summary - The summary of the conversation.
    Summary Conversationsummary `json:"summary"`


    // SessionSummaries - All the summaries of the session.
    SessionSummaries []Conversationsessionsummary `json:"sessionSummaries"`

}

// String returns a JSON representation of the model
func (o *Conversationsummariesgetresponse) String() string {
    
    
     o.SessionSummaries = []Conversationsessionsummary{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationsummariesgetresponse) MarshalJSON() ([]byte, error) {
    type Alias Conversationsummariesgetresponse

    if ConversationsummariesgetresponseMarshalled {
        return []byte("{}"), nil
    }
    ConversationsummariesgetresponseMarshalled = true

    return json.Marshal(&struct {
        
        Conversation Addressableentityref `json:"conversation"`
        
        Summary Conversationsummary `json:"summary"`
        
        SessionSummaries []Conversationsessionsummary `json:"sessionSummaries"`
        *Alias
    }{

        


        


        
        SessionSummaries: []Conversationsessionsummary{{}},
        

        Alias: (*Alias)(u),
    })
}

