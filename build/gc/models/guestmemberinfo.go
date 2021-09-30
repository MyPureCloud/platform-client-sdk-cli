package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GuestmemberinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GuestmemberinfoDud struct { 
    


    


    


    


    


    


    

}

// Guestmemberinfo
type Guestmemberinfo struct { 
    // DisplayName - The display name to use for the guest member in the conversation.
    DisplayName string `json:"displayName"`


    // FirstName - The first name to use for the guest member in the conversation.
    FirstName string `json:"firstName"`


    // LastName - The last name to use for the guest member in the conversation.
    LastName string `json:"lastName"`


    // Email - The email address to use for the guest member in the conversation.
    Email string `json:"email"`


    // PhoneNumber - The phone number to use for the guest member in the conversation.
    PhoneNumber string `json:"phoneNumber"`


    // AvatarImageUrl - The URL to the avatar image to use for the guest member in the conversation, if any.
    AvatarImageUrl string `json:"avatarImageUrl"`


    // CustomFields - Any custom fields of information, in key-value format, to attach to the guest member in the conversation.
    CustomFields map[string]string `json:"customFields"`

}

// String returns a JSON representation of the model
func (o *Guestmemberinfo) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.CustomFields = map[string]string{"": ""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Guestmemberinfo) MarshalJSON() ([]byte, error) {
    type Alias Guestmemberinfo

    if GuestmemberinfoMarshalled {
        return []byte("{}"), nil
    }
    GuestmemberinfoMarshalled = true

    return json.Marshal(&struct { 
        DisplayName string `json:"displayName"`
        
        FirstName string `json:"firstName"`
        
        LastName string `json:"lastName"`
        
        Email string `json:"email"`
        
        PhoneNumber string `json:"phoneNumber"`
        
        AvatarImageUrl string `json:"avatarImageUrl"`
        
        CustomFields map[string]string `json:"customFields"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        
        CustomFields: map[string]string{"": ""},
        

        
        Alias: (*Alias)(u),
    })
}

