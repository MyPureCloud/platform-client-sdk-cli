package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SearchshifttradesresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SearchshifttradesresponseDud struct { 
    


    

}

// Searchshifttradesresponse
type Searchshifttradesresponse struct { 
    // Trades - The shift trades that match the search criteria
    Trades []Searchshifttraderesponse `json:"trades"`


    // DownloadUrl - URL from which to fetch results for requests with a large result set. If populated, the downloaded data will conform to the same schema as would normally be returned, excepting downloaded data will never itself contain a downloadUrl
    DownloadUrl string `json:"downloadUrl"`

}

// String returns a JSON representation of the model
func (o *Searchshifttradesresponse) String() string {
     o.Trades = []Searchshifttraderesponse{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Searchshifttradesresponse) MarshalJSON() ([]byte, error) {
    type Alias Searchshifttradesresponse

    if SearchshifttradesresponseMarshalled {
        return []byte("{}"), nil
    }
    SearchshifttradesresponseMarshalled = true

    return json.Marshal(&struct {
        
        Trades []Searchshifttraderesponse `json:"trades"`
        
        DownloadUrl string `json:"downloadUrl"`
        *Alias
    }{

        
        Trades: []Searchshifttraderesponse{{}},
        


        

        Alias: (*Alias)(u),
    })
}

