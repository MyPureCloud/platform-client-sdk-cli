package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactidentifierMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactidentifierDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Contactidentifier
type Contactidentifier struct { 
    


    // Division - The division to which this entity belongs.
    Division Writablestarrabledivision `json:"division"`


    // VarType - The type of this identifier
    VarType string `json:"type"`


    // Value - The string value of the identifier. Will vary in syntax by type.
    Value string `json:"value"`


    // DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // ExternalSource - The External Source ID of the identifier
    ExternalSource Externalsource `json:"externalSource"`


    

}

// String returns a JSON representation of the model
func (o *Contactidentifier) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactidentifier) MarshalJSON() ([]byte, error) {
    type Alias Contactidentifier

    if ContactidentifierMarshalled {
        return []byte("{}"), nil
    }
    ContactidentifierMarshalled = true

    return json.Marshal(&struct {
        
        Division Writablestarrabledivision `json:"division"`
        
        VarType string `json:"type"`
        
        Value string `json:"value"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        ExternalSource Externalsource `json:"externalSource"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

