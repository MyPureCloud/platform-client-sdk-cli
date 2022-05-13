package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SegmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SegmentDud struct { 
    


    


    


    


    

}

// Segment
type Segment struct { 
    // StartTime - The timestamp when this segment began. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartTime time.Time `json:"startTime"`


    // EndTime - The timestamp when this segment ended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndTime time.Time `json:"endTime"`


    // VarType - The activity taking place for the participant in the segment.
    VarType string `json:"type"`


    // HowEnded - A description of the event that ended the segment.
    HowEnded string `json:"howEnded"`


    // DisconnectType - A description of the event that disconnected the segment
    DisconnectType string `json:"disconnectType"`

}

// String returns a JSON representation of the model
func (o *Segment) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Segment) MarshalJSON() ([]byte, error) {
    type Alias Segment

    if SegmentMarshalled {
        return []byte("{}"), nil
    }
    SegmentMarshalled = true

    return json.Marshal(&struct {
        
        StartTime time.Time `json:"startTime"`
        
        EndTime time.Time `json:"endTime"`
        
        VarType string `json:"type"`
        
        HowEnded string `json:"howEnded"`
        
        DisconnectType string `json:"disconnectType"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

