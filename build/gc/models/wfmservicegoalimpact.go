package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmservicegoalimpactMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmservicegoalimpactDud struct { 
    


    

}

// Wfmservicegoalimpact
type Wfmservicegoalimpact struct { 
    // IncreaseByPercent - The maximum allowed percent increase from the configured goal
    IncreaseByPercent float64 `json:"increaseByPercent"`


    // DecreaseByPercent - The maximum allowed percent decrease from the configured goal
    DecreaseByPercent float64 `json:"decreaseByPercent"`

}

// String returns a JSON representation of the model
func (o *Wfmservicegoalimpact) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmservicegoalimpact) MarshalJSON() ([]byte, error) {
    type Alias Wfmservicegoalimpact

    if WfmservicegoalimpactMarshalled {
        return []byte("{}"), nil
    }
    WfmservicegoalimpactMarshalled = true

    return json.Marshal(&struct {
        
        IncreaseByPercent float64 `json:"increaseByPercent"`
        
        DecreaseByPercent float64 `json:"decreaseByPercent"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

