package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReportingturnguardraileventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReportingturnguardraileventDud struct { 
    


    


    


    

}

// Reportingturnguardrailevent
type Reportingturnguardrailevent struct { 
    // VarType - The type of guardrail violation
    VarType string `json:"type"`


    // Instruction - The attached instruction to the guardrail
    Instruction string `json:"instruction"`


    // ViolationsThreshold - The number of violations allowed before an exit occurs.
    ViolationsThreshold int `json:"violationsThreshold"`


    // ViolationsTriggered - The current amount of violations that have been triggered in the current action.
    ViolationsTriggered int `json:"violationsTriggered"`

}

// String returns a JSON representation of the model
func (o *Reportingturnguardrailevent) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reportingturnguardrailevent) MarshalJSON() ([]byte, error) {
    type Alias Reportingturnguardrailevent

    if ReportingturnguardraileventMarshalled {
        return []byte("{}"), nil
    }
    ReportingturnguardraileventMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Instruction string `json:"instruction"`
        
        ViolationsThreshold int `json:"violationsThreshold"`
        
        ViolationsTriggered int `json:"violationsTriggered"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

