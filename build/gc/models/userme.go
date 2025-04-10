package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsermeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsermeDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    PrimaryContactInfo []Contact `json:"primaryContactInfo"`


    


    State string `json:"state"`


    


    


    


    


    


    


    


    


    


    RoutingStatus Routingstatus `json:"routingStatus"`


    Presence Userpresence `json:"presence"`


    IntegrationPresence Userpresence `json:"integrationPresence"`


    ConversationSummary Userconversationsummary `json:"conversationSummary"`


    OutOfOffice Outofoffice `json:"outOfOffice"`


    Geolocation Geolocation `json:"geolocation"`


    Station Userstations `json:"station"`


    Authorization Userauthorization `json:"authorization"`


    ProfileSkills []string `json:"profileSkills"`


    Locations []Location `json:"locations"`


    Groups []Group `json:"groups"`


    Team Team `json:"team"`


    WorkPlanBidRanks Workplanbidranks `json:"workPlanBidRanks"`


    Skills []Userroutingskill `json:"skills"`


    Languages []Userroutinglanguage `json:"languages"`


    AutoAnswerSettings Autoanswersettings `json:"autoAnswerSettings"`


    


    LanguagePreference string `json:"languagePreference"`


    


    DateLastLogin time.Time `json:"dateLastLogin"`


    DateWelcomeSent time.Time `json:"dateWelcomeSent"`


    Date Serverdate `json:"date"`


    GeolocationSettings Geolocationsettings `json:"geolocationSettings"`


    Organization Organization `json:"organization"`


    PresenceDefinitions []Organizationpresence `json:"presenceDefinitions"`


    DivisionedPresenceDefinitions []Organizationpresencedefinition `json:"divisionedPresenceDefinitions"`


    LocationDefinitions []Locationdefinition `json:"locationDefinitions"`


    OrgAuthorization []Domainorganizationrole `json:"orgAuthorization"`


    Favorites []User `json:"favorites"`


    Superiors []User `json:"superiors"`


    DirectReports []User `json:"directReports"`


    Adjacents Adjacents `json:"adjacents"`


    RoutingSkills []Routingskill `json:"routingSkills"`


    FieldConfigs Fieldconfigs `json:"fieldConfigs"`


    Token Tokeninfo `json:"token"`


    Trustors []Trustor `json:"trustors"`


    OrgProducts []Domainorganizationproduct `json:"orgProducts"`


    SelfUri string `json:"selfUri"`

}

// Userme
type Userme struct { 
    


    // Name
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // Chat
    Chat Chat `json:"chat"`


    // Department
    Department string `json:"department"`


    // Email
    Email string `json:"email"`


    


    // Addresses - Email addresses and phone numbers for this user
    Addresses []Contact `json:"addresses"`


    


    // Title
    Title string `json:"title"`


    // Username
    Username string `json:"username"`


    // Manager
    Manager User `json:"manager"`


    // Images
    Images []Image `json:"images"`


    // Version - Required when updating a user, this value should be the current version of the user.  The current version can be obtained with a GET on the user before doing a PATCH.
    Version int `json:"version"`


    // Certifications
    Certifications []string `json:"certifications"`


    // Biography
    Biography Biography `json:"biography"`


    // EmployerInfo
    EmployerInfo Employerinfo `json:"employerInfo"`


    // PreferredName - Preferred full name of the agent
    PreferredName string `json:"preferredName"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    // AcdAutoAnswer - acd auto answer
    AcdAutoAnswer bool `json:"acdAutoAnswer"`


    


    // LastTokenIssued
    LastTokenIssued Oauthlasttokenissued `json:"lastTokenIssued"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Userme) String() string {
    
    
    
    
    
     o.Addresses = []Contact{{}} 
    
    
    
     o.Images = []Image{{}} 
    
     o.Certifications = []string{""} 
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userme) MarshalJSON() ([]byte, error) {
    type Alias Userme

    if UsermeMarshalled {
        return []byte("{}"), nil
    }
    UsermeMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        Chat Chat `json:"chat"`
        
        Department string `json:"department"`
        
        Email string `json:"email"`
        
        Addresses []Contact `json:"addresses"`
        
        Title string `json:"title"`
        
        Username string `json:"username"`
        
        Manager User `json:"manager"`
        
        Images []Image `json:"images"`
        
        Version int `json:"version"`
        
        Certifications []string `json:"certifications"`
        
        Biography Biography `json:"biography"`
        
        EmployerInfo Employerinfo `json:"employerInfo"`
        
        PreferredName string `json:"preferredName"`
        
        AcdAutoAnswer bool `json:"acdAutoAnswer"`
        
        LastTokenIssued Oauthlasttokenissued `json:"lastTokenIssued"`
        *Alias
    }{

        


        


        


        


        


        


        


        
        Addresses: []Contact{{}},
        


        


        


        


        


        
        Images: []Image{{}},
        


        


        
        Certifications: []string{""},
        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

