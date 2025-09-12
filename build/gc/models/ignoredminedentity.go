package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IgnoredminedentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IgnoredminedentityDud struct { 
    


    


    


    


    


    

}

// Ignoredminedentity
type Ignoredminedentity struct { 
    // Id - Unique identifier for the ignored entity
    Id string `json:"id"`


    // Text - Text of the ignored entity
    Text string `json:"text"`


    // Participant - Type of participant
    Participant string `json:"participant"`


    // DateCreated - Date when the ignored entity was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Date when the ignored entity was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // MediaType - Media Type for the entity (Optional)
    MediaType string `json:"mediaType"`

}

// String returns a JSON representation of the model
func (o *Ignoredminedentity) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ignoredminedentity) MarshalJSON() ([]byte, error) {
    type Alias Ignoredminedentity

    if IgnoredminedentityMarshalled {
        return []byte("{}"), nil
    }
    IgnoredminedentityMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Text string `json:"text"`
        
        Participant string `json:"participant"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        MediaType string `json:"mediaType"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

