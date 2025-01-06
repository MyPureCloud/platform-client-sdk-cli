package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingtemplatebuttonMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingtemplatebuttonDud struct { 
    


    


    


    


    


    

}

// Recordingtemplatebutton
type Recordingtemplatebutton struct { 
    // VarType - Specifies the type of the button.
    VarType string `json:"type"`


    // Text - Button text message.
    Text string `json:"text"`


    // Index - Index of the button in the list.
    Index int `json:"index"`


    // PhoneNumber - Button phone number.
    PhoneNumber string `json:"phoneNumber"`


    // Url - Button URL link.
    Url string `json:"url"`


    // IsSelected - Indicates if the button is selected by end customer.
    IsSelected bool `json:"isSelected"`

}

// String returns a JSON representation of the model
func (o *Recordingtemplatebutton) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingtemplatebutton) MarshalJSON() ([]byte, error) {
    type Alias Recordingtemplatebutton

    if RecordingtemplatebuttonMarshalled {
        return []byte("{}"), nil
    }
    RecordingtemplatebuttonMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Text string `json:"text"`
        
        Index int `json:"index"`
        
        PhoneNumber string `json:"phoneNumber"`
        
        Url string `json:"url"`
        
        IsSelected bool `json:"isSelected"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

