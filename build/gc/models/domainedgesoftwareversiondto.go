package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DomainedgesoftwareversiondtoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DomainedgesoftwareversiondtoDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Domainedgesoftwareversiondto
type Domainedgesoftwareversiondto struct { 
    


    // Name
    Name string `json:"name"`


    // EdgeVersion
    EdgeVersion string `json:"edgeVersion"`


    // PublishDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    PublishDate time.Time `json:"publishDate"`


    // EdgeUri
    EdgeUri string `json:"edgeUri"`


    // Current
    Current bool `json:"current"`


    // LatestRelease
    LatestRelease bool `json:"latestRelease"`


    

}

// String returns a JSON representation of the model
func (o *Domainedgesoftwareversiondto) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Domainedgesoftwareversiondto) MarshalJSON() ([]byte, error) {
    type Alias Domainedgesoftwareversiondto

    if DomainedgesoftwareversiondtoMarshalled {
        return []byte("{}"), nil
    }
    DomainedgesoftwareversiondtoMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        EdgeVersion string `json:"edgeVersion"`
        
        PublishDate time.Time `json:"publishDate"`
        
        EdgeUri string `json:"edgeUri"`
        
        Current bool `json:"current"`
        
        LatestRelease bool `json:"latestRelease"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

