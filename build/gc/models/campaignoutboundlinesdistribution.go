package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignoutboundlinesdistributionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignoutboundlinesdistributionDud struct { 
    


    


    


    


    


    


    

}

// Campaignoutboundlinesdistribution - Lines distribution information or Campaign's Edge Group or Site
type Campaignoutboundlinesdistribution struct { 
    // Campaign - The Campaign for which dialing group distribution information was requested
    Campaign Addressableentityref `json:"campaign"`


    // MaxOutboundLineCount - Maximum outbound calls that can be placed for Campaign's Edge Group or Site
    MaxOutboundLineCount int `json:"maxOutboundLineCount"`


    // MaxLineUtilization - Maximum ratio of dialer calls to Campaign's Edge Group or Site capacity
    MaxLineUtilization float32 `json:"maxLineUtilization"`


    // AvailableOutboundLines - Number of available outbound lines in Campaign's Edge Group or Site
    AvailableOutboundLines int `json:"availableOutboundLines"`


    // ReservedLines - Number of reserved outbound lines in Campaign's Edge Group or Site
    ReservedLines int `json:"reservedLines"`


    // CampaignsWithReservedLines - Information about campaigns with reserving lines in Campaign's Edge Group or Site
    CampaignsWithReservedLines []Campaignoutboundlinesreservation `json:"campaignsWithReservedLines"`


    // CampaignsWithDynamicallyAllocatedLines - Information about campaigns using dynamic lines allocation in Campaign's Edge Group or Site
    CampaignsWithDynamicallyAllocatedLines []Campaignoutboundlinesallocation `json:"campaignsWithDynamicallyAllocatedLines"`

}

// String returns a JSON representation of the model
func (o *Campaignoutboundlinesdistribution) String() string {
    
    
    
    
    
     o.CampaignsWithReservedLines = []Campaignoutboundlinesreservation{{}} 
     o.CampaignsWithDynamicallyAllocatedLines = []Campaignoutboundlinesallocation{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignoutboundlinesdistribution) MarshalJSON() ([]byte, error) {
    type Alias Campaignoutboundlinesdistribution

    if CampaignoutboundlinesdistributionMarshalled {
        return []byte("{}"), nil
    }
    CampaignoutboundlinesdistributionMarshalled = true

    return json.Marshal(&struct {
        
        Campaign Addressableentityref `json:"campaign"`
        
        MaxOutboundLineCount int `json:"maxOutboundLineCount"`
        
        MaxLineUtilization float32 `json:"maxLineUtilization"`
        
        AvailableOutboundLines int `json:"availableOutboundLines"`
        
        ReservedLines int `json:"reservedLines"`
        
        CampaignsWithReservedLines []Campaignoutboundlinesreservation `json:"campaignsWithReservedLines"`
        
        CampaignsWithDynamicallyAllocatedLines []Campaignoutboundlinesallocation `json:"campaignsWithDynamicallyAllocatedLines"`
        *Alias
    }{

        


        


        


        


        


        
        CampaignsWithReservedLines: []Campaignoutboundlinesreservation{{}},
        


        
        CampaignsWithDynamicallyAllocatedLines: []Campaignoutboundlinesallocation{{}},
        

        Alias: (*Alias)(u),
    })
}

