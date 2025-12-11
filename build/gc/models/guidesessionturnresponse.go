package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GuidesessionturnresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GuidesessionturnresponseDud struct { 
    


    


    


    


    


    

}

// Guidesessionturnresponse - Response for a guide session turn
type Guidesessionturnresponse struct { 
    // Response - The response content for this turn.
    Response Guidesessionturnresponsedata `json:"response"`


    // Status - The status of the turn.
    Status string `json:"status"`


    // Result - The result of the turn.
    Result string `json:"result"`


    // OutputVariables - The output variables for this turn.
    OutputVariables []Guidesessionvariable `json:"outputVariables"`


    // InvocationId - Invocation ID for this turn.
    InvocationId string `json:"invocationId"`


    // Invocations - The invocations for this turn.
    Invocations []Guidesessionturninvocationresponse `json:"invocations"`

}

// String returns a JSON representation of the model
func (o *Guidesessionturnresponse) String() string {
    
    
    
     o.OutputVariables = []Guidesessionvariable{{}} 
    
     o.Invocations = []Guidesessionturninvocationresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Guidesessionturnresponse) MarshalJSON() ([]byte, error) {
    type Alias Guidesessionturnresponse

    if GuidesessionturnresponseMarshalled {
        return []byte("{}"), nil
    }
    GuidesessionturnresponseMarshalled = true

    return json.Marshal(&struct {
        
        Response Guidesessionturnresponsedata `json:"response"`
        
        Status string `json:"status"`
        
        Result string `json:"result"`
        
        OutputVariables []Guidesessionvariable `json:"outputVariables"`
        
        InvocationId string `json:"invocationId"`
        
        Invocations []Guidesessionturninvocationresponse `json:"invocations"`
        *Alias
    }{

        


        


        


        
        OutputVariables: []Guidesessionvariable{{}},
        


        


        
        Invocations: []Guidesessionturninvocationresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

