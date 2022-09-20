package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BusinessunitresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BusinessunitresponseDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Businessunitresponse
type Businessunitresponse struct { 
    


    // Name
    Name string `json:"name"`


    // Settings - Settings for this business unit
    Settings Businessunitsettingsresponse `json:"settings"`


    // Division - The division to which this entity belongs.
    Division Divisionreference `json:"division"`


    

}

// String returns a JSON representation of the model
func (o *Businessunitresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Businessunitresponse) MarshalJSON() ([]byte, error) {
    type Alias Businessunitresponse

    if BusinessunitresponseMarshalled {
        return []byte("{}"), nil
    }
    BusinessunitresponseMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Settings Businessunitsettingsresponse `json:"settings"`
        
        Division Divisionreference `json:"division"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

