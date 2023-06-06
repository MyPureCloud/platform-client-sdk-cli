package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AdjustablelivespeakerdetectionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AdjustablelivespeakerdetectionDud struct { 
    


    


    


    


    


    

}

// Adjustablelivespeakerdetection
type Adjustablelivespeakerdetection struct { 
    // Mode - Modes to tune between speed to live speaker detection vs accuracy.
    Mode string `json:"mode"`


    // PreconnectDuration - ISO 8601 formatted relative duration (e.g., PT30.8427419S for 30.8 seconds), calculated on line connect.
    PreconnectDuration string `json:"preconnectDuration"`


    // EventName - The name of the event that triggered the ALSD evaluation (e.g., line.connect, speech.generic).
    EventName string `json:"eventName"`


    // IsPersonLikely - The output of the ALSD detector, evaluating whether there is likely a person on the call based on the above inputs, and if so, a person is detected early (person disposition name and speech.person analyzer result) and the associated action taken (e.g., speech.person postconnect entry in the disposition table has the action to transfer to a queue).
    IsPersonLikely bool `json:"isPersonLikely"`


    // TotalRingbacks - Number of tone.ring.* analyzer events detected during the call (expected mostly during pre-connect but the last ringback tone detection could potentially complete after line connect, which will increment totalRingbacks still).
    TotalRingbacks int `json:"totalRingbacks"`


    // LineConnected - Protocol line connect received (answered by a person, machine, busy, fax).
    LineConnected bool `json:"lineConnected"`

}

// String returns a JSON representation of the model
func (o *Adjustablelivespeakerdetection) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Adjustablelivespeakerdetection) MarshalJSON() ([]byte, error) {
    type Alias Adjustablelivespeakerdetection

    if AdjustablelivespeakerdetectionMarshalled {
        return []byte("{}"), nil
    }
    AdjustablelivespeakerdetectionMarshalled = true

    return json.Marshal(&struct {
        
        Mode string `json:"mode"`
        
        PreconnectDuration string `json:"preconnectDuration"`
        
        EventName string `json:"eventName"`
        
        IsPersonLikely bool `json:"isPersonLikely"`
        
        TotalRingbacks int `json:"totalRingbacks"`
        
        LineConnected bool `json:"lineConnected"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

