package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpensocialmediarecipientMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpensocialmediarecipientDud struct { 
    


    


    


    


    


    


    

}

// Opensocialmediarecipient - Information about the recipient the message is sent to or received from.
type Opensocialmediarecipient struct { 
    // Id - The recipient ID specific to the provider.
    Id string `json:"id"`


    // IdType - The recipient ID type. This is used to indicate the format used for the ID.
    IdType string `json:"idType"`


    // FirstName - First name of the recipient.
    FirstName string `json:"firstName"`


    // LastName - Last name of the recipient.
    LastName string `json:"lastName"`


    // Nickname - Nickname or display name of the recipient.
    Nickname string `json:"nickname"`


    // Image - URL of an image that represents the recipient.
    Image string `json:"image"`


    // AdditionalIds - List of recipient additional identifiers
    AdditionalIds []Opensocialmediarecipientadditionalidentifier `json:"additionalIds"`

}

// String returns a JSON representation of the model
func (o *Opensocialmediarecipient) String() string {
    
    
    
    
    
    
     o.AdditionalIds = []Opensocialmediarecipientadditionalidentifier{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Opensocialmediarecipient) MarshalJSON() ([]byte, error) {
    type Alias Opensocialmediarecipient

    if OpensocialmediarecipientMarshalled {
        return []byte("{}"), nil
    }
    OpensocialmediarecipientMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        IdType string `json:"idType"`
        
        FirstName string `json:"firstName"`
        
        LastName string `json:"lastName"`
        
        Nickname string `json:"nickname"`
        
        Image string `json:"image"`
        
        AdditionalIds []Opensocialmediarecipientadditionalidentifier `json:"additionalIds"`
        *Alias
    }{

        


        


        


        


        


        


        
        AdditionalIds: []Opensocialmediarecipientadditionalidentifier{{}},
        

        Alias: (*Alias)(u),
    })
}

