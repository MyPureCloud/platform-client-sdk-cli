package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowpathselementMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowpathselementDud struct { 
    


    


    


    


    


    


    

}

// Flowpathselement
type Flowpathselement struct { 
    // ParentId - Unique identifier of the parent element. Will be null for the root element.
    ParentId string `json:"parentId"`


    // VarType - Type of the element.
    VarType string `json:"type"`


    // Count - Count of all journeys that include this element.
    Count int `json:"count"`


    // Flows - Details of flows involved in journeys that include this element.
    Flows []Flowpathsflowdetails `json:"flows"`


    // FlowMilestone - The flow milestone, set if the element type is Milestone.
    FlowMilestone Addressableentityref `json:"flowMilestone"`


    // FlowOutcome - The flow outcome, set if the element type is Outcome or Milestone.
    FlowOutcome Addressableentityref `json:"flowOutcome"`


    // FlowOutcomeValue - The value of the flow outcome, if the element type is Outcome.
    FlowOutcomeValue string `json:"flowOutcomeValue"`

}

// String returns a JSON representation of the model
func (o *Flowpathselement) String() string {
    
    
    
     o.Flows = []Flowpathsflowdetails{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowpathselement) MarshalJSON() ([]byte, error) {
    type Alias Flowpathselement

    if FlowpathselementMarshalled {
        return []byte("{}"), nil
    }
    FlowpathselementMarshalled = true

    return json.Marshal(&struct {
        
        ParentId string `json:"parentId"`
        
        VarType string `json:"type"`
        
        Count int `json:"count"`
        
        Flows []Flowpathsflowdetails `json:"flows"`
        
        FlowMilestone Addressableentityref `json:"flowMilestone"`
        
        FlowOutcome Addressableentityref `json:"flowOutcome"`
        
        FlowOutcomeValue string `json:"flowOutcomeValue"`
        *Alias
    }{

        


        


        


        
        Flows: []Flowpathsflowdetails{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

