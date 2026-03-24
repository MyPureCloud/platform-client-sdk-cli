package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CaseassociationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CaseassociationDud struct { 
    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`


    

}

// Caseassociation - Represents an association between a case and an interaction
type Caseassociation struct { 
    // Id - The ID of the association.
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // AssociationType - Association type.
    AssociationType string `json:"associationType"`


    // DateAssociated - Interaction association date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateAssociated time.Time `json:"dateAssociated"`


    // Workitem - Associated workitem ID.
    Workitem Workitemreference `json:"workitem"`


    // Conversation - Associated conversation ID.
    Conversation Conversationreference `json:"conversation"`


    // Stage - The stage related to this association.
    Stage Stagereference `json:"stage"`


    // Step - The step related to this association.
    Step Stepreference `json:"step"`


    


    // VarCase - Case ID
    VarCase Casereference `json:"case"`

}

// String returns a JSON representation of the model
func (o *Caseassociation) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Caseassociation) MarshalJSON() ([]byte, error) {
    type Alias Caseassociation

    if CaseassociationMarshalled {
        return []byte("{}"), nil
    }
    CaseassociationMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        AssociationType string `json:"associationType"`
        
        DateAssociated time.Time `json:"dateAssociated"`
        
        Workitem Workitemreference `json:"workitem"`
        
        Conversation Conversationreference `json:"conversation"`
        
        Stage Stagereference `json:"stage"`
        
        Step Stepreference `json:"step"`
        
        VarCase Casereference `json:"case"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

