package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DomainpermissioncollectionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DomainpermissioncollectionDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Domainpermissioncollection
type Domainpermissioncollection struct { 
    


    // Name
    Name string `json:"name"`


    // Domain
    Domain string `json:"domain"`


    // PermissionMap
    PermissionMap map[string][]Domainpermission `json:"permissionMap"`


    

}

// String returns a JSON representation of the model
func (o *Domainpermissioncollection) String() string {
    
    
     o.PermissionMap = map[string][]Domainpermission{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Domainpermissioncollection) MarshalJSON() ([]byte, error) {
    type Alias Domainpermissioncollection

    if DomainpermissioncollectionMarshalled {
        return []byte("{}"), nil
    }
    DomainpermissioncollectionMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Domain string `json:"domain"`
        
        PermissionMap map[string][]Domainpermission `json:"permissionMap"`
        *Alias
    }{

        


        


        


        
        PermissionMap: map[string][]Domainpermission{"": {}},
        


        

        Alias: (*Alias)(u),
    })
}

