package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GdprsubjectMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GdprsubjectDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Gdprsubject
type Gdprsubject struct { 
    // Name
    Name string `json:"name"`


    // UserId
    UserId string `json:"userId"`


    // ExternalContactId
    ExternalContactId string `json:"externalContactId"`


    // DialerContactId
    DialerContactId Dialercontactid `json:"dialerContactId"`


    // JourneyCustomer
    JourneyCustomer Gdprjourneycustomer `json:"journeyCustomer"`


    // SocialHandle
    SocialHandle Socialhandle `json:"socialHandle"`


    // ExternalId
    ExternalId string `json:"externalId"`


    // Addresses
    Addresses []string `json:"addresses"`


    // PhoneNumbers
    PhoneNumbers []string `json:"phoneNumbers"`


    // EmailAddresses
    EmailAddresses []string `json:"emailAddresses"`

}

// String returns a JSON representation of the model
func (o *Gdprsubject) String() string {
    
    
    
    
    
    
    
     o.Addresses = []string{""} 
     o.PhoneNumbers = []string{""} 
     o.EmailAddresses = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Gdprsubject) MarshalJSON() ([]byte, error) {
    type Alias Gdprsubject

    if GdprsubjectMarshalled {
        return []byte("{}"), nil
    }
    GdprsubjectMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        UserId string `json:"userId"`
        
        ExternalContactId string `json:"externalContactId"`
        
        DialerContactId Dialercontactid `json:"dialerContactId"`
        
        JourneyCustomer Gdprjourneycustomer `json:"journeyCustomer"`
        
        SocialHandle Socialhandle `json:"socialHandle"`
        
        ExternalId string `json:"externalId"`
        
        Addresses []string `json:"addresses"`
        
        PhoneNumbers []string `json:"phoneNumbers"`
        
        EmailAddresses []string `json:"emailAddresses"`
        *Alias
    }{

        


        


        


        


        


        


        


        
        Addresses: []string{""},
        


        
        PhoneNumbers: []string{""},
        


        
        EmailAddresses: []string{""},
        

        Alias: (*Alias)(u),
    })
}

