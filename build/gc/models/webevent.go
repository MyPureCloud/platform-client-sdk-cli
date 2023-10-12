package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebeventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebeventDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Webevent
type Webevent struct { 
    // EventName - Represents the action the customer performed. A good event name is typically an object followed by the action performed in past tense (e.g. page_viewed, order_completed, user_registered).
    EventName string `json:"eventName"`


    // TotalEventCount - The total count of events performed by the customer across all sessions.
    TotalEventCount int `json:"totalEventCount"`


    // TotalPageviewCount - The total count of pageviews performed by the customer across all sessions.
    TotalPageviewCount int `json:"totalPageviewCount"`


    // Page - The webpage where the user interaction occurred.
    Page Journeypage `json:"page"`


    // UserAgentString - HTTP User-Agent string (see https://tools.ietf.org/html/rfc1945#section-10.15).
    UserAgentString string `json:"userAgentString"`


    // Browser - Customer's browser.
    Browser Browser `json:"browser"`


    // Device - Customer's device.
    Device Device `json:"device"`


    // Geolocation - Customer's geolocation.
    Geolocation Journeygeolocation `json:"geolocation"`


    // IpAddress - Customer's IP address. May be null if the business configures the tracker to not collect IP addresses.
    IpAddress string `json:"ipAddress"`


    // IpOrganization - Customer's IP-based organization or ISP name.
    IpOrganization string `json:"ipOrganization"`


    // MktCampaign - Marketing / traffic source information.
    MktCampaign Journeycampaign `json:"mktCampaign"`


    // Referrer - Identifies the page URL that originally generated the request for the current page being viewed.
    Referrer Referrer `json:"referrer"`


    // Attributes - User-defined attributes associated with a particular event.
    Attributes map[string]Customeventattribute `json:"attributes"`


    // Traits - User-defined traits associated with a particular event.
    Traits map[string]Customeventattribute `json:"traits"`


    // SearchQuery - Represents the keywords in a customer search query.
    SearchQuery string `json:"searchQuery"`


    // Authenticated - Indicates whether the event was produced during an authenticated session.
    Authenticated bool `json:"authenticated"`

}

// String returns a JSON representation of the model
func (o *Webevent) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
     o.Attributes = map[string]Customeventattribute{"": {}} 
     o.Traits = map[string]Customeventattribute{"": {}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webevent) MarshalJSON() ([]byte, error) {
    type Alias Webevent

    if WebeventMarshalled {
        return []byte("{}"), nil
    }
    WebeventMarshalled = true

    return json.Marshal(&struct {
        
        EventName string `json:"eventName"`
        
        TotalEventCount int `json:"totalEventCount"`
        
        TotalPageviewCount int `json:"totalPageviewCount"`
        
        Page Journeypage `json:"page"`
        
        UserAgentString string `json:"userAgentString"`
        
        Browser Browser `json:"browser"`
        
        Device Device `json:"device"`
        
        Geolocation Journeygeolocation `json:"geolocation"`
        
        IpAddress string `json:"ipAddress"`
        
        IpOrganization string `json:"ipOrganization"`
        
        MktCampaign Journeycampaign `json:"mktCampaign"`
        
        Referrer Referrer `json:"referrer"`
        
        Attributes map[string]Customeventattribute `json:"attributes"`
        
        Traits map[string]Customeventattribute `json:"traits"`
        
        SearchQuery string `json:"searchQuery"`
        
        Authenticated bool `json:"authenticated"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        
        Attributes: map[string]Customeventattribute{"": {}},
        


        
        Traits: map[string]Customeventattribute{"": {}},
        


        


        

        Alias: (*Alias)(u),
    })
}

