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


    


    


    


    


    


    


    


    State string `json:"state"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Domainphysicalinterface
type Domainphysicalinterface struct { 
    


    // Name - The name of the entity.
    Name string `json:"name"`


    // Description - The resource's description.
    Description string `json:"description"`


    // Version - The current version of the resource.
    Version int `json:"version"`


    // DateCreated - The date the resource was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - The date of the last modification to the resource. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // ModifiedBy - The ID of the user that last modified the resource.
    ModifiedBy string `json:"modifiedBy"`


    // CreatedBy - The ID of the user that created the resource.
    CreatedBy string `json:"createdBy"`


    


    // ModifiedByApp - The application that last modified the resource.
    ModifiedByApp string `json:"modifiedByApp"`


    // CreatedByApp - The application that created the resource.
    CreatedByApp string `json:"createdByApp"`


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
        
        Description string `json:"description"`
        
        Version int `json:"version"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        ModifiedBy string `json:"modifiedBy"`
        
        CreatedBy string `json:"createdBy"`
        
        
        
        ModifiedByApp string `json:"modifiedByApp"`
        
        CreatedByApp string `json:"createdByApp"`
        
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

