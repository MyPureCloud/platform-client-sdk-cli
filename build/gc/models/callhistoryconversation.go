package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CallhistoryconversationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CallhistoryconversationDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Callhistoryconversation
type Callhistoryconversation struct { 
    


    // Name
    Name string `json:"name"`


    // Participants - The list of participants involved in the conversation.
    Participants []Callhistoryparticipant `json:"participants"`


    // Direction - The direction of the call relating to the current user
    Direction string `json:"direction"`


    // WentToVoicemail - Did the call end in the current user's voicemail
    WentToVoicemail bool `json:"wentToVoicemail"`


    // MissedCall - Did the user not answer this conversation
    MissedCall bool `json:"missedCall"`


    // StartTime - The time the user joined the conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartTime time.Time `json:"startTime"`


    // WasConference - Was this conversation a conference
    WasConference bool `json:"wasConference"`


    // WasCallback - Was this conversation a callback
    WasCallback bool `json:"wasCallback"`


    // HadScreenShare - Did this conversation have a screen share session
    HadScreenShare bool `json:"hadScreenShare"`


    // HadCobrowse - Did this conversation have a cobrowse session
    HadCobrowse bool `json:"hadCobrowse"`


    // WasOutboundCampaign - Was this conversation associated with an outbound campaign
    WasOutboundCampaign bool `json:"wasOutboundCampaign"`


    

}

// String returns a JSON representation of the model
func (o *Callhistoryconversation) String() string {
    
    
    
    
    
    
    
    
     o.Participants = []Callhistoryparticipant{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Callhistoryconversation) MarshalJSON() ([]byte, error) {
    type Alias Callhistoryconversation

    if CallhistoryconversationMarshalled {
        return []byte("{}"), nil
    }
    CallhistoryconversationMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Participants []Callhistoryparticipant `json:"participants"`
        
        Direction string `json:"direction"`
        
        WentToVoicemail bool `json:"wentToVoicemail"`
        
        MissedCall bool `json:"missedCall"`
        
        StartTime time.Time `json:"startTime"`
        
        WasConference bool `json:"wasConference"`
        
        WasCallback bool `json:"wasCallback"`
        
        HadScreenShare bool `json:"hadScreenShare"`
        
        HadCobrowse bool `json:"hadCobrowse"`
        
        WasOutboundCampaign bool `json:"wasOutboundCampaign"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        
        Participants: []Callhistoryparticipant{{}},
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

