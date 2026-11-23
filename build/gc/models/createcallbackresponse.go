package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatecallbackresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatecallbackresponseDud struct { 
    


    


    

}

// Createcallbackresponse
type Createcallbackresponse struct { 
    // Conversation - The conversation associated with the callback
    Conversation Domainentityref `json:"conversation"`


    // CallbackIdentifiers - The list of communication identifiers for the callback participants
    CallbackIdentifiers []Callbackidentifier `json:"callbackIdentifiers"`


    // ScheduleEvaluationResult - The outcome of the operating schedule group evaluation. Present only when schedule group validation is enabled for the organization
    ScheduleEvaluationResult string `json:"scheduleEvaluationResult"`

}

// String returns a JSON representation of the model
func (o *Createcallbackresponse) String() string {
    
     o.CallbackIdentifiers = []Callbackidentifier{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createcallbackresponse) MarshalJSON() ([]byte, error) {
    type Alias Createcallbackresponse

    if CreatecallbackresponseMarshalled {
        return []byte("{}"), nil
    }
    CreatecallbackresponseMarshalled = true

    return json.Marshal(&struct {
        
        Conversation Domainentityref `json:"conversation"`
        
        CallbackIdentifiers []Callbackidentifier `json:"callbackIdentifiers"`
        
        ScheduleEvaluationResult string `json:"scheduleEvaluationResult"`
        *Alias
    }{

        


        
        CallbackIdentifiers: []Callbackidentifier{{}},
        


        

        Alias: (*Alias)(u),
    })
}

