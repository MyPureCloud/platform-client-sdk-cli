package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AllocationoutputstemplateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AllocationoutputstemplateDud struct { 
    


    


    

}

// Allocationoutputstemplate
type Allocationoutputstemplate struct { 
    // CalculationStartDate - The beginning of the allocation results, in ISO-8601 format
    CalculationStartDate time.Time `json:"calculationStartDate"`


    // CalculationIntervalLengthMinutes - Interval length of the response metrics
    CalculationIntervalLengthMinutes int `json:"calculationIntervalLengthMinutes"`


    // PlanningGroupAllocationResults - Planning group level allocation results
    PlanningGroupAllocationResults []Allocationresultstemplate `json:"planningGroupAllocationResults"`

}

// String returns a JSON representation of the model
func (o *Allocationoutputstemplate) String() string {
    
    
     o.PlanningGroupAllocationResults = []Allocationresultstemplate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Allocationoutputstemplate) MarshalJSON() ([]byte, error) {
    type Alias Allocationoutputstemplate

    if AllocationoutputstemplateMarshalled {
        return []byte("{}"), nil
    }
    AllocationoutputstemplateMarshalled = true

    return json.Marshal(&struct {
        
        CalculationStartDate time.Time `json:"calculationStartDate"`
        
        CalculationIntervalLengthMinutes int `json:"calculationIntervalLengthMinutes"`
        
        PlanningGroupAllocationResults []Allocationresultstemplate `json:"planningGroupAllocationResults"`
        *Alias
    }{

        


        


        
        PlanningGroupAllocationResults: []Allocationresultstemplate{{}},
        

        Alias: (*Alias)(u),
    })
}

