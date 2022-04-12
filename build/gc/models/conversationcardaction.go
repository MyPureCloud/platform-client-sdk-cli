package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcardactionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcardactionDud struct { 
    


    


    


    

}

// Conversationcardaction - CardAction Object
type Conversationcardaction struct { 
    // VarType - Describes the type of action.
    VarType string `json:"type"`


    // Text - The response text from the button click.
    Text string `json:"text"`


    // Payload - Text to be returned as the payload from a ButtonResponse when a button is clicked.
    Payload string `json:"payload"`


    // Url - A URL of a web page to direct the user to.
    Url string `json:"url"`

}

// String returns a JSON representation of the model
func (o *Conversationcardaction) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcardaction) MarshalJSON() ([]byte, error) {
    type Alias Conversationcardaction

    if ConversationcardactionMarshalled {
        return []byte("{}"), nil
    }
    ConversationcardactionMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Text string `json:"text"`
        
        Payload string `json:"payload"`
        
        Url string `json:"url"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

