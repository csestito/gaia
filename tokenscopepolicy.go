package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// TokenScopePolicyIdentity represents the Identity of the object.
var TokenScopePolicyIdentity = elemental.Identity{
	Name:     "tokenscopepolicy",
	Category: "tokenscopepolicies",
	Package:  "squall",
	Private:  false,
}

// TokenScopePoliciesList represents a list of TokenScopePolicies
type TokenScopePoliciesList []*TokenScopePolicy

// Identity returns the identity of the objects in the list.
func (o TokenScopePoliciesList) Identity() elemental.Identity {

	return TokenScopePolicyIdentity
}

// Copy returns a pointer to a copy the TokenScopePoliciesList.
func (o TokenScopePoliciesList) Copy() elemental.Identifiables {

	copy := append(TokenScopePoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the TokenScopePoliciesList.
func (o TokenScopePoliciesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(TokenScopePoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*TokenScopePolicy))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o TokenScopePoliciesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o TokenScopePoliciesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the TokenScopePoliciesList converted to SparseTokenScopePoliciesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o TokenScopePoliciesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseTokenScopePoliciesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseTokenScopePolicy)
	}

	return out
}

// Version returns the version of the content.
func (o TokenScopePoliciesList) Version() int {

	return 1
}

// TokenScopePolicy represents the model of a tokenscopepolicy
type TokenScopePolicy struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Defines for how long the policy will be active according to the
	// `activeSchedule`.
	ActiveDuration string `json:"activeDuration" msgpack:"activeDuration" bson:"activeduration" mapstructure:"activeDuration,omitempty"`

	// Defines when the policy should be active using the cron notation.
	// The policy will be active for the given `activeDuration`.
	ActiveSchedule string `json:"activeSchedule" msgpack:"activeSchedule" bson:"activeschedule" mapstructure:"activeSchedule,omitempty"`

	// A list of audience values that are allowed when issuing a service token. An
	// empty list will allow any audience values.
	AllowedAudiences []string `json:"allowedAudiences" msgpack:"allowedAudiences" bson:"allowedaudiences" mapstructure:"allowedAudiences,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// The audience that should be assigned to a request if the caller is not
	// requesting any specific audience.
	AssignedAudience string `json:"assignedAudience" msgpack:"assignedAudience" bson:"assignedaudience" mapstructure:"assignedAudience,omitempty"`

	// The list of scopes that the policy will assign.
	AssignedScopes []string `json:"assignedScopes" msgpack:"assignedScopes" bson:"assignedscopes" mapstructure:"assignedScopes,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Defines if the property is disabled.
	Disabled bool `json:"disabled" msgpack:"disabled" bson:"disabled" mapstructure:"disabled,omitempty"`

	// If set the policy will be automatically deleted after the given time.
	ExpirationTime time.Time `json:"expirationTime" msgpack:"expirationTime" bson:"expirationtime" mapstructure:"expirationTime,omitempty"`

	// Indicates that this is fallback policy. It will only be
	// applied if no other policies have been resolved. If the policy is also
	// propagated it will become a fallback for children namespaces.
	Fallback bool `json:"fallback" msgpack:"fallback" bson:"fallback" mapstructure:"fallback,omitempty"`

	// A list of claim keys that should be inherited from the claims of the caller to
	// the assigned token. In this case, some of the caller claims will be propagated
	// to resolved token.
	InheritedClaimKeys []string `json:"inheritedClaimKeys" msgpack:"inheritedClaimKeys" bson:"inheritedclaimkeys" mapstructure:"inheritedClaimKeys,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" msgpack:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Propagates the policy to all of its children.
	Propagate bool `json:"propagate" msgpack:"propagate" bson:"propagate" mapstructure:"propagate,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Defines the selection criteria that this policy must match on identity
	// and scope request information.
	Subject [][]string `json:"subject" msgpack:"subject" bson:"subject" mapstructure:"subject,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewTokenScopePolicy returns a new *TokenScopePolicy
func NewTokenScopePolicy() *TokenScopePolicy {

	return &TokenScopePolicy{
		ModelVersion:       1,
		AssignedScopes:     []string{},
		AllowedAudiences:   []string{},
		Annotations:        map[string][]string{},
		AssociatedTags:     []string{},
		InheritedClaimKeys: []string{},
		Metadata:           []string{},
		NormalizedTags:     []string{},
		Subject:            [][]string{},
	}
}

// Identity returns the Identity of the object.
func (o *TokenScopePolicy) Identity() elemental.Identity {

	return TokenScopePolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *TokenScopePolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *TokenScopePolicy) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *TokenScopePolicy) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesTokenScopePolicy{}

	s.ActiveDuration = o.ActiveDuration
	s.ActiveSchedule = o.ActiveSchedule
	s.AllowedAudiences = o.AllowedAudiences
	s.Annotations = o.Annotations
	s.AssignedAudience = o.AssignedAudience
	s.AssignedScopes = o.AssignedScopes
	s.AssociatedTags = o.AssociatedTags
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CreateTime = o.CreateTime
	s.Description = o.Description
	s.Disabled = o.Disabled
	s.ExpirationTime = o.ExpirationTime
	s.Fallback = o.Fallback
	s.InheritedClaimKeys = o.InheritedClaimKeys
	s.Metadata = o.Metadata
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NormalizedTags = o.NormalizedTags
	s.Propagate = o.Propagate
	s.Protected = o.Protected
	s.Subject = o.Subject
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.UpdateTime = o.UpdateTime

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *TokenScopePolicy) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesTokenScopePolicy{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ActiveDuration = s.ActiveDuration
	o.ActiveSchedule = s.ActiveSchedule
	o.AllowedAudiences = s.AllowedAudiences
	o.Annotations = s.Annotations
	o.AssignedAudience = s.AssignedAudience
	o.AssignedScopes = s.AssignedScopes
	o.AssociatedTags = s.AssociatedTags
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CreateTime = s.CreateTime
	o.Description = s.Description
	o.Disabled = s.Disabled
	o.ExpirationTime = s.ExpirationTime
	o.Fallback = s.Fallback
	o.InheritedClaimKeys = s.InheritedClaimKeys
	o.Metadata = s.Metadata
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NormalizedTags = s.NormalizedTags
	o.Propagate = s.Propagate
	o.Protected = s.Protected
	o.Subject = s.Subject
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.UpdateTime = s.UpdateTime

	return nil
}

// Version returns the hardcoded version of the model.
func (o *TokenScopePolicy) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *TokenScopePolicy) BleveType() string {

	return "tokenscopepolicy"
}

// DefaultOrder returns the list of default ordering fields.
func (o *TokenScopePolicy) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *TokenScopePolicy) Doc() string {

	return `Defines a set of policies that allow customization of the
