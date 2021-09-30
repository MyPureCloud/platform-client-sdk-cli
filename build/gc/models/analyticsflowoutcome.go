package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticsflowoutcomeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticsflowoutcomeDud struct { 
    


    


    


    


    

}

// Analyticsflowoutcome
type Analyticsflowoutcome struct { 
    // FlowOutcome - Combination of unique flow outcome identifier and its value separated by colon
    FlowOutcome string `json:"flowOutcome"`


    // FlowOutcomeEndTimestamp - The outcome ending timestamp in ISO 8601 format. This may be null if the outcome did not succeed.
    FlowOutcomeEndTimestamp time.Time `json:"flowOutcomeEndTimestamp"`


    // FlowOutcomeId - Unique identifier of a flow outcome
    FlowOutcomeId string `json:"flowOutcomeId"`


    // FlowOutcomeStartTimestamp - The outcome starting timestamp in ISO 8601 format
    FlowOutcomeStartTimestamp time.Time `json:"flowOutcomeStartTimestamp"`


    // FlowOutcomeValue - Flow outcome value, e.g. SUCCESS
    FlowOutcomeValue string `json:"flowOutcomeValue"`

}

// String returns a JSON representation of the model
func (o *Analyticsflowoutcome) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticsflowoutcome) MarshalJSON() ([]byte, error) {
    type Alias Analyticsflowoutcome

    if AnalyticsflowoutcomeMarshalled {
        return []byte("{}"), nil
    }
    AnalyticsflowoutcomeMarshalled = true

    return json.Marshal(&struct { 
        FlowOutcome string `json:"flowOutcome"`
        
        FlowOutcomeEndTimestamp time.Time `json:"flowOutcomeEndTimestamp"`
        
        FlowOutcomeId string `json:"flowOutcomeId"`
        
        FlowOutcomeStartTimestamp time.Time `json:"flowOutcomeStartTimestamp"`
        
        FlowOutcomeValue string `json:"flowOutcomeValue"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

