package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TextbotflowturnresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TextbotflowturnresponseDud struct { 
    


    


    


    


    


    


    

}

// Textbotflowturnresponse - Information related to a success bot flow turn request.
type Textbotflowturnresponse struct { 
    // Id - The ID of the bot flow turn. If additional turns are needed, supply this ID as the previous turn in your next turn request.
    Id string `json:"id"`


    // PreviousTurn - The reference to a previous turn, if applicable.
    PreviousTurn Textbotturnreference `json:"previousTurn"`


    // Prompts - The output prompts for this turn.
    Prompts Textbotoutputprompts `json:"prompts"`


    // NextActionType - Indicates the suggested next action. If appropriate, the matching output event object includes additional information.
    NextActionType string `json:"nextActionType"`


    // NextActionDisconnect - The next action directive for this turn if it is a Disconnect type.
    NextActionDisconnect Textbotdisconnectaction `json:"nextActionDisconnect"`


    // NextActionWaitForInput - The next action directive for this turn if it is a WaitForInput type.
    NextActionWaitForInput Textbotwaitforinputaction `json:"nextActionWaitForInput"`


    // NextActionExit - The next action directive for this turn if it is an Exit type.
    NextActionExit Textbotexitaction `json:"nextActionExit"`

}

// String returns a JSON representation of the model
func (o *Textbotflowturnresponse) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Textbotflowturnresponse) MarshalJSON() ([]byte, error) {
    type Alias Textbotflowturnresponse

    if TextbotflowturnresponseMarshalled {
        return []byte("{}"), nil
    }
    TextbotflowturnresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        PreviousTurn Textbotturnreference `json:"previousTurn"`
        
        Prompts Textbotoutputprompts `json:"prompts"`
        
        NextActionType string `json:"nextActionType"`
        
        NextActionDisconnect Textbotdisconnectaction `json:"nextActionDisconnect"`
        
        NextActionWaitForInput Textbotwaitforinputaction `json:"nextActionWaitForInput"`
        
        NextActionExit Textbotexitaction `json:"nextActionExit"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

