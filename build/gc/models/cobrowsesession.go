package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CobrowsesessionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CobrowsesessionDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Cobrowsesession
type Cobrowsesession struct { 
    // State - The connection state of this communication.
    State string `json:"state"`


    // Id - A globally unique identifier for this communication.
    Id string `json:"id"`


    // DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
    DisconnectType string `json:"disconnectType"`


    // Self - Address and name data for a call endpoint.
    Self Address `json:"self"`


    // CobrowseSessionId - The co-browse session ID.
    CobrowseSessionId string `json:"cobrowseSessionId"`


    // CobrowseRole - This value identifies the role of the co-browse client within the co-browse session (a client is a sharer or a viewer).
    CobrowseRole string `json:"cobrowseRole"`


    // Controlling - ID of co-browse participants for which this client has been granted control (list is empty if this client cannot control any shared pages).
    Controlling []string `json:"controlling"`


    // ViewerUrl - The URL that can be used to open co-browse session in web browser.
    ViewerUrl string `json:"viewerUrl"`


    // ProviderEventTime - The time when the provider event which triggered this conversation update happened in the corrected provider clock (milliseconds since 1970-01-01 00:00:00 UTC). Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ProviderEventTime time.Time `json:"providerEventTime"`


    // StartAlertingTime - The timestamp the communication has when it is first put into an alerting state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartAlertingTime time.Time `json:"startAlertingTime"`


    // ConnectedTime - The timestamp when this communication was connected in the cloud clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ConnectedTime time.Time `json:"connectedTime"`


    // DisconnectedTime - The timestamp when this communication disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DisconnectedTime time.Time `json:"disconnectedTime"`


    // Provider - The source provider for the co-browse session.
    Provider string `json:"provider"`


    // PeerId - The id of the peer communication corresponding to a matching leg for this communication.
    PeerId string `json:"peerId"`


    // Segments - The time line of the participant's call, divided into activity segments.
    Segments []Segment `json:"segments"`


    // Wrapup - Call wrap up or disposition data.
    Wrapup Wrapup `json:"wrapup"`


    // AfterCallWork - After-call work for the communication.
    AfterCallWork Aftercallwork `json:"afterCallWork"`


    // AfterCallWorkRequired - Indicates if after-call work is required for a communication. Only used when the ACW Setting is Agent Requested.
    AfterCallWorkRequired bool `json:"afterCallWorkRequired"`

}

// String returns a JSON representation of the model
func (o *Cobrowsesession) String() string {
    
    
    
    
    
    
     o.Controlling = []string{""} 
    
    
    
    
    
    
    
     o.Segments = []Segment{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Cobrowsesession) MarshalJSON() ([]byte, error) {
    type Alias Cobrowsesession

    if CobrowsesessionMarshalled {
        return []byte("{}"), nil
    }
    CobrowsesessionMarshalled = true

    return json.Marshal(&struct {
        
        State string `json:"state"`
        
        Id string `json:"id"`
        
        DisconnectType string `json:"disconnectType"`
        
        Self Address `json:"self"`
        
        CobrowseSessionId string `json:"cobrowseSessionId"`
        
        CobrowseRole string `json:"cobrowseRole"`
        
        Controlling []string `json:"controlling"`
        
        ViewerUrl string `json:"viewerUrl"`
        
        ProviderEventTime time.Time `json:"providerEventTime"`
        
        StartAlertingTime time.Time `json:"startAlertingTime"`
        
        ConnectedTime time.Time `json:"connectedTime"`
        
        DisconnectedTime time.Time `json:"disconnectedTime"`
        
        Provider string `json:"provider"`
        
        PeerId string `json:"peerId"`
        
        Segments []Segment `json:"segments"`
        
        Wrapup Wrapup `json:"wrapup"`
        
        AfterCallWork Aftercallwork `json:"afterCallWork"`
        
        AfterCallWorkRequired bool `json:"afterCallWorkRequired"`
        *Alias
    }{

        


        


        


        


        


        


        
        Controlling: []string{""},
        


        


        


        


        


        


        


        


        
        Segments: []Segment{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

