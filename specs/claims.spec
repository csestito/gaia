# Model
model:
  rest_name: claims
  resource_name: claims
  entity_name: Claims
  package: guy
  group: policy/networking
  description: Represents the claims in the token used to access a service.
  get:
    description: Retrieves the object with the given ID.
  extends:
  - '@zoned'
  - '@migratable'
  - '@namespaced'
  - '@identifiable-stored'

# Indexes
indexes:
- - namespace
  - hash

# Attributes
attributes:
  v1:
  - name: content
    description: Contains the raw JSON web token (JWT) claims.
    type: external
    exposed: true
    subtype: map[string]string
    stored: true
    creation_only: true
    example_value:
      exp: 1553899021
      iat: 1553888221
      iss: https://accounts.acme.com
      sub: alice@acme.com
    omit_empty: true

  - name: firstSeen
    description: Contains the date of the first appearance of the claims.
    type: time
    stored: true
    read_only: true
    autogenerated: true
    omit_empty: true

  - name: hash
    description: |-
      XXH64 hash of the claims content. It will be used as ID. To compute a correct
      hash,
      you must first clob `content` as an string array in the form `key=value`, sort
      it
      then apply the XXH64 function.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: "1134423925458173049"

  - name: lastSeen
    description: Contains the date of the last appearance of the claims.
    type: time
    stored: true
    read_only: true
    autogenerated: true
