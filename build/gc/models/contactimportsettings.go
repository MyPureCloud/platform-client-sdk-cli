package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactimportsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactimportsettingsDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    SelfUri string `json:"selfUri"`

}

// Contactimportsettings
type Contactimportsettings struct { 
    


    // Name - Display name for the settings
    Name string `json:"name"`


    // MatchingCriteria - Which fields you want to identity resolution based on. (e.g.: Email, Phone). It can be empty, populated only one of them or it can be both for now. The order of the items is important for identity resolution
    MatchingCriteria []string `json:"matchingCriteria"`


    // MergeContacts - Decides what happens when a record already found in the system. Action will be Upsert or Merge
    MergeContacts bool `json:"mergeContacts"`


    // ExternalSourceId - Define the corresponding source system by the customer, the customer can have different externalId source, they can collect this id from contact service
    ExternalSourceId string `json:"externalSourceId"`


    // ImportFields - Decides which field we need to send towards contact service
    ImportFields []Contactimportfield `json:"importFields"`


    


    

}

// String returns a JSON representation of the model
func (o *Contactimportsettings) String() string {
    
     o.MatchingCriteria = []string{""} 
    
    
     o.ImportFields = []Contactimportfield{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactimportsettings) MarshalJSON() ([]byte, error) {
    type Alias Contactimportsettings

    if ContactimportsettingsMarshalled {
        return []byte("{}"), nil
    }
    ContactimportsettingsMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        MatchingCriteria []string `json:"matchingCriteria"`
        
        MergeContacts bool `json:"mergeContacts"`
        
        ExternalSourceId string `json:"externalSourceId"`
        
        ImportFields []Contactimportfield `json:"importFields"`
        *Alias
    }{

        


        


        
        MatchingCriteria: []string{""},
        


        


        


        
        ImportFields: []Contactimportfield{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