authorization tokens issued by the Aporeto service. This allows Aporeto
generated tokens to be used by external applications.`
}

func (o *TokenScopePolicy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetActiveDuration returns the ActiveDuration of the receiver.
func (o *TokenScopePolicy) GetActiveDuration() string {

	return o.ActiveDuration
}

// SetActiveDuration sets the property ActiveDuration of the receiver using the given value.
func (o *TokenScopePolicy) SetActiveDuration(activeDuration string) {

	o.ActiveDuration = activeDuration
}

// GetActiveSchedule returns the ActiveSchedule of the receiver.
func (o *TokenScopePolicy) GetActiveSchedule() string {

	return o.ActiveSchedule
}

// SetActiveSchedule sets the property ActiveSchedule of the receiver using the given value.
func (o *TokenScopePolicy) SetActiveSchedule(activeSchedule string) {

	o.ActiveSchedule = activeSchedule
}

// GetAnnotations returns the Annotations of the receiver.
func (o *TokenScopePolicy) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *TokenScopePolicy) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *TokenScopePolicy) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *TokenScopePolicy) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *TokenScopePolicy) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *TokenScopePolicy) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *TokenScopePolicy) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *TokenScopePolicy) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *TokenScopePolicy) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *TokenScopePolicy) SetDescription(description string) {

	o.Description = description
}

// GetDisabled returns the Disabled of the receiver.
func (o *TokenScopePolicy) GetDisabled() bool {

	return o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the given value.
func (o *TokenScopePolicy) SetDisabled(disabled bool) {

	o.Disabled = disabled
}

// GetExpirationTime returns the ExpirationTime of the receiver.
func (o *TokenScopePolicy) GetExpirationTime() time.Time {

	return o.ExpirationTime
}

// SetExpirationTime sets the property ExpirationTime of the receiver using the given value.
func (o *TokenScopePolicy) SetExpirationTime(expirationTime time.Time) {

	o.ExpirationTime = expirationTime
}

// GetFallback returns the Fallback of the receiver.
func (o *TokenScopePolicy) GetFallback() bool {

	return o.Fallback
}

// SetFallback sets the property Fallback of the receiver using the given value.
func (o *TokenScopePolicy) SetFallback(fallback bool) {

	o.Fallback = fallback
}

// GetMetadata returns the Metadata of the receiver.
func (o *TokenScopePolicy) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *TokenScopePolicy) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *TokenScopePolicy) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *TokenScopePolicy) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *TokenScopePolicy) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *TokenScopePolicy) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *TokenScopePolicy) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *TokenScopePolicy) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *TokenScopePolicy) GetPropagate() bool {

	return o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the given value.
func (o *TokenScopePolicy) SetPropagate(propagate bool) {

	o.Propagate = propagate
}

// GetProtected returns the Protected of the receiver.
func (o *TokenScopePolicy) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *TokenScopePolicy) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *TokenScopePolicy) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *TokenScopePolicy) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *TokenScopePolicy) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *TokenScopePolicy) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *TokenScopePolicy) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseTokenScopePolicy{
			ID:                   &o.ID,
			ActiveDuration:       &o.ActiveDuration,
			ActiveSchedule:       &o.ActiveSchedule,
			AllowedAudiences:     &o.AllowedAudiences,
			Annotations:          &o.Annotations,
			AssignedAudience:     &o.AssignedAudience,
			AssignedScopes:       &o.AssignedScopes,
			AssociatedTags:       &o.AssociatedTags,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			CreateTime:           &o.CreateTime,
			Description:          &o.Description,
			Disabled:             &o.Disabled,
			ExpirationTime:       &o.ExpirationTime,
			Fallback:             &o.Fallback,
			InheritedClaimKeys:   &o.InheritedClaimKeys,
			Metadata:             &o.Metadata,
			Name:                 &o.Name,
			Namespace:            &o.Namespace,
			NormalizedTags:       &o.NormalizedTags,
			Propagate:            &o.Propagate,
			Protected:            &o.Protected,
			Subject:              &o.Subject,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			UpdateTime:           &o.UpdateTime,
		}
	}

	sp := &SparseTokenScopePolicy{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "activeDuration":
			sp.ActiveDuration = &(o.ActiveDuration)
		case "activeSchedule":
			sp.ActiveSchedule = &(o.ActiveSchedule)
		case "allowedAudiences":
			sp.AllowedAudiences = &(o.AllowedAudiences)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "assignedAudience":
			sp.AssignedAudience = &(o.AssignedAudience)
		case "assignedScopes":
			sp.AssignedScopes = &(o.AssignedScopes)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "description":
			sp.Description = &(o.Description)
		case "disabled":
			sp.Disabled = &(o.Disabled)
		case "expirationTime":
			sp.ExpirationTime = &(o.ExpirationTime)
		case "fallback":
			sp.Fallback = &(o.Fallback)
		case "inheritedClaimKeys":
			sp.InheritedClaimKeys = &(o.InheritedClaimKeys)
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "propagate":
			sp.Propagate = &(o.Propagate)
		case "protected":
			sp.Protected = &(o.Protected)
		case "subject":
			sp.Subject = &(o.Subject)
		case "updateIdempotencyKey":
			sp.UpdateIdempotencyKey = &(o.UpdateIdempotencyKey)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseTokenScopePolicy to the object.
func (o *TokenScopePolicy) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseTokenScopePolicy)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.ActiveDuration != nil {
		o.ActiveDuration = *so.ActiveDuration
	}
	if so.ActiveSchedule != nil {
		o.ActiveSchedule = *so.ActiveSchedule
	}
	if so.AllowedAudiences != nil {
		o.AllowedAudiences = *so.AllowedAudiences
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssignedAudience != nil {
		o.AssignedAudience = *so.AssignedAudience
	}
	if so.AssignedScopes != nil {
		o.AssignedScopes = *so.AssignedScopes
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Disabled != nil {
		o.Disabled = *so.Disabled
	}
	if so.ExpirationTime != nil {
		o.ExpirationTime = *so.ExpirationTime
	}
	if so.Fallback != nil {
		o.Fallback = *so.Fallback
	}
	if so.InheritedClaimKeys != nil {
		o.InheritedClaimKeys = *so.InheritedClaimKeys
	}
	if so.Metadata != nil {
		o.Metadata = *so.Metadata
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.Propagate != nil {
		o.Propagate = *so.Propagate
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.Subject != nil {
		o.Subject = *so.Subject
	}
	if so.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = *so.UpdateIdempotencyKey
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
}

// DeepCopy returns a deep copy if the TokenScopePolicy.
func (o *TokenScopePolicy) DeepCopy() *TokenScopePolicy {

	if o == nil {
		return nil
	}

	out := &TokenScopePolicy{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *TokenScopePolicy.
func (o *TokenScopePolicy) DeepCopyInto(out *TokenScopePolicy) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy TokenScopePolicy: %s", err))
	}

	*out = *target.(*TokenScopePolicy)
}

// Validate valides the current information stored into the structure.
func (o *TokenScopePolicy) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidatePattern("activeDuration", o.ActiveDuration, `^[0-9]+[smh]$`, `must be a valid duration like <n>s or <n>s or <n>h`, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateStringListNotEmpty("assignedScopes", o.AssignedScopes); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateMetadata("metadata", o.Metadata); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTagsExpression("subject", o.Subject); err != nil {
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

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*TokenScopePolicy) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := TokenScopePolicyAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return TokenScopePolicyLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*TokenScopePolicy) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return TokenScopePolicyAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *TokenScopePolicy) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "activeDuration":
		return o.ActiveDuration
	case "activeSchedule":
		return o.ActiveSchedule
	case "allowedAudiences":
		return o.AllowedAudiences
	case "annotations":
		return o.Annotations
	case "assignedAudience":
		return o.AssignedAudience
	case "assignedScopes":
		return o.AssignedScopes
	case "associatedTags":
		return o.AssociatedTags
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "description":
		return o.Description
	case "disabled":
		return o.Disabled
	case "expirationTime":
		return o.ExpirationTime
	case "fallback":
		return o.Fallback
	case "inheritedClaimKeys":
		return o.InheritedClaimKeys
	case "metadata":
		return o.Metadata
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "propagate":
		return o.Propagate
	case "protected":
		return o.Protected
	case "subject":
		return o.Subject
	case "updateIdempotencyKey":
		return o.UpdateIdempotencyKey
	case "updateTime":
		return o.UpdateTime
	}

	return nil
}

// TokenScopePolicyAttributesMap represents the map of attribute for TokenScopePolicy.
var TokenScopePolicyAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"ActiveDuration": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		ConvertedName:  "ActiveDuration",
		Description: `Defines for how long the policy will be active according to the
