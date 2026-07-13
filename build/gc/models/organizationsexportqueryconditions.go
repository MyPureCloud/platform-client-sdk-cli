package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OrganizationsexportqueryconditionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OrganizationsexportqueryconditionsDud struct { 
    


    

}

// Organizationsexportqueryconditions
type Organizationsexportqueryconditions struct { 
    // Filters - Filters to apply on export
    Filters *Organizationsexportfilter `json:"filters"`


    // Limit - Maximum result count in export, default is 180 000 000
    Limit int `json:"limit"`

}

// String returns a JSON representation of the model
func (o *Organizationsexportqueryconditions) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Organizationsexportqueryconditions) MarshalJSON() ([]byte, error) {
    type Alias Organizationsexportqueryconditions

    if OrganizationsexportqueryconditionsMarshalled {
        return []byte("{}"), nil
    }
    OrganizationsexportqueryconditionsMarshalled = true

    return json.Marshal(&struct {
        
        Filters *Organizationsexportfilter `json:"filters"`
        
        Limit int `json:"limit"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

