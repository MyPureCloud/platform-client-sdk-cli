package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowruntimeexecutionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowruntimeexecutionDud struct { 
    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Flowruntimeexecution - Details about the current state of a Flow execution
type Flowruntimeexecution struct { 
    // Id - The flow execution ID
    Id string `json:"id"`


    // Name - The flow execution name.
    Name string `json:"name"`


    // FlowVersion - The Version of the flow definition of the flow execution.
    FlowVersion Flowversion `json:"flowVersion"`


    // DateLaunched - The time the flow was launched. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateLaunched time.Time `json:"dateLaunched"`


    // Status - The flow's running status, which indicates whether the flow is running normally or completed, etc.
    Status string `json:"status"`


    // DateCompleted - The time the flow completed, if applicable. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCompleted time.Time `json:"dateCompleted"`


    // CompletionReason - The completion reason set at the flow completion time, if applicable.
    CompletionReason string `json:"completionReason"`


    // FlowErrorInfo - Additional information if the flow is in error
    FlowErrorInfo Errorbody `json:"flowErrorInfo"`


    // OutputData - List of the flow's output variables, if any. Output variables are only supplied for Completed flows.
    OutputData map[string]interface{} `json:"outputData"`


    // Conversation - The conversation to which this Flow execution is related
    Conversation Domainentityref `json:"conversation"`


    

}

// String returns a JSON representation of the model
func (o *Flowruntimeexecution) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.OutputData = map[string]interface{}{"": Interface{}} 
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowruntimeexecution) MarshalJSON() ([]byte, error) {
    type Alias Flowruntimeexecution

    if FlowruntimeexecutionMarshalled {
        return []byte("{}"), nil
    }
    FlowruntimeexecutionMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        FlowVersion Flowversion `json:"flowVersion"`
        
        DateLaunched time.Time `json:"dateLaunched"`
        
        Status string `json:"status"`
        
        DateCompleted time.Time `json:"dateCompleted"`
        
        CompletionReason string `json:"completionReason"`
        
        FlowErrorInfo Errorbody `json:"flowErrorInfo"`
        
        OutputData map[string]interface{} `json:"outputData"`
        
        Conversation Domainentityref `json:"conversation"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        OutputData: map[string]interface{}{"": Interface{}},
        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