` + "`" + `activeSchedule` + "`" + `.`,
		Exposed: true,
		Getter:  true,
		Name:    "activeDuration",
		Setter:  true,
		Stored:  true,
		Type:    "string",
	},
	"ActiveSchedule": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ActiveSchedule",
		Description: `Defines when the policy should be active using the cron notation.
The policy will be active for the given ` + "`" + `activeDuration` + "`" + `.`,
		Exposed: true,
		Getter:  true,
		Name:    "activeSchedule",
		Setter:  true,
		Stored:  true,
		Type:    "string",
	},
	"AllowedAudiences": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllowedAudiences",
		Description: `A list of audience values that are allowed when issuing a service token. An
empty list will allow any audience values.`,
		Exposed:   true,
		Name:      "allowedAudiences",
		Orderable: true,
		Stored:    true,
		SubType:   "string",
		Type:      "list",
	},
	"Annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"AssignedAudience": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssignedAudience",
		Description: `The audience that should be assigned to a request if the caller is not
requesting any specific audience.`,
		Exposed:   true,
		Name:      "assignedAudience",
		Orderable: true,
		Stored:    true,
		Type:      "string",
	},
	"AssignedScopes": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssignedScopes",
		Description:    `The list of scopes that the policy will assign.`,
		Exposed:        true,
		Name:           "assignedScopes",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"AssociatedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"CreateIdempotencyKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Disabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Disabled",
		Description:    `Defines if the property is disabled.`,
		Exposed:        true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"ExpirationTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExpirationTime",
		Description:    `If set the policy will be automatically deleted after the given time.`,
		Exposed:        true,
		Getter:         true,
		Name:           "expirationTime",
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Fallback": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Fallback",
		Description: `Indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.`,
		Exposed:   true,
		Getter:    true,
		Name:      "fallback",
		Orderable: true,
		Setter:    true,
		Stored:    true,
		Type:      "boolean",
	},
	"InheritedClaimKeys": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "InheritedClaimKeys",
		Description: `A list of claim keys that should be inherited from the claims of the caller to
