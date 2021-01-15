package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudEndpointDataTypeValue represents the possible values for attribute "type".
type CloudEndpointDataTypeValue string

const (
	// CloudEndpointDataTypeGateway represents the value Gateway.
	CloudEndpointDataTypeGateway CloudEndpointDataTypeValue = "Gateway"

	// CloudEndpointDataTypeInstance represents the value Instance.
	CloudEndpointDataTypeInstance CloudEndpointDataTypeValue = "Instance"

	// CloudEndpointDataTypeLoadBalancer represents the value LoadBalancer.
	CloudEndpointDataTypeLoadBalancer CloudEndpointDataTypeValue = "LoadBalancer"

	// CloudEndpointDataTypePeeringConnection represents the value PeeringConnection.
	CloudEndpointDataTypePeeringConnection CloudEndpointDataTypeValue = "PeeringConnection"

	// CloudEndpointDataTypeService represents the value Service.
	CloudEndpointDataTypeService CloudEndpointDataTypeValue = "Service"

	// CloudEndpointDataTypeTransitGateway represents the value TransitGateway.
	CloudEndpointDataTypeTransitGateway CloudEndpointDataTypeValue = "TransitGateway"
)

// CloudEndpointData represents the model of a cloudendpointdata
type CloudEndpointData struct {
	// Indicates that the endpoint is directly attached to the VPC. In this case the
	// attachedInterfaces is empty. In general this is only valid for endpoint type
	// Gateway and Peering Connection.
	VPCAttached bool `json:"VPCAttached" msgpack:"VPCAttached" bson:"vpcattached" mapstructure:"VPCAttached,omitempty"`

	// The list of VPCs that this endpoint is directly attached to.
	VPCAttachments []string `json:"VPCAttachments" msgpack:"VPCAttachments" bson:"vpcattachments" mapstructure:"VPCAttachments,omitempty"`

	// List of route tables associated with this endpoint. Depending on cloud provider
	// it can apply in some gateways.
	AssociatedRouteTables []string `json:"associatedRouteTables" msgpack:"associatedRouteTables" bson:"associatedroutetables" mapstructure:"associatedRouteTables,omitempty"`

	// A list of interfaces attached with the endpoint. In some cases endpoints can
	// have more than one interface.
	AttachedInterfaces []string `json:"attachedInterfaces" msgpack:"attachedInterfaces" bson:"attachedinterfaces" mapstructure:"attachedInterfaces,omitempty"`

	// If the endpoint has multiple connections and forwarding can be enabled between
	// them.
	ForwardingEnabled bool `json:"forwardingEnabled" msgpack:"forwardingEnabled" bson:"forwardingenabled" mapstructure:"forwardingEnabled,omitempty"`

	// Type of the endpoint.
	Type CloudEndpointDataTypeValue `json:"type" msgpack:"type" bson:"type" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudEndpointData returns a new *CloudEndpointData
func NewCloudEndpointData() *CloudEndpointData {

	return &CloudEndpointData{
		ModelVersion:          1,
		AssociatedRouteTables: []string{},
		VPCAttachments:        []string{},
		AttachedInterfaces:    []string{},
		ForwardingEnabled:     true,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudEndpointData) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudEndpointData{}

	s.VPCAttached = o.VPCAttached
	s.VPCAttachments = o.VPCAttachments
	s.AssociatedRouteTables = o.AssociatedRouteTables
	s.AttachedInterfaces = o.AttachedInterfaces
	s.ForwardingEnabled = o.ForwardingEnabled
	s.Type = o.Type

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudEndpointData) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudEndpointData{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.VPCAttached = s.VPCAttached
	o.VPCAttachments = s.VPCAttachments
	o.AssociatedRouteTables = s.AssociatedRouteTables
	o.AttachedInterfaces = s.AttachedInterfaces
	o.ForwardingEnabled = s.ForwardingEnabled
	o.Type = s.Type

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudEndpointData) BleveType() string {

	return "cloudendpointdata"
}

// DeepCopy returns a deep copy if the CloudEndpointData.
func (o *CloudEndpointData) DeepCopy() *CloudEndpointData {

	if o == nil {
		return nil
	}

	out := &CloudEndpointData{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudEndpointData.
func (o *CloudEndpointData) DeepCopyInto(out *CloudEndpointData) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudEndpointData: %s", err))
	}

	*out = *target.(*CloudEndpointData)
}

// Validate valides the current information stored into the structure.
func (o *CloudEndpointData) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Instance", "LoadBalancer", "PeeringConnection", "Service", "Gateway", "TransitGateway"}, false); err != nil {
		errors = errors.Append(err)
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

type mongoAttributesCloudEndpointData struct {
	VPCAttached           bool                       `bson:"vpcattached"`
	VPCAttachments        []string                   `bson:"vpcattachments"`
	AssociatedRouteTables []string                   `bson:"associatedroutetables"`
	AttachedInterfaces    []string                   `bson:"attachedinterfaces"`
	ForwardingEnabled     bool                       `bson:"forwardingenabled"`
	Type                  CloudEndpointDataTypeValue `bson:"type"`
}