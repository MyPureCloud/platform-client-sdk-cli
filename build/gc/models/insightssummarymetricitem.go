package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InsightssummarymetricitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InsightssummarymetricitemDud struct { 
    


    


    


    


    

}

// Insightssummarymetricitem
type Insightssummarymetricitem struct { 
    // Metric - The gamification metric for the data
    Metric Addressableentityref `json:"metric"`


    // ComparativePeriod - Insights data in the comparative period
    ComparativePeriod Insightssummarymetricperiodpoints `json:"comparativePeriod"`


    // PrimaryPeriod - Insights data in the primary period
    PrimaryPeriod Insightssummarymetricperiodpoints `json:"primaryPeriod"`


    // PercentOfGoalChange - Percent of goal change
    PercentOfGoalChange float64 `json:"percentOfGoalChange"`


    // ValueChange - Value change
    ValueChange float64 `json:"valueChange"`

}

// String returns a JSON representation of the model
func (o *Insightssummarymetricitem) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Insightssummarymetricitem) MarshalJSON() ([]byte, error) {
    type Alias Insightssummarymetricitem

    if InsightssummarymetricitemMarshalled {
        return []byte("{}"), nil
    }
    InsightssummarymetricitemMarshalled = true

    return json.Marshal(&struct {
        
        Metric Addressableentityref `json:"metric"`
        
        ComparativePeriod Insightssummarymetricperiodpoints `json:"comparativePeriod"`
        
        PrimaryPeriod Insightssummarymetricperiodpoints `json:"primaryPeriod"`
        
        PercentOfGoalChange float64 `json:"percentOfGoalChange"`
        
        ValueChange float64 `json:"valueChange"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

