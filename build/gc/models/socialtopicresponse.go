package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SocialtopicresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SocialtopicresponseDud struct { 
    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Socialtopicresponse
type Socialtopicresponse struct { 
    // Id - ID of the social topic.
    Id string `json:"id"`


    // Name - The name of the social topic.
    Name string `json:"name"`


    // Description - A description of the social topic.
    Description string `json:"description"`


    // DateCreated - Timestamp indicating when the social topic was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Timestamp indicating when the social topic was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // DivisionId - The ID of the division to which the social topic belongs to.
    DivisionId string `json:"divisionId"`


    // Status - The status of the social topic.
    Status string `json:"status"`


    

}

// String returns a JSON representation of the model
func (o *Socialtopicresponse) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Socialtopicresponse) MarshalJSON() ([]byte, error) {
    type Alias Socialtopicresponse

    if SocialtopicresponseMarshalled {
        return []byte("{}"), nil
    }
    SocialtopicresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        DivisionId string `json:"divisionId"`
        
        Status string `json:"status"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

