package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SocialmediaasyncdetailqueryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SocialmediaasyncdetailqueryresponseDud struct { 
    


    

}

// Socialmediaasyncdetailqueryresponse
type Socialmediaasyncdetailqueryresponse struct { 
    // Results
    Results []Socialmediadetaildatacontainer `json:"results"`


    // Cursor - Cursor token to retrieve next or previous page
    Cursor string `json:"cursor"`

}

// String returns a JSON representation of the model
func (o *Socialmediaasyncdetailqueryresponse) String() string {
     o.Results = []Socialmediadetaildatacontainer{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Socialmediaasyncdetailqueryresponse) MarshalJSON() ([]byte, error) {
    type Alias Socialmediaasyncdetailqueryresponse

    if SocialmediaasyncdetailqueryresponseMarshalled {
        return []byte("{}"), nil
    }
    SocialmediaasyncdetailqueryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Socialmediadetaildatacontainer `json:"results"`
        
        Cursor string `json:"cursor"`
        *Alias
    }{

        
        Results: []Socialmediadetaildatacontainer{{}},
        


        

        Alias: (*Alias)(u),
    })
}

