package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RelationshipMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RelationshipDud struct { 
    Id string `json:"id"`


    


    


    


    


    ExternalDataSources []Externaldatasource `json:"externalDataSources"`


    SelfUri string `json:"selfUri"`

}

// Relationship
type Relationship struct { 
    


    // Division - The division to which this entity belongs.
    Division Writablestarrabledivision `json:"division"`


    // User - The user associated with the external organization. When creating or updating a relationship, only User.id is required. User object is fully populated when expanding a note.
    User User `json:"user"`


    // ExternalOrganization - The external organization this relationship is attached to
    ExternalOrganization Externalorganization `json:"externalOrganization"`


    // Relationship - The relationship or role of the user to this external organization.Examples: Account Manager, Sales Engineer, Implementation Consultant
    Relationship string `json:"relationship"`


    


    

}

// String returns a JSON representation of the model
func (o *Relationship) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Relationship) MarshalJSON() ([]byte, error) {
    type Alias Relationship

    if RelationshipMarshalled {
        return []byte("{}"), nil
    }
    RelationshipMarshalled = true

    return json.Marshal(&struct {
        
        Division Writablestarrabledivision `json:"division"`
        
        User User `json:"user"`
        
        ExternalOrganization Externalorganization `json:"externalOrganization"`
        
        Relationship string `json:"relationship"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

