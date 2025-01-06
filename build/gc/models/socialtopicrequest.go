package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SocialtopicrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SocialtopicrequestDud struct { 
    


    


    

}

// Socialtopicrequest
type Socialtopicrequest struct { 
    // Name - Name of the social topic.
    Name string `json:"name"`


    // Description - A description of the social topic.
    Description string `json:"description"`


    // DivisionId - The ID of the division the social topic belongs to.
    DivisionId string `json:"divisionId"`

}

// String returns a JSON representation of the model
func (o *Socialtopicrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Socialtopicrequest) MarshalJSON() ([]byte, error) {
    type Alias Socialtopicrequest

    if SocialtopicrequestMarshalled {
        return []byte("{}"), nil
    }
    SocialtopicrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        DivisionId string `json:"divisionId"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

