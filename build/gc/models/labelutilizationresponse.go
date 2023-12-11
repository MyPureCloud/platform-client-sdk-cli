package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LabelutilizationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LabelutilizationresponseDud struct { 
    


    


    

}

// Labelutilizationresponse
type Labelutilizationresponse struct { 
    // MaximumCapacity - Defines the maximum number of interactions with this label that an agent can handle at one time.
    MaximumCapacity int `json:"maximumCapacity"`


    // InterruptingLabelIds - Defines other labels that can interrupt an interaction with this label.
    InterruptingLabelIds []string `json:"interruptingLabelIds"`


    // LabelName - Name of the label this utilization relates to.
    LabelName string `json:"labelName"`

}

// String returns a JSON representation of the model
func (o *Labelutilizationresponse) String() string {
    
     o.InterruptingLabelIds = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Labelutilizationresponse) MarshalJSON() ([]byte, error) {
    type Alias Labelutilizationresponse

    if LabelutilizationresponseMarshalled {
        return []byte("{}"), nil
    }
    LabelutilizationresponseMarshalled = true

    return json.Marshal(&struct {
        
        MaximumCapacity int `json:"maximumCapacity"`
        
        InterruptingLabelIds []string `json:"interruptingLabelIds"`
        
        LabelName string `json:"labelName"`
        *Alias
    }{

        


        
        InterruptingLabelIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

