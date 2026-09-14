package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UseractivityroutingstatusMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UseractivityroutingstatusDud struct { 
    


    

}

// Useractivityroutingstatus
type Useractivityroutingstatus struct { 
    // Status - Indicates the current routing status of the agent
    Status string `json:"status"`


    // StartTime - The timestamp when the agent went into this state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartTime time.Time `json:"startTime"`

}

// String returns a JSON representation of the model
func (o *Useractivityroutingstatus) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Useractivityroutingstatus) MarshalJSON() ([]byte, error) {
    type Alias Useractivityroutingstatus

    if UseractivityroutingstatusMarshalled {
        return []byte("{}"), nil
    }
    UseractivityroutingstatusMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        StartTime time.Time `json:"startTime"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

