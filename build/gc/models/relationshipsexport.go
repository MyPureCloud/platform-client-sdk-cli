package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RelationshipsexportMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RelationshipsexportDud struct { 
    Id string `json:"id"`


    


    CreatedBy Domainentityref `json:"createdBy"`


    DateCreated time.Time `json:"dateCreated"`


    DateCompletion time.Time `json:"dateCompletion"`


    Status string `json:"status"`


    DownloadUrl string `json:"downloadUrl"`


    ResultRowCount int `json:"resultRowCount"`


    


    SelfUri string `json:"selfUri"`

}

// Relationshipsexport
type Relationshipsexport struct { 
    


    // DivisionIds - Division IDs of entities
    DivisionIds []string `json:"divisionIds"`


    


    


    


    


    


    


    // QueryConditions - Query conditions to apply on export
    QueryConditions Relationshipsexportqueryconditions `json:"queryConditions"`


    

}

// String returns a JSON representation of the model
func (o *Relationshipsexport) String() string {
     o.DivisionIds = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Relationshipsexport) MarshalJSON() ([]byte, error) {
    type Alias Relationshipsexport

    if RelationshipsexportMarshalled {
        return []byte("{}"), nil
    }
    RelationshipsexportMarshalled = true

    return json.Marshal(&struct {
        
        DivisionIds []string `json:"divisionIds"`
        
        QueryConditions Relationshipsexportqueryconditions `json:"queryConditions"`
        *Alias
    }{

        


        
        DivisionIds: []string{""},
        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

