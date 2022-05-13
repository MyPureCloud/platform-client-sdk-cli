package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowexecutionlaunchresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowexecutionlaunchresponseDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Flowexecutionlaunchresponse - Response object from launching a flow.
type Flowexecutionlaunchresponse struct { 
    // Id - The flow execution ID
    Id string `json:"id"`


    // Name - The flow execution name.
    Name string `json:"name"`


    // FlowVersion - The version of the flow that launched
    FlowVersion Domainentityref `json:"flowVersion"`


    

}

// String returns a JSON representation of the model
func (o *Flowexecutionlaunchresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowexecutionlaunchresponse) MarshalJSON() ([]byte, error) {
    type Alias Flowexecutionlaunchresponse

    if FlowexecutionlaunchresponseMarshalled {
        return []byte("{}"), nil
    }
    FlowexecutionlaunchresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        FlowVersion Domainentityref `json:"flowVersion"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

