package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowexecutionlaunchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowexecutionlaunchrequestDud struct { 
    


    


    


    

}

// Flowexecutionlaunchrequest - Parameters for launching a flow.
type Flowexecutionlaunchrequest struct { 
    // FlowId - ID of the flow to launch.
    FlowId string `json:"flowId"`


    // FlowVersion - The version of the flow to launch. Omit this value (or supply null/empty) to use the latest published version.
    FlowVersion string `json:"flowVersion"`


    // InputData - Input values to the flow. Valid values are defined by a flow's input JSON schema.
    InputData map[string]interface{} `json:"inputData"`


    // Name - A displayable name to assign to the new flow execution
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Flowexecutionlaunchrequest) String() string {
    
    
     o.InputData = map[string]interface{}{"": Interface{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowexecutionlaunchrequest) MarshalJSON() ([]byte, error) {
    type Alias Flowexecutionlaunchrequest

    if FlowexecutionlaunchrequestMarshalled {
        return []byte("{}"), nil
    }
    FlowexecutionlaunchrequestMarshalled = true

    return json.Marshal(&struct {
        
        FlowId string `json:"flowId"`
        
        FlowVersion string `json:"flowVersion"`
        
        InputData map[string]interface{} `json:"inputData"`
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        
        InputData: map[string]interface{}{"": Interface{}},
        


        

        Alias: (*Alias)(u),
    })
}

