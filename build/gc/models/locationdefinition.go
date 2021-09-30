package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LocationdefinitionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LocationdefinitionDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    Path []string `json:"path"`


    ProfileImage []Locationimage `json:"profileImage"`


    FloorplanImage []Locationimage `json:"floorplanImage"`


    AddressVerificationDetails Locationaddressverificationdetails `json:"addressVerificationDetails"`


    AddressVerified bool `json:"addressVerified"`


    AddressStored bool `json:"addressStored"`


    


    SelfUri string `json:"selfUri"`

}

// Locationdefinition
type Locationdefinition struct { 
    


    // Name
    Name string `json:"name"`


    // ContactUser - Site contact for the location entity
    ContactUser Addressableentityref `json:"contactUser"`


    // EmergencyNumber - Emergency number for the location entity
    EmergencyNumber Locationemergencynumber `json:"emergencyNumber"`


    // Address
    Address Locationaddress `json:"address"`


    // State - Current state of the location entity
    State string `json:"state"`


    // Notes - Notes for the location entity
    Notes string `json:"notes"`


    // Version - Current version of the location entity, value to be supplied should be retrieved by a GET or on create/update response
    Version int `json:"version"`


    


    


    


    


    


    


    // Images
    Images string `json:"images"`


    

}

// String returns a JSON representation of the model
func (o *Locationdefinition) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Locationdefinition) MarshalJSON() ([]byte, error) {
    type Alias Locationdefinition

    if LocationdefinitionMarshalled {
        return []byte("{}"), nil
    }
    LocationdefinitionMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        ContactUser Addressableentityref `json:"contactUser"`
        
        EmergencyNumber Locationemergencynumber `json:"emergencyNumber"`
        
        Address Locationaddress `json:"address"`
        
        State string `json:"state"`
        
        Notes string `json:"notes"`
        
        Version int `json:"version"`
        
        
        
        
        
        
        
        
        
        
        
        
        
        Images string `json:"images"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

