package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpendataingestionruleversionresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpendataingestionruleversionresponseDud struct { 
    


    


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    Platform string `json:"platform"`


    Countries []string `json:"countries"`


    


    SelfUri string `json:"selfUri"`

}

// Opendataingestionruleversionresponse
type Opendataingestionruleversionresponse struct { 
    // Id - ID of the open data ingestion rule.
    Id string `json:"id"`


    // Name - The name of the data ingestion rule.
    Name string `json:"name"`


    // Description - A description of the data ingestion rule.
    Description string `json:"description"`


    // Status - The status of the data ingestion rule.
    Status string `json:"status"`


    // Version - The version number of the data ingestion rule.
    Version int `json:"version"`


    


    


    


    


    // ExternalSource - The external source associated with this open data ingestion rule, which is used when performing identity resolution
    ExternalSource Domainentityref `json:"externalSource"`


    

}

// String returns a JSON representation of the model
func (o *Opendataingestionruleversionresponse) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Opendataingestionruleversionresponse) MarshalJSON() ([]byte, error) {
    type Alias Opendataingestionruleversionresponse

    if OpendataingestionruleversionresponseMarshalled {
        return []byte("{}"), nil
    }
    OpendataingestionruleversionresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Status string `json:"status"`
        
        Version int `json:"version"`
        
        ExternalSource Domainentityref `json:"externalSource"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

