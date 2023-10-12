package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CriteriagroupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CriteriagroupDud struct { 
    


    


    


    

}

// Criteriagroup - A group of logical or a singular criteria used to create a query of executionData
type Criteriagroup struct { 
    // And - These criteriaItems will be AND'd together to find a match.
    And []Criteriaitem `json:"and"`


    // Or - These criteriaItems will be OR'd together to find a match.
    Or []Criteriaitem `json:"or"`


    // Not - These criteriaItems must all be false to find a match.
    Not []Criteriaitem `json:"not"`


    // Criteria - A singular critieriaItem to match.
    Criteria Criteriaitem `json:"criteria"`

}

// String returns a JSON representation of the model
func (o *Criteriagroup) String() string {
     o.And = []Criteriaitem{{}} 
     o.Or = []Criteriaitem{{}} 
     o.Not = []Criteriaitem{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Criteriagroup) MarshalJSON() ([]byte, error) {
    type Alias Criteriagroup

    if CriteriagroupMarshalled {
        return []byte("{}"), nil
    }
    CriteriagroupMarshalled = true

    return json.Marshal(&struct {
        
        And []Criteriaitem `json:"and"`
        
        Or []Criteriaitem `json:"or"`
        
        Not []Criteriaitem `json:"not"`
        
        Criteria Criteriaitem `json:"criteria"`
        *Alias
    }{

        
        And: []Criteriaitem{{}},
        


        
        Or: []Criteriaitem{{}},
        


        
        Not: []Criteriaitem{{}},
        


        

        Alias: (*Alias)(u),
    })
}

