package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScimuserextensionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScimuserextensionsDud struct { 
    


    


    

}

// Scimuserextensions - Genesys Cloud user extensions to SCIM RFC.
type Scimuserextensions struct { 
    // RoutingSkills - The list of routing skills assigned to a user. Maximum 50 skills.
    RoutingSkills []Scimuserroutingskill `json:"routingSkills"`


    // RoutingLanguages - The list of routing languages assigned to a user. Maximum 50 languages.
    RoutingLanguages []Scimuserroutinglanguage `json:"routingLanguages"`


    // ExternalIds - The list of external identifiers assigned to user. Always includes an immutable SCIM authority prefixed with \"x-pc:scimv2:v1\". ExternalIds are searchable with complex filter query parameter using 'authority' and 'value', e.g., filter=urn:ietf:params:scim:schemas:extension:genesys:purecloud:2.0:User:externalIds[authority eq \"matchAuthName\" and value eq \"matchingExternalKeyValue\"].
    ExternalIds []Scimgenesysuserexternalid `json:"externalIds"`

}

// String returns a JSON representation of the model
func (o *Scimuserextensions) String() string {
    
    
     o.RoutingSkills = []Scimuserroutingskill{{}} 
    
    
    
     o.RoutingLanguages = []Scimuserroutinglanguage{{}} 
    
    
    
     o.ExternalIds = []Scimgenesysuserexternalid{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scimuserextensions) MarshalJSON() ([]byte, error) {
    type Alias Scimuserextensions

    if ScimuserextensionsMarshalled {
        return []byte("{}"), nil
    }
    ScimuserextensionsMarshalled = true

    return json.Marshal(&struct { 
        RoutingSkills []Scimuserroutingskill `json:"routingSkills"`
        
        RoutingLanguages []Scimuserroutinglanguage `json:"routingLanguages"`
        
        ExternalIds []Scimgenesysuserexternalid `json:"externalIds"`
        
        *Alias
    }{
        

        
        RoutingSkills: []Scimuserroutingskill{{}},
        

        

        
        RoutingLanguages: []Scimuserroutinglanguage{{}},
        

        

        
        ExternalIds: []Scimgenesysuserexternalid{{}},
        

        
        Alias: (*Alias)(u),
    })
}

