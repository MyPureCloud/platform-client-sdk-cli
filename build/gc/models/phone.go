package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PhoneMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PhoneDud struct { 
    Id string `json:"id"`


    


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    ModifiedBy string `json:"modifiedBy"`


    CreatedBy string `json:"createdBy"`


    State string `json:"state"`


    ModifiedByApp string `json:"modifiedByApp"`


    CreatedByApp string `json:"createdByApp"`


    


    


    


    


    


    


    


    UserAgentInfo Useragentinfo `json:"userAgentInfo"`


    


    


    


    StandAlone bool `json:"standAlone"`


    


    


    SelfUri string `json:"selfUri"`

}

// Phone
type Phone struct { 
    


    // Name - The name of the entity.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // Description - The resource's description.
    Description string `json:"description"`


    // Version - The current version of the resource.
    Version int `json:"version"`


    


    


    


    


    


    


    


    // Site - The site associated to the phone.
    Site Domainentityref `json:"site"`


    // PhoneBaseSettings - Phone Base Settings
    PhoneBaseSettings Phonebasesettings `json:"phoneBaseSettings"`


    // LineBaseSettings
    LineBaseSettings Domainentityref `json:"lineBaseSettings"`


    // PhoneMetaBase
    PhoneMetaBase Domainentityref `json:"phoneMetaBase"`


    // Lines - Lines
    Lines []Line `json:"lines"`


    // Status - The status of the phone and lines from the primary Edge.
    Status Phonestatus `json:"status"`


    // SecondaryStatus - The status of the phone and lines from the secondary Edge.
    SecondaryStatus Phonestatus `json:"secondaryStatus"`


    


    // Properties
    Properties map[string]interface{} `json:"properties"`


    // Capabilities
    Capabilities Phonecapabilities `json:"capabilities"`


    // WebRtcUser - This is the user associated with a WebRTC type phone.  It is required for all WebRTC phones.
    WebRtcUser Domainentityref `json:"webRtcUser"`


    


    // PrimaryEdge
    PrimaryEdge Edge `json:"primaryEdge"`


    // SecondaryEdge
    SecondaryEdge Edge `json:"secondaryEdge"`


    

}

// String returns a JSON representation of the model
func (o *Phone) String() string {
    
    
    
    
    
    
    
    
     o.Lines = []Line{{}} 
    
    
     o.Properties = map[string]interface{}{"": Interface{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Phone) MarshalJSON() ([]byte, error) {
    type Alias Phone

    if PhoneMarshalled {
        return []byte("{}"), nil
    }
    PhoneMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        Description string `json:"description"`
        
        Version int `json:"version"`
        
        Site Domainentityref `json:"site"`
        
        PhoneBaseSettings Phonebasesettings `json:"phoneBaseSettings"`
        
        LineBaseSettings Domainentityref `json:"lineBaseSettings"`
        
        PhoneMetaBase Domainentityref `json:"phoneMetaBase"`
        
        Lines []Line `json:"lines"`
        
        Status Phonestatus `json:"status"`
        
        SecondaryStatus Phonestatus `json:"secondaryStatus"`
        
        Properties map[string]interface{} `json:"properties"`
        
        Capabilities Phonecapabilities `json:"capabilities"`
        
        WebRtcUser Domainentityref `json:"webRtcUser"`
        
        PrimaryEdge Edge `json:"primaryEdge"`
        
        SecondaryEdge Edge `json:"secondaryEdge"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        
        Lines: []Line{{}},
        


        


        


        


        
        Properties: map[string]interface{}{"": Interface{}},
        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

