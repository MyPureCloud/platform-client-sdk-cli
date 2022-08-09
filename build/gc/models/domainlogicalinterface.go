package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DomainlogicalinterfaceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DomainlogicalinterfaceDud struct { 
    Id string `json:"id"`


    


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    ModifiedBy string `json:"modifiedBy"`


    CreatedBy string `json:"createdBy"`


    State string `json:"state"`


    ModifiedByApp string `json:"modifiedByApp"`


    CreatedByApp string `json:"createdByApp"`


    


    


    


    


    


    


    


    InterfaceType string `json:"interfaceType"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    UseForWanInterface bool `json:"useForWanInterface"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Domainlogicalinterface
type Domainlogicalinterface struct { 
    


    // Name - The name of the entity.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // Description - The resource's description.
    Description string `json:"description"`


    // Version - The current version of the resource.
    Version int `json:"version"`


    


    


    


    


    


    


    


    // EdgeUri
    EdgeUri string `json:"edgeUri"`


    // EdgeAssignedId
    EdgeAssignedId string `json:"edgeAssignedId"`


    // FriendlyName - Friendly Name
    FriendlyName string `json:"friendlyName"`


    // VlanTagId
    VlanTagId int `json:"vlanTagId"`


    // HardwareAddress - Hardware Address
    HardwareAddress string `json:"hardwareAddress"`


    // PhysicalAdapterId - Physical Adapter Id
    PhysicalAdapterId string `json:"physicalAdapterId"`


    // IfStatus
    IfStatus string `json:"ifStatus"`


    


    // PublicNatAddressIpV4 - IPv4 NENT IP Address
    PublicNatAddressIpV4 string `json:"publicNatAddressIpV4"`


    // PublicNatAddressIpV6 - IPv6 NENT IP Address
    PublicNatAddressIpV6 string `json:"publicNatAddressIpV6"`


    // Routes - The list of routes assigned to this interface.
    Routes []Domainnetworkroute `json:"routes"`


    // Addresses - The list of IP addresses on this interface.  Priority of dns addresses are based on order in the list.
    Addresses []Domainnetworkaddress `json:"addresses"`


    // Ipv4Capabilities - IPv4 interface settings.
    Ipv4Capabilities Domaincapabilities `json:"ipv4Capabilities"`


    // Ipv6Capabilities - IPv6 interface settings.
    Ipv6Capabilities Domaincapabilities `json:"ipv6Capabilities"`


    // CurrentState
    CurrentState string `json:"currentState"`


    // LastModifiedUserId
    LastModifiedUserId string `json:"lastModifiedUserId"`


    // LastModifiedCorrelationId
    LastModifiedCorrelationId string `json:"lastModifiedCorrelationId"`


    // CommandResponses
    CommandResponses []Domainnetworkcommandresponse `json:"commandResponses"`


    // InheritPhoneTrunkBasesIPv4 - The IPv4 phone trunk base assignment will be inherited from the Edge Group.
    InheritPhoneTrunkBasesIPv4 bool `json:"inheritPhoneTrunkBasesIPv4"`


    // InheritPhoneTrunkBasesIPv6 - The IPv6 phone trunk base assignment will be inherited from the Edge Group.
    InheritPhoneTrunkBasesIPv6 bool `json:"inheritPhoneTrunkBasesIPv6"`


    // UseForInternalEdgeCommunication - This interface will be used for all internal edge-to-edge communication using settings from the edgeTrunkBaseAssignment on the Edge Group.
    UseForInternalEdgeCommunication bool `json:"useForInternalEdgeCommunication"`


    // UseForIndirectEdgeCommunication - Site Interconnects using the \"Indirect\" method will communicate using the Public IP Address specified on the interface. Use this option when a NAT enabled firewall is between the Edge and the far end.
    UseForIndirectEdgeCommunication bool `json:"useForIndirectEdgeCommunication"`


    // UseForCloudProxyEdgeCommunication - Site Interconnects using the \"Cloud Proxy\" method will broker the connection between them with a Cloud Proxy. This method is required for connections between one or more Sites using Cloud Media, but can optionally be used between two premises Sites if Direct or Indirect are not an option.
    UseForCloudProxyEdgeCommunication bool `json:"useForCloudProxyEdgeCommunication"`


    


    // ExternalTrunkBaseAssignments - External trunk base settings to use for external communication from this interface.
    ExternalTrunkBaseAssignments []Trunkbaseassignment `json:"externalTrunkBaseAssignments"`


    // PhoneTrunkBaseAssignments - Phone trunk base settings to use for phone communication from this interface.  These settings will be ignored when \"inheritPhoneTrunkBases\" is true.
    PhoneTrunkBaseAssignments []Trunkbaseassignment `json:"phoneTrunkBaseAssignments"`


    // TraceEnabled
    TraceEnabled bool `json:"traceEnabled"`


    // StartDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartDate time.Time `json:"startDate"`


    // EndDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndDate time.Time `json:"endDate"`


    

}

// String returns a JSON representation of the model
func (o *Domainlogicalinterface) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Routes = []Domainnetworkroute{{}} 
     o.Addresses = []Domainnetworkaddress{{}} 
    
    
    
    
    
     o.CommandResponses = []Domainnetworkcommandresponse{{}} 
    
    
    
    
    
     o.ExternalTrunkBaseAssignments = []Trunkbaseassignment{{}} 
     o.PhoneTrunkBaseAssignments = []Trunkbaseassignment{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Domainlogicalinterface) MarshalJSON() ([]byte, error) {
    type Alias Domainlogicalinterface

    if DomainlogicalinterfaceMarshalled {
        return []byte("{}"), nil
    }
    DomainlogicalinterfaceMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        Description string `json:"description"`
        
        Version int `json:"version"`
        
        EdgeUri string `json:"edgeUri"`
        
        EdgeAssignedId string `json:"edgeAssignedId"`
        
        FriendlyName string `json:"friendlyName"`
        
        VlanTagId int `json:"vlanTagId"`
        
        HardwareAddress string `json:"hardwareAddress"`
        
        PhysicalAdapterId string `json:"physicalAdapterId"`
        
        IfStatus string `json:"ifStatus"`
        
        PublicNatAddressIpV4 string `json:"publicNatAddressIpV4"`
        
        PublicNatAddressIpV6 string `json:"publicNatAddressIpV6"`
        
        Routes []Domainnetworkroute `json:"routes"`
        
        Addresses []Domainnetworkaddress `json:"addresses"`
        
        Ipv4Capabilities Domaincapabilities `json:"ipv4Capabilities"`
        
        Ipv6Capabilities Domaincapabilities `json:"ipv6Capabilities"`
        
        CurrentState string `json:"currentState"`
        
        LastModifiedUserId string `json:"lastModifiedUserId"`
        
        LastModifiedCorrelationId string `json:"lastModifiedCorrelationId"`
        
        CommandResponses []Domainnetworkcommandresponse `json:"commandResponses"`
        
        InheritPhoneTrunkBasesIPv4 bool `json:"inheritPhoneTrunkBasesIPv4"`
        
        InheritPhoneTrunkBasesIPv6 bool `json:"inheritPhoneTrunkBasesIPv6"`
        
        UseForInternalEdgeCommunication bool `json:"useForInternalEdgeCommunication"`
        
        UseForIndirectEdgeCommunication bool `json:"useForIndirectEdgeCommunication"`
        
        UseForCloudProxyEdgeCommunication bool `json:"useForCloudProxyEdgeCommunication"`
        
        ExternalTrunkBaseAssignments []Trunkbaseassignment `json:"externalTrunkBaseAssignments"`
        
        PhoneTrunkBaseAssignments []Trunkbaseassignment `json:"phoneTrunkBaseAssignments"`
        
        TraceEnabled bool `json:"traceEnabled"`
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        
        Routes: []Domainnetworkroute{{}},
        


        
        Addresses: []Domainnetworkaddress{{}},
        


        


        


        


        


        


        
        CommandResponses: []Domainnetworkcommandresponse{{}},
        


        


        


        


        


        


        


        
        ExternalTrunkBaseAssignments: []Trunkbaseassignment{{}},
        


        
        PhoneTrunkBaseAssignments: []Trunkbaseassignment{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

