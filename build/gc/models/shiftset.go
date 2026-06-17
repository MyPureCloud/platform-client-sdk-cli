package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShiftsetMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShiftsetDud struct { 
    


    


    


    


    


    

}

// Shiftset
type Shiftset struct { 
    // Id - The ID of the shift set
    Id string `json:"id"`


    // Name - The name given for the shift set
    Name string `json:"name"`


    // EffectiveWorkPlan - The work plan or work plan rotation used for generating the shift set
    EffectiveWorkPlan Shiftseteffectiveworkplan `json:"effectiveWorkPlan"`


    // Shifts - The scheduled shifts
    Shifts []Schedulebidscheduledshift `json:"shifts"`


    // SuggestedAgentCount - The suggested agent count
    SuggestedAgentCount int `json:"suggestedAgentCount"`


    // OverrideAgentCount - The override agent count. If it is null, it falls back to using the suggestedAgentCount
    OverrideAgentCount int `json:"overrideAgentCount"`

}

// String returns a JSON representation of the model
func (o *Shiftset) String() string {
    
    
    
     o.Shifts = []Schedulebidscheduledshift{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shiftset) MarshalJSON() ([]byte, error) {
    type Alias Shiftset

    if ShiftsetMarshalled {
        return []byte("{}"), nil
    }
    ShiftsetMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        EffectiveWorkPlan Shiftseteffectiveworkplan `json:"effectiveWorkPlan"`
        
        Shifts []Schedulebidscheduledshift `json:"shifts"`
        
        SuggestedAgentCount int `json:"suggestedAgentCount"`
        
        OverrideAgentCount int `json:"overrideAgentCount"`
        *Alias
    }{

        


        


        


        
        Shifts: []Schedulebidscheduledshift{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

