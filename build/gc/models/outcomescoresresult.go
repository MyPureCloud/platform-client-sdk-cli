package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutcomescoresresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutcomescoresresultDud struct { 
    


    ModifiedDate time.Time `json:"modifiedDate"`

}

// Outcomescoresresult
type Outcomescoresresult struct { 
    // OutcomeScores - List of scored outcomes in the session.
    OutcomeScores []Outcomescore `json:"outcomeScores"`


    

}

// String returns a JSON representation of the model
func (o *Outcomescoresresult) String() string {
     o.OutcomeScores = []Outcomescore{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outcomescoresresult) MarshalJSON() ([]byte, error) {
    type Alias Outcomescoresresult

    if OutcomescoresresultMarshalled {
        return []byte("{}"), nil
    }
    OutcomescoresresultMarshalled = true

    return json.Marshal(&struct {
        
        OutcomeScores []Outcomescore `json:"outcomeScores"`
        *Alias
    }{

        
        OutcomeScores: []Outcomescore{{}},
        


        

        Alias: (*Alias)(u),
    })
}

