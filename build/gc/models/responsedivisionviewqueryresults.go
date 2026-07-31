package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ResponsedivisionviewqueryresultsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ResponsedivisionviewqueryresultsDud struct { 
    

}

// Responsedivisionviewqueryresults - Used to return response division view query results.
type Responsedivisionviewqueryresults struct { 
    // Results - Contains the query results
    Results Domainentitylistingresponsedivisionview `json:"results"`

}

// String returns a JSON representation of the model
func (o *Responsedivisionviewqueryresults) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Responsedivisionviewqueryresults) MarshalJSON() ([]byte, error) {
    type Alias Responsedivisionviewqueryresults

    if ResponsedivisionviewqueryresultsMarshalled {
        return []byte("{}"), nil
    }
    ResponsedivisionviewqueryresultsMarshalled = true

    return json.Marshal(&struct {
        
        Results Domainentitylistingresponsedivisionview `json:"results"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

