package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NotificationtemplatebodyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NotificationtemplatebodyDud struct { 
    


    

}

// Notificationtemplatebody - Template body object.
type Notificationtemplatebody struct { 
    // Text - Body text. For WhatsApp, ignored.
    Text string `json:"text"`


    // Parameters - Template parameters for placeholders in template.
    Parameters []Notificationtemplateparameter `json:"parameters"`

}

// String returns a JSON representation of the model
func (o *Notificationtemplatebody) String() string {
    
    
    
    
    
    
     o.Parameters = []Notificationtemplateparameter{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Notificationtemplatebody) MarshalJSON() ([]byte, error) {
    type Alias Notificationtemplatebody

    if NotificationtemplatebodyMarshalled {
        return []byte("{}"), nil
    }
    NotificationtemplatebodyMarshalled = true

    return json.Marshal(&struct { 
        Text string `json:"text"`
        
        Parameters []Notificationtemplateparameter `json:"parameters"`
        
        *Alias
    }{
        

        

        

        
        Parameters: []Notificationtemplateparameter{{}},
        

        
        Alias: (*Alias)(u),
    })
}

