package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactcallbackrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactcallbackrequestDud struct { 
    


    


    


    


    

}

// Contactcallbackrequest
type Contactcallbackrequest struct { 
    // CampaignId - Campaign identifier
    CampaignId string `json:"campaignId"`


    // ContactListId - Contact list identifier
    ContactListId string `json:"contactListId"`


    // ContactId - Contact identifier
    ContactId string `json:"contactId"`


    // PhoneColumn - Name of the phone column containing the number to be called
    PhoneColumn string `json:"phoneColumn"`


    // Schedule - The scheduled time for the callback as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ\", example = \"2016-01-02T16:59:59\"
    Schedule string `json:"schedule"`

}

// String returns a JSON representation of the model
func (o *Contactcallbackrequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactcallbackrequest) MarshalJSON() ([]byte, error) {
    type Alias Contactcallbackrequest

    if ContactcallbackrequestMarshalled {
        return []byte("{}"), nil
    }
    ContactcallbackrequestMarshalled = true

    return json.Marshal(&struct { 
        CampaignId string `json:"campaignId"`
        
        ContactListId string `json:"contactListId"`
        
        ContactId string `json:"contactId"`
        
        PhoneColumn string `json:"phoneColumn"`
        
        Schedule string `json:"schedule"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

