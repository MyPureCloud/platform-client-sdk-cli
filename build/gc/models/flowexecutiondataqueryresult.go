package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowexecutiondataqueryresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowexecutiondataqueryresultDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Flowexecutiondataqueryresult - This is the metadata of an executionData entry for a flow.
type Flowexecutiondataqueryresult struct { 
    


    // Name
    Name string `json:"name"`


    // StartDateTime - The start time for the execution of this flow. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartDateTime time.Time `json:"startDateTime"`


    // EndDateTime - The end time for the execution of this flow. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndDateTime time.Time `json:"endDateTime"`


    // FlowId - The id of the flow that was executed.
    FlowId string `json:"flowId"`


    // FlowVersion - The version of the flow that was executed.
    FlowVersion string `json:"flowVersion"`


    // ConversationId - The id of the conversation that executed this flow.
    ConversationId string `json:"conversationId"`


    // WorkitemId - The id of the workitem that executed this flow.
    WorkitemId string `json:"workitemId"`


    // FlowType - The type of flow.
    FlowType string `json:"flowType"`


    // FlowErrorReason - If the flow errored out this is the reason.
    FlowErrorReason string `json:"flowErrorReason"`


    // FlowWarningReason - If the flow had a warning, this is the reason.
    FlowWarningReason string `json:"flowWarningReason"`


    // FlowName - The name of the flow.
    FlowName string `json:"flowName"`


    

}

// String returns a JSON representation of the model
func (o *Flowexecutiondataqueryresult) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowexecutiondataqueryresult) MarshalJSON() ([]byte, error) {
    type Alias Flowexecutiondataqueryresult

    if FlowexecutiondataqueryresultMarshalled {
        return []byte("{}"), nil
    }
    FlowexecutiondataqueryresultMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        StartDateTime time.Time `json:"startDateTime"`
        
        EndDateTime time.Time `json:"endDateTime"`
        
        FlowId string `json:"flowId"`
        
        FlowVersion string `json:"flowVersion"`
        
        ConversationId string `json:"conversationId"`
        
        WorkitemId string `json:"workitemId"`
        
        FlowType string `json:"flowType"`
        
        FlowErrorReason string `json:"flowErrorReason"`
        
        FlowWarningReason string `json:"flowWarningReason"`
        
        FlowName string `json:"flowName"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

