package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DataingestionrulesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DataingestionrulesDud struct { 
    


    


    

}

// Dataingestionrules
type Dataingestionrules struct { 
    // Twitter - A list of X (formally Twitter) data ingestion rules.
    Twitter []Twitterdataingestionruleresponse `json:"twitter"`


    // Open - A list of Open data ingestion rules.
    Open []Opendataingestionruleresponse `json:"open"`


    // Facebook - A list of Facebook data ingestion rules.
    Facebook []Facebookdataingestionruleresponse `json:"facebook"`

}

// String returns a JSON representation of the model
func (o *Dataingestionrules) String() string {
     o.Twitter = []Twitterdataingestionruleresponse{{}} 
     o.Open = []Opendataingestionruleresponse{{}} 
     o.Facebook = []Facebookdataingestionruleresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dataingestionrules) MarshalJSON() ([]byte, error) {
    type Alias Dataingestionrules

    if DataingestionrulesMarshalled {
        return []byte("{}"), nil
    }
    DataingestionrulesMarshalled = true

    return json.Marshal(&struct {
        
        Twitter []Twitterdataingestionruleresponse `json:"twitter"`
        
        Open []Opendataingestionruleresponse `json:"open"`
        
        Facebook []Facebookdataingestionruleresponse `json:"facebook"`
        *Alias
    }{

        
        Twitter: []Twitterdataingestionruleresponse{{}},
        


        
        Open: []Opendataingestionruleresponse{{}},
        


        
        Facebook: []Facebookdataingestionruleresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

