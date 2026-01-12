package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IntentscategoryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IntentscategoryDud struct { 
    Id string `json:"id"`


    


    


    DateCreated time.Time `json:"dateCreated"`


    SelfUri string `json:"selfUri"`

}

// Intentscategory
type Intentscategory struct { 
    


    // Name - Name of the category
    Name string `json:"name"`


    // Description - Description of the category
    Description string `json:"description"`


    


    

}

// String returns a JSON representation of the model
func (o *Intentscategory) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Intentscategory) MarshalJSON() ([]byte, error) {
    type Alias Intentscategory

    if IntentscategoryMarshalled {
        return []byte("{}"), nil
    }
    IntentscategoryMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

