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
    // Id - Unique identifier for the segment assignment.
    Id string `json:"id"`


    // State - State of the segment assignment.
    State string `json:"state"`


    // DateAssigned - Date when the segment was assigned. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateAssigned time.Time `json:"dateAssigned"`


    // DateUnassigned - Date when the segment was unassigned. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateUnassigned time.Time `json:"dateUnassigned"`


    // DateModified - Date when the segment assignment was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // Segment - The segment the assignment is for.
    Segment Segmentassignmentsegment `json:"segment"`


    // CustomerId - ID of the customer to which the segment is assigned.
    CustomerId string `json:"customerId"`


    // CustomerIdType - Type of customer ID (e.g. cookie, email, phone).
    CustomerIdType string `json:"customerIdType"`


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
        
        Id string `json:"id"`
        
        State string `json:"state"`
        
        DateAssigned time.Time `json:"dateAssigned"`
        
        DateUnassigned time.Time `json:"dateUnassigned"`
        
        DateModified time.Time `json:"dateModified"`
        
        Segment Segmentassignmentsegment `json:"segment"`
        
        CustomerId string `json:"customerId"`
        
        CustomerIdType string `json:"customerIdType"`
        
        Session Segmentassignmentsession `json:"session"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

