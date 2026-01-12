package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowpathsflowdetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowpathsflowdetailsDud struct { 
    


    


    


    

}

// Flowpathsflowdetails
type Flowpathsflowdetails struct { 
    // Flow - The identifier of the flow.
    Flow Addressableentityref `json:"flow"`


    // Version - The version of the flow.
    Version string `json:"version"`


    // VarType - The type of the flow.
    VarType string `json:"type"`


    // Count - Count of all journeys that include this element in the given flow.
    Count int `json:"count"`

}

// String returns a JSON representation of the model
func (o *Flowpathsflowdetails) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowpathsflowdetails) MarshalJSON() ([]byte, error) {
    type Alias Flowpathsflowdetails

    if FlowpathsflowdetailsMarshalled {
        return []byte("{}"), nil
    }
    FlowpathsflowdetailsMarshalled = true

    return json.Marshal(&struct {
        
        Flow Addressableentityref `json:"flow"`
        
        Version string `json:"version"`
        
        VarType string `json:"type"`
        
        Count int `json:"count"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

