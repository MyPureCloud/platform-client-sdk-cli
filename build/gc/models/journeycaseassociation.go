package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneycaseassociationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneycaseassociationDud struct { 
    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Journeycaseassociation - The representation of a case association on a journey session.
type Journeycaseassociation struct { 
    // Id - The ID of the association.
    Id string `json:"id"`


    // AssociatedCase - The case that was associated with the journey session.
    AssociatedCase Addressableentityref `json:"associatedCase"`


    // CaseReference - The reference for the case that was associated with the journey session.
    CaseReference string `json:"caseReference"`


    // DateAssociated - The date of the association. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateAssociated time.Time `json:"dateAssociated"`


    

}

// String returns a JSON representation of the model
func (o *Journeycaseassociation) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeycaseassociation) MarshalJSON() ([]byte, error) {
    type Alias Journeycaseassociation

    if JourneycaseassociationMarshalled {
        return []byte("{}"), nil
    }
    JourneycaseassociationMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        AssociatedCase Addressableentityref `json:"associatedCase"`
        
        CaseReference string `json:"caseReference"`
        
        DateAssociated time.Time `json:"dateAssociated"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

