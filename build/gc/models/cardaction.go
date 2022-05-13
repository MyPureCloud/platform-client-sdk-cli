package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CardactionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CardactionDud struct { 
    


    


    


    


    

}

// Cardaction - A card action that a user can take.
type Cardaction struct { 
    // VarType - Describes the type of action.
    VarType string `json:"type"`


    // Text - The response text from the button click.
    Text string `json:"text"`


    // Payload - Content of the textback payload after clicking a button.
    Payload string `json:"payload"`


    // Url - The location of the image file associated with action.
    Url string `json:"url"`


    // IsSelected - Indicates if the card option is selected by end customer.
    IsSelected bool `json:"isSelected"`

}

// String returns a JSON representation of the model
func (o *Cardaction) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Cardaction) MarshalJSON() ([]byte, error) {
    type Alias Cardaction

    if CardactionMarshalled {
        return []byte("{}"), nil
    }
    CardactionMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Text string `json:"text"`
        
        Payload string `json:"payload"`
        
        Url string `json:"url"`
        
        IsSelected bool `json:"isSelected"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

