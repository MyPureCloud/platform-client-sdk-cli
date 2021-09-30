package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContentnotificationtemplateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContentnotificationtemplateDud struct { 
    


    


    


    


    

}

// Contentnotificationtemplate - Template notification object.
type Contentnotificationtemplate struct { 
    // Id - The messaging provider template ID. For WhatsApp, 'namespace@name'.
    Id string `json:"id"`


    // Language - Template language.
    Language string `json:"language"`


    // Header - The template header.
    Header Notificationtemplateheader `json:"header"`


    // Body - The template body.
    Body Notificationtemplatebody `json:"body"`


    // Footer - The template footer.
    Footer Notificationtemplatefooter `json:"footer"`

}

// String returns a JSON representation of the model
func (o *Contentnotificationtemplate) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contentnotificationtemplate) MarshalJSON() ([]byte, error) {
    type Alias Contentnotificationtemplate

    if ContentnotificationtemplateMarshalled {
        return []byte("{}"), nil
    }
    ContentnotificationtemplateMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Language string `json:"language"`
        
        Header Notificationtemplateheader `json:"header"`
        
        Body Notificationtemplatebody `json:"body"`
        
        Footer Notificationtemplatefooter `json:"footer"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

