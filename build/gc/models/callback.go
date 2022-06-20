package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CallbackMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CallbackDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Callback
type Callback struct { 
    // State - The connection state of this communication.
    State string `json:"state"`


    // Id - A globally unique identifier for this communication.
    Id string `json:"id"`


    // Segments - The time line of the participant's callback, divided into activity segments.
    Segments []Segment `json:"segments"`


    // Direction - The direction of the call
    Direction string `json:"direction"`


    // Held - True if this call is held and the person on this side hears silence.
    Held bool `json:"held"`


    // DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
    DisconnectType string `json:"disconnectType"`


    // StartHoldTime - The timestamp the callback was placed on hold in the cloud clock if the callback is currently on hold. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartHoldTime time.Time `json:"startHoldTime"`


    // DialerPreview - The preview data to be used when this callback is a Preview.
    DialerPreview Dialerpreview `json:"dialerPreview"`


    // Voicemail - The voicemail data to be used when this callback is an ACD voicemail.
    Voicemail Voicemail `json:"voicemail"`


    // CallbackNumbers - The phone number(s) to use to place the callback.
    CallbackNumbers []string `json:"callbackNumbers"`


    // CallbackUserName - The name of the user requesting a callback.
    CallbackUserName string `json:"callbackUserName"`


    // ScriptId - The UUID of the script to use.
    ScriptId string `json:"scriptId"`


    // ExternalCampaign - True if the call for the callback uses external dialing.
    ExternalCampaign bool `json:"externalCampaign"`


    // SkipEnabled - True if the ability to skip a callback should be enabled.
    SkipEnabled bool `json:"skipEnabled"`


    // TimeoutSeconds - The number of seconds before the system automatically places a call for a callback.  0 means the automatic placement is disabled.
    TimeoutSeconds int `json:"timeoutSeconds"`


    // StartAlertingTime - The timestamp the communication has when it is first put into an alerting state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartAlertingTime time.Time `json:"startAlertingTime"`


    // ConnectedTime - The timestamp when this communication was connected in the cloud clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ConnectedTime time.Time `json:"connectedTime"`


    // DisconnectedTime - The timestamp when this communication disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DisconnectedTime time.Time `json:"disconnectedTime"`


    // CallbackScheduledTime - The timestamp when this communication is scheduled in the provider clock. If this value is missing it indicates the callback will be placed immediately. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CallbackScheduledTime time.Time `json:"callbackScheduledTime"`


    // AutomatedCallbackConfigId - The id of the config for automatically placing the callback (and handling the disposition). If null, the callback will not be placed automatically but routed to an agent as per normal.
    AutomatedCallbackConfigId string `json:"automatedCallbackConfigId"`


    // Provider - The source provider for the callback.
    Provider string `json:"provider"`


    // PeerId - The id of the peer communication corresponding to a matching leg for this communication.
    PeerId string `json:"peerId"`


    // Wrapup - Call wrap up or disposition data.
    Wrapup Wrapup `json:"wrapup"`


    // AfterCallWork - After-call work for the communication.
    AfterCallWork Aftercallwork `json:"afterCallWork"`


    // AfterCallWorkRequired - Indicates if after-call work is required for a communication. Only used when the ACW Setting is Agent Requested.
    AfterCallWorkRequired bool `json:"afterCallWorkRequired"`


    // CallerId - The phone number displayed to recipients of the phone call. The value should conform to the E164 format.
    CallerId string `json:"callerId"`


    // CallerIdName - The name displayed to recipients of the phone call.
    CallerIdName string `json:"callerIdName"`


    // InitialState - The initial connection state of this communication.
    InitialState string `json:"initialState"`

}

// String returns a JSON representation of the model
func (o *Callback) String() string {
    
    
     o.Segments = []Segment{{}} 
    
    
    
    
    
    
     o.CallbackNumbers = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Callback) MarshalJSON() ([]byte, error) {
    type Alias Callback

    if CallbackMarshalled {
        return []byte("{}"), nil
    }
    CallbackMarshalled = true

    return json.Marshal(&struct {
        
        State string `json:"state"`
        
        Id string `json:"id"`
        
        Segments []Segment `json:"segments"`
        
        Direction string `json:"direction"`
        
        Held bool `json:"held"`
        
        DisconnectType string `json:"disconnectType"`
        
        StartHoldTime time.Time `json:"startHoldTime"`
        
        DialerPreview Dialerpreview `json:"dialerPreview"`
        
        Voicemail Voicemail `json:"voicemail"`
        
        CallbackNumbers []string `json:"callbackNumbers"`
        
        CallbackUserName string `json:"callbackUserName"`
        
        ScriptId string `json:"scriptId"`
        
        ExternalCampaign bool `json:"externalCampaign"`
        
        SkipEnabled bool `json:"skipEnabled"`
        
        TimeoutSeconds int `json:"timeoutSeconds"`
        
        StartAlertingTime time.Time `json:"startAlertingTime"`
        
        ConnectedTime time.Time `json:"connectedTime"`
        
        DisconnectedTime time.Time `json:"disconnectedTime"`
        
        CallbackScheduledTime time.Time `json:"callbackScheduledTime"`
        
        AutomatedCallbackConfigId string `json:"automatedCallbackConfigId"`
        
        Provider string `json:"provider"`
        
        PeerId string `json:"peerId"`
        
        Wrapup Wrapup `json:"wrapup"`
        
        AfterCallWork Aftercallwork `json:"afterCallWork"`
        
        AfterCallWorkRequired bool `json:"afterCallWorkRequired"`
        
        CallerId string `json:"callerId"`
        
        CallerIdName string `json:"callerIdName"`
        
        InitialState string `json:"initialState"`
        *Alias
    }{

        


        


        
        Segments: []Segment{{}},
        


        


        


        


        


        


        


        
        CallbackNumbers: []string{""},
        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

