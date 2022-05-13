package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryresultsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryresultsDud struct { 
    


    

}

// Queryresults
type Queryresults struct { 
    // Results
    Results Domainentitylistingqueryresult `json:"results"`


    // FacetInfo
    FacetInfo Queryfacetinfo `json:"facetInfo"`

}

// String returns a JSON representation of the model
func (o *Queryresults) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryresults) MarshalJSON() ([]byte, error) {
    type Alias Queryresults

    if QueryresultsMarshalled {
        return []byte("{}"), nil
    }
    QueryresultsMarshalled = true

    return json.Marshal(&struct {
        
        Results Domainentitylistingqueryresult `json:"results"`
        
        FacetInfo Queryfacetinfo `json:"facetInfo"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

