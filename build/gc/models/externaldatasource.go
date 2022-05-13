package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExternaldatasourceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExternaldatasourceDud struct { 
    


    

}

// Externaldatasource - Describes a link to a record in an external system that contributed data to a Relate record
type Externaldatasource struct { 
    // Platform - The platform that was the source of the data.  Example: a CRM like SALESFORCE.
    Platform string `json:"platform"`


    // Url - An URL that links to the source record that contributed data to the associated entity.
    Url string `json:"url"`

}

// String returns a JSON representation of the model
func (o *Externaldatasource) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Externaldatasource) MarshalJSON() ([]byte, error) {
    type Alias Externaldatasource

    if ExternaldatasourceMarshalled {
        return []byte("{}"), nil
    }
    ExternaldatasourceMarshalled = true

    return json.Marshal(&struct {
        
        Platform string `json:"platform"`
        
        Url string `json:"url"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

