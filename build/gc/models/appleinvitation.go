package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AppleinvitationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AppleinvitationDud struct { 
    


    


    


    

}

// Appleinvitation - Apple Messages for Business invitation template configuration
type Appleinvitation struct { 
    // BusinessName - The business name displayed in the invitation
    BusinessName string `json:"businessName"`


    // TranscriptMessage - The transcript message displayed in the invitation
    TranscriptMessage string `json:"transcriptMessage"`


    // TemplateType - The template type for the invitation
    TemplateType string `json:"templateType"`


    // Locale - The locale for the invitation
    Locale string `json:"locale"`

}

// String returns a JSON representation of the model
func (o *Appleinvitation) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Appleinvitation) MarshalJSON() ([]byte, error) {
    type Alias Appleinvitation

    if AppleinvitationMarshalled {
        return []byte("{}"), nil
    }
    AppleinvitationMarshalled = true

    return json.Marshal(&struct {
        
        BusinessName string `json:"businessName"`
        
        TranscriptMessage string `json:"transcriptMessage"`
        
        TemplateType string `json:"templateType"`
        
        Locale string `json:"locale"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

