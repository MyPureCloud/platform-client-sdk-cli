package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InsightsdetailsoverallitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InsightsdetailsoverallitemDud struct { 
    


    


    

}

// Insightsdetailsoverallitem
type Insightsdetailsoverallitem struct { 
    // ComparativePeriod - Insights data in the comparative period
    ComparativePeriod Insightsdetailsoverallperiodpoints `json:"comparativePeriod"`


    // PrimaryPeriod - Insights data in the primary period
    PrimaryPeriod Insightsdetailsoverallperiodpoints `json:"primaryPeriod"`


    // PercentOfGoalChange - Percent of goal change
    PercentOfGoalChange float64 `json:"percentOfGoalChange"`

}

// String returns a JSON representation of the model
func (o *Insightsdetailsoverallitem) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Insightsdetailsoverallitem) MarshalJSON() ([]byte, error) {
    type Alias Insightsdetailsoverallitem

    if InsightsdetailsoverallitemMarshalled {
        return []byte("{}"), nil
    }
    InsightsdetailsoverallitemMarshalled = true

    return json.Marshal(&struct {
        
        ComparativePeriod Insightsdetailsoverallperiodpoints `json:"comparativePeriod"`
        
        PrimaryPeriod Insightsdetailsoverallperiodpoints `json:"primaryPeriod"`
        
        PercentOfGoalChange float64 `json:"percentOfGoalChange"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

