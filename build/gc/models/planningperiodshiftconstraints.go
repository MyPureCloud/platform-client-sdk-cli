package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PlanningperiodshiftconstraintsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PlanningperiodshiftconstraintsDud struct { 
    


    


    

}

// Planningperiodshiftconstraints
type Planningperiodshiftconstraints struct { 
    // Enabled - Whether shifts per planning period is enabled. This field is non-nullable on the response
    Enabled bool `json:"enabled"`


    // MinimumCount - The minimum number of shifts required per planning period. This field is non-nullable on the response
    MinimumCount int `json:"minimumCount"`


    // MaximumCount - The maximum number of shifts allowed per planning period. This field is non-nullable on the response
    MaximumCount int `json:"maximumCount"`

}

// String returns a JSON representation of the model
func (o *Planningperiodshiftconstraints) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Planningperiodshiftconstraints) MarshalJSON() ([]byte, error) {
    type Alias Planningperiodshiftconstraints

    if PlanningperiodshiftconstraintsMarshalled {
        return []byte("{}"), nil
    }
    PlanningperiodshiftconstraintsMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        MinimumCount int `json:"minimumCount"`
        
        MaximumCount int `json:"maximumCount"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

