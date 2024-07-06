package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignpatchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignpatchrequestDud struct { 
    


    


    


    


    

}

// Campaignpatchrequest - Campaign patch request
type Campaignpatchrequest struct { 
    // OutboundLineCount - The number of outbound lines to be concurrently dialed.
    OutboundLineCount int `json:"outboundLineCount"`


    // AbandonRate - The targeted compliance abandon rate percentage
    AbandonRate float32 `json:"abandonRate"`


    // MaxCallsPerAgent - The maximum number of calls that can be placed per agent on this campaign
    MaxCallsPerAgent float32 `json:"maxCallsPerAgent"`


    // DynamicLineBalancingSettings - Dynamic line balancing settings
    DynamicLineBalancingSettings Dynamiclinebalancingsettingspatchrequest `json:"dynamicLineBalancingSettings"`


    // Queue - The Queue for this Campaign to route calls to.
    Queue Addressableentityref `json:"queue"`

}

// String returns a JSON representation of the model
func (o *Campaignpatchrequest) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignpatchrequest) MarshalJSON() ([]byte, error) {
    type Alias Campaignpatchrequest

    if CampaignpatchrequestMarshalled {
        return []byte("{}"), nil
    }
    CampaignpatchrequestMarshalled = true

    return json.Marshal(&struct {
        
        OutboundLineCount int `json:"outboundLineCount"`
        
        AbandonRate float32 `json:"abandonRate"`
        
        MaxCallsPerAgent float32 `json:"maxCallsPerAgent"`
        
        DynamicLineBalancingSettings Dynamiclinebalancingsettingspatchrequest `json:"dynamicLineBalancingSettings"`
        
        Queue Addressableentityref `json:"queue"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

