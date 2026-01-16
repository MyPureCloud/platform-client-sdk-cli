package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactenrichrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactenrichrequestDud struct { 
    


    


    


    


    


    


    

}

// Contactenrichrequest
type Contactenrichrequest struct { 
    // Id - A user-specified tracker string, only useful in the Bulk-Enrich API. If one Bulk-Enrich operation in a request fails, the requested operation will be repeated in the Bulk API response, including this id field, allowing associating of request and response operations.
    Id string `json:"id"`


    // Division - The division to which this entity belongs.
    Division Writablestarrabledivision `json:"division"`


    // MatchingIdentifiers - An ordered list of one or more Identifiers which might each be claimed by a Contact. `action` describes what to do with any possibly matching Contacts. Identifier lookups will occur in the order specified here. Between 1 and 25.
    MatchingIdentifiers []Contactidentifier `json:"matchingIdentifiers"`


    // Action - The action that should be taken based on any Contacts found by `matchingIdentifiers`.
    Action string `json:"action"`


    // Contact - Data to be added, either as an update to an existing Contact or the body of a new Contact. Omitting a field in this contract means that it will be treated as null in the `fieldRules` logic.
    Contact Externalcontact `json:"contact"`


    // FieldRules - Logic describing how to combine data from the submitted request with data found in the database.
    FieldRules Enrichfieldrules `json:"fieldRules"`


    // Options - Additional options modifying the behavior of the API operation.
    Options Contactenrichoptions `json:"options"`

}

// String returns a JSON representation of the model
func (o *Contactenrichrequest) String() string {
    
    
     o.MatchingIdentifiers = []Contactidentifier{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactenrichrequest) MarshalJSON() ([]byte, error) {
    type Alias Contactenrichrequest

    if ContactenrichrequestMarshalled {
        return []byte("{}"), nil
    }
    ContactenrichrequestMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Division Writablestarrabledivision `json:"division"`
        
        MatchingIdentifiers []Contactidentifier `json:"matchingIdentifiers"`
        
        Action string `json:"action"`
        
        Contact Externalcontact `json:"contact"`
        
        FieldRules Enrichfieldrules `json:"fieldRules"`
        
        Options Contactenrichoptions `json:"options"`
        *Alias
    }{

        


        


        
        MatchingIdentifiers: []Contactidentifier{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

