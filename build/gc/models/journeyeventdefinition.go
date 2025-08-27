package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyeventdefinitionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyeventdefinitionDud struct { 
    Id string `json:"id"`


    Name string `json:"name"`


    Source string `json:"source"`


    Rank int `json:"rank"`


    DisplayName string `json:"displayName"`


    Description string `json:"description"`


    JsonSchema Jsonschemadocument `json:"jsonSchema"`


    SelfUri string `json:"selfUri"`

}

// Journeyeventdefinition - An event definition used when creating journey views
type Journeyeventdefinition struct { 
    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Journeyeventdefinition) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyeventdefinition) MarshalJSON() ([]byte, error) {
    type Alias Journeyeventdefinition

    if JourneyeventdefinitionMarshalled {
        return []byte("{}"), nil
    }
    JourneyeventdefinitionMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

