package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpportunitiesresultwithpaginationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpportunitiesresultwithpaginationDud struct { 
    


    

}

// Opportunitiesresultwithpagination
type Opportunitiesresultwithpagination struct { 
    // NextStartDate - The start date to use for the next query to retrieve additional results in ISO-8601 format. Null if there are no more results
    NextStartDate time.Time `json:"nextStartDate"`


    // Opportunities - The list of opportunities
    Opportunities []Opportunityresult `json:"opportunities"`

}

// String returns a JSON representation of the model
func (o *Opportunitiesresultwithpagination) String() string {
    
     o.Opportunities = []Opportunityresult{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Opportunitiesresultwithpagination) MarshalJSON() ([]byte, error) {
    type Alias Opportunitiesresultwithpagination

    if OpportunitiesresultwithpaginationMarshalled {
        return []byte("{}"), nil
    }
    OpportunitiesresultwithpaginationMarshalled = true

    return json.Marshal(&struct {
        
        NextStartDate time.Time `json:"nextStartDate"`
        
        Opportunities []Opportunityresult `json:"opportunities"`
        *Alias
    }{

        


        
        Opportunities: []Opportunityresult{{}},
        

        Alias: (*Alias)(u),
    })
}

