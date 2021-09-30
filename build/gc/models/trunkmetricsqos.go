package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrunkmetricsqosMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrunkmetricsqosDud struct { 
    

}

// Trunkmetricsqos
type Trunkmetricsqos struct { 
    // MismatchCount - Total number of QoS mismatches over the course of the last 24-hour period (sliding window).
    MismatchCount int `json:"mismatchCount"`

}

// String returns a JSON representation of the model
func (o *Trunkmetricsqos) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trunkmetricsqos) MarshalJSON() ([]byte, error) {
    type Alias Trunkmetricsqos

    if TrunkmetricsqosMarshalled {
        return []byte("{}"), nil
    }
    TrunkmetricsqosMarshalled = true

    return json.Marshal(&struct { 
        MismatchCount int `json:"mismatchCount"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

