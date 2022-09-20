package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActionmapestimateresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActionmapestimateresultDud struct { 
    


    


    


    

}

// Actionmapestimateresult
type Actionmapestimateresult struct { 
    // QualifiedSessionCount - Number of sessions qualified for Action map.
    QualifiedSessionCount int `json:"qualifiedSessionCount"`


    // TotalSessionCount - Total number of sessions.
    TotalSessionCount int `json:"totalSessionCount"`


    // PerSegmentCounts - Number of sessions qualified for Action map per segment.
    PerSegmentCounts []Segmentestimatecount `json:"perSegmentCounts"`


    // OutcomesScoresCount - Difference made by outcome criteria to number of sessions qualified for Action map.
    OutcomesScoresCount int `json:"outcomesScoresCount"`

}

// String returns a JSON representation of the model
func (o *Actionmapestimateresult) String() string {
    
    
     o.PerSegmentCounts = []Segmentestimatecount{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Actionmapestimateresult) MarshalJSON() ([]byte, error) {
    type Alias Actionmapestimateresult

    if ActionmapestimateresultMarshalled {
        return []byte("{}"), nil
    }
    ActionmapestimateresultMarshalled = true

    return json.Marshal(&struct {
        
        QualifiedSessionCount int `json:"qualifiedSessionCount"`
        
        TotalSessionCount int `json:"totalSessionCount"`
        
        PerSegmentCounts []Segmentestimatecount `json:"perSegmentCounts"`
        
        OutcomesScoresCount int `json:"outcomesScoresCount"`
        *Alias
    }{

        


        


        
        PerSegmentCounts: []Segmentestimatecount{{}},
        


        

        Alias: (*Alias)(u),
    })
}

