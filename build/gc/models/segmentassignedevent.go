package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SegmentassignedeventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SegmentassignedeventDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Segmentassignedevent
type Segmentassignedevent struct { 
    // Segment - The segment that was matched.
    Segment Segmentassignedeventsegment `json:"segment"`


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


    // VisitCreatedDate - When visit was created (e.g. timestamp of the first event in visit). Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    VisitCreatedDate time.Time `json:"visitCreatedDate"`

}

// String returns a JSON representation of the model
func (o *Segmentassignedevent) String() string {
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Segmentassignedevent) MarshalJSON() ([]byte, error) {
    type Alias Segmentassignedevent

    if SegmentassignedeventMarshalled {
        return []byte("{}"), nil
    }
    SegmentassignedeventMarshalled = true

    return json.Marshal(&struct {
        
        Segment Segmentassignedeventsegment `json:"segment"`
        
        UserAgentString string `json:"userAgentString"`
        
        Browser Browser `json:"browser"`
        
        Device Device `json:"device"`
        
        Geolocation Journeygeolocation `json:"geolocation"`
        
        IpAddress string `json:"ipAddress"`
        
        IpOrganization string `json:"ipOrganization"`
        
        MktCampaign Journeycampaign `json:"mktCampaign"`
        
        VisitReferrer Referrer `json:"visitReferrer"`
        
        VisitCreatedDate time.Time `json:"visitCreatedDate"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

