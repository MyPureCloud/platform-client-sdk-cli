package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SimilarcategoryresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SimilarcategoryresultDud struct { 
    Id string `json:"id"`


    


    


    DateCreated time.Time `json:"dateCreated"`


    


    SelfUri string `json:"selfUri"`

}

// Similarcategoryresult
type Similarcategoryresult struct { 
    


    // Name - Name of the category
    Name string `json:"name"`


    // Description - Description of the category
    Description string `json:"description"`


    


    // Distance - The distance score indicating similarity (lower is more similar)
    Distance float64 `json:"distance"`


    

}

// String returns a JSON representation of the model
func (o *Similarcategoryresult) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Similarcategoryresult) MarshalJSON() ([]byte, error) {
    type Alias Similarcategoryresult

    if SimilarcategoryresultMarshalled {
        return []byte("{}"), nil
    }
    SimilarcategoryresultMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Distance float64 `json:"distance"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

