package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DidnumberMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DidnumberDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Didnumber - Represents an unassigned or assigned DID in a DID Pool.
type Didnumber struct { 
    


    // Name
    Name string `json:"name"`


    // Number - The number of the DID formatted as E164.
    Number string `json:"number"`


    // Assigned - True if this DID is assigned to an entity.  False otherwise.
    Assigned bool `json:"assigned"`


    // DidPool - A Uri reference to the DID Pool this DID is a part of.
    DidPool Addressableentityref `json:"didPool"`


    // Owner - A Uri reference to the owner of this DID.  The owner's type can be found in ownerType.  If the DID is unassigned, this will be NULL.
    Owner Domainentityref `json:"owner"`


    // OwnerType - The type of the entity that owns this DID.  If the DID is unassigned, this will be NULL.
    OwnerType string `json:"ownerType"`


    

}

// String returns a JSON representation of the model
func (o *Didnumber) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Didnumber) MarshalJSON() ([]byte, error) {
    type Alias Didnumber

    if DidnumberMarshalled {
        return []byte("{}"), nil
    }
    DidnumberMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Number string `json:"number"`
        
        Assigned bool `json:"assigned"`
        
        DidPool Addressableentityref `json:"didPool"`
        
        Owner Domainentityref `json:"owner"`
        
        OwnerType string `json:"ownerType"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