the assigned token. In this case, some of the caller claims will be propagated
to resolved token.`,
		Exposed:   true,
		Name:      "inheritedClaimKeys",
		Orderable: true,
		Stored:    true,
		SubType:   "string",
		Type:      "list",
	},
	"Metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Contains tags that can only be set during creation, must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "string",
		Type:       "list",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"NormalizedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"Propagate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Propagate",
		Description:    `Propagates the policy to all of its children.`,
		Exposed:        true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		Description: `Defines the selection criteria that this policy must match on identity
and scope request information.`,
		Exposed:   true,
		Name:      "subject",
		Orderable: true,
		Stored:    true,
		SubType:   "[][]string",
		Type:      "external",
	},
	"UpdateIdempotencyKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
}

// TokenScopePolicyLowerCaseAttributesMap represents the map of attribute for TokenScopePolicy.
var TokenScopePolicyLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"activeduration": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		ConvertedName:  "ActiveDuration",
		Description: `Defines for how long the policy will be active according to the
` + "`" + `activeSchedule` + "`" + `.`,
		Exposed: true,
		Getter:  true,
		Name:    "activeDuration",
		Setter:  true,
		Stored:  true,
		Type:    "string",
	},
	"activeschedule": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ActiveSchedule",
		Description: `Defines when the policy should be active using the cron notation.
