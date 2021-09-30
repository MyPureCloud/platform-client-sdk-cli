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


    


    


    


    


    


    


    


    State string `json:"state"`


    


    


    


    


    


    


    


    


    


    UserAgentInfo Useragentinfo `json:"userAgentInfo"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Phone
type Phone struct { 
    


    // Name - The name of the entity.
    Name string `json:"name"`


    // Description - The resource's description.
    Description string `json:"description"`


    // Version - The current version of the resource.
    Version int `json:"version"`


    // DateCreated - The date the resource was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - The date of the last modification to the resource. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // ModifiedBy - The ID of the user that last modified the resource.
    ModifiedBy string `json:"modifiedBy"`


    // CreatedBy - The ID of the user that created the resource.
    CreatedBy string `json:"createdBy"`


    


    // ModifiedByApp - The application that last modified the resource.
    ModifiedByApp string `json:"modifiedByApp"`


    // CreatedByApp - The application that created the resource.
    CreatedByApp string `json:"createdByApp"`


    // Site - The site associated to the phone.
    Site Domainentityref `json:"site"`


    // PhoneBaseSettings - Phone Base Settings
    PhoneBaseSettings Domainentityref `json:"phoneBaseSettings"`


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
        
        Description string `json:"description"`
        
        Version int `json:"version"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        ModifiedBy string `json:"modifiedBy"`
        
        CreatedBy string `json:"createdBy"`
        
        
        
        ModifiedByApp string `json:"modifiedByApp"`
        
        CreatedByApp string `json:"createdByApp"`
        
        Site Domainentityref `json:"site"`
        
        PhoneBaseSettings Domainentityref `json:"phoneBaseSettings"`
        
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

