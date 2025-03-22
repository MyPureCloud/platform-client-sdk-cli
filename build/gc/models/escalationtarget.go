package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EscalationtargetMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EscalationtargetDud struct { 
    


    


    

}

// Escalationtarget
type Escalationtarget struct { 
    // TargetType - Defines the target that the message will be escalated to.
    TargetType string `json:"targetType"`


    // Destination - Defines the destination of the escalation.SourceIntegration means use the SocialMedia Source Integration as the destination.OverrideIntegration means the set integration will be used regardless of the source.
    Destination string `json:"destination"`


    // Override - Set the integration ID.Only valid when type is OverrideIntegration.
    Override Overrideescalationtarget `json:"override"`

}

// String returns a JSON representation of the model
func (o *Escalationtarget) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Escalationtarget) MarshalJSON() ([]byte, error) {
    type Alias Escalationtarget

    if EscalationtargetMarshalled {
        return []byte("{}"), nil
    }
    EscalationtargetMarshalled = true

    return json.Marshal(&struct {
        
        TargetType string `json:"targetType"`
        
        Destination string `json:"destination"`
        
        Override Overrideescalationtarget `json:"override"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

