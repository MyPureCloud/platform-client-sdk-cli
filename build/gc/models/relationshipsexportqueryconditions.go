package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RelationshipsexportqueryconditionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RelationshipsexportqueryconditionsDud struct { 
    


    

}

// Relationshipsexportqueryconditions
type Relationshipsexportqueryconditions struct { 
    // Filters - Filters to apply on export
    Filters *Relationshipsexportfilter `json:"filters"`


    // Limit - Maximum result count in export, default is 180 000 000
    Limit int `json:"limit"`

}

// String returns a JSON representation of the model
func (o *Relationshipsexportqueryconditions) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Relationshipsexportqueryconditions) MarshalJSON() ([]byte, error) {
    type Alias Relationshipsexportqueryconditions

    if RelationshipsexportqueryconditionsMarshalled {
        return []byte("{}"), nil
    }
    RelationshipsexportqueryconditionsMarshalled = true

    return json.Marshal(&struct {
        
        Filters *Relationshipsexportfilter `json:"filters"`
        
        Limit int `json:"limit"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

