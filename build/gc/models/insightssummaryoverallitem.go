package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InsightssummaryoverallitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InsightssummaryoverallitemDud struct { 
    


    


    

}

// Insightssummaryoverallitem
type Insightssummaryoverallitem struct { 
    // ComparativePeriod - Insights data in the comparative period
    ComparativePeriod Insightssummaryoverallperiodpoints `json:"comparativePeriod"`


    // PrimaryPeriod - Insights data in the primary period
    PrimaryPeriod Insightssummaryoverallperiodpoints `json:"primaryPeriod"`


    // PercentOfGoalChange - Percent of goal change
    PercentOfGoalChange float64 `json:"percentOfGoalChange"`

}

// String returns a JSON representation of the model
func (o *Insightssummaryoverallitem) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Insightssummaryoverallitem) MarshalJSON() ([]byte, error) {
    type Alias Insightssummaryoverallitem

    if InsightssummaryoverallitemMarshalled {
        return []byte("{}"), nil
    }
    InsightssummaryoverallitemMarshalled = true

    return json.Marshal(&struct {
        
        ComparativePeriod Insightssummaryoverallperiodpoints `json:"comparativePeriod"`
        
        PrimaryPeriod Insightssummaryoverallperiodpoints `json:"primaryPeriod"`
        
        PercentOfGoalChange float64 `json:"percentOfGoalChange"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

