package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DataingestionruleresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DataingestionruleresponseDud struct { 
    


    


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    Platform string `json:"platform"`


    SelfUri string `json:"selfUri"`

}

// Dataingestionruleresponse
type Dataingestionruleresponse struct { 
    // Id - ID of the data ingestion rule.
    Id string `json:"id"`


    // Name - The name of the data ingestion rule.
    Name string `json:"name"`


    // Description - A description of the data ingestion rule.
    Description string `json:"description"`


    // Status - The status of the data ingestion rule.
    Status string `json:"status"`


    // Version - The version number of the data ingestion rule.
    Version int `json:"version"`


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Dataingestionruleresponse) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dataingestionruleresponse) MarshalJSON() ([]byte, error) {
    type Alias Dataingestionruleresponse

    if DataingestionruleresponseMarshalled {
        return []byte("{}"), nil
    }
    DataingestionruleresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Status string `json:"status"`
        
        Version int `json:"version"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

