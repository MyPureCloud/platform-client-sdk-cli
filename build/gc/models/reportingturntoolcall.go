package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReportingturntoolcallMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReportingturntoolcallDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Reportingturntoolcall
type Reportingturntoolcall struct { 
    // ToolId - Represents the identifier of the tool called.
    ToolId string `json:"toolId"`


    // ToolName - Represents the name of the tool used in the event.
    ToolName string `json:"toolName"`


    // ToolType - Represents the type of tool used in the event.
    ToolType string `json:"toolType"`


    // TargetId - Represents the identifier of the target that the tool is using.
    TargetId string `json:"targetId"`


    // Status - Represents whether the tool call was successful or not.
    Status string `json:"status"`


    // ErrorText - Represents the error returned by the tool in the event of a failure.
    ErrorText string `json:"errorText"`


    // DateInvoked - Represents the starting time of the tool call. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateInvoked time.Time `json:"dateInvoked"`


    // LatencyMs - Represents the time it took the tool call to execute.
    LatencyMs int `json:"latencyMs"`


    // Origin - Represents the origin of the tool call.
    Origin string `json:"origin"`


    // KnowledgeMetadata - Represents various metadata of knowledge calls used by the tool if the tool is configured to use knowledge.
    KnowledgeMetadata Reportingturnknowledgemetadata `json:"knowledgeMetadata"`

}

// String returns a JSON representation of the model
func (o *Reportingturntoolcall) String() string {
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reportingturntoolcall) MarshalJSON() ([]byte, error) {
    type Alias Reportingturntoolcall

    if ReportingturntoolcallMarshalled {
        return []byte("{}"), nil
    }
    ReportingturntoolcallMarshalled = true

    return json.Marshal(&struct {
        
        ToolId string `json:"toolId"`
        
        ToolName string `json:"toolName"`
        
        ToolType string `json:"toolType"`
        
        TargetId string `json:"targetId"`
        
        Status string `json:"status"`
        
        ErrorText string `json:"errorText"`
        
        DateInvoked time.Time `json:"dateInvoked"`
        
        LatencyMs int `json:"latencyMs"`
        
        Origin string `json:"origin"`
        
        KnowledgeMetadata Reportingturnknowledgemetadata `json:"knowledgeMetadata"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

