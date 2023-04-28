package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WritabledialercontactMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WritabledialercontactDud struct { 
    


    


    


    LatestSmsEvaluations map[string]Messageevaluation `json:"latestSmsEvaluations"`


    LatestEmailEvaluations map[string]Messageevaluation `json:"latestEmailEvaluations"`


    


    


    

}

// Writabledialercontact
type Writabledialercontact struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // ContactListId - The identifier of the contact list containing this contact.
    ContactListId string `json:"contactListId"`


    // Data - An ordered map of the contact's columns and corresponding values.
    Data map[string]string `json:"data"`


    


    


    // Callable - Indicates whether or not the contact can be called.
    Callable bool `json:"callable"`


    // PhoneNumberStatus - A map of phone number columns to PhoneNumberStatuses, which indicate if the phone number is callable or not.
    PhoneNumberStatus map[string]Phonenumberstatus `json:"phoneNumberStatus"`


    // ContactableStatus - A map of media types (Voice, SMS and Email) to ContactableStatus, which indicates if the contact can be contacted using the specified media type.
    ContactableStatus map[string]Contactablestatus `json:"contactableStatus"`

}

// String returns a JSON representation of the model
func (o *Writabledialercontact) String() string {
    
    
     o.Data = map[string]string{"": ""} 
    
     o.PhoneNumberStatus = map[string]Phonenumberstatus{"": {}} 
     o.ContactableStatus = map[string]Contactablestatus{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Writabledialercontact) MarshalJSON() ([]byte, error) {
    type Alias Writabledialercontact

    if WritabledialercontactMarshalled {
        return []byte("{}"), nil
    }
    WritabledialercontactMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        ContactListId string `json:"contactListId"`
        
        Data map[string]string `json:"data"`
        
        Callable bool `json:"callable"`
        
        PhoneNumberStatus map[string]Phonenumberstatus `json:"phoneNumberStatus"`
        
        ContactableStatus map[string]Contactablestatus `json:"contactableStatus"`
        *Alias
    }{

        


        


        
        Data: map[string]string{"": ""},
        


        


        


        


        
        PhoneNumberStatus: map[string]Phonenumberstatus{"": {}},
        


        
        ContactableStatus: map[string]Contactablestatus{"": {}},
        

        Alias: (*Alias)(u),
    })
}

