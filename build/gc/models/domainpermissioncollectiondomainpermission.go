package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DomainpermissioncollectiondomainpermissionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DomainpermissioncollectiondomainpermissionDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Domainpermissioncollectiondomainpermission
type Domainpermissioncollectiondomainpermission struct { 
    


    // Name
    Name string `json:"name"`


    // Domain
    Domain string `json:"domain"`


    // PermissionMap
    PermissionMap map[string][]Domainpermission `json:"permissionMap"`


    

}

// String returns a JSON representation of the model
func (o *Domainpermissioncollectiondomainpermission) String() string {
    
    
     o.PermissionMap = map[string][]Domainpermission{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Domainpermissioncollectiondomainpermission) MarshalJSON() ([]byte, error) {
    type Alias Domainpermissioncollectiondomainpermission

    if DomainpermissioncollectiondomainpermissionMarshalled {
        return []byte("{}"), nil
    }
    DomainpermissioncollectiondomainpermissionMarshalled = true

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

