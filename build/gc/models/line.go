package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LineMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LineDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    State string `json:"state"`


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Line
type Line struct { 
    


    // Name - The name of the entity.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


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


    // Properties
    Properties map[string]interface{} `json:"properties"`


    // EdgeGroup
    EdgeGroup Domainentityref `json:"edgeGroup"`


    // Template
    Template Domainentityref `json:"template"`


    // Site
    Site Domainentityref `json:"site"`


    // LineBaseSettings
    LineBaseSettings Domainentityref `json:"lineBaseSettings"`


    // PrimaryEdge - The primary edge associated to the line. (Deprecated)
    PrimaryEdge Edge `json:"primaryEdge"`


    // SecondaryEdge - The secondary edge associated to the line. (Deprecated)
    SecondaryEdge Edge `json:"secondaryEdge"`


    // LoggedInUser
    LoggedInUser Domainentityref `json:"loggedInUser"`


    // DefaultForUser
    DefaultForUser Domainentityref `json:"defaultForUser"`


    

}

// String returns a JSON representation of the model
func (o *Line) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Properties = map[string]interface{}{"": Interface{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Line) MarshalJSON() ([]byte, error) {
    type Alias Line

    if LineMarshalled {
        return []byte("{}"), nil
    }
    LineMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        Description string `json:"description"`
        
        Version int `json:"version"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        ModifiedBy string `json:"modifiedBy"`
        
        CreatedBy string `json:"createdBy"`
        
        
        
        ModifiedByApp string `json:"modifiedByApp"`
        
        CreatedByApp string `json:"createdByApp"`
        
        Properties map[string]interface{} `json:"properties"`
        
        EdgeGroup Domainentityref `json:"edgeGroup"`
        
        Template Domainentityref `json:"template"`
        
        Site Domainentityref `json:"site"`
        
        LineBaseSettings Domainentityref `json:"lineBaseSettings"`
        
        PrimaryEdge Edge `json:"primaryEdge"`
        
        SecondaryEdge Edge `json:"secondaryEdge"`
        
        LoggedInUser Domainentityref `json:"loggedInUser"`
        
        DefaultForUser Domainentityref `json:"defaultForUser"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Properties: map[string]interface{}{"": Interface{}},
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}
