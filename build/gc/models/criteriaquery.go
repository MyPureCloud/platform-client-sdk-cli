package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CriteriaqueryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CriteriaqueryDud struct { 
    

}

// Criteriaquery - Used to retrieve executionData based upon certain criteria
type Criteriaquery struct { 
    // Query - A list of CriteriaGroups which will be AND'd together to generate a result set.
    Query []Criteriagroup `json:"query"`

}

// String returns a JSON representation of the model
func (o *Criteriaquery) String() string {
     o.Query = []Criteriagroup{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Criteriaquery) MarshalJSON() ([]byte, error) {
    type Alias Criteriaquery

    if CriteriaqueryMarshalled {
        return []byte("{}"), nil
    }
    CriteriaqueryMarshalled = true

    return json.Marshal(&struct {
        
        Query []Criteriagroup `json:"query"`
        *Alias
    }{

        
        Query: []Criteriagroup{{}},
        

        Alias: (*Alias)(u),
    })
}

