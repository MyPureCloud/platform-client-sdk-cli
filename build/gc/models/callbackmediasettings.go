package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CallbackmediasettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CallbackmediasettingsDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Callbackmediasettings
type Callbackmediasettings struct { 
    // EnableAutoAnswer - Indicates if auto-answer is enabled for the given media type or subtype (default is false).  Subtype settings take precedence over media type settings.
    EnableAutoAnswer bool `json:"enableAutoAnswer"`


    // AlertingTimeoutSeconds - The alerting timeout for the media type, in seconds
    AlertingTimeoutSeconds int `json:"alertingTimeoutSeconds"`


    // ServiceLevel - The targeted service level for the media type
    ServiceLevel Servicelevel `json:"serviceLevel"`


    // AutoAnswerAlertToneSeconds - How long to play the alerting tone for an auto-answer interaction
    AutoAnswerAlertToneSeconds float64 `json:"autoAnswerAlertToneSeconds"`


    // ManualAnswerAlertToneSeconds - How long to play the alerting tone for a manual-answer interaction
    ManualAnswerAlertToneSeconds float64 `json:"manualAnswerAlertToneSeconds"`


    // Mode - The mode callbacks will use on this queue.
    Mode string `json:"mode"`


    // EnableAutoDialAndEnd - Flag to enable Auto-Dial and Auto-End automation for callbacks on this queue.
    EnableAutoDialAndEnd bool `json:"enableAutoDialAndEnd"`


    // AutoDialDelaySeconds - Time in seconds after agent connects to callback before outgoing call is auto-dialed. Allowable values in range 0 - 1200 seconds. Defaults to 300 seconds.
    AutoDialDelaySeconds int `json:"autoDialDelaySeconds"`


    // AutoEndDelaySeconds - Time in seconds after agent disconnects from the outgoing call before the encasing callback is auto-ended. Allowable values in range 0 - 1200 seconds. Defaults to 300 seconds.
    AutoEndDelaySeconds int `json:"autoEndDelaySeconds"`


    // PacingModifier - Controls the maximum number of outbound calls at one time when mode is CustomerFirst.
    PacingModifier float64 `json:"pacingModifier"`


    // MaxRetryCount - Maximum number of retries that should be attempted to try and connect a customer first callback to a customer when the initial callback attempt did not connect.
    MaxRetryCount int `json:"maxRetryCount"`


    // RetryDelaySeconds - Delay in seconds between each retry of a customer first callback.
    RetryDelaySeconds int `json:"retryDelaySeconds"`


    // LiveVoiceReactionType - The action to take if a live voice is detected during the outbound call of a customer first callback.
    LiveVoiceReactionType string `json:"liveVoiceReactionType"`


    // LiveVoiceFlow - The inbound flow to transfer to if a live voice is detected during the outbound call of a customer first callback.
    LiveVoiceFlow Domainentityref `json:"liveVoiceFlow"`


    // AnsweringMachineReactionType - The action to take if an answering machine is detected during the outbound call of a customer first callback.
    AnsweringMachineReactionType string `json:"answeringMachineReactionType"`


    // AnsweringMachineFlow - The inbound flow to transfer to if an answering machine is detected during the outbound call of a customer first callback when answeringMachineReactionType is set to TransferToFlow.
    AnsweringMachineFlow Domainentityref `json:"answeringMachineFlow"`


    // EdgeGroup - The identifier of the edge group that will place the calls. Can be set to specify custom edge group instead of default one.
    EdgeGroup Domainentityref `json:"edgeGroup"`


    // Site - The identifier of the site to be used for dialing; can be set in place of an edge group.
    Site Domainentityref `json:"site"`

}

// String returns a JSON representation of the model
func (o *Callbackmediasettings) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Callbackmediasettings) MarshalJSON() ([]byte, error) {
    type Alias Callbackmediasettings

    if CallbackmediasettingsMarshalled {
        return []byte("{}"), nil
    }
    CallbackmediasettingsMarshalled = true

    return json.Marshal(&struct {
        
        EnableAutoAnswer bool `json:"enableAutoAnswer"`
        
        AlertingTimeoutSeconds int `json:"alertingTimeoutSeconds"`
        
        ServiceLevel Servicelevel `json:"serviceLevel"`
        
        AutoAnswerAlertToneSeconds float64 `json:"autoAnswerAlertToneSeconds"`
        
        ManualAnswerAlertToneSeconds float64 `json:"manualAnswerAlertToneSeconds"`
        
        Mode string `json:"mode"`
        
        EnableAutoDialAndEnd bool `json:"enableAutoDialAndEnd"`
        
        AutoDialDelaySeconds int `json:"autoDialDelaySeconds"`
        
        AutoEndDelaySeconds int `json:"autoEndDelaySeconds"`
        
        PacingModifier float64 `json:"pacingModifier"`
        
        MaxRetryCount int `json:"maxRetryCount"`
        
        RetryDelaySeconds int `json:"retryDelaySeconds"`
        
        LiveVoiceReactionType string `json:"liveVoiceReactionType"`
        
        LiveVoiceFlow Domainentityref `json:"liveVoiceFlow"`
        
        AnsweringMachineReactionType string `json:"answeringMachineReactionType"`
        
        AnsweringMachineFlow Domainentityref `json:"answeringMachineFlow"`
        
        EdgeGroup Domainentityref `json:"edgeGroup"`
        
        Site Domainentityref `json:"site"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

