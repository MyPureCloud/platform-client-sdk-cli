package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateuserMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateuserDud struct { 
    Id string `json:"id"`


    


    


    


    


    PrimaryContactInfo []Contact `json:"primaryContactInfo"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Updateuser
type Updateuser struct { 
    


    // Name
    Name string `json:"name"`


    // Chat
    Chat Chat `json:"chat"`


    // Department
    Department string `json:"department"`


    // Email
    Email string `json:"email"`


    


    // Addresses - Email address, phone number, and/or extension for this user. One entry is allowed per media type
    Addresses []Contact `json:"addresses"`


    // Title
    Title string `json:"title"`


    // Username
    Username string `json:"username"`


    // Manager
    Manager string `json:"manager"`


    // Images
    Images []Userimage `json:"images"`


    // Version - This value should be the current version of the user. The current version can be obtained with a GET on the user before doing a PATCH.
    Version int `json:"version"`


    // ProfileSkills - Profile skills possessed by the user
    ProfileSkills []string `json:"profileSkills"`


    // Locations - The user placement at each site location.
    Locations []Location `json:"locations"`


    // Groups - The groups the user is a member of
    Groups []Group `json:"groups"`


    // State - The state of the user. This property can be used to restore a deleted user or transition between active and inactive. If specified, it is the only modifiable field.
    State string `json:"state"`


    // AcdAutoAnswer - The value that denotes if acdAutoAnswer is set on the user
    AcdAutoAnswer bool `json:"acdAutoAnswer"`


    // Certifications
    Certifications []string `json:"certifications"`


    // Biography
    Biography Biography `json:"biography"`


    // EmployerInfo
    EmployerInfo Employerinfo `json:"employerInfo"`


    

}

// String returns a JSON representation of the model
func (o *Updateuser) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Addresses = []Contact{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Images = []Userimage{{}} 
    
    
    
    
    
    
    
     o.ProfileSkills = []string{""} 
    
    
    
     o.Locations = []Location{{}} 
    
    
    
     o.Groups = []Group{{}} 
    
    
    
    
    
    
    
    
    
    
    
     o.Certifications = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateuser) MarshalJSON() ([]byte, error) {
    type Alias Updateuser

    if UpdateuserMarshalled {
        return []byte("{}"), nil
    }
    UpdateuserMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Chat Chat `json:"chat"`
        
        Department string `json:"department"`
        
        Email string `json:"email"`
        
        
        
        Addresses []Contact `json:"addresses"`
        
        Title string `json:"title"`
        
        Username string `json:"username"`
        
        Manager string `json:"manager"`
        
        Images []Userimage `json:"images"`
        
        Version int `json:"version"`
        
        ProfileSkills []string `json:"profileSkills"`
        
        Locations []Location `json:"locations"`
        
        Groups []Group `json:"groups"`
        
        State string `json:"state"`
        
        AcdAutoAnswer bool `json:"acdAutoAnswer"`
        
        Certifications []string `json:"certifications"`
        
        Biography Biography `json:"biography"`
        
        EmployerInfo Employerinfo `json:"employerInfo"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Addresses: []Contact{{}},
        

        

        

        

        

        

        

        

        
        Images: []Userimage{{}},
        

        

        

        

        
        ProfileSkills: []string{""},
        

        

        
        Locations: []Location{{}},
        

        

        
        Groups: []Group{{}},
        

        

        

        

        

        

        
        Certifications: []string{""},
        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

