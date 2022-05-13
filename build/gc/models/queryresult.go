package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryresultDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Queryresult
type Queryresult struct { 
    


    // Name
    Name string `json:"name"`


    // Body
    Body Domainentity `json:"body"`


    

}

// String returns a JSON representation of the model
func (o *Queryresult) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryresult) MarshalJSON() ([]byte, error) {
    type Alias Queryresult

    if QueryresultMarshalled {
        return []byte("{}"), nil
    }
    QueryresultMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Body Domainentity `json:"body"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

