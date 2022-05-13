package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    VideoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type VideoDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Video
type Video struct { 
    // State - The connection state of this communication.
    State string `json:"state"`


    // Id - A globally unique identifier for this communication.
    Id string `json:"id"`


    // Context - The room id context (xmpp jid) for the conference session.
    Context string `json:"context"`


    // AudioMuted - Indicates whether this participant has muted their outgoing audio.
    AudioMuted bool `json:"audioMuted"`


    // VideoMuted - Indicates whether this participant has muted/paused their outgoing video.
    VideoMuted bool `json:"videoMuted"`


    // SharingScreen - Indicates whether this participant is sharing their screen to the session.
    SharingScreen bool `json:"sharingScreen"`


    // PeerCount - The number of peer participants from the perspective of the participant in the conference.
    PeerCount int `json:"peerCount"`


    // DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
    DisconnectType string `json:"disconnectType"`


    // StartAlertingTime - The timestamp the communication has when it is first put into an alerting state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartAlertingTime time.Time `json:"startAlertingTime"`


    // ConnectedTime - The timestamp when this communication was connected in the cloud clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ConnectedTime time.Time `json:"connectedTime"`


    // DisconnectedTime - The timestamp when this communication disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DisconnectedTime time.Time `json:"disconnectedTime"`


    // Provider - The source provider for the video.
    Provider string `json:"provider"`


    // PeerId - The id of the peer communication corresponding to a matching leg for this communication.
    PeerId string `json:"peerId"`


    // Msids - List of media stream ids
    Msids []string `json:"msids"`


    // Self - Address and name data for a call endpoint.
    Self Address `json:"self"`


    // Wrapup - Call wrap up or disposition data.
    Wrapup Wrapup `json:"wrapup"`


    // AfterCallWork - After-call work for the communication.
    AfterCallWork Aftercallwork `json:"afterCallWork"`


    // AfterCallWorkRequired - Indicates if after-call work is required for a communication. Only used when the ACW Setting is Agent Requested.
    AfterCallWorkRequired bool `json:"afterCallWorkRequired"`

}

// String returns a JSON representation of the model
func (o *Video) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Msids = []string{""} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Video) MarshalJSON() ([]byte, error) {
    type Alias Video

    if VideoMarshalled {
        return []byte("{}"), nil
    }
    VideoMarshalled = true

    return json.Marshal(&struct {
        
        State string `json:"state"`
        
        Id string `json:"id"`
        
        Context string `json:"context"`
        
        AudioMuted bool `json:"audioMuted"`
        
        VideoMuted bool `json:"videoMuted"`
        
        SharingScreen bool `json:"sharingScreen"`
        
        PeerCount int `json:"peerCount"`
        
        DisconnectType string `json:"disconnectType"`
        
        StartAlertingTime time.Time `json:"startAlertingTime"`
        
        ConnectedTime time.Time `json:"connectedTime"`
        
        DisconnectedTime time.Time `json:"disconnectedTime"`
        
        Provider string `json:"provider"`
        
        PeerId string `json:"peerId"`
        
        Msids []string `json:"msids"`
        
        Self Address `json:"self"`
        
        Wrapup Wrapup `json:"wrapup"`
        
        AfterCallWork Aftercallwork `json:"afterCallWork"`
        
        AfterCallWorkRequired bool `json:"afterCallWorkRequired"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        
        Msids: []string{""},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

