package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SegmentassignmentsegmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SegmentassignmentsegmentDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Segmentassignmentsegment
type Segmentassignmentsegment struct { 
    


    // Scope - The target entity that a segment applies to.
    Scope string `json:"scope"`


    // Version - The version of the segment.
    Version int `json:"version"`


    

}

// String returns a JSON representation of the model
func (o *Segmentassignmentsegment) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Segmentassignmentsegment) MarshalJSON() ([]byte, error) {
    type Alias Segmentassignmentsegment

    if SegmentassignmentsegmentMarshalled {
        return []byte("{}"), nil
    }
    SegmentassignmentsegmentMarshalled = true

    return json.Marshal(&struct {
        
        Scope string `json:"scope"`
        
        Version int `json:"version"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

