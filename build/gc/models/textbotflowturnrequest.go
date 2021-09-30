package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TextbotflowturnrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TextbotflowturnrequestDud struct { 
    


    


    


    

}

// Textbotflowturnrequest - Settings for a turn request to a bot flow.
type Textbotflowturnrequest struct { 
    // PreviousTurn - The reference to a previous turn if appropriate, used to avoid race conditions.
    PreviousTurn Textbotturnreference `json:"previousTurn"`


    // InputEventType - Indicates the type of input event being requested. If appropriate, fill out the matching user input object details on this request.
    InputEventType string `json:"inputEventType"`


    // InputEventUserInput - The data for the input event of this turn if it is a user input event. Only one inputEvent may be set.
    InputEventUserInput Textbotuserinputevent `json:"inputEventUserInput"`


    // InputEventError - The data for the input event of this turn if it is an error event. Only one inputEvent may be set.
    InputEventError Textboterrorinputevent `json:"inputEventError"`

}

// String returns a JSON representation of the model
func (o *Textbotflowturnrequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Textbotflowturnrequest) MarshalJSON() ([]byte, error) {
    type Alias Textbotflowturnrequest

    if TextbotflowturnrequestMarshalled {
        return []byte("{}"), nil
    }
    TextbotflowturnrequestMarshalled = true

    return json.Marshal(&struct { 
        PreviousTurn Textbotturnreference `json:"previousTurn"`
        
        InputEventType string `json:"inputEventType"`
        
        InputEventUserInput Textbotuserinputevent `json:"inputEventUserInput"`
        
        InputEventError Textboterrorinputevent `json:"inputEventError"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

