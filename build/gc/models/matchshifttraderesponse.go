package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MatchshifttraderesponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MatchshifttraderesponseDud struct { 
    


    


    


    

}

// Matchshifttraderesponse
type Matchshifttraderesponse struct { 
    // Trade - The associated shift trade
    Trade Shifttraderesponse `json:"trade"`


    // Violations - Constraint violations which disallow this shift trade
    Violations []Shifttradematchviolation `json:"violations"`


    // AdminReviewViolations - Constraint violations for this shift trade which require shift trade administrator review
    AdminReviewViolations []Shifttradematchviolation `json:"adminReviewViolations"`


    // UnevaluatedRules - Unevaluated rules for this shift trade which require shift trade administrator review
    UnevaluatedRules []string `json:"unevaluatedRules"`

}

// String returns a JSON representation of the model
func (o *Matchshifttraderesponse) String() string {
    
     o.Violations = []Shifttradematchviolation{{}} 
     o.AdminReviewViolations = []Shifttradematchviolation{{}} 
     o.UnevaluatedRules = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Matchshifttraderesponse) MarshalJSON() ([]byte, error) {
    type Alias Matchshifttraderesponse

    if MatchshifttraderesponseMarshalled {
        return []byte("{}"), nil
    }
    MatchshifttraderesponseMarshalled = true

    return json.Marshal(&struct {
        
        Trade Shifttraderesponse `json:"trade"`
        
        Violations []Shifttradematchviolation `json:"violations"`
        
        AdminReviewViolations []Shifttradematchviolation `json:"adminReviewViolations"`
        
        UnevaluatedRules []string `json:"unevaluatedRules"`
        *Alias
    }{

        


        
        Violations: []Shifttradematchviolation{{}},
        


        
        AdminReviewViolations: []Shifttradematchviolation{{}},
        


        
        UnevaluatedRules: []string{""},
        

        Alias: (*Alias)(u),
    })
}

