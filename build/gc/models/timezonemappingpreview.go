package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimezonemappingpreviewMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimezonemappingpreviewDud struct { 
    


    


    


    


    


    


    


    


    

}

// Timezonemappingpreview
type Timezonemappingpreview struct { 
    // ContactList - The associated ContactList
    ContactList Domainentityref `json:"contactList"`


    // ContactsPerTimeZone - The number of contacts per time zone that mapped to only that time zone
    ContactsPerTimeZone map[string]int `json:"contactsPerTimeZone"`


    // ContactsMappedUsingZipCode - The number of contacts per time zone that mapped to only that time zone and were mapped using the zip code column
    ContactsMappedUsingZipCode map[string]int `json:"contactsMappedUsingZipCode"`


    // ContactsMappedToASingleZone - The total number of contacts that mapped to a single time zone
    ContactsMappedToASingleZone int `json:"contactsMappedToASingleZone"`


    // ContactsMappedToASingleZoneUsingZipCode - The total number of contacts that mapped to a single time zone and were mapped using the zip code column
    ContactsMappedToASingleZoneUsingZipCode int `json:"contactsMappedToASingleZoneUsingZipCode"`


    // ContactsMappedToMultipleZones - The total number of contacts that mapped to multiple time zones
    ContactsMappedToMultipleZones int `json:"contactsMappedToMultipleZones"`


    // ContactsMappedToMultipleZonesUsingZipCode - The total number of contacts that mapped to multiple time zones and were mapped using the zip code column
    ContactsMappedToMultipleZonesUsingZipCode int `json:"contactsMappedToMultipleZonesUsingZipCode"`


    // ContactsInDefaultWindow - The total number of contacts that will be dialed during the default window
    ContactsInDefaultWindow int `json:"contactsInDefaultWindow"`


    // ContactListSize - The total number of contacts in the contact list
    ContactListSize int `json:"contactListSize"`

}

// String returns a JSON representation of the model
func (o *Timezonemappingpreview) String() string {
    
    
    
    
    
    
     o.ContactsPerTimeZone = map[string]int{"": 0} 
    
    
    
     o.ContactsMappedUsingZipCode = map[string]int{"": 0} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timezonemappingpreview) MarshalJSON() ([]byte, error) {
    type Alias Timezonemappingpreview

    if TimezonemappingpreviewMarshalled {
        return []byte("{}"), nil
    }
    TimezonemappingpreviewMarshalled = true

    return json.Marshal(&struct { 
        ContactList Domainentityref `json:"contactList"`
        
        ContactsPerTimeZone map[string]int `json:"contactsPerTimeZone"`
        
        ContactsMappedUsingZipCode map[string]int `json:"contactsMappedUsingZipCode"`
        
        ContactsMappedToASingleZone int `json:"contactsMappedToASingleZone"`
        
        ContactsMappedToASingleZoneUsingZipCode int `json:"contactsMappedToASingleZoneUsingZipCode"`
        
        ContactsMappedToMultipleZones int `json:"contactsMappedToMultipleZones"`
        
        ContactsMappedToMultipleZonesUsingZipCode int `json:"contactsMappedToMultipleZonesUsingZipCode"`
        
        ContactsInDefaultWindow int `json:"contactsInDefaultWindow"`
        
        ContactListSize int `json:"contactListSize"`
        
        *Alias
    }{
        

        

        

        
        ContactsPerTimeZone: map[string]int{"": 0},
        

        

        
        ContactsMappedUsingZipCode: map[string]int{"": 0},
        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

