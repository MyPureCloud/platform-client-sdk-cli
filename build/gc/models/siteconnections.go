package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SiteconnectionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SiteconnectionsDud struct { 
    Id string `json:"id"`


    


    SiteConnections []Siteconnection `json:"siteConnections"`


    SelfUri string `json:"selfUri"`

}

// Siteconnections
type Siteconnections struct { 
    


    // Name
    Name string `json:"name"`


    


    

}

// String returns a JSON representation of the model
func (o *Siteconnections) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Siteconnections) MarshalJSON() ([]byte, error) {
    type Alias Siteconnections

    if SiteconnectionsMarshalled {
        return []byte("{}"), nil
    }
    SiteconnectionsMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