The policy will be active for the given ` + "`" + `activeDuration` + "`" + `.`,
		Exposed: true,
		Getter:  true,
		Name:    "activeSchedule",
		Setter:  true,
		Stored:  true,
		Type:    "string",
	},
	"allowedaudiences": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllowedAudiences",
		Description: `A list of audience values that are allowed when issuing a service token. An
empty list will allow any audience values.`,
		Exposed:   true,
		Name:      "allowedAudiences",
		Orderable: true,
		Stored:    true,
		SubType:   "string",
		Type:      "list",
	},
	"annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"assignedaudience": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssignedAudience",
		Description: `The audience that should be assigned to a request if the caller is not
requesting any specific audience.`,
		Exposed:   true,
		Name:      "assignedAudience",
		Orderable: true,
		Stored:    true,
		Type:      "string",
	},
	"assignedscopes": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssignedScopes",
		Description:    `The list of scopes that the policy will assign.`,
		Exposed:        true,
		Name:           "assignedScopes",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"associatedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"createidempotencykey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"disabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Disabled",
		Description:    `Defines if the property is disabled.`,
		Exposed:        true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"expirationtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExpirationTime",
		Description:    `If set the policy will be automatically deleted after the given time.`,
		Exposed:        true,
		Getter:         true,
		Name:           "expirationTime",
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"fallback": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Fallback",
		Description: `Indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.`,
		Exposed:   true,
		Getter:    true,
		Name:      "fallback",
		Orderable: true,
		Setter:    true,
		Stored:    true,
		Type:      "boolean",
	},
	"inheritedclaimkeys": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "InheritedClaimKeys",
		Description: `A list of claim keys that should be inherited from the claims of the caller to
the assigned token. In this case, some of the caller claims will be propagated
to resolved token.`,
		Exposed:   true,
		Name:      "inheritedClaimKeys",
		Orderable: true,
		Stored:    true,
		SubType:   "string",
		Type:      "list",
	},
	"metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Contains tags that can only be set during creation, must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "string",
		Type:       "list",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"normalizedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"propagate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Propagate",
		Description:    `Propagates the policy to all of its children.`,
		Exposed:        true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		Description: `Defines the selection criteria that this policy must match on identity
and scope request information.`,
		Exposed:   true,
		Name:      "subject",
		Orderable: true,
		Stored:    true,
		SubType:   "[][]string",
		Type:      "external",
	},
	"updateidempotencykey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"updatetime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
}

// SparseTokenScopePoliciesList represents a list of SparseTokenScopePolicies
type SparseTokenScopePoliciesList []*SparseTokenScopePolicy

// Identity returns the identity of the objects in the list.
func (o SparseTokenScopePoliciesList) Identity() elemental.Identity {

	return TokenScopePolicyIdentity
}

// Copy returns a pointer to a copy the SparseTokenScopePoliciesList.
func (o SparseTokenScopePoliciesList) Copy() elemental.Identifiables {

	copy := append(SparseTokenScopePoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseTokenScopePoliciesList.
func (o SparseTokenScopePoliciesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseTokenScopePoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseTokenScopePolicy))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseTokenScopePoliciesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseTokenScopePoliciesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseTokenScopePoliciesList converted to TokenScopePoliciesList.
func (o SparseTokenScopePoliciesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseTokenScopePoliciesList) Version() int {

	return 1
}

// SparseTokenScopePolicy represents the sparse version of a tokenscopepolicy.
type SparseTokenScopePolicy struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Defines for how long the policy will be active according to the
	// `activeSchedule`.
	ActiveDuration *string `json:"activeDuration,omitempty" msgpack:"activeDuration,omitempty" bson:"activeduration,omitempty" mapstructure:"activeDuration,omitempty"`

	// Defines when the policy should be active using the cron notation.
	// The policy will be active for the given `activeDuration`.
	ActiveSchedule *string `json:"activeSchedule,omitempty" msgpack:"activeSchedule,omitempty" bson:"activeschedule,omitempty" mapstructure:"activeSchedule,omitempty"`

	// A list of audience values that are allowed when issuing a service token. An
	// empty list will allow any audience values.
	AllowedAudiences *[]string `json:"allowedAudiences,omitempty" msgpack:"allowedAudiences,omitempty" bson:"allowedaudiences,omitempty" mapstructure:"allowedAudiences,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// The audience that should be assigned to a request if the caller is not
	// requesting any specific audience.
	AssignedAudience *string `json:"assignedAudience,omitempty" msgpack:"assignedAudience,omitempty" bson:"assignedaudience,omitempty" mapstructure:"assignedAudience,omitempty"`

	// The list of scopes that the policy will assign.
	AssignedScopes *[]string `json:"assignedScopes,omitempty" msgpack:"assignedScopes,omitempty" bson:"assignedscopes,omitempty" mapstructure:"assignedScopes,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Defines if the property is disabled.
	Disabled *bool `json:"disabled,omitempty" msgpack:"disabled,omitempty" bson:"disabled,omitempty" mapstructure:"disabled,omitempty"`

	// If set the policy will be automatically deleted after the given time.
	ExpirationTime *time.Time `json:"expirationTime,omitempty" msgpack:"expirationTime,omitempty" bson:"expirationtime,omitempty" mapstructure:"expirationTime,omitempty"`

	// Indicates that this is fallback policy. It will only be
	// applied if no other policies have been resolved. If the policy is also
	// propagated it will become a fallback for children namespaces.
	Fallback *bool `json:"fallback,omitempty" msgpack:"fallback,omitempty" bson:"fallback,omitempty" mapstructure:"fallback,omitempty"`

	// A list of claim keys that should be inherited from the claims of the caller to
	// the assigned token. In this case, some of the caller claims will be propagated
	// to resolved token.
	InheritedClaimKeys *[]string `json:"inheritedClaimKeys,omitempty" msgpack:"inheritedClaimKeys,omitempty" bson:"inheritedclaimkeys,omitempty" mapstructure:"inheritedClaimKeys,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" msgpack:"metadata,omitempty" bson:"metadata,omitempty" mapstructure:"metadata,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Propagates the policy to all of its children.
	Propagate *bool `json:"propagate,omitempty" msgpack:"propagate,omitempty" bson:"propagate,omitempty" mapstructure:"propagate,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// Defines the selection criteria that this policy must match on identity
	// and scope request information.
	Subject *[][]string `json:"subject,omitempty" msgpack:"subject,omitempty" bson:"subject,omitempty" mapstructure:"subject,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseTokenScopePolicy returns a new  SparseTokenScopePolicy.
