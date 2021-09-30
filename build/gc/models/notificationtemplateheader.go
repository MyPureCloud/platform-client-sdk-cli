package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NotificationtemplateheaderMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NotificationtemplateheaderDud struct { 
    


    


    


    

}

// Notificationtemplateheader - Template header object.
type Notificationtemplateheader struct { 
    // VarType - Template header type.
    VarType string `json:"type"`


    // Text - Header text. For WhatsApp, ignored.
    Text string `json:"text"`


    // Media - Media template header image.
    Media Contentattachment `json:"media"`


    // Parameters - Template parameters for placeholders in template.
    Parameters []Notificationtemplateparameter `json:"parameters"`

}

// String returns a JSON representation of the model
func (o *Notificationtemplateheader) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Parameters = []Notificationtemplateparameter{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Notificationtemplateheader) MarshalJSON() ([]byte, error) {
    type Alias Notificationtemplateheader

    if NotificationtemplateheaderMarshalled {
        return []byte("{}"), nil
    }
    NotificationtemplateheaderMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Text string `json:"text"`
        
        Media Contentattachment `json:"media"`
        
        Parameters []Notificationtemplateparameter `json:"parameters"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        Parameters: []Notificationtemplateparameter{{}},
        

        
        Alias: (*Alias)(u),
    })
}

