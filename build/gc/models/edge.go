package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgeDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    State string `json:"state"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    CallDrainingState string `json:"callDrainingState"`


    ConversationCount int `json:"conversationCount"`


    


    OfflineConfigCalled bool `json:"offlineConfigCalled"`


    OsName string `json:"osName"`


    SelfUri string `json:"selfUri"`

}

// Edge
type Edge struct { 
    


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


    // Interfaces - The list of interfaces for the edge. (Deprecated) Replaced by configuring trunks/ip info on the logical interface instead
    Interfaces []Edgeinterface `json:"interfaces"`


    // Make
    Make string `json:"make"`


    // Model
    Model string `json:"model"`


    // ApiVersion
    ApiVersion string `json:"apiVersion"`


    // SoftwareVersion
    SoftwareVersion string `json:"softwareVersion"`


    // SoftwareVersionTimestamp
    SoftwareVersionTimestamp string `json:"softwareVersionTimestamp"`


    // SoftwareVersionPlatform
    SoftwareVersionPlatform string `json:"softwareVersionPlatform"`


    // SoftwareVersionConfiguration
    SoftwareVersionConfiguration string `json:"softwareVersionConfiguration"`


    // FullSoftwareVersion
    FullSoftwareVersion string `json:"fullSoftwareVersion"`


    // PairingId - The pairing Id for a hardware Edge in the format: 00000-00000-00000-00000-00000. This field is only required when creating an Edge with a deployment type of HARDWARE.
    PairingId string `json:"pairingId"`


    // Fingerprint
    Fingerprint string `json:"fingerprint"`


    // FingerprintHint
    FingerprintHint string `json:"fingerprintHint"`


    // CurrentVersion
    CurrentVersion string `json:"currentVersion"`


    // StagedVersion
    StagedVersion string `json:"stagedVersion"`


    // Patch
    Patch string `json:"patch"`


    // StatusCode - The current status of the Edge.
    StatusCode string `json:"statusCode"`


    // EdgeGroup
    EdgeGroup Edgegroup `json:"edgeGroup"`


    // Site - The Site to which the Edge is assigned.
    Site Site `json:"site"`


    // SoftwareStatus - Details about an in-progress or recently in-progress Edge software upgrade. This node appears only if a software upgrade was recently initiated for this Edge.
    SoftwareStatus Domainedgesoftwareupdatedto `json:"softwareStatus"`


    // OnlineStatus
    OnlineStatus string `json:"onlineStatus"`


    // SerialNumber
    SerialNumber string `json:"serialNumber"`


    // PhysicalEdge
    PhysicalEdge bool `json:"physicalEdge"`


    // Managed
    Managed bool `json:"managed"`


    // EdgeDeploymentType
    EdgeDeploymentType string `json:"edgeDeploymentType"`


    


    


    // Proxy - Edge HTTP proxy configuration for the WAN port. The field can be a hostname, FQDN, IPv4 or IPv6 address. If port is not included, port 80 is assumed.
    Proxy string `json:"proxy"`


    


    


    

}

// String returns a JSON representation of the model
func (o *Edge) String() string {
    
    
    
    
    
    
    
    
    
    
     o.Interfaces = []Edgeinterface{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edge) MarshalJSON() ([]byte, error) {
    type Alias Edge

    if EdgeMarshalled {
        return []byte("{}"), nil
    }
    EdgeMarshalled = true

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
        
        Interfaces []Edgeinterface `json:"interfaces"`
        
        Make string `json:"make"`
        
        Model string `json:"model"`
        
        ApiVersion string `json:"apiVersion"`
        
        SoftwareVersion string `json:"softwareVersion"`
        
        SoftwareVersionTimestamp string `json:"softwareVersionTimestamp"`
        
        SoftwareVersionPlatform string `json:"softwareVersionPlatform"`
        
        SoftwareVersionConfiguration string `json:"softwareVersionConfiguration"`
        
        FullSoftwareVersion string `json:"fullSoftwareVersion"`
        
        PairingId string `json:"pairingId"`
        
        Fingerprint string `json:"fingerprint"`
        
        FingerprintHint string `json:"fingerprintHint"`
        
        CurrentVersion string `json:"currentVersion"`
        
        StagedVersion string `json:"stagedVersion"`
        
        Patch string `json:"patch"`
        
        StatusCode string `json:"statusCode"`
        
        EdgeGroup Edgegroup `json:"edgeGroup"`
        
        Site Site `json:"site"`
        
        SoftwareStatus Domainedgesoftwareupdatedto `json:"softwareStatus"`
        
        OnlineStatus string `json:"onlineStatus"`
        
        SerialNumber string `json:"serialNumber"`
        
        PhysicalEdge bool `json:"physicalEdge"`
        
        Managed bool `json:"managed"`
        
        EdgeDeploymentType string `json:"edgeDeploymentType"`
        
        Proxy string `json:"proxy"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        
        Interfaces: []Edgeinterface{{}},
        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

