package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulingnoforecastoptionsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulingnoforecastoptionsrequestDud struct { 
    


    

}

// Schedulingnoforecastoptionsrequest
type Schedulingnoforecastoptionsrequest struct { 
    // ShiftLength - The shift length option to apply if no forecast is supplied
    ShiftLength string `json:"shiftLength"`


    // ShiftStart - The shift start option to apply if no forecast is supplied
    ShiftStart string `json:"shiftStart"`

}

// String returns a JSON representation of the model
func (o *Schedulingnoforecastoptionsrequest) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulingnoforecastoptionsrequest) MarshalJSON() ([]byte, error) {
    type Alias Schedulingnoforecastoptionsrequest

    if SchedulingnoforecastoptionsrequestMarshalled {
        return []byte("{}"), nil
    }
    SchedulingnoforecastoptionsrequestMarshalled = true

    return json.Marshal(&struct { 
        ShiftLength string `json:"shiftLength"`
        
        ShiftStart string `json:"shiftStart"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

