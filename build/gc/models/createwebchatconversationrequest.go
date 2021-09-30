package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatewebchatconversationrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatewebchatconversationrequestDud struct { 
    


    


    


    


    


    

}

// Createwebchatconversationrequest
type Createwebchatconversationrequest struct { 
    // OrganizationId - The organization identifier.
    OrganizationId string `json:"organizationId"`


    // DeploymentId - The web chat Deployment ID which contains the appropriate settings for this chat conversation.
    DeploymentId string `json:"deploymentId"`


    // RoutingTarget - The routing information to use for the new chat conversation.
    RoutingTarget Webchatroutingtarget `json:"routingTarget"`


    // MemberInfo - The guest member info to use for the new chat conversation.
    MemberInfo Guestmemberinfo `json:"memberInfo"`


    // MemberAuthToken - If the guest member is an authenticated member (ie, not anonymous) his JWT is provided here. The token will have been previously generated with the \"POST /api/v2/signeddata\" resource.
    MemberAuthToken string `json:"memberAuthToken"`


    // JourneyContext - A subset of the Journey System's data relevant to this conversation/session request (for external linkage and internal usage/context).
    JourneyContext Journeycontext `json:"journeyContext"`

}

// String returns a JSON representation of the model
func (o *Createwebchatconversationrequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createwebchatconversationrequest) MarshalJSON() ([]byte, error) {
    type Alias Createwebchatconversationrequest

    if CreatewebchatconversationrequestMarshalled {
        return []byte("{}"), nil
    }
    CreatewebchatconversationrequestMarshalled = true

    return json.Marshal(&struct { 
        OrganizationId string `json:"organizationId"`
        
        DeploymentId string `json:"deploymentId"`
        
        RoutingTarget Webchatroutingtarget `json:"routingTarget"`
        
        MemberInfo Guestmemberinfo `json:"memberInfo"`
        
        MemberAuthToken string `json:"memberAuthToken"`
        
        JourneyContext Journeycontext `json:"journeyContext"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

