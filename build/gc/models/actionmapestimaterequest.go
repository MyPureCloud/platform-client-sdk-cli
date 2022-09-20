package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActionmapestimaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActionmapestimaterequestDud struct { 
    


    

}

// Actionmapestimaterequest
type Actionmapestimaterequest struct { 
    // SegmentIds - List of Segment IDs.
    SegmentIds []string `json:"segmentIds"`


    // OutcomeCriteria - Outcome Criteria containing outcomeId and probability thresholds.
    OutcomeCriteria Actionmapestimateoutcomecriteria `json:"outcomeCriteria"`

}

// String returns a JSON representation of the model
func (o *Actionmapestimaterequest) String() string {
     o.SegmentIds = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Actionmapestimaterequest) MarshalJSON() ([]byte, error) {
    type Alias Actionmapestimaterequest

    if ActionmapestimaterequestMarshalled {
        return []byte("{}"), nil
    }
    ActionmapestimaterequestMarshalled = true

    return json.Marshal(&struct {
        
        SegmentIds []string `json:"segmentIds"`
        
        OutcomeCriteria Actionmapestimateoutcomecriteria `json:"outcomeCriteria"`
        *Alias
    }{

        
        SegmentIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

