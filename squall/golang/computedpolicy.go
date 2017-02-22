package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

// ComputedPolicyIdentity represents the Identity of the object
var ComputedPolicyIdentity = elemental.Identity{
	Name:     "computedpolicy",
	Category: "computedpolicies",
}

// ComputedPoliciesList represents a list of ComputedPolicies
type ComputedPoliciesList []*ComputedPolicy

// ComputedPolicy represents the model of a computedpolicy
type ComputedPolicy struct {
	// ID is the identifier of the object.
	ID string `json:"ID" cql:"-" bson:"-"`

	// Array of netowrk access policies computed
	NetworkAccessPolicies []*NetworkAccessPolicy `json:"networkAccessPolicies" cql:"networkaccesspolicies,omitempty" bson:"networkaccesspolicies"`
}

// NewComputedPolicy returns a new *ComputedPolicy
func NewComputedPolicy() *ComputedPolicy {

	return &ComputedPolicy{}
}

// Identity returns the Identity of the object.
func (o *ComputedPolicy) Identity() elemental.Identity {

	return ComputedPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ComputedPolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ComputedPolicy) SetIdentifier(ID string) {

	o.ID = ID
}

func (o *ComputedPolicy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *ComputedPolicy) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (ComputedPolicy) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return ComputedPolicyAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (ComputedPolicy) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ComputedPolicyAttributesMap
}

// ComputedPolicyAttributesMap represents the map of attribute for ComputedPolicy.
var ComputedPolicyAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
		Unique:         true,
	},
	"NetworkAccessPolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Array of netowrk access policies computed`,
		Exposed:        true,
		Name:           "networkAccessPolicies",
		Orderable:      true,
		Stored:         true,
		SubType:        "network_access_policies_list",
		Type:           "external",
	},
}