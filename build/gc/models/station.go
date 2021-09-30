package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StationDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    WebRtcMediaDscp int `json:"webRtcMediaDscp"`


    WebRtcPersistentEnabled bool `json:"webRtcPersistentEnabled"`


    WebRtcForceTurn bool `json:"webRtcForceTurn"`


    SelfUri string `json:"selfUri"`

}

// Station
type Station struct { 
    


    // Name
    Name string `json:"name"`


    // Description
    Description string `json:"description"`


    // Status
    Status string `json:"status"`


    // UserId - The Id of the user currently logged in and associated with the station.
    UserId string `json:"userId"`


    // WebRtcUserId - The Id of the user configured for the station if it is of type inin_webrtc_softphone. Empty if station type is not inin_webrtc_softphone.
    WebRtcUserId string `json:"webRtcUserId"`


    // PrimaryEdge
    PrimaryEdge Domainentityref `json:"primaryEdge"`


    // SecondaryEdge
    SecondaryEdge Domainentityref `json:"secondaryEdge"`


    // VarType
    VarType string `json:"type"`


    // LineAppearanceId
    LineAppearanceId string `json:"lineAppearanceId"`


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Station) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Station) MarshalJSON() ([]byte, error) {
    type Alias Station

    if StationMarshalled {
        return []byte("{}"), nil
    }
    StationMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Status string `json:"status"`
        
        UserId string `json:"userId"`
        
        WebRtcUserId string `json:"webRtcUserId"`
        
        PrimaryEdge Domainentityref `json:"primaryEdge"`
        
        SecondaryEdge Domainentityref `json:"secondaryEdge"`
        
        VarType string `json:"type"`
        
        LineAppearanceId string `json:"lineAppearanceId"`
        
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

