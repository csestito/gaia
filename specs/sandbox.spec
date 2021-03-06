# Model
model:
  rest_name: sandbox
  resource_name: sandboxes
  entity_name: Sandbox
  package: service
  group: ext/documentation
  description: |-
    This APIs allows to create a temporary namespace to experiment with Microsegmentation.
    This API is not authenticated, and contains small quotas. After one hour,
    everything will be deleted.

# Attributes
attributes:
  v1:
  - name: URL
    description: Contains a link to directly connect the UI to your API sandbox.
    type: string
    exposed: true
    read_only: true
    autogenerated: true

  - name: credentials
    description: The app credential data.
    type: ref
    exposed: true
    subtype: credential
    read_only: true
    autogenerated: true
    orderable: true
    transient: true
    extensions:
      noInit: true
      refMode: pointer

  - name: lifetime
    description: Contains the lifetime of the sandbox namespace.
    type: string
    exposed: true
    read_only: true
    autogenerated: true

  - name: namespace
    description: Contains the name of the sandbox namespace that has been created.
    type: string
    exposed: true
    read_only: true
    autogenerated: true
