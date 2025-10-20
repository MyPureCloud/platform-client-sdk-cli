package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignperformancedataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignperformancedataDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Campaignperformancedata
type Campaignperformancedata struct { 
    // Campaign - Identifier of the campaign
    Campaign Domainentityref `json:"campaign"`


    // Division - The division the campaign belongs to
    Division Addressableentityref `json:"division"`


    // ContactRate - Information regarding the campaign's connect rate
    ContactRate Campaignperformancedatacontactrate `json:"contactRate"`


    // IdleAgents - Number of available agents not currently being utilized
    IdleAgents int `json:"idleAgents"`


    // EffectiveIdleAgents - Number of effective available agents not currently being utilized
    EffectiveIdleAgents float32 `json:"effectiveIdleAgents"`


    // AdjustedCallsPerAgent - Calls per agent adjusted by pace
    AdjustedCallsPerAgent float32 `json:"adjustedCallsPerAgent"`


    // OutstandingCalls - Number of campaign calls currently ongoing
    OutstandingCalls int `json:"outstandingCalls"`


    // ScheduledCalls - Number of campaign calls currently scheduled
    ScheduledCalls int `json:"scheduledCalls"`


    // RightPartyContactsCount - Information on the campaign's number of Right Party Contacts
    RightPartyContactsCount int `json:"rightPartyContactsCount"`


    // CampaignStatus - The status of the campaign
    CampaignStatus string `json:"campaignStatus"`


    // DialingMode - The strategy this Campaign will use for dialing
    DialingMode string `json:"dialingMode"`


    // Progress - Information on the campaign's progress
    Progress Campaignperformancedataprogress `json:"progress"`


    // LinesUtilization - Information on the campaign's lines utilization
    LinesUtilization Campaignlinesutilization `json:"linesUtilization"`


    // BusinessCategoryMetrics - Information on the campaign's business category metrics
    BusinessCategoryMetrics Campaignbusinesscategorymetrics `json:"businessCategoryMetrics"`

}

// String returns a JSON representation of the model
func (o *Campaignperformancedata) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaignperformancedata) MarshalJSON() ([]byte, error) {
    type Alias Campaignperformancedata

    if CampaignperformancedataMarshalled {
        return []byte("{}"), nil
    }
    CampaignperformancedataMarshalled = true

    return json.Marshal(&struct {
        
        Campaign Domainentityref `json:"campaign"`
        
        Division Addressableentityref `json:"division"`
        
        ContactRate Campaignperformancedatacontactrate `json:"contactRate"`
        
        IdleAgents int `json:"idleAgents"`
        
        EffectiveIdleAgents float32 `json:"effectiveIdleAgents"`
        
        AdjustedCallsPerAgent float32 `json:"adjustedCallsPerAgent"`
        
        OutstandingCalls int `json:"outstandingCalls"`
        
        ScheduledCalls int `json:"scheduledCalls"`
        
        RightPartyContactsCount int `json:"rightPartyContactsCount"`
        
        CampaignStatus string `json:"campaignStatus"`
        
        DialingMode string `json:"dialingMode"`
        
        Progress Campaignperformancedataprogress `json:"progress"`
        
        LinesUtilization Campaignlinesutilization `json:"linesUtilization"`
        
        BusinessCategoryMetrics Campaignbusinesscategorymetrics `json:"businessCategoryMetrics"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

