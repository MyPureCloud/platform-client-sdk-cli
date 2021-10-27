package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AvailabletopicMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AvailabletopicDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Availabletopic
type Availabletopic struct { 
    // Description
    Description string `json:"description"`


    // Id
    Id string `json:"id"`


    // PermissionDetails - Full detailed permissions required to subscribe to the topic
    PermissionDetails []Permissiondetails `json:"permissionDetails"`


    // RequiresPermissions - Permissions required to subscribe to the topic
    RequiresPermissions []string `json:"requiresPermissions"`


    // RequiresDivisionPermissions - True if the subscribing user must belong to the same division as the topic object ID
    RequiresDivisionPermissions bool `json:"requiresDivisionPermissions"`


    // RequiresAnyValidator - If multiple permissions are required for this topic, such as both requiresCurrentUser and requiresDivisionPermissions, then true here indicates that meeting any one condition will satisfy the requirements; false indicates all conditions must be met.
    RequiresAnyValidator bool `json:"requiresAnyValidator"`


    // Enforced - Whether or not the permissions on this topic are enforced
    Enforced bool `json:"enforced"`


    // Visibility - Visibility of this topic (Public or Preview)
    Visibility string `json:"visibility"`


    // Schema
    Schema map[string]interface{} `json:"schema"`


    // RequiresCurrentUser - True if the topic user ID is required to match the subscribing user ID
    RequiresCurrentUser bool `json:"requiresCurrentUser"`


    // RequiresCurrentUserOrPermission - True if permissions are only required when the topic user ID does not match the subscribing user ID
    RequiresCurrentUserOrPermission bool `json:"requiresCurrentUserOrPermission"`


    // Transports - Transports that support events for the topic
    Transports []string `json:"transports"`


    // PublicApiTemplateUriPaths
    PublicApiTemplateUriPaths []string `json:"publicApiTemplateUriPaths"`


    // TopicParameters - Parameters in the topic name that can be substituted, in the order they appear in the topic name
    TopicParameters []string `json:"topicParameters"`

}

// String returns a JSON representation of the model
func (o *Availabletopic) String() string {
    
    
    
    
    
    
    
    
    
    
     o.PermissionDetails = []Permissiondetails{{}} 
    
    
    
     o.RequiresPermissions = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Schema = map[string]interface{}{"": Interface{}} 
    
    
    
    
    
    
    
    
    
    
    
     o.Transports = []string{""} 
    
    
    
     o.PublicApiTemplateUriPaths = []string{""} 
    
    
    
     o.TopicParameters = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Availabletopic) MarshalJSON() ([]byte, error) {
    type Alias Availabletopic

    if AvailabletopicMarshalled {
        return []byte("{}"), nil
    }
    AvailabletopicMarshalled = true

    return json.Marshal(&struct { 
        Description string `json:"description"`
        
        Id string `json:"id"`
        
        PermissionDetails []Permissiondetails `json:"permissionDetails"`
        
        RequiresPermissions []string `json:"requiresPermissions"`
        
        RequiresDivisionPermissions bool `json:"requiresDivisionPermissions"`
        
        RequiresAnyValidator bool `json:"requiresAnyValidator"`
        
        Enforced bool `json:"enforced"`
        
        Visibility string `json:"visibility"`
        
        Schema map[string]interface{} `json:"schema"`
        
        RequiresCurrentUser bool `json:"requiresCurrentUser"`
        
        RequiresCurrentUserOrPermission bool `json:"requiresCurrentUserOrPermission"`
        
        Transports []string `json:"transports"`
        
        PublicApiTemplateUriPaths []string `json:"publicApiTemplateUriPaths"`
        
        TopicParameters []string `json:"topicParameters"`
        
        *Alias
    }{
        

        

        

        

        

        
        PermissionDetails: []Permissiondetails{{}},
        

        

        
        RequiresPermissions: []string{""},
        

        

        

        

        

        

        

        

        

        

        
        Schema: map[string]interface{}{"": Interface{}},
        

        

        

        

        

        

        
        Transports: []string{""},
        

        

        
        PublicApiTemplateUriPaths: []string{""},
        

        

        
        TopicParameters: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}

