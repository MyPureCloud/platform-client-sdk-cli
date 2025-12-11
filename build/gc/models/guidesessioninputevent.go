package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GuidesessioninputeventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GuidesessioninputeventDud struct { 
    


    


    


    


    

}

// Guidesessioninputevent
type Guidesessioninputevent struct { 
    // VarType - The type of input event.
    VarType string `json:"type"`


    // Text - The text content of the input event.
    Text string `json:"text"`


    // Mode - The input mode for this event.
    Mode string `json:"mode"`


    // InvocationId - The invocation ID of the input event.
    InvocationId string `json:"invocationId"`


    // Invocations - The invocation data of the input event.
    Invocations []Guidesessionturninvocationdata `json:"invocations"`

}

// String returns a JSON representation of the model
func (o *Guidesessioninputevent) String() string {
    
    
    
    
     o.Invocations = []Guidesessionturninvocationdata{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Guidesessioninputevent) MarshalJSON() ([]byte, error) {
    type Alias Guidesessioninputevent

    if GuidesessioninputeventMarshalled {
        return []byte("{}"), nil
    }
    GuidesessioninputeventMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Text string `json:"text"`
        
        Mode string `json:"mode"`
        
        InvocationId string `json:"invocationId"`
        
        Invocations []Guidesessionturninvocationdata `json:"invocations"`
        *Alias
    }{

        


        


        


        


        
        Invocations: []Guidesessionturninvocationdata{{}},
        

        Alias: (*Alias)(u),
    })
}

