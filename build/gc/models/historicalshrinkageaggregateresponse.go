package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HistoricalshrinkageaggregateresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HistoricalshrinkageaggregateresponseDud struct { 
    


    


    


    


    


    


    


    

}

// Historicalshrinkageaggregateresponse
type Historicalshrinkageaggregateresponse struct { 
    // ScheduledShrinkageSeconds - Aggregated shrinkage value in seconds for scheduled activities
    ScheduledShrinkageSeconds int `json:"scheduledShrinkageSeconds"`


    // ScheduledShrinkagePercent - Aggregated shrinkage value in percent from 0.0 to 100.0 for scheduled activities
    ScheduledShrinkagePercent float64 `json:"scheduledShrinkagePercent"`


    // ActualShrinkageSeconds - Aggregated actual value in seconds for scheduled activities
    ActualShrinkageSeconds int `json:"actualShrinkageSeconds"`


    // ActualShrinkagePercent - Aggregated actual value in percent from 0.0 to 100.0 for scheduled activities
    ActualShrinkagePercent float64 `json:"actualShrinkagePercent"`


    // PaidShrinkageSeconds - Aggregated shrinkage value in seconds for paid activities
    PaidShrinkageSeconds int `json:"paidShrinkageSeconds"`


    // UnpaidShrinkageSeconds - Aggregated shrinkage value in seconds for unpaid activities
    UnpaidShrinkageSeconds int `json:"unpaidShrinkageSeconds"`


    // PlannedShrinkageSeconds - Aggregated shrinkage value in seconds for planned activities
    PlannedShrinkageSeconds int `json:"plannedShrinkageSeconds"`


    // UnplannedShrinkageSeconds - Aggregated shrinkage value in seconds for unplanned activities
    UnplannedShrinkageSeconds int `json:"unplannedShrinkageSeconds"`

}

// String returns a JSON representation of the model
func (o *Historicalshrinkageaggregateresponse) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Historicalshrinkageaggregateresponse) MarshalJSON() ([]byte, error) {
    type Alias Historicalshrinkageaggregateresponse

    if HistoricalshrinkageaggregateresponseMarshalled {
        return []byte("{}"), nil
    }
    HistoricalshrinkageaggregateresponseMarshalled = true

    return json.Marshal(&struct {
        
        ScheduledShrinkageSeconds int `json:"scheduledShrinkageSeconds"`
        
        ScheduledShrinkagePercent float64 `json:"scheduledShrinkagePercent"`
        
        ActualShrinkageSeconds int `json:"actualShrinkageSeconds"`
        
        ActualShrinkagePercent float64 `json:"actualShrinkagePercent"`
        
        PaidShrinkageSeconds int `json:"paidShrinkageSeconds"`
        
        UnpaidShrinkageSeconds int `json:"unpaidShrinkageSeconds"`
        
        PlannedShrinkageSeconds int `json:"plannedShrinkageSeconds"`
        
        UnplannedShrinkageSeconds int `json:"unplannedShrinkageSeconds"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

