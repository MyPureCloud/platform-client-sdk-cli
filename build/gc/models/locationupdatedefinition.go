package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LocationupdatedefinitionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LocationupdatedefinitionDud struct { 
    


    


    


    


    


    


    


    

}

// Locationupdatedefinition
type Locationupdatedefinition struct { 
    // Name - The name of the Location. Required for creates, not required for updates
    Name string `json:"name"`


    // Version - Current version of the location
    Version int `json:"version"`


    // State - Current activity status of the location.
    State string `json:"state"`


    // Path - A list of ancestor ids
    Path []string `json:"path"`


    // Notes - Notes for the location
    Notes string `json:"notes"`


    // ContactUser - The user id of the location contact
    ContactUser string `json:"contactUser"`


    // EmergencyNumber - Emergency number for the location
    EmergencyNumber Locationemergencynumber `json:"emergencyNumber"`


    // Address - Address of the location
    Address Locationaddress `json:"address"`

}

// String returns a JSON representation of the model
func (o *Locationupdatedefinition) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Path = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Locationupdatedefinition) MarshalJSON() ([]byte, error) {
    type Alias Locationupdatedefinition

    if LocationupdatedefinitionMarshalled {
        return []byte("{}"), nil
    }
    LocationupdatedefinitionMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        Version int `json:"version"`
        
        State string `json:"state"`
        
        Path []string `json:"path"`
        
        Notes string `json:"notes"`
        
        ContactUser string `json:"contactUser"`
        
        EmergencyNumber Locationemergencynumber `json:"emergencyNumber"`
        
        Address Locationaddress `json:"address"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        Path: []string{""},
        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

