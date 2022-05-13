package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TextbotflowlocationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TextbotflowlocationDud struct { 
    


    


    

}

// Textbotflowlocation - Describes a flow location.
type Textbotflowlocation struct { 
    // ActionName - The name of the action that was active when the event of interest happened.
    ActionName string `json:"actionName"`


    // ActionNumber - The number of the action that was active when the event of interest happened.
    ActionNumber int `json:"actionNumber"`


    // SequenceName - The name of the state or task which was active when the event of interest happened.
    SequenceName string `json:"sequenceName"`

}

// String returns a JSON representation of the model
func (o *Textbotflowlocation) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Textbotflowlocation) MarshalJSON() ([]byte, error) {
    type Alias Textbotflowlocation

    if TextbotflowlocationMarshalled {
        return []byte("{}"), nil
    }
    TextbotflowlocationMarshalled = true

    return json.Marshal(&struct {
        
        ActionName string `json:"actionName"`
        
        ActionNumber int `json:"actionNumber"`
        
        SequenceName string `json:"sequenceName"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

