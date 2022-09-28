package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DialercontactMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DialercontactDud struct { 
    Id string `json:"id"`


    


    


    


    CallRecords map[string]Callrecord `json:"callRecords"`


    LatestSmsEvaluations map[string]Messageevaluation `json:"latestSmsEvaluations"`


    


    


    


    ContactColumnTimeZones map[string]Contactcolumntimezone `json:"contactColumnTimeZones"`


    ConfigurationOverrides Configurationoverrides `json:"configurationOverrides"`


    SelfUri string `json:"selfUri"`

}

// Dialercontact
type Dialercontact struct { 
    


    // Name
    Name string `json:"name"`


    // ContactListId - The identifier of the contact list containing this contact.
    ContactListId string `json:"contactListId"`


    // Data - An ordered map of the contact's columns and corresponding values.
    Data map[string]interface{} `json:"data"`


    


    


    // Callable - Indicates whether or not the contact can be called.
    Callable bool `json:"callable"`


    // PhoneNumberStatus - A map of phone number columns to PhoneNumberStatuses, which indicate if the phone number is callable or not.
    PhoneNumberStatus map[string]Phonenumberstatus `json:"phoneNumberStatus"`


    // ContactableStatus - A map of media types (Voice, SMS and Email) to ContactableStatus, which indicates if the contact can be contacted using the specified media type.
    ContactableStatus map[string]Contactablestatus `json:"contactableStatus"`


    


    


    

}

// String returns a JSON representation of the model
func (o *Dialercontact) String() string {
    
    
     o.Data = map[string]interface{}{"": Interface{}} 
    
     o.PhoneNumberStatus = map[string]Phonenumberstatus{"": {}} 
     o.ContactableStatus = map[string]Contactablestatus{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dialercontact) MarshalJSON() ([]byte, error) {
    type Alias Dialercontact

    if DialercontactMarshalled {
        return []byte("{}"), nil
    }
    DialercontactMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ContactListId string `json:"contactListId"`
        
        Data map[string]interface{} `json:"data"`
        
        Callable bool `json:"callable"`
        
        PhoneNumberStatus map[string]Phonenumberstatus `json:"phoneNumberStatus"`
        
        ContactableStatus map[string]Contactablestatus `json:"contactableStatus"`
        *Alias
    }{

        


        


        


        
        Data: map[string]interface{}{"": Interface{}},
        


        


        


        


        
        PhoneNumberStatus: map[string]Phonenumberstatus{"": {}},
        


        
        ContactableStatus: map[string]Contactablestatus{"": {}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

