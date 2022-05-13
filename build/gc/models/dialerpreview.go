package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DialerpreviewMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DialerpreviewDud struct { 
    


    


    


    


    

}

// Dialerpreview
type Dialerpreview struct { 
    // Id
    Id string `json:"id"`


    // ContactId - The contact associated with this preview data pop
    ContactId string `json:"contactId"`


    // ContactListId - The contactList associated with this preview data pop.
    ContactListId string `json:"contactListId"`


    // CampaignId - The campaignId associated with this preview data pop.
    CampaignId string `json:"campaignId"`


    // PhoneNumberColumns - The phone number columns associated with this campaign
    PhoneNumberColumns []Phonenumbercolumn `json:"phoneNumberColumns"`

}

// String returns a JSON representation of the model
func (o *Dialerpreview) String() string {
    
    
    
    
     o.PhoneNumberColumns = []Phonenumbercolumn{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dialerpreview) MarshalJSON() ([]byte, error) {
    type Alias Dialerpreview

    if DialerpreviewMarshalled {
        return []byte("{}"), nil
    }
    DialerpreviewMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        ContactId string `json:"contactId"`
        
        ContactListId string `json:"contactListId"`
        
        CampaignId string `json:"campaignId"`
        
        PhoneNumberColumns []Phonenumbercolumn `json:"phoneNumberColumns"`
        *Alias
    }{

        


        


        


        


        
        PhoneNumberColumns: []Phonenumbercolumn{{}},
        

        Alias: (*Alias)(u),
    })
}

