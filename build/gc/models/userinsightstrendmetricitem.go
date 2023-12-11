package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserinsightstrendmetricitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserinsightstrendmetricitemDud struct { 
    


    

}

// Userinsightstrendmetricitem
type Userinsightstrendmetricitem struct { 
    // Metric - The gamification metric for the trend
    Metric Addressableentityref `json:"metric"`


    // Trends - Trends for the metric
    Trends Userinsightstrends `json:"trends"`

}

// String returns a JSON representation of the model
func (o *Userinsightstrendmetricitem) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userinsightstrendmetricitem) MarshalJSON() ([]byte, error) {
    type Alias Userinsightstrendmetricitem

    if UserinsightstrendmetricitemMarshalled {
        return []byte("{}"), nil
    }
    UserinsightstrendmetricitemMarshalled = true

    return json.Marshal(&struct {
        
        Metric Addressableentityref `json:"metric"`
        
        Trends Userinsightstrends `json:"trends"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

