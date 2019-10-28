# Model
model:
  rest_name: enforcer
  resource_name: enforcers
  entity_name: Enforcer
  package: squall
  group: core/enforcer
  description: |-
    Contains all parameters associated with a registered enforcer. The
    object is mainly maintained by the enforcers themselves. Users can read the
    object in order to understand the current status of the enforcers.
  get:
    description: Retrieves the enforcer with the given ID.
  update:
    description: Updates the enforcer with the given ID.
  delete:
    description: Deletes the enforcer with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@zoned'
  - '@base'
  - '@migratable'
  - '@namespaced'
  - '@described'
  - '@identifiable-stored'
  - '@metadatable'
  - '@named'
  - '@timeable'

# Indexes
indexes:
- - namespace
  - operationalStatus
- - namespace
  - machineID

# Attributes
attributes:
  v1:
  - name: FQDN
    description: |-
      Contains the fully qualified domain name (FQDN) of the server where the
      enforcer is running.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: server1.domain.com
    orderable: true

  - name: certificate
    description: The certificate of the enforcer.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    orderable: true

  - name: certificateExpirationDate
    description: |-
      The expiration date of the certificate. This is an
      internal attribute, not exposed in the API.
    type: time
    read_only: true
    autogenerated: true

  - name: certificateKey
    description: |-
      The certificate key of the enforcer. This is an internal
      attribute, not exposed in the API.
    type: string
    read_only: true
    autogenerated: true

  - name: certificateRequest
    description: |-
      If not empty during a create or update operation, the provided certificate
      signing request (CSR) will be validated and signed by the control plane,
      providing a renewed certificate.
    type: string
    exposed: true
    example_value: |-
      -----BEGIN CERTIFICATE REQUEST-----
      MIICvDCCAaQCAQAwdzELMAkGA1UEBhMCVVMxDTALBgNVBAgMBFV0YWgxDzANBgNV
      BAcMBkxpbmRvbjEWMBQGA1UECgwNRGlnaUNlcnQgSW5jLjERMA8GA1UECwwIRGln
      aUNlcnQxHTAbBgNVBAMMFGV4YW1wbGUuZGlnaWNlcnQuY29tMIIBIjANBgkqhkiG
      9w0BAQEFAAOCAQ8AMIIBCgKCAQEA8+To7d+2kPWeBv/orU3LVbJwDrSQbeKamCmo
      wp5bqDxIwV20zqRb7APUOKYoVEFFOEQs6T6gImnIolhbiH6m4zgZ/CPvWBOkZc+c
      1Po2EmvBz+AD5sBdT5kzGQA6NbWyZGldxRthNLOs1efOhdnWFuhI162qmcflgpiI
      WDuwq4C9f+YkeJhNn9dF5+owm8cOQmDrV8NNdiTqin8q3qYAHHJRW28glJUCZkTZ
      wIaSR6crBQ8TbYNE0dc+Caa3DOIkz1EOsHWzTx+n0zKfqcbgXi4DJx+C1bjptYPR
      BPZL8DAeWuA8ebudVT44yEp82G96/Ggcf7F33xMxe0yc+Xa6owIDAQABoAAwDQYJ
      KoZIhvcNAQEFBQADggEBAB0kcrFccSmFDmxox0Ne01UIqSsDqHgL+XmHTXJwre6D
      hJSZwbvEtOK0G3+dr4Fs11WuUNt5qcLsx5a8uk4G6AKHMzuhLsJ7XZjgmQXGECpY
      Q4mC3yT3ZoCGpIXbw+iP3lmEEXgaQL0Tx5LFl/okKbKYwIqNiyKWOMj7ZR/wxWg/
      ZDGRs55xuoeLDJ/ZRFf9bI+IaCUd1YrfYcHIl3G87Av+r49YVwqRDT0VDV7uLgqn
      29XI1PpVUNCPQGn9p/eX6Qo7vpDaPybRtA2R7XLKjQaF9oXWeCUqy1hvJac9QFO2
      97Ob1alpHPoZ7mWiEuJwjBPii6a9M9G30nUo39lBi1w=
      -----END CERTIFICATE REQUEST-----
    transient: true

  - name: collectInfo
    description: Indicates to the enforcer whether or not it needs to collect information.
    type: boolean
    exposed: true
    stored: true

  - name: collectedInfo
    description: Represents the latest information collected by the enforcer.
    type: external
    exposed: true
    subtype: map[string]string
    stored: true

  - name: currentVersion
    description: The version number of the installed enforcer binary.
    type: string
    exposed: true
    stored: true
    filterable: true
    orderable: true

  - name: enforcementStatus
    description: Status of the enforcement for host services.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Inactive
    - Active
    - Failed
    default_value: Inactive
    filterable: true

  - name: lastCollectionTime
    description: Identifies when the information was collected.
    type: time
    exposed: true
    stored: true

  - name: lastPokeTime
    description: The time and date of the last poke.
    type: time
    stored: true

  - name: lastSyncTime
    description: The time and date of the last heartbeat.
    type: time
    exposed: true
    stored: true
    orderable: true

  - name: lastValidHostServices
    description: |-
      LastValidHostServices is a read only attribute that stores the list valid host
      services that have been applied to this enforcer. This list might be different
      from the list retrieved through policy, if the dynamically calculated list leads
      into conflicts.
    type: refList
    subtype: hostservice
    stored: true
    autogenerated: true

  - name: localCA
    description: |-
      Contains the initial chain of trust for the enforcer. This value is only
      given when you retrieve a single enforcer.
    type: string
    exposed: true
    autogenerated: true
    transient: true

  - name: logLevel
    description: Log level of the enforcer.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Info
    - Debug
    - Warn
    - Error
    - Trace
    default_value: Info

  - name: logLevelDuration
    description: |-
      Determines the duration of which the log level will be active, using [Golang
      duration syntax](https://golang.org/pkg/time/#example_Duration).
    type: string
    exposed: true
    stored: true
    default_value: 10s
    validations:
    - $timeDuration

  - name: machineID
    description: |-
      A unique identifier for every machine as detected by the enforcer. It is
      based on hardware information such as the SMBIOS UUID, MAC addresses of
      interfaces, or cloud provider IDs.
    type: string
    exposed: true
    stored: true
    example_value: 3F23E8DF-C56D-45CF-89B8-A867F3956409
    filterable: true

  - name: operationalStatus
    description: The status of the enforcer.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Registered
    - Connected
    - Disconnected
    - Initialized
    default_value: Registered
    filterable: true

  - name: publicToken
    description: |-
      The public token of the server that will be included in the datapath and
      is signed by the private certificate authority.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    transient: true

  - name: startTime
    description: |-
      The time and date on which this enforcer was started. The enforcer reports
      this and the value is preserved across disconnects.
    type: time
    exposed: true
    stored: true
    orderable: true

  - name: subnets
    description: Local subnets of this enforcer.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: unreachable
    description: |-
      The Aporeto control plane sets this value to `true` if it hasn't heard from
      the enforcer in the last five minutes.
    type: boolean
    exposed: true
    stored: true
    read_only: true
    autogenerated: true

  - name: updateAvailable
    description: If `true`, the enforcer version is outdated and should be updated.
    type: boolean
    exposed: true
    stored: true
    orderable: true

# Relations
relations:
- rest_name: auditprofile
  get:
    description: Returns a list of the audit profiles that must be applied to this
      enforcer.

- rest_name: enforcerprofile
  get:
    description: Returns the enforcer profile that must be used by an enforcer.

- rest_name: hostservice
  get:
    description: Returns a list of the host services policies that apply to this enforcer.
    parameters:
      entries:
      - name: appliedServices
        description: Valid when retrieved for a given enforcer and returns the applied
          services.
        type: boolean

      - name: setServices
        description: Instructs the backend to cache the services that were resolved.
          services.
        type: boolean

- rest_name: poke
  get:
    description: Sends a poke empty object. This is used to ensure an enforcer is
      up and running.
    parameters:
      entries:
      - name: cpuload
        description: Deprecated.
        type: float
        example_value: 1000

      - name: enforcementStatus
        description: If set, changes the enforcement status of the enforcer alongside
          with the poke.
        type: enum
        allowed_choices:
        - Failed
        - Inactive
        - Active

      - name: forceFullPoke
        description: If set, it will trigger a full poke (slower).
        type: boolean

      - name: memory
        description: Deprecated.
        type: integer
        example_value: 1000

      - name: processes
        description: Deprecated.
        type: integer
        example_value: 10

      - name: sessionClose
        description: If set, terminates a session for an enforcer.
        type: boolean

      - name: sessionID
        description: If set, sends the current session ID of an enforcer.
        type: string
        example_value: "1233"

      - name: status
        description: If set, changes the status of the enforcer alongside with the
          poke.
        type: enum
        allowed_choices:
        - Registered
        - Connected
        - Disconnected

      - name: ts
        description: time of report. If not set, local server time will be used.
        type: time

      - name: version
        description: If set, version of the current running enforcer.
        type: string
        example_value: v1.10

      - name: zhash
        description: Can be set to help backend target the correct shard where the
          enforcer is stored.
        type: integer

- rest_name: trustedca
  get:
    description: |-
      Returns the list of certificate authorities that should be trusted by this
      enforcer.
    parameters:
      entries:
      - name: type
        description: Type of certificate to get.
        type: enum
        allowed_choices:
        - Any
        - X509
        - SSH
        default_value: Any
