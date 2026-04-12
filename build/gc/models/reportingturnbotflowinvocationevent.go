package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReportingturnbotflowinvocationeventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReportingturnbotflowinvocationeventDud struct { 
    


    


    

}

// Reportingturnbotflowinvocationevent
type Reportingturnbotflowinvocationevent struct { 
    // VarType - Represents the type of invocation event which occurred.
    VarType string `json:"type"`


    // Action - The action in which the event occurred.
    Action Reportingturnaction `json:"action"`


    // Flow - The details relating to the invoking or invoked flow.
    Flow Reportingturnflow `json:"flow"`

}

// String returns a JSON representation of the model
func (o *Reportingturnbotflowinvocationevent) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reportingturnbotflowinvocationevent) MarshalJSON() ([]byte, error) {
    type Alias Reportingturnbotflowinvocationevent

    if ReportingturnbotflowinvocationeventMarshalled {
        return []byte("{}"), nil
    }
    ReportingturnbotflowinvocationeventMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Action Reportingturnaction `json:"action"`
        
        Flow Reportingturnflow `json:"flow"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

