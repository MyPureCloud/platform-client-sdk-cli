package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Lexv2intentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Lexv2intentDud struct { 
    


    


    


    

}

// Lexv2intent
type Lexv2intent struct { 
    // IntentName - The intent name
    IntentName string `json:"intentName"`


    // Description - A description of the intent
    Description string `json:"description"`


    // Slots - An object mapping slot names to Slot objects
    Slots map[string]Lexv2slot `json:"slots"`


    // IntentId - The intent id
    IntentId string `json:"intentId"`

}

// String returns a JSON representation of the model
func (o *Lexv2intent) String() string {
    
    
     o.Slots = map[string]Lexv2slot{"": {}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Lexv2intent) MarshalJSON() ([]byte, error) {
    type Alias Lexv2intent

    if Lexv2intentMarshalled {
        return []byte("{}"), nil
    }
    Lexv2intentMarshalled = true

    return json.Marshal(&struct {
        
        IntentName string `json:"intentName"`
        
        Description string `json:"description"`
        
        Slots map[string]Lexv2slot `json:"slots"`
        
        IntentId string `json:"intentId"`
        *Alias
    }{

        


        


        
        Slots: map[string]Lexv2slot{"": {}},
        


        

        Alias: (*Alias)(u),
    })
}

