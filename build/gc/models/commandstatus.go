package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CommandstatusMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CommandstatusDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Commandstatus
type Commandstatus struct { 
    


    // Name
    Name string `json:"name"`


    // Expiration - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Expiration time.Time `json:"expiration"`


    // UserId
    UserId string `json:"userId"`


    // StatusCode
    StatusCode string `json:"statusCode"`


    // CommandType
    CommandType string `json:"commandType"`


    // Document
    Document Document `json:"document"`


    

}

// String returns a JSON representation of the model
func (o *Commandstatus) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Commandstatus) MarshalJSON() ([]byte, error) {
    type Alias Commandstatus

    if CommandstatusMarshalled {
        return []byte("{}"), nil
    }
    CommandstatusMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Expiration time.Time `json:"expiration"`
        
        UserId string `json:"userId"`
        
        StatusCode string `json:"statusCode"`
        
        CommandType string `json:"commandType"`
        
        Document Document `json:"document"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

