package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpendataingestionrulerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpendataingestionrulerequestDud struct { 
    


    


    

}

// Opendataingestionrulerequest
type Opendataingestionrulerequest struct { 
    // Name - The name of the data ingestion rule.
    Name string `json:"name"`


    // Description - A description of the data ingestion rule.
    Description string `json:"description"`


    // ExternalSource - The external source associated with this open data ingestion rule, which will be used when performing identity resolution
    ExternalSource Domainentityref `json:"externalSource"`

}

// String returns a JSON representation of the model
func (o *Opendataingestionrulerequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Opendataingestionrulerequest) MarshalJSON() ([]byte, error) {
    type Alias Opendataingestionrulerequest

    if OpendataingestionrulerequestMarshalled {
        return []byte("{}"), nil
    }
    OpendataingestionrulerequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        ExternalSource Domainentityref `json:"externalSource"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

