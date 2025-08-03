package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ClientusagequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ClientusagequeryresponseDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Clientusagequeryresponse
type Clientusagequeryresponse struct { 
    // Id - The jobId of the query.
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // ResultsUri - The uri to get the results.
    ResultsUri string `json:"resultsUri"`


    

}

// String returns a JSON representation of the model
func (o *Clientusagequeryresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Clientusagequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Clientusagequeryresponse

    if ClientusagequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    ClientusagequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        ResultsUri string `json:"resultsUri"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

