package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QuerycriteriagroupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QuerycriteriagroupDud struct { 
    


    


    


    

}

// Querycriteriagroup - A group of logical items for library queries
type Querycriteriagroup struct { 
    // And - Items that will be AND'd together
    And []Querycriteriaitem `json:"and"`


    // Or - Items that will be OR'd together
    Or []Querycriteriaitem `json:"or"`


    // Not - Items that must all be false
    Not []Querycriteriaitem `json:"not"`


    // Criteria - A single item
    Criteria Querycriteriaitem `json:"criteria"`

}

// String returns a JSON representation of the model
func (o *Querycriteriagroup) String() string {
     o.And = []Querycriteriaitem{{}} 
     o.Or = []Querycriteriaitem{{}} 
     o.Not = []Querycriteriaitem{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Querycriteriagroup) MarshalJSON() ([]byte, error) {
    type Alias Querycriteriagroup

    if QuerycriteriagroupMarshalled {
        return []byte("{}"), nil
    }
    QuerycriteriagroupMarshalled = true

    return json.Marshal(&struct {
        
        And []Querycriteriaitem `json:"and"`
        
        Or []Querycriteriaitem `json:"or"`
        
        Not []Querycriteriaitem `json:"not"`
        
        Criteria Querycriteriaitem `json:"criteria"`
        *Alias
    }{

        
        And: []Querycriteriaitem{{}},
        


        
        Or: []Querycriteriaitem{{}},
        


        
        Not: []Querycriteriaitem{{}},
        


        

        Alias: (*Alias)(u),
    })
}

