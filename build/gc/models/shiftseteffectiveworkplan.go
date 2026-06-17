package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShiftseteffectiveworkplanMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShiftseteffectiveworkplanDud struct { 
    


    


    

}

// Shiftseteffectiveworkplan
type Shiftseteffectiveworkplan struct { 
    // WorkPlanId - The ID of the work plan
    WorkPlanId string `json:"workPlanId"`


    // WorkPlanRotationId - The ID of the work plan rotation
    WorkPlanRotationId string `json:"workPlanRotationId"`


    // PositionInRotation - The position in rotation
    PositionInRotation int `json:"positionInRotation"`

}

// String returns a JSON representation of the model
func (o *Shiftseteffectiveworkplan) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shiftseteffectiveworkplan) MarshalJSON() ([]byte, error) {
    type Alias Shiftseteffectiveworkplan

    if ShiftseteffectiveworkplanMarshalled {
        return []byte("{}"), nil
    }
    ShiftseteffectiveworkplanMarshalled = true

    return json.Marshal(&struct {
        
        WorkPlanId string `json:"workPlanId"`
        
        WorkPlanRotationId string `json:"workPlanRotationId"`
        
        PositionInRotation int `json:"positionInRotation"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