func NewSparseTokenScopePolicy() *SparseTokenScopePolicy {
	return &SparseTokenScopePolicy{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseTokenScopePolicy) Identity() elemental.Identity {

	return TokenScopePolicyIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseTokenScopePolicy) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseTokenScopePolicy) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseTokenScopePolicy) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseTokenScopePolicy{}

	if o.ActiveDuration != nil {
		s.ActiveDuration = o.ActiveDuration
	}
	if o.ActiveSchedule != nil {
		s.ActiveSchedule = o.ActiveSchedule
	}
	if o.AllowedAudiences != nil {
		s.AllowedAudiences = o.AllowedAudiences
	}
	if o.Annotations != nil {
		s.Annotations = o.Annotations
	}
	if o.AssignedAudience != nil {
		s.AssignedAudience = o.AssignedAudience
	}
	if o.AssignedScopes != nil {
		s.AssignedScopes = o.AssignedScopes
	}
	if o.AssociatedTags != nil {
		s.AssociatedTags = o.AssociatedTags
	}
	if o.CreateIdempotencyKey != nil {
		s.CreateIdempotencyKey = o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
	}
	if o.Description != nil {
		s.Description = o.Description
	}
	if o.Disabled != nil {
		s.Disabled = o.Disabled
	}
	if o.ExpirationTime != nil {
		s.ExpirationTime = o.ExpirationTime
	}
	if o.Fallback != nil {
		s.Fallback = o.Fallback
	}
	if o.InheritedClaimKeys != nil {
		s.InheritedClaimKeys = o.InheritedClaimKeys
	}
	if o.Metadata != nil {
		s.Metadata = o.Metadata
	}
	if o.Name != nil {
		s.Name = o.Name
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.NormalizedTags != nil {
		s.NormalizedTags = o.NormalizedTags
	}
	if o.Propagate != nil {
		s.Propagate = o.Propagate
	}
	if o.Protected != nil {
		s.Protected = o.Protected
	}
	if o.Subject != nil {
		s.Subject = o.Subject
	}
	if o.UpdateIdempotencyKey != nil {
		s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		s.UpdateTime = o.UpdateTime
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseTokenScopePolicy) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseTokenScopePolicy{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.ActiveDuration != nil {
		o.ActiveDuration = s.ActiveDuration
	}
	if s.ActiveSchedule != nil {
		o.ActiveSchedule = s.ActiveSchedule
	}
	if s.AllowedAudiences != nil {
		o.AllowedAudiences = s.AllowedAudiences
	}
	if s.Annotations != nil {
		o.Annotations = s.Annotations
	}
	if s.AssignedAudience != nil {
		o.AssignedAudience = s.AssignedAudience
	}
	if s.AssignedScopes != nil {
		o.AssignedScopes = s.AssignedScopes
	}
	if s.AssociatedTags != nil {
		o.AssociatedTags = s.AssociatedTags
	}
	if s.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = s.CreateIdempotencyKey
	}
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.Description != nil {
		o.Description = s.Description
	}
	if s.Disabled != nil {
		o.Disabled = s.Disabled
	}
	if s.ExpirationTime != nil {
		o.ExpirationTime = s.ExpirationTime
	}
	if s.Fallback != nil {
		o.Fallback = s.Fallback
	}
	if s.InheritedClaimKeys != nil {
		o.InheritedClaimKeys = s.InheritedClaimKeys
	}
	if s.Metadata != nil {
		o.Metadata = s.Metadata
	}
	if s.Name != nil {
		o.Name = s.Name
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.NormalizedTags != nil {
		o.NormalizedTags = s.NormalizedTags
	}
	if s.Propagate != nil {
		o.Propagate = s.Propagate
	}
	if s.Protected != nil {
		o.Protected = s.Protected
	}
	if s.Subject != nil {
		o.Subject = s.Subject
	}
	if s.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	}
	if s.UpdateTime != nil {
		o.UpdateTime = s.UpdateTime
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseTokenScopePolicy) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseTokenScopePolicy) ToPlain() elemental.PlainIdentifiable {

	out := NewTokenScopePolicy()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.ActiveDuration != nil {
		out.ActiveDuration = *o.ActiveDuration
	}
	if o.ActiveSchedule != nil {
		out.ActiveSchedule = *o.ActiveSchedule
	}
	if o.AllowedAudiences != nil {
		out.AllowedAudiences = *o.AllowedAudiences
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssignedAudience != nil {
		out.AssignedAudience = *o.AssignedAudience
	}
	if o.AssignedScopes != nil {
		out.AssignedScopes = *o.AssignedScopes
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Disabled != nil {
		out.Disabled = *o.Disabled
	}
	if o.ExpirationTime != nil {
		out.ExpirationTime = *o.ExpirationTime
	}
	if o.Fallback != nil {
		out.Fallback = *o.Fallback
	}
	if o.InheritedClaimKeys != nil {
		out.InheritedClaimKeys = *o.InheritedClaimKeys
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.Propagate != nil {
		out.Propagate = *o.Propagate
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.Subject != nil {
		out.Subject = *o.Subject
	}
	if o.UpdateIdempotencyKey != nil {
		out.UpdateIdempotencyKey = *o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}

	return out
}

// GetActiveDuration returns the ActiveDuration of the receiver.
func (o *SparseTokenScopePolicy) GetActiveDuration() (out string) {

	if o.ActiveDuration == nil {
		return
	}

	return *o.ActiveDuration
}

// SetActiveDuration sets the property ActiveDuration of the receiver using the address of the given value.
func (o *SparseTokenScopePolicy) SetActiveDuration(activeDuration string) {

	o.ActiveDuration = &activeDuration
}

// GetActiveSchedule returns the ActiveSchedule of the receiver.
func (o *SparseTokenScopePolicy) GetActiveSchedule() (out string) {

	if o.ActiveSchedule == nil {
		return
	}

	return *o.ActiveSchedule
}

// SetActiveSchedule sets the property ActiveSchedule of the receiver using the address of the given value.
func (o *SparseTokenScopePolicy) SetActiveSchedule(activeSchedule string) {

	o.ActiveSchedule = &activeSchedule
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseTokenScopePolicy) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseTokenScopePolicy) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseTokenScopePolicy) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseTokenScopePolicy) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseTokenScopePolicy) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseTokenScopePolicy) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseTokenScopePolicy) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseTokenScopePolicy) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseTokenScopePolicy) GetDescription() (out string) {

	if o.Description == nil {
		return
	}

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseTokenScopePolicy) SetDescription(description string) {

	o.Description = &description
}

