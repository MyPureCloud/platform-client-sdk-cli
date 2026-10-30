package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SimilarcustomerintentsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SimilarcustomerintentsrequestDud struct { 
    


    


    

}

// Similarcustomerintentsrequest
type Similarcustomerintentsrequest struct { 
    // Name - Name to find similar customer intents for
    Name string `json:"name"`


    // Description - Description to find similar customer intents for
    Description string `json:"description"`


    // CategoryId - Optional category ID to scope the similarity search
    CategoryId string `json:"categoryId"`

}

// String returns a JSON representation of the model
func (o *Similarcustomerintentsrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Similarcustomerintentsrequest) MarshalJSON() ([]byte, error) {
    type Alias Similarcustomerintentsrequest

    if SimilarcustomerintentsrequestMarshalled {
        return []byte("{}"), nil
    }
    SimilarcustomerintentsrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        CategoryId string `json:"categoryId"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

