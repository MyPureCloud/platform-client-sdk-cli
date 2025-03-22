package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SocialmediaasyncaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SocialmediaasyncaggregatequeryresponseDud struct { 
    


    

}

// Socialmediaasyncaggregatequeryresponse
type Socialmediaasyncaggregatequeryresponse struct { 
    // Results
    Results []Socialmediaaggregatedatacontainer `json:"results"`


    // Cursor - Cursor token to retrieve next page
    Cursor string `json:"cursor"`

}

// String returns a JSON representation of the model
func (o *Socialmediaasyncaggregatequeryresponse) String() string {
     o.Results = []Socialmediaaggregatedatacontainer{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Socialmediaasyncaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Socialmediaasyncaggregatequeryresponse

    if SocialmediaasyncaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    SocialmediaasyncaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Socialmediaaggregatedatacontainer `json:"results"`
        
        Cursor string `json:"cursor"`
        *Alias
    }{

        
        Results: []Socialmediaaggregatedatacontainer{{}},
        


        

        Alias: (*Alias)(u),
    })
}

