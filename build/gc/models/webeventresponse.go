package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebeventresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebeventresponseDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Webeventresponse
type Webeventresponse struct { 
    // CustomerId - Identifier of the customer in the source of the event.
    CustomerId string `json:"customerId"`


    // EventName - Represents the action the customer performed. Event types are created for each unique event name and can be faceted on in segment and outcome conditions. A valid event name must only contain alphanumeric characters and underscores. A good event name is typically an object followed by the action performed in past tense, e.g. page_viewed, order_completed, user_registered.
    EventName string `json:"eventName"`


    // CustomerIdType - Type of identifier for the customer ID (e.g., cookie).
    CustomerIdType string `json:"customerIdType"`


    // Page - The webpage where the user interaction occurred.
    Page Responsepage `json:"page"`


    // UserAgentString - HTTP User-Agent string (see https://tools.ietf.org/html/rfc1945#section-10.15).
    UserAgentString string `json:"userAgentString"`


    // Browser - Customer's browser.
    Browser Webeventbrowser `json:"browser"`


    // Device - Customer's device.
    Device Webeventdevice `json:"device"`


    // SearchQuery - Represents the keywords in a customer search query.
    SearchQuery string `json:"searchQuery"`


    // IpOrganization - Customer's IP-based organization or ISP name.
    IpOrganization string `json:"ipOrganization"`


    // Geolocation - Customer's geolocation.
    Geolocation Journeygeolocation `json:"geolocation"`


    // MktCampaign - Urchin Tracking Module (UTM) parameters used to track the effectiveness of online marketing campaigns.
    MktCampaign Journeycampaign `json:"mktCampaign"`


    // Session - The session that the event belongs to.
    Session Webeventresponsesession `json:"session"`


    // Referrer - Identifies the web page that originally generated the request for the current page being viewed.
    Referrer Referrer `json:"referrer"`


    // Attributes - User-defined attributes associated with a particular event. These attributes provide additional context about the event. For example, items_in_cart or subscription_level.
    Attributes map[string]Customeventattribute `json:"attributes"`


    // Traits - Traits are attributes intrinsic to the customer that may be sent in selected events, (e.g. email, givenName, cellPhone). Traits are used to collect information for identity resolution. For example, the same person might be using an application on different devices which might create two sessions with different customerIds. Additional information can be provided as traits to help link those two sessions and customers to a single external contact through common identifiers that were submitted via a form fill, message, or other input in both sessions.
    Traits map[string]Customeventattribute `json:"traits"`


    // Authenticated - Indicates whether the event was produced during an authenticated session.
    Authenticated bool `json:"authenticated"`


    // CreatedDate - UTC timestamp indicating when the event actually took place, events older than an hour will be rejected. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreatedDate time.Time `json:"createdDate"`

}

// String returns a JSON representation of the model
func (o *Webeventresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Attributes = map[string]Customeventattribute{"": {}} 
     o.Traits = map[string]Customeventattribute{"": {}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webeventresponse) MarshalJSON() ([]byte, error) {
    type Alias Webeventresponse

    if WebeventresponseMarshalled {
        return []byte("{}"), nil
    }
    WebeventresponseMarshalled = true

    return json.Marshal(&struct {
        
        CustomerId string `json:"customerId"`
        
        EventName string `json:"eventName"`
        
        CustomerIdType string `json:"customerIdType"`
        
        Page Responsepage `json:"page"`
        
        UserAgentString string `json:"userAgentString"`
        
        Browser Webeventbrowser `json:"browser"`
        
        Device Webeventdevice `json:"device"`
        
        SearchQuery string `json:"searchQuery"`
        
        IpOrganization string `json:"ipOrganization"`
        
        Geolocation Journeygeolocation `json:"geolocation"`
        
        MktCampaign Journeycampaign `json:"mktCampaign"`
        
        Session Webeventresponsesession `json:"session"`
        
        Referrer Referrer `json:"referrer"`
        
        Attributes map[string]Customeventattribute `json:"attributes"`
        
        Traits map[string]Customeventattribute `json:"traits"`
        
        Authenticated bool `json:"authenticated"`
        
        CreatedDate time.Time `json:"createdDate"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        
        Attributes: map[string]Customeventattribute{"": {}},
        


        
        Traits: map[string]Customeventattribute{"": {}},
        


        


        

        Alias: (*Alias)(u),
    })
}

