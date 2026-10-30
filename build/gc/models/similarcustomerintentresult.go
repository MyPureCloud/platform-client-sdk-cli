package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SimilarcustomerintentresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SimilarcustomerintentresultDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Similarcustomerintentresult
type Similarcustomerintentresult struct { 
    


    // Name - Name of the customer intent
    Name string `json:"name"`


    // Description - Description of the customer intent
    Description string `json:"description"`


    // ExpiryTime - Expiry time in hours of the customer intent
    ExpiryTime int `json:"expiryTime"`


    // Distance - The distance score indicating similarity (lower is more similar)
    Distance float64 `json:"distance"`


    // Category - The category of the similar customer intent
    Category Similarentityref `json:"category"`


    // DateCreated - Creation date of the customer intent. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    

}

// String returns a JSON representation of the model
func (o *Similarcustomerintentresult) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Similarcustomerintentresult) MarshalJSON() ([]byte, error) {
    type Alias Similarcustomerintentresult

    if SimilarcustomerintentresultMarshalled {
        return []byte("{}"), nil
    }
    SimilarcustomerintentresultMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        ExpiryTime int `json:"expiryTime"`
        
        Distance float64 `json:"distance"`
        
        Category Similarentityref `json:"category"`
        
        DateCreated time.Time `json:"dateCreated"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

