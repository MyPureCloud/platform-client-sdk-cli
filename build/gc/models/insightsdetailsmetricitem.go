package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InsightsdetailsmetricitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InsightsdetailsmetricitemDud struct { 
    


    


    


    


    

}

// Insightsdetailsmetricitem
type Insightsdetailsmetricitem struct { 
    // Metric - The gamification metric for the data
    Metric Addressableentityref `json:"metric"`


    // ComparativePeriod - Insights data in the comparative period
    ComparativePeriod Insightsdetailsmetricperiodpoints `json:"comparativePeriod"`


    // PrimaryPeriod - Insights data in the primary period
    PrimaryPeriod Insightsdetailsmetricperiodpoints `json:"primaryPeriod"`


    // PercentOfGoalChange - Percent of goal change
    PercentOfGoalChange float64 `json:"percentOfGoalChange"`


    // ValueChange - Value change
    ValueChange float64 `json:"valueChange"`

}

// String returns a JSON representation of the model
func (o *Insightsdetailsmetricitem) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Insightsdetailsmetricitem) MarshalJSON() ([]byte, error) {
    type Alias Insightsdetailsmetricitem

    if InsightsdetailsmetricitemMarshalled {
        return []byte("{}"), nil
    }
    InsightsdetailsmetricitemMarshalled = true

    return json.Marshal(&struct {
        
        Metric Addressableentityref `json:"metric"`
        
        ComparativePeriod Insightsdetailsmetricperiodpoints `json:"comparativePeriod"`
        
        PrimaryPeriod Insightsdetailsmetricperiodpoints `json:"primaryPeriod"`
        
        PercentOfGoalChange float64 `json:"percentOfGoalChange"`
        
        ValueChange float64 `json:"valueChange"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

