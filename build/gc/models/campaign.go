package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CampaignMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CampaignDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    Errors []Resterrordetail `json:"errors"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Campaign
type Campaign struct { 
    


    // Name - The name of the Campaign.
    Name string `json:"name"`


    


    


    // Version - Required for updates, must match the version number of the most recent update
    Version int `json:"version"`


    // ContactList - The ContactList for this Campaign to dial.
    ContactList Domainentityref `json:"contactList"`


    // Queue - The Queue for this Campaign to route calls to. Required for all dialing modes except agentless.
    Queue Domainentityref `json:"queue"`


    // DialingMode - The strategy this Campaign will use for dialing.
    DialingMode string `json:"dialingMode"`


    // Script - The Script to be displayed to agents that are handling outbound calls. Required for all dialing modes except agentless.
    Script Domainentityref `json:"script"`


    // EdgeGroup - The EdgeGroup that will place the calls. Required for all dialing modes except preview.
    EdgeGroup Domainentityref `json:"edgeGroup"`


    // Site - The identifier of the site to be used for dialing; can be set in place of an edge group.
    Site Domainentityref `json:"site"`


    // CampaignStatus - The current status of the Campaign. A Campaign may be turned 'on' or 'off'. Required for updates.
    CampaignStatus string `json:"campaignStatus"`


    // PhoneColumns - The ContactPhoneNumberColumns on the ContactList that this Campaign should dial.
    PhoneColumns []Phonecolumn `json:"phoneColumns"`


    // AbandonRate - The targeted compliance abandon rate percentage. Required for power and predictive campaigns.
    AbandonRate float64 `json:"abandonRate"`


    // DncLists - DncLists for this Campaign to check before placing a call.
    DncLists []Domainentityref `json:"dncLists"`


    // CallableTimeSet - The callable time set for this campaign to check before placing a call.
    CallableTimeSet Domainentityref `json:"callableTimeSet"`


    // CallAnalysisResponseSet - The call analysis response set to handle call analysis results from the edge. Required for all dialing modes except preview.
    CallAnalysisResponseSet Domainentityref `json:"callAnalysisResponseSet"`


    


    // CallerName - The caller id name to be displayed on the outbound call.
    CallerName string `json:"callerName"`


    // CallerAddress - The caller id phone number to be displayed on the outbound call.
    CallerAddress string `json:"callerAddress"`


    // OutboundLineCount - The number of outbound lines to be concurrently dialed. Only applicable to non-preview campaigns; only required for agentless.
    OutboundLineCount int `json:"outboundLineCount"`


    // RuleSets - Rule sets to be applied while this campaign is dialing.
    RuleSets []Domainentityref `json:"ruleSets"`


    // SkipPreviewDisabled - Whether or not agents can skip previews without placing a call. Only applicable for preview campaigns.
    SkipPreviewDisabled bool `json:"skipPreviewDisabled"`


    // PreviewTimeOutSeconds - The number of seconds before a call will be automatically placed on a preview. A value of 0 indicates no automatic placement of calls. Only applicable to preview campaigns.
    PreviewTimeOutSeconds int `json:"previewTimeOutSeconds"`


    // AlwaysRunning - Indicates (when true) that the campaign will remain on after contacts are depleted, allowing additional contacts to be appended/added to the contact list and processed by the still-running campaign. The campaign can still be turned off manually.
    AlwaysRunning bool `json:"alwaysRunning"`


    // ContactSort - The order in which to sort contacts for dialing, based on a column.
    ContactSort Contactsort `json:"contactSort"`


    // ContactSorts - The order in which to sort contacts for dialing, based on up to four columns.
    ContactSorts []Contactsort `json:"contactSorts"`


    // NoAnswerTimeout - How long to wait before dispositioning a call as 'no-answer'. Default 30 seconds. Only applicable to non-preview campaigns.
    NoAnswerTimeout int `json:"noAnswerTimeout"`


    // CallAnalysisLanguage - The language the edge will use to analyze the call.
    CallAnalysisLanguage string `json:"callAnalysisLanguage"`


    // Priority - The priority of this campaign relative to other campaigns that are running on the same queue. 5 is the highest priority, 1 the lowest.
    Priority int `json:"priority"`


    // ContactListFilters - Filter to apply to the contact list before dialing. Currently a campaign can only have one filter applied.
    ContactListFilters []Domainentityref `json:"contactListFilters"`


    // Division - The division this campaign belongs to.
    Division Domainentityref `json:"division"`


    // DynamicContactQueueingSettings - Settings for dynamic queueing of contacts.
    DynamicContactQueueingSettings Dynamiccontactqueueingsettings `json:"dynamicContactQueueingSettings"`


    

}

// String returns a JSON representation of the model
func (o *Campaign) String() string {
    
    
    
    
    
    
    
    
    
     o.PhoneColumns = []Phonecolumn{{}} 
    
     o.DncLists = []Domainentityref{{}} 
    
    
    
    
    
     o.RuleSets = []Domainentityref{{}} 
    
    
    
    
     o.ContactSorts = []Contactsort{{}} 
    
    
    
     o.ContactListFilters = []Domainentityref{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Campaign) MarshalJSON() ([]byte, error) {
    type Alias Campaign

    if CampaignMarshalled {
        return []byte("{}"), nil
    }
    CampaignMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Version int `json:"version"`
        
        ContactList Domainentityref `json:"contactList"`
        
        Queue Domainentityref `json:"queue"`
        
        DialingMode string `json:"dialingMode"`
        
        Script Domainentityref `json:"script"`
        
        EdgeGroup Domainentityref `json:"edgeGroup"`
        
        Site Domainentityref `json:"site"`
        
        CampaignStatus string `json:"campaignStatus"`
        
        PhoneColumns []Phonecolumn `json:"phoneColumns"`
        
        AbandonRate float64 `json:"abandonRate"`
        
        DncLists []Domainentityref `json:"dncLists"`
        
        CallableTimeSet Domainentityref `json:"callableTimeSet"`
        
        CallAnalysisResponseSet Domainentityref `json:"callAnalysisResponseSet"`
        
        CallerName string `json:"callerName"`
        
        CallerAddress string `json:"callerAddress"`
        
        OutboundLineCount int `json:"outboundLineCount"`
        
        RuleSets []Domainentityref `json:"ruleSets"`
        
        SkipPreviewDisabled bool `json:"skipPreviewDisabled"`
        
        PreviewTimeOutSeconds int `json:"previewTimeOutSeconds"`
        
        AlwaysRunning bool `json:"alwaysRunning"`
        
        ContactSort Contactsort `json:"contactSort"`
        
        ContactSorts []Contactsort `json:"contactSorts"`
        
        NoAnswerTimeout int `json:"noAnswerTimeout"`
        
        CallAnalysisLanguage string `json:"callAnalysisLanguage"`
        
        Priority int `json:"priority"`
        
        ContactListFilters []Domainentityref `json:"contactListFilters"`
        
        Division Domainentityref `json:"division"`
        
        DynamicContactQueueingSettings Dynamiccontactqueueingsettings `json:"dynamicContactQueueingSettings"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        
        PhoneColumns: []Phonecolumn{{}},
        


        


        
        DncLists: []Domainentityref{{}},
        


        


        


        


        


        


        


        
        RuleSets: []Domainentityref{{}},
        


        


        


        


        


        
        ContactSorts: []Contactsort{{}},
        


        


        


        


        
        ContactListFilters: []Domainentityref{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

