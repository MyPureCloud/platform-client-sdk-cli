package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SelfschedulejobMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SelfschedulejobDud struct { 
    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Selfschedulejob
type Selfschedulejob struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Status - The status of the job
    Status string `json:"status"`


    // ActivityMoves - The self-schedule activity moves, if status == 'Complete'
    ActivityMoves []Selfscheduleactivitymovelistitem `json:"activityMoves"`


    // VarError - Any error information, only set if the status == 'Error'
    VarError Errorbody `json:"error"`


    

}

// String returns a JSON representation of the model
func (o *Selfschedulejob) String() string {
    
    
     o.ActivityMoves = []Selfscheduleactivitymovelistitem{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Selfschedulejob) MarshalJSON() ([]byte, error) {
    type Alias Selfschedulejob

    if SelfschedulejobMarshalled {
        return []byte("{}"), nil
    }
    SelfschedulejobMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Status string `json:"status"`
        
        ActivityMoves []Selfscheduleactivitymovelistitem `json:"activityMoves"`
        
        VarError Errorbody `json:"error"`
        *Alias
    }{

        


        


        
        ActivityMoves: []Selfscheduleactivitymovelistitem{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

