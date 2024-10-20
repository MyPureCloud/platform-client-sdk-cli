package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemqueryjobqueryfiltersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemqueryjobqueryfiltersDud struct { 
    


    


    

}

// Workitemqueryjobqueryfilters
type Workitemqueryjobqueryfilters struct { 
    // Name - Name of the attribute to filter.
    Name string `json:"name"`


    // Operator - Query filter logical operator to join criteria.
    Operator string `json:"operator"`


    // Criteria - Query filter criteria.
    Criteria []Workitemqueryjobqueryfilterscriteria `json:"criteria"`

}

// String returns a JSON representation of the model
func (o *Workitemqueryjobqueryfilters) String() string {
    
    
     o.Criteria = []Workitemqueryjobqueryfilterscriteria{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemqueryjobqueryfilters) MarshalJSON() ([]byte, error) {
    type Alias Workitemqueryjobqueryfilters

    if WorkitemqueryjobqueryfiltersMarshalled {
        return []byte("{}"), nil
    }
    WorkitemqueryjobqueryfiltersMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Operator string `json:"operator"`
        
        Criteria []Workitemqueryjobqueryfilterscriteria `json:"criteria"`
        *Alias
    }{

        


        


        
        Criteria: []Workitemqueryjobqueryfilterscriteria{{}},
        

        Alias: (*Alias)(u),
    })
}