// GetDisabled returns the Disabled of the receiver.
func (o *SparseTokenScopePolicy) GetDisabled() (out bool) {

	if o.Disabled == nil {
		return
	}

	return *o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the address of the given value.
func (o *SparseTokenScopePolicy) SetDisabled(disabled bool) {

	o.Disabled = &disabled
}

// GetExpirationTime returns the ExpirationTime of the receiver.
func (o *SparseTokenScopePolicy) GetExpirationTime() (out time.Time) {

	if o.ExpirationTime == nil {
		return
	}

	return *o.ExpirationTime
}

// SetExpirationTime sets the property ExpirationTime of the receiver using the address of the given value.
func (o *SparseTokenScopePolicy) SetExpirationTime(expirationTime time.Time) {

	o.ExpirationTime = &expirationTime
}

// GetFallback returns the Fallback of the receiver.
func (o *SparseTokenScopePolicy) GetFallback() (out bool) {

	if o.Fallback == nil {
		return
	}

	return *o.Fallback
}

// SetFallback sets the property Fallback of the receiver using the address of the given value.
func (o *SparseTokenScopePolicy) SetFallback(fallback bool) {

	o.Fallback = &fallback
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseTokenScopePolicy) GetMetadata() (out []string) {

	if o.Metadata == nil {
		return
	}

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseTokenScopePolicy) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetName returns the Name of the receiver.
func (o *SparseTokenScopePolicy) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseTokenScopePolicy) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseTokenScopePolicy) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseTokenScopePolicy) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseTokenScopePolicy) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseTokenScopePolicy) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *SparseTokenScopePolicy) GetPropagate() (out bool) {

	if o.Propagate == nil {
		return
	}

	return *o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the address of the given value.
func (o *SparseTokenScopePolicy) SetPropagate(propagate bool) {

	o.Propagate = &propagate
}

// GetProtected returns the Protected of the receiver.
func (o *SparseTokenScopePolicy) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseTokenScopePolicy) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseTokenScopePolicy) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseTokenScopePolicy) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseTokenScopePolicy) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseTokenScopePolicy) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// DeepCopy returns a deep copy if the SparseTokenScopePolicy.
func (o *SparseTokenScopePolicy) DeepCopy() *SparseTokenScopePolicy {

	if o == nil {
		return nil
	}

	out := &SparseTokenScopePolicy{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseTokenScopePolicy.
func (o *SparseTokenScopePolicy) DeepCopyInto(out *SparseTokenScopePolicy) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseTokenScopePolicy: %s", err))
	}

	*out = *target.(*SparseTokenScopePolicy)
}

