package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TextbotexitactionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TextbotexitactionDud struct { 
    


    


    


    


    


    

}

// Textbotexitaction - Settings for a next-action of exiting the bot flow. Any output variables are available in the details.
type Textbotexitaction struct { 
    // Reason - The reason for the exit.
    Reason string `json:"reason"`


    // ReasonExtendedInfo - Extended information related to the reason, if available.
    ReasonExtendedInfo string `json:"reasonExtendedInfo"`


    // ActiveIntent - The active intent at the time of the exit.
    ActiveIntent string `json:"activeIntent"`


    // FlowLocation - Describes where in the Bot Flow the user was when the exit occurred.
    FlowLocation Textbotflowlocation `json:"flowLocation"`


    // OutputData - The output data for the bot flow.
    OutputData Textbotinputoutputdata `json:"outputData"`


    // FlowOutcomes - The list of Flow Outcomes for the bot flow and their details.
    FlowOutcomes []Textbotflowoutcome `json:"flowOutcomes"`

}

// String returns a JSON representation of the model
func (o *Textbotexitaction) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.FlowOutcomes = []Textbotflowoutcome{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Textbotexitaction) MarshalJSON() ([]byte, error) {
    type Alias Textbotexitaction

    if TextbotexitactionMarshalled {
        return []byte("{}"), nil
    }
    TextbotexitactionMarshalled = true

    return json.Marshal(&struct { 
        Reason string `json:"reason"`
        
        ReasonExtendedInfo string `json:"reasonExtendedInfo"`
        
        ActiveIntent string `json:"activeIntent"`
        
        FlowLocation Textbotflowlocation `json:"flowLocation"`
        
        OutputData Textbotinputoutputdata `json:"outputData"`
        
        FlowOutcomes []Textbotflowoutcome `json:"flowOutcomes"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        
        FlowOutcomes: []Textbotflowoutcome{{}},
        

        
        Alias: (*Alias)(u),
    })
}

