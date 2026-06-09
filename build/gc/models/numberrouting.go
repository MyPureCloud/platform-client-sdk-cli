package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NumberroutingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NumberroutingDud struct { 
    


    


    


    


    


    


    


    

}

// Numberrouting
type Numberrouting struct { 
    // NumberId - Phone number Id that has a disaster recovery linking
    NumberId string `json:"numberId"`


    // OwnerOrganizationId - Owner organization of numberId
    OwnerOrganizationId string `json:"ownerOrganizationId"`


    // CarrierCode - Code that indicates which carrier manages the number ie. VERIZON
    CarrierCode string `json:"carrierCode"`


    // PendingOrganizationId - OrganizationId where the number will be routed to during a change routing event
    PendingOrganizationId string `json:"pendingOrganizationId"`


    // Region - The current region where the number is located
    Region string `json:"region"`


    // Status - The current status of the number routing
    Status string `json:"status"`


    // ActiveOrganizationId - The orgId where the number is currently routing to
    ActiveOrganizationId string `json:"activeOrganizationId"`


    // LinkedOrganizationIds - List of linked organizations ids
    LinkedOrganizationIds []string `json:"linkedOrganizationIds"`

}

// String returns a JSON representation of the model
func (o *Numberrouting) String() string {
    
    
    
    
    
    
    
     o.LinkedOrganizationIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Numberrouting) MarshalJSON() ([]byte, error) {
    type Alias Numberrouting

    if NumberroutingMarshalled {
        return []byte("{}"), nil
    }
    NumberroutingMarshalled = true

    return json.Marshal(&struct {
        
        NumberId string `json:"numberId"`
        
        OwnerOrganizationId string `json:"ownerOrganizationId"`
        
        CarrierCode string `json:"carrierCode"`
        
        PendingOrganizationId string `json:"pendingOrganizationId"`
        
        Region string `json:"region"`
        
        Status string `json:"status"`
        
        ActiveOrganizationId string `json:"activeOrganizationId"`
        
        LinkedOrganizationIds []string `json:"linkedOrganizationIds"`
        *Alias
    }{

        


        


        


        


        


        


        


        
        LinkedOrganizationIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

