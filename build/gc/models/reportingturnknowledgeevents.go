package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReportingturnknowledgeeventsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReportingturnknowledgeeventsDud struct { 
    


    

}

// Reportingturnknowledgeevents
type Reportingturnknowledgeevents struct { 
    // Search - The knowledge search data captured during this reporting turn.
    Search []Reportingturnknowledgesearchevent `json:"search"`


    // Feedback - The knowledge feedback data captured during this reporting turn.
    Feedback []Reportingturnknowledgefeedbackevent `json:"feedback"`

}

// String returns a JSON representation of the model
func (o *Reportingturnknowledgeevents) String() string {
     o.Search = []Reportingturnknowledgesearchevent{{}} 
     o.Feedback = []Reportingturnknowledgefeedbackevent{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reportingturnknowledgeevents) MarshalJSON() ([]byte, error) {
    type Alias Reportingturnknowledgeevents

    if ReportingturnknowledgeeventsMarshalled {
        return []byte("{}"), nil
    }
    ReportingturnknowledgeeventsMarshalled = true

    return json.Marshal(&struct {
        
        Search []Reportingturnknowledgesearchevent `json:"search"`
        
        Feedback []Reportingturnknowledgefeedbackevent `json:"feedback"`
        *Alias
    }{

        
        Search: []Reportingturnknowledgesearchevent{{}},
        


        
        Feedback: []Reportingturnknowledgefeedbackevent{{}},
        

        Alias: (*Alias)(u),
    })
}

