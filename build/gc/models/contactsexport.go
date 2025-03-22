package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactsexportMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactsexportDud struct { 
    Id string `json:"id"`


    


    


    CreatedBy Domainentityref `json:"createdBy"`


    DateCreated time.Time `json:"dateCreated"`


    Status string `json:"status"`


    DownloadUrl string `json:"downloadUrl"`


    SelfUri string `json:"selfUri"`

}

// Contactsexport
type Contactsexport struct { 
    


    // DivisionIds - Division IDs of entities
    DivisionIds []string `json:"divisionIds"`


    // QueryConditions - Query conditions to apply on export
    QueryConditions Contactsexportqueryconditions `json:"queryConditions"`


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Contactsexport) String() string {
     o.DivisionIds = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactsexport) MarshalJSON() ([]byte, error) {
    type Alias Contactsexport

    if ContactsexportMarshalled {
        return []byte("{}"), nil
    }
    ContactsexportMarshalled = true

    return json.Marshal(&struct {
        
        DivisionIds []string `json:"divisionIds"`
        
        QueryConditions Contactsexportqueryconditions `json:"queryConditions"`
        *Alias
    }{

        


        
        DivisionIds: []string{""},
        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

