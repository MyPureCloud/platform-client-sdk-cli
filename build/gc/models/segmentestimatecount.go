package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SegmentestimatecountMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SegmentestimatecountDud struct { 
    


    

}

// Segmentestimatecount
type Segmentestimatecount struct { 
    // SegmentId - ID of Segment.
    SegmentId string `json:"segmentId"`


    // Count - Estimate count per segment.
    Count int `json:"count"`

}

// String returns a JSON representation of the model
func (o *Segmentestimatecount) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Segmentestimatecount) MarshalJSON() ([]byte, error) {
    type Alias Segmentestimatecount

    if SegmentestimatecountMarshalled {
        return []byte("{}"), nil
    }
    SegmentestimatecountMarshalled = true

    return json.Marshal(&struct {
        
        SegmentId string `json:"segmentId"`
        
        Count int `json:"count"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

