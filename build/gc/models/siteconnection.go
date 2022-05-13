package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SiteconnectionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SiteconnectionDud struct { 
    


    


    


    


    


    


    MediaModel string `json:"mediaModel"`


    EdgeList []Connectededge `json:"edgeList"`


    CoreSite bool `json:"coreSite"`


    PrimaryCoreSites []Domainentityref `json:"primaryCoreSites"`


    SecondaryCoreSites []Domainentityref `json:"secondaryCoreSites"`

}

// Siteconnection
type Siteconnection struct { 
    // Id
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // Managed
    Managed bool `json:"managed"`


    // VarType - Connection method from site to site (Direct, Indirect, CloudProxy
    VarType string `json:"type"`


    // Enabled - Indicates if the current site is linked
    Enabled bool `json:"enabled"`


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Siteconnection) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Siteconnection) MarshalJSON() ([]byte, error) {
    type Alias Siteconnection

    if SiteconnectionMarshalled {
        return []byte("{}"), nil
    }
    SiteconnectionMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        SelfUri string `json:"selfUri"`
        
        Managed bool `json:"managed"`
        
        VarType string `json:"type"`
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

