package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TextbotdisconnectactionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TextbotdisconnectactionDud struct { 
    


    


    


    

}

// Textbotdisconnectaction - Settings for a next-action of disconnecting, including the reason code for the disconnect.
type Textbotdisconnectaction struct { 
    // Reason - The reason for the disconnect.
    Reason string `json:"reason"`


    // ReasonExtendedInfo - Extended information related to the reason, if available.
    ReasonExtendedInfo string `json:"reasonExtendedInfo"`


    // FlowLocation - Describes where in the Bot Flow the user was when the disconnect occurred.
    FlowLocation Textbotflowlocation `json:"flowLocation"`


    // FlowOutcomes - The list of Flow Outcomes for the bot flow and their details.
    FlowOutcomes []Textbotflowoutcome `json:"flowOutcomes"`

}

// String returns a JSON representation of the model
func (o *Textbotdisconnectaction) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.FlowOutcomes = []Textbotflowoutcome{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Textbotdisconnectaction) MarshalJSON() ([]byte, error) {
    type Alias Textbotdisconnectaction

    if TextbotdisconnectactionMarshalled {
        return []byte("{}"), nil
    }
    TextbotdisconnectactionMarshalled = true

    return json.Marshal(&struct { 
        Reason string `json:"reason"`
        
        ReasonExtendedInfo string `json:"reasonExtendedInfo"`
        
        FlowLocation Textbotflowlocation `json:"flowLocation"`
        
        FlowOutcomes []Textbotflowoutcome `json:"flowOutcomes"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        FlowOutcomes: []Textbotflowoutcome{{}},
        

        
        Alias: (*Alias)(u),
    })
}

