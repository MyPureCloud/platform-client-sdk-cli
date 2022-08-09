package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SiteMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SiteDud struct { 
    Id string `json:"id"`


    


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    ModifiedBy string `json:"modifiedBy"`


    CreatedBy string `json:"createdBy"`


    State string `json:"state"`


    ModifiedByApp string `json:"modifiedByApp"`


    CreatedByApp string `json:"createdByApp"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Site
type Site struct { 
    


    // Name - The name of the entity.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // Description - The resource's description.
    Description string `json:"description"`


    // Version - The current version of the resource.
    Version int `json:"version"`


    


    


    


    


    


    


    


    // PrimarySites
    PrimarySites []Domainentityref `json:"primarySites"`


    // SecondarySites
    SecondarySites []Domainentityref `json:"secondarySites"`


    // PrimaryEdges
    PrimaryEdges []Edge `json:"primaryEdges"`


    // SecondaryEdges
    SecondaryEdges []Edge `json:"secondaryEdges"`


    // Addresses
    Addresses []Contact `json:"addresses"`


    // Edges
    Edges []Edge `json:"edges"`


    // EdgeAutoUpdateConfig - Recurrance rule, time zone, and start/end settings for automatic edge updates for this site
    EdgeAutoUpdateConfig Edgeautoupdateconfig `json:"edgeAutoUpdateConfig"`


    // MediaRegionsUseLatencyBased
    MediaRegionsUseLatencyBased bool `json:"mediaRegionsUseLatencyBased"`


    // Location - Location
    Location Locationdefinition `json:"location"`


    // Managed
    Managed bool `json:"managed"`


    // NtpSettings - Network Time Protocol settings for the site
    NtpSettings Ntpsettings `json:"ntpSettings"`


    // MediaModel - Media model for the site
    MediaModel string `json:"mediaModel"`


    // CoreSite - Is this site a core site
    CoreSite bool `json:"coreSite"`


    // SiteConnections - The site connections
    SiteConnections []Siteconnection `json:"siteConnections"`


    

}

// String returns a JSON representation of the model
func (o *Site) String() string {
    
    
    
    
     o.PrimarySites = []Domainentityref{{}} 
     o.SecondarySites = []Domainentityref{{}} 
     o.PrimaryEdges = []Edge{{}} 
     o.SecondaryEdges = []Edge{{}} 
     o.Addresses = []Contact{{}} 
     o.Edges = []Edge{{}} 
    
    
    
    
    
    
    
     o.SiteConnections = []Siteconnection{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Site) MarshalJSON() ([]byte, error) {
    type Alias Site

    if SiteMarshalled {
        return []byte("{}"), nil
    }
    SiteMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        Description string `json:"description"`
        
        Version int `json:"version"`
        
        PrimarySites []Domainentityref `json:"primarySites"`
        
        SecondarySites []Domainentityref `json:"secondarySites"`
        
        PrimaryEdges []Edge `json:"primaryEdges"`
        
        SecondaryEdges []Edge `json:"secondaryEdges"`
        
        Addresses []Contact `json:"addresses"`
        
        Edges []Edge `json:"edges"`
        
        EdgeAutoUpdateConfig Edgeautoupdateconfig `json:"edgeAutoUpdateConfig"`
        
        MediaRegionsUseLatencyBased bool `json:"mediaRegionsUseLatencyBased"`
        
        Location Locationdefinition `json:"location"`
        
        Managed bool `json:"managed"`
        
        NtpSettings Ntpsettings `json:"ntpSettings"`
        
        MediaModel string `json:"mediaModel"`
        
        CoreSite bool `json:"coreSite"`
        
        SiteConnections []Siteconnection `json:"siteConnections"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        
        PrimarySites: []Domainentityref{{}},
        


        
        SecondarySites: []Domainentityref{{}},
        


        
        PrimaryEdges: []Edge{{}},
        


        
        SecondaryEdges: []Edge{{}},
        


        
        Addresses: []Contact{{}},
        


        
        Edges: []Edge{{}},
        


        


        


        


        


        


        


        


        
        SiteConnections: []Siteconnection{{}},
        


        

        Alias: (*Alias)(u),
    })
}