type mongoAttributesTokenScopePolicy struct {
	ActiveDuration       string              `bson:"activeduration"`
	ActiveSchedule       string              `bson:"activeschedule"`
	AllowedAudiences     []string            `bson:"allowedaudiences"`
	Annotations          map[string][]string `bson:"annotations"`
	AssignedAudience     string              `bson:"assignedaudience"`
	AssignedScopes       []string            `bson:"assignedscopes"`
	AssociatedTags       []string            `bson:"associatedtags"`
	CreateIdempotencyKey string              `bson:"createidempotencykey"`
	CreateTime           time.Time           `bson:"createtime"`
	Description          string              `bson:"description"`
	Disabled             bool                `bson:"disabled"`
	ExpirationTime       time.Time           `bson:"expirationtime"`
	Fallback             bool                `bson:"fallback"`
	InheritedClaimKeys   []string            `bson:"inheritedclaimkeys"`
	Metadata             []string            `bson:"metadata"`
	Name                 string              `bson:"name"`
	Namespace            string              `bson:"namespace"`
	NormalizedTags       []string            `bson:"normalizedtags"`
	Propagate            bool                `bson:"propagate"`
	Protected            bool                `bson:"protected"`
	Subject              [][]string          `bson:"subject"`
	UpdateIdempotencyKey string              `bson:"updateidempotencykey"`
	UpdateTime           time.Time           `bson:"updatetime"`
}
type mongoAttributesSparseTokenScopePolicy struct {
	ActiveDuration       *string              `bson:"activeduration,omitempty"`
	ActiveSchedule       *string              `bson:"activeschedule,omitempty"`
	AllowedAudiences     *[]string            `bson:"allowedaudiences,omitempty"`
	Annotations          *map[string][]string `bson:"annotations,omitempty"`
	AssignedAudience     *string              `bson:"assignedaudience,omitempty"`
	AssignedScopes       *[]string            `bson:"assignedscopes,omitempty"`
	AssociatedTags       *[]string            `bson:"associatedtags,omitempty"`
	CreateIdempotencyKey *string              `bson:"createidempotencykey,omitempty"`
	CreateTime           *time.Time           `bson:"createtime,omitempty"`
	Description          *string              `bson:"description,omitempty"`
	Disabled             *bool                `bson:"disabled,omitempty"`
	ExpirationTime       *time.Time           `bson:"expirationtime,omitempty"`
	Fallback             *bool                `bson:"fallback,omitempty"`
	InheritedClaimKeys   *[]string            `bson:"inheritedclaimkeys,omitempty"`
	Metadata             *[]string            `bson:"metadata,omitempty"`
	Name                 *string              `bson:"name,omitempty"`
	Namespace            *string              `bson:"namespace,omitempty"`
	NormalizedTags       *[]string            `bson:"normalizedtags,omitempty"`
	Propagate            *bool                `bson:"propagate,omitempty"`
	Protected            *bool                `bson:"protected,omitempty"`
	Subject              *[][]string          `bson:"subject,omitempty"`
	UpdateIdempotencyKey *string              `bson:"updateidempotencykey,omitempty"`
	UpdateTime           *time.Time           `bson:"updatetime,omitempty"`
}
