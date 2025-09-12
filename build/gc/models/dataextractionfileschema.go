package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DataextractionfileschemaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DataextractionfileschemaDud struct { 
    


    


    


    

}

// Dataextractionfileschema
type Dataextractionfileschema struct { 
    // Id - The unique identifier for the file
    Id string `json:"id"`


    // DataSchema - The data schema the file belongs to
    DataSchema string `json:"dataSchema"`


    // DateCreated - The date and time when this file was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateExpires - The date and time when this file becomes unavailable. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateExpires time.Time `json:"dateExpires"`

}

// String returns a JSON representation of the model
func (o *Dataextractionfileschema) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dataextractionfileschema) MarshalJSON() ([]byte, error) {
    type Alias Dataextractionfileschema

    if DataextractionfileschemaMarshalled {
        return []byte("{}"), nil
    }
    DataextractionfileschemaMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        DataSchema string `json:"dataSchema"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateExpires time.Time `json:"dateExpires"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

