package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebchatmemberinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebchatmemberinfoDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Webchatmemberinfo
type Webchatmemberinfo struct { 
    // Id - The communicationId of this member.
    Id string `json:"id"`


    // DisplayName - The display name of the member.
    DisplayName string `json:"displayName"`


    // FirstName - The first name of the member.
    FirstName string `json:"firstName"`


    // LastName - The last name of the member.
    LastName string `json:"lastName"`


    // Email - The email address of the member.
    Email string `json:"email"`


    // PhoneNumber - The phone number of the member.
    PhoneNumber string `json:"phoneNumber"`


    // AvatarImageUrl - The url to the avatar image of the member.
    AvatarImageUrl string `json:"avatarImageUrl"`


    // Role - The role of the member, one of [agent, customer, acd, workflow]
    Role string `json:"role"`


    // JoinDate - The time the member joined the conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    JoinDate time.Time `json:"joinDate"`


    // LeaveDate - The time the member left the conversation, or null if the member is still active in the conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    LeaveDate time.Time `json:"leaveDate"`


    // AuthenticatedGuest - If true, the guest member is an authenticated guest.
    AuthenticatedGuest bool `json:"authenticatedGuest"`


    // CustomFields - Any custom fields of information pertaining to this member.
    CustomFields map[string]string `json:"customFields"`


    // State - The connection state of this member.
    State string `json:"state"`

}

// String returns a JSON representation of the model
func (o *Webchatmemberinfo) String() string {
    
    
    
    
    
    
    
    
    
    
    
     o.CustomFields = map[string]string{"": ""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webchatmemberinfo) MarshalJSON() ([]byte, error) {
    type Alias Webchatmemberinfo

    if WebchatmemberinfoMarshalled {
        return []byte("{}"), nil
    }
    WebchatmemberinfoMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        DisplayName string `json:"displayName"`
        
        FirstName string `json:"firstName"`
        
        LastName string `json:"lastName"`
        
        Email string `json:"email"`
        
        PhoneNumber string `json:"phoneNumber"`
        
        AvatarImageUrl string `json:"avatarImageUrl"`
        
        Role string `json:"role"`
        
        JoinDate time.Time `json:"joinDate"`
        
        LeaveDate time.Time `json:"leaveDate"`
        
        AuthenticatedGuest bool `json:"authenticatedGuest"`
        
        CustomFields map[string]string `json:"customFields"`
        
        State string `json:"state"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        
        CustomFields: map[string]string{"": ""},
        


        

        Alias: (*Alias)(u),
    })
}

