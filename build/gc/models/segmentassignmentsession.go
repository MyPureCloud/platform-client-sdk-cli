package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SegmentassignmentsessionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SegmentassignmentsessionDud struct { 
    


    


    SelfUri string `json:"selfUri"`

}

// Segmentassignmentsession
type Segmentassignmentsession struct { 
    // Id - The ID of the session.
    Id string `json:"id"`


    // VarType - The type or category of session (e.g. web, app).
    VarType string `json:"type"`


    

}

// String returns a JSON representation of the model
func (o *Segmentassignmentsession) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Segmentassignmentsession) MarshalJSON() ([]byte, error) {
    type Alias Segmentassignmentsession

    if SegmentassignmentsessionMarshalled {
        return []byte("{}"), nil
    }
    SegmentassignmentsessionMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        VarType string `json:"type"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

