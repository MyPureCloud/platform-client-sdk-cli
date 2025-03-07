package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NotificationtemplatebuttonMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NotificationtemplatebuttonDud struct { 
    


    


    


    


    


    

}

// Notificationtemplatebutton - Template button object
type Notificationtemplatebutton struct { 
    // VarType - Specifies the type of the button.
    VarType string `json:"type"`


    // Text - Button text message.
    Text string `json:"text"`


    // Index - index of the button in the list.
    Index int `json:"index"`


    // PhoneNumber - Button phone number.
    PhoneNumber string `json:"phoneNumber"`


    // Url - Button URL link.
    Url string `json:"url"`


    // Parameters - Template parameters for placeholders in the button.
    Parameters []Notificationtemplateparameter `json:"parameters"`

}

// String returns a JSON representation of the model
func (o *Notificationtemplatebutton) String() string {
    
    
    
    
    
     o.Parameters = []Notificationtemplateparameter{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Notificationtemplatebutton) MarshalJSON() ([]byte, error) {
    type Alias Notificationtemplatebutton

    if NotificationtemplatebuttonMarshalled {
        return []byte("{}"), nil
    }
    NotificationtemplatebuttonMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Text string `json:"text"`
        
        Index int `json:"index"`
        
        PhoneNumber string `json:"phoneNumber"`
        
        Url string `json:"url"`
        
        Parameters []Notificationtemplateparameter `json:"parameters"`
        *Alias
    }{

        


        


        


        


        


        
        Parameters: []Notificationtemplateparameter{{}},
        

        Alias: (*Alias)(u),
    })
}

