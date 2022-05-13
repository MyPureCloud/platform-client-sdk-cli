package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DevelopmentactivityaggregatequeryrequestfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DevelopmentactivityaggregatequeryrequestfilterDud struct { 
    


    

}

// Developmentactivityaggregatequeryrequestfilter
type Developmentactivityaggregatequeryrequestfilter struct { 
    // VarType - The logic used to combine the clauses
    VarType string `json:"type"`


    // Clauses - The list of clauses used to filter the data. Note that clauses must filter by attendeeId and a maximum of 100 user IDs are allowed
    Clauses []Developmentactivityaggregatequeryrequestclause `json:"clauses"`

}

// String returns a JSON representation of the model
func (o *Developmentactivityaggregatequeryrequestfilter) String() string {
    
     o.Clauses = []Developmentactivityaggregatequeryrequestclause{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Developmentactivityaggregatequeryrequestfilter) MarshalJSON() ([]byte, error) {
    type Alias Developmentactivityaggregatequeryrequestfilter

    if DevelopmentactivityaggregatequeryrequestfilterMarshalled {
        return []byte("{}"), nil
    }
    DevelopmentactivityaggregatequeryrequestfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Developmentactivityaggregatequeryrequestclause `json:"clauses"`
        *Alias
    }{

        


        
        Clauses: []Developmentactivityaggregatequeryrequestclause{{}},
        

        Alias: (*Alias)(u),
    })
}

