package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExternalorganizationenrichrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExternalorganizationenrichrequestDud struct { 
    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Externalorganizationenrichrequest
type Externalorganizationenrichrequest struct { 
    // Id - A user-specified tracker string, only useful in the Bulk-Enrich API. If one Bulk-Enrich operation in a request fails, the requested operation will be repeated in the Bulk API response, including this id field, allowing associating request and response operations.
    Id string `json:"id"`


    // Division - The division to which this entity belongs.
    Division Writablestarrabledivision `json:"division"`


    // MatchingIdentifiers - An ordered list of one or more Identifiers which might each be claimed by an External Organization. `action` describes what to do with any possibly matching External Organization. Identifier lookups will occur in the order specified here.
    MatchingIdentifiers []Externalorganizationidentifier `json:"matchingIdentifiers"`


    // Action - The action that should be taken based on any External Organization found by `matchingIdentifiers`.
    Action string `json:"action"`


    // ExternalOrganization - Data to be added, either as an update to an existing External Organization or the body of a new External Organization. Omitting a field in this contract means that it will be treated as null in the `fieldRules` logic.
    ExternalOrganization Externalorganization `json:"externalOrganization"`


    // FieldRules - Logic describing how to combine data from the submitted request with data found in the database.
    FieldRules Enrichfieldrules `json:"fieldRules"`


    

}

// String returns a JSON representation of the model
func (o *Externalorganizationenrichrequest) String() string {
    
    
     o.MatchingIdentifiers = []Externalorganizationidentifier{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Externalorganizationenrichrequest) MarshalJSON() ([]byte, error) {
    type Alias Externalorganizationenrichrequest

    if ExternalorganizationenrichrequestMarshalled {
        return []byte("{}"), nil
    }
    ExternalorganizationenrichrequestMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Division Writablestarrabledivision `json:"division"`
        
        MatchingIdentifiers []Externalorganizationidentifier `json:"matchingIdentifiers"`
        
        Action string `json:"action"`
        
        ExternalOrganization Externalorganization `json:"externalOrganization"`
        
        FieldRules Enrichfieldrules `json:"fieldRules"`
        *Alias
    }{

        


        


        
        MatchingIdentifiers: []Externalorganizationidentifier{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

