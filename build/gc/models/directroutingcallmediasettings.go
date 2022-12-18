package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DirectroutingcallmediasettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DirectroutingcallmediasettingsDud struct { 
    


    


    

}

// Directroutingcallmediasettings
type Directroutingcallmediasettings struct { 
    // Enabled - Toggle that enables Direct Routing for this media type.
    Enabled bool `json:"enabled"`


    // InboundFlow - The Direct Routing inbound flow id for this media type.
    InboundFlow Addressableentityref `json:"inboundFlow"`


    // VoicemailFlow - ID of the in-queue flow that handles voicemails for Direct Routing and to transfer calls to ACD voicemail.
    VoicemailFlow Addressableentityref `json:"voicemailFlow"`

}

// String returns a JSON representation of the model
func (o *Directroutingcallmediasettings) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Directroutingcallmediasettings) MarshalJSON() ([]byte, error) {
    type Alias Directroutingcallmediasettings

    if DirectroutingcallmediasettingsMarshalled {
        return []byte("{}"), nil
    }
    DirectroutingcallmediasettingsMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        InboundFlow Addressableentityref `json:"inboundFlow"`
        
        VoicemailFlow Addressableentityref `json:"voicemailFlow"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

