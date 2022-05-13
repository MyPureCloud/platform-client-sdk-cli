package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticsflowMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticsflowDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Analyticsflow
type Analyticsflow struct { 
    // EndingLanguage - Flow ending language, e.g. en-us
    EndingLanguage string `json:"endingLanguage"`


    // EntryReason - The particular entry reason for this flow, e.g. an address, userId, or flowId
    EntryReason string `json:"entryReason"`


    // EntryType - The entry type for this flow, e.g. dnis, dialer, agent, flow, or direct
    EntryType string `json:"entryType"`


    // ExitReason - The exit reason for this flow, e.g. DISCONNECT
    ExitReason string `json:"exitReason"`


    // FlowId - The unique identifier of this flow
    FlowId string `json:"flowId"`


    // FlowName - The name of this flow at the time of flow execution
    FlowName string `json:"flowName"`


    // FlowType - The type of this flow
    FlowType string `json:"flowType"`


    // FlowVersion - The version of this flow
    FlowVersion string `json:"flowVersion"`


    // IssuedCallback - Flag indicating whether the flow issued a callback
    IssuedCallback bool `json:"issuedCallback"`


    // RecognitionFailureReason - The recognition failure reason causing to exit/disconnect
    RecognitionFailureReason string `json:"recognitionFailureReason"`


    // StartingLanguage - Flow starting language, e.g. en-us
    StartingLanguage string `json:"startingLanguage"`


    // TransferTargetAddress - The address of a flow transfer target, e.g. a phone number, an email address, or a queueId
    TransferTargetAddress string `json:"transferTargetAddress"`


    // TransferTargetName - The name of a flow transfer target
    TransferTargetName string `json:"transferTargetName"`


    // TransferType - The type of transfer for flows that ended with a transfer
    TransferType string `json:"transferType"`


    // Outcomes - Flow outcomes
    Outcomes []Analyticsflowoutcome `json:"outcomes"`

}

// String returns a JSON representation of the model
func (o *Analyticsflow) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Outcomes = []Analyticsflowoutcome{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticsflow) MarshalJSON() ([]byte, error) {
    type Alias Analyticsflow

    if AnalyticsflowMarshalled {
        return []byte("{}"), nil
    }
    AnalyticsflowMarshalled = true

    return json.Marshal(&struct {
        
        EndingLanguage string `json:"endingLanguage"`
        
        EntryReason string `json:"entryReason"`
        
        EntryType string `json:"entryType"`
        
        ExitReason string `json:"exitReason"`
        
        FlowId string `json:"flowId"`
        
        FlowName string `json:"flowName"`
        
        FlowType string `json:"flowType"`
        
        FlowVersion string `json:"flowVersion"`
        
        IssuedCallback bool `json:"issuedCallback"`
        
        RecognitionFailureReason string `json:"recognitionFailureReason"`
        
        StartingLanguage string `json:"startingLanguage"`
        
        TransferTargetAddress string `json:"transferTargetAddress"`
        
        TransferTargetName string `json:"transferTargetName"`
        
        TransferType string `json:"transferType"`
        
        Outcomes []Analyticsflowoutcome `json:"outcomes"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        
        Outcomes: []Analyticsflowoutcome{{}},
        

        Alias: (*Alias)(u),
    })
}

