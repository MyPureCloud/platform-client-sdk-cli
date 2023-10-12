package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebactioneventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebactioneventDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Webactionevent
type Webactionevent struct { 
    // Action - The action that triggered the event.
    Action Eventaction `json:"action"`


    // ActionMap - The action map that triggered the action.
    ActionMap Actioneventactionmap `json:"actionMap"`


    // ActionTarget - The target for engagement actions.
    ActionTarget Addressableentityref `json:"actionTarget"`


    // TimeToDisposition - Milliseconds elapsed until the action is disposed.
    TimeToDisposition int `json:"timeToDisposition"`


    // ErrorCode - Code of the error returned when the action fails.
    ErrorCode string `json:"errorCode"`


    // ErrorMessage - Message of the error returned when the action fails.
    ErrorMessage string `json:"errorMessage"`


    // UserAgentString - HTTP User-Agent string (see https://tools.ietf.org/html/rfc1945#section-10.15).
    UserAgentString string `json:"userAgentString"`


    // Browser - Customer's browser.
    Browser Browser `json:"browser"`


    // Device - Customer's device.
    Device Device `json:"device"`


    // Geolocation - Customer's geolocation.
    Geolocation Journeygeolocation `json:"geolocation"`


    // IpAddress - Visitor's IP address.
    IpAddress string `json:"ipAddress"`


    // IpOrganization - Visitor's IP-based organization or ISP name.
    IpOrganization string `json:"ipOrganization"`


    // MktCampaign - Marketing / traffic source information.
    MktCampaign Journeycampaign `json:"mktCampaign"`


    // VisitReferrer - Visit's referrer.
    VisitReferrer Referrer `json:"visitReferrer"`

}

// String returns a JSON representation of the model
func (o *Webactionevent) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webactionevent) MarshalJSON() ([]byte, error) {
    type Alias Webactionevent

    if WebactioneventMarshalled {
        return []byte("{}"), nil
    }
    WebactioneventMarshalled = true

    return json.Marshal(&struct {
        
        Action Eventaction `json:"action"`
        
        ActionMap Actioneventactionmap `json:"actionMap"`
        
        ActionTarget Addressableentityref `json:"actionTarget"`
        
        TimeToDisposition int `json:"timeToDisposition"`
        
        ErrorCode string `json:"errorCode"`
        
        ErrorMessage string `json:"errorMessage"`
        
        UserAgentString string `json:"userAgentString"`
        
        Browser Browser `json:"browser"`
        
        Device Device `json:"device"`
        
        Geolocation Journeygeolocation `json:"geolocation"`
        
        IpAddress string `json:"ipAddress"`
        
        IpOrganization string `json:"ipOrganization"`
        
        MktCampaign Journeycampaign `json:"mktCampaign"`
        
        VisitReferrer Referrer `json:"visitReferrer"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

