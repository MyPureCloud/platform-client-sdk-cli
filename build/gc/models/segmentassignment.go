package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SegmentassignmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SegmentassignmentDud struct { 
    


    


    


    


    ExternalContact Addressableentityref `json:"externalContact"`

}

// Segmentassignment
type Segmentassignment struct { 
    // DateAssigned - Date when the segment was assigned. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateAssigned time.Time `json:"dateAssigned"`


    // DateForUnassignment - Date indicating when a segment is scheduled to be unassigned. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateForUnassignment time.Time `json:"dateForUnassignment"`


    // Segment - The segment the assignment is for.
    Segment Segmentassignmentsegment `json:"segment"`


    // Session - For session-scoped segments, the session for which the segment was assigned.
    Session Segmentassignmentsession `json:"session"`


    

}

// String returns a JSON representation of the model
func (o *Segmentassignment) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Segmentassignment) MarshalJSON() ([]byte, error) {
    type Alias Segmentassignment

    if SegmentassignmentMarshalled {
        return []byte("{}"), nil
    }
    SegmentassignmentMarshalled = true

    return json.Marshal(&struct {
        
        DateAssigned time.Time `json:"dateAssigned"`
        
        DateForUnassignment time.Time `json:"dateForUnassignment"`
        
        Segment Segmentassignmentsegment `json:"segment"`
        
        Session Segmentassignmentsession `json:"session"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

