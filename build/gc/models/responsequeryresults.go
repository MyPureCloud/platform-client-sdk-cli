package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ResponsequeryresultsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ResponsequeryresultsDud struct { 
    

}

// Responsequeryresults - Used to return response query results
type Responsequeryresults struct { 
    // Results - Contains the query results
    Results Responseentitylist `json:"results"`

}

// String returns a JSON representation of the model
func (o *Responsequeryresults) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Responsequeryresults) MarshalJSON() ([]byte, error) {
    type Alias Responsequeryresults

    if ResponsequeryresultsMarshalled {
        return []byte("{}"), nil
    }
    ResponsequeryresultsMarshalled = true

    return json.Marshal(&struct {
        
        Results Responseentitylist `json:"results"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

