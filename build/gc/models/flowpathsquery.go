package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowpathsqueryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowpathsqueryDud struct { 
    


    

}

// Flowpathsquery
type Flowpathsquery struct { 
    // Category - Category (use case) of the paths within a given domain.
    Category string `json:"category"`


    // Flows - List of flows to query the paths result.
    Flows []Flowpathsflowfilter `json:"flows"`

}

// String returns a JSON representation of the model
func (o *Flowpathsquery) String() string {
    
     o.Flows = []Flowpathsflowfilter{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowpathsquery) MarshalJSON() ([]byte, error) {
    type Alias Flowpathsquery

    if FlowpathsqueryMarshalled {
        return []byte("{}"), nil
    }
    FlowpathsqueryMarshalled = true

    return json.Marshal(&struct {
        
        Category string `json:"category"`
        
        Flows []Flowpathsflowfilter `json:"flows"`
        *Alias
    }{

        


        
        Flows: []Flowpathsflowfilter{{}},
        

        Alias: (*Alias)(u),
    })
}

