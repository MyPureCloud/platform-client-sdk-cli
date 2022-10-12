package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InitialconfigurationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InitialconfigurationDud struct { 
    


    


    


    


    


    

}

// Initialconfiguration
type Initialconfiguration struct { 
    // AudioState - Indicates the initial audio state for the communication.
    AudioState Audiostate `json:"audioState"`


    // Alerting - Indicates that this communication's initial state is alerting. If false, the communication started in a connected state.
    Alerting bool `json:"alerting"`


    // Inbound - Indicates the direction of this communication with respect to the contact center. `true` means the communication is INBOUND. `false` means the communication is OUTBOUND.
    Inbound bool `json:"inbound"`


    // InvitedBy - The id of the communication (the \"peer\") that \"invited\" this communication, if this occurred.
    InvitedBy string `json:"invitedBy"`


    // RecordingActive - Indicates whether recording is active for this communication at creation.
    RecordingActive bool `json:"recordingActive"`


    // AdditionalInfo - Additional metadata about this session which should be recorded by the platform but which will not be indexed or searchable. Primarily for diagnostic value. Any information that needs to be accessible through other components like Analytics should be moved to dedicated fields.
    AdditionalInfo map[string]string `json:"additionalInfo"`

}

// String returns a JSON representation of the model
func (o *Initialconfiguration) String() string {
    
    
    
    
    
     o.AdditionalInfo = map[string]string{"": ""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Initialconfiguration) MarshalJSON() ([]byte, error) {
    type Alias Initialconfiguration

    if InitialconfigurationMarshalled {
        return []byte("{}"), nil
    }
    InitialconfigurationMarshalled = true

    return json.Marshal(&struct {
        
        AudioState Audiostate `json:"audioState"`
        
        Alerting bool `json:"alerting"`
        
        Inbound bool `json:"inbound"`
        
        InvitedBy string `json:"invitedBy"`
        
        RecordingActive bool `json:"recordingActive"`
        
        AdditionalInfo map[string]string `json:"additionalInfo"`
        *Alias
    }{

        


        


        


        


        


        
        AdditionalInfo: map[string]string{"": ""},
        

        Alias: (*Alias)(u),
    })
}

