package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TextbotflowoutcomeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TextbotflowoutcomeDud struct { 
    


    


    


    


    

}

// Textbotflowoutcome - Flow Outcome data related to a bot flow which is exiting gracefully.
type Textbotflowoutcome struct { 
    // OutcomeId - The Flow Outcome ID.
    OutcomeId string `json:"outcomeId"`


    // OutcomeValue - The value of the FlowOutcome.
    OutcomeValue string `json:"outcomeValue"`


    // DateStart - The timestamp for when the Flow Outcome began. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStart time.Time `json:"dateStart"`


    // DateEnd - The timestamp for when the Flow Outcome finished. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateEnd time.Time `json:"dateEnd"`


    // Milestones - The Flow Milestones for the Flow Outcome.
    Milestones []Textbotflowmilestone `json:"milestones"`

}

// String returns a JSON representation of the model
func (o *Textbotflowoutcome) String() string {
    
    
    
    
     o.Milestones = []Textbotflowmilestone{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Textbotflowoutcome) MarshalJSON() ([]byte, error) {
    type Alias Textbotflowoutcome

    if TextbotflowoutcomeMarshalled {
        return []byte("{}"), nil
    }
    TextbotflowoutcomeMarshalled = true

    return json.Marshal(&struct {
        
        OutcomeId string `json:"outcomeId"`
        
        OutcomeValue string `json:"outcomeValue"`
        
        DateStart time.Time `json:"dateStart"`
        
        DateEnd time.Time `json:"dateEnd"`
        
        Milestones []Textbotflowmilestone `json:"milestones"`
        *Alias
    }{

        


        


        


        


        
        Milestones: []Textbotflowmilestone{{}},
        

        Alias: (*Alias)(u),
    })
}

