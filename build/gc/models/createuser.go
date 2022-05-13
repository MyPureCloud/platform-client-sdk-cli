package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateuserMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateuserDud struct { 
    


    


    


    


    


    


    


    

}

// Createuser
type Createuser struct { 
    // Name - User's full name
    Name string `json:"name"`


    // Department
    Department string `json:"department"`


    // Email - User's email and username
    Email string `json:"email"`


    // Addresses - Email addresses and phone numbers for this user
    Addresses []Contact `json:"addresses"`


    // Title
    Title string `json:"title"`


    // Password - User's password
    Password string `json:"password"`


    // DivisionId - The division to which this user will belong
    DivisionId string `json:"divisionId"`


    // State - Optional initialized state of the user. If not specified, state will be Active if invites are sent, otherwise Inactive.
    State string `json:"state"`

}

// String returns a JSON representation of the model
func (o *Createuser) String() string {
    
    
    
     o.Addresses = []Contact{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createuser) MarshalJSON() ([]byte, error) {
    type Alias Createuser

    if CreateuserMarshalled {
        return []byte("{}"), nil
    }
    CreateuserMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Department string `json:"department"`
        
        Email string `json:"email"`
        
        Addresses []Contact `json:"addresses"`
        
        Title string `json:"title"`
        
        Password string `json:"password"`
        
        DivisionId string `json:"divisionId"`
        
        State string `json:"state"`
        *Alias
    }{

        


        


        


        
        Addresses: []Contact{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

