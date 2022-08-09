package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DomainphysicalinterfaceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DomainphysicalinterfaceDud struct { 
    Id string `json:"id"`


    


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    ModifiedBy string `json:"modifiedBy"`


    CreatedBy string `json:"createdBy"`


    State string `json:"state"`


    ModifiedByApp string `json:"modifiedByApp"`


    CreatedByApp string `json:"createdByApp"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Domainphysicalinterface
type Domainphysicalinterface struct { 
    


    // Name - The name of the entity.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // Description - The resource's description.
    Description string `json:"description"`


    // Version - The current version of the resource.
    Version int `json:"version"`


    


    


    


    


    


    


    


    // EdgeUri
    EdgeUri string `json:"edgeUri"`


    // FriendlyName
    FriendlyName string `json:"friendlyName"`


    // HardwareAddress
    HardwareAddress string `json:"hardwareAddress"`


    // PortLabel
    PortLabel string `json:"portLabel"`


    // PhysicalCapabilities
    PhysicalCapabilities Domainphysicalcapabilities `json:"physicalCapabilities"`


    

}

// String returns a JSON representation of the model
func (o *Domainphysicalinterface) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Domainphysicalinterface) MarshalJSON() ([]byte, error) {
    type Alias Domainphysicalinterface

    if DomainphysicalinterfaceMarshalled {
        return []byte("{}"), nil
    }
    DomainphysicalinterfaceMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        Description string `json:"description"`
        
        Version int `json:"version"`
        
        EdgeUri string `json:"edgeUri"`
        
        FriendlyName string `json:"friendlyName"`
        
        HardwareAddress string `json:"hardwareAddress"`
        
        PortLabel string `json:"portLabel"`
        
        PhysicalCapabilities Domainphysicalcapabilities `json:"physicalCapabilities"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

