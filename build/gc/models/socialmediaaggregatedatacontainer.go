package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SocialmediaaggregatedatacontainerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SocialmediaaggregatedatacontainerDud struct { 
    


    

}

// Socialmediaaggregatedatacontainer
type Socialmediaaggregatedatacontainer struct { 
    // Group - A mapping from dimension to value
    Group map[string]string `json:"group"`


    // Data
    Data []Socialmediastatisticalresponse `json:"data"`

}

// String returns a JSON representation of the model
func (o *Socialmediaaggregatedatacontainer) String() string {
     o.Group = map[string]string{"": ""} 
     o.Data = []Socialmediastatisticalresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Socialmediaaggregatedatacontainer) MarshalJSON() ([]byte, error) {
    type Alias Socialmediaaggregatedatacontainer

    if SocialmediaaggregatedatacontainerMarshalled {
        return []byte("{}"), nil
    }
    SocialmediaaggregatedatacontainerMarshalled = true

    return json.Marshal(&struct {
        
        Group map[string]string `json:"group"`
        
        Data []Socialmediastatisticalresponse `json:"data"`
        *Alias
    }{

        
        Group: map[string]string{"": ""},
        


        
        Data: []Socialmediastatisticalresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

