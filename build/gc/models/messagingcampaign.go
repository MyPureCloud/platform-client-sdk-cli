package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagingcampaignMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagingcampaignDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Messagingcampaign
type Messagingcampaign struct { 
    


    // Name
    Name string `json:"name"`


    


    


    // Version - Required for updates, must match the version number of the most recent update
    Version int `json:"version"`


    // Division - The division this entity belongs to.
    Division Domainentityref `json:"division"`


    // CampaignStatus - The current status of the messaging campaign. A messaging campaign may be turned 'on' or 'off'.
    CampaignStatus string `json:"campaignStatus"`


    // CallableTimeSet - The callable time set for this messaging campaign.
    CallableTimeSet Domainentityref `json:"callableTimeSet"`


    // ContactList - The contact list that this messaging campaign will send messages for.
    ContactList Domainentityref `json:"contactList"`


    // DncLists - The dnc lists to check before sending a message for this messaging campaign.
    DncLists []Domainentityref `json:"dncLists"`


    // AlwaysRunning - Whether this messaging campaign is always running
    AlwaysRunning bool `json:"alwaysRunning"`


    // ContactSorts - The order in which to sort contacts for dialing, based on up to four columns.
    ContactSorts []Contactsort `json:"contactSorts"`


    // MessagesPerMinute - How many messages this messaging campaign will send per minute.
    MessagesPerMinute int `json:"messagesPerMinute"`


    // RuleSets - Rule Sets to be applied while this campaign is sending messages
    RuleSets []Domainentityref `json:"ruleSets"`


    // ContactListFilters - The contact list filter to check before sending a message for this messaging campaign.
    ContactListFilters []Domainentityref `json:"contactListFilters"`


    // Errors - A list of current error conditions associated with this messaging campaign.
    Errors []Resterrordetail `json:"errors"`


    // DynamicContactQueueingSettings - Indicates (when true) that the campaign supports dynamic queueing of the contact list at the time of a request for contacts.
    DynamicContactQueueingSettings Dynamiccontactqueueingsettings `json:"dynamicContactQueueingSettings"`


    // EmailConfig - Configuration for this messaging campaign to send Email messages.
    EmailConfig Emailconfig `json:"emailConfig"`


    // SmsConfig - Configuration for this messaging campaign to send SMS messages.
    SmsConfig Smsconfig `json:"smsConfig"`


    // WhatsAppConfig - Configuration for this messaging campaign to send WhatsApp messages.
    WhatsAppConfig Whatsappconfig `json:"whatsAppConfig"`


    

}

// String returns a JSON representation of the model
func (o *Messagingcampaign) String() string {
    
    
    
    
    
    
     o.DncLists = []Domainentityref{{}} 
    
     o.ContactSorts = []Contactsort{{}} 
    
     o.RuleSets = []Domainentityref{{}} 
     o.ContactListFilters = []Domainentityref{{}} 
     o.Errors = []Resterrordetail{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagingcampaign) MarshalJSON() ([]byte, error) {
    type Alias Messagingcampaign

    if MessagingcampaignMarshalled {
        return []byte("{}"), nil
    }
    MessagingcampaignMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Version int `json:"version"`
        
        Division Domainentityref `json:"division"`
        
        CampaignStatus string `json:"campaignStatus"`
        
        CallableTimeSet Domainentityref `json:"callableTimeSet"`
        
        ContactList Domainentityref `json:"contactList"`
        
        DncLists []Domainentityref `json:"dncLists"`
        
        AlwaysRunning bool `json:"alwaysRunning"`
        
        ContactSorts []Contactsort `json:"contactSorts"`
        
        MessagesPerMinute int `json:"messagesPerMinute"`
        
        RuleSets []Domainentityref `json:"ruleSets"`
        
        ContactListFilters []Domainentityref `json:"contactListFilters"`
        
        Errors []Resterrordetail `json:"errors"`
        
        DynamicContactQueueingSettings Dynamiccontactqueueingsettings `json:"dynamicContactQueueingSettings"`
        
        EmailConfig Emailconfig `json:"emailConfig"`
        
        SmsConfig Smsconfig `json:"smsConfig"`
        
        WhatsAppConfig Whatsappconfig `json:"whatsAppConfig"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        
        DncLists: []Domainentityref{{}},
        


        


        
        ContactSorts: []Contactsort{{}},
        


        


        
        RuleSets: []Domainentityref{{}},
        


        
        ContactListFilters: []Domainentityref{{}},
        


        
        Errors: []Resterrordetail{{}},
        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

