# Model
model:
  rest_name: revocation
  resource_name: revocations
  entity_name: Revocation
  package: barret
  description: Used to revoke a certificate
  update: true

# Attributes
attributes:
- name: ID
  description: ID contains the ID of the revocation.
  type: string
  stored: true
  read_only: true
  autogenerated: true
  format: free
  identifier: true
  primary_key: true

- name: expirationDate
  description: Contains the certificate expiration date. This will be used to clean
    up revoked certificates that have expired.
  type: time
  exposed: true
  stored: true
  creation_only: true

- name: revokeDate
  description: Set time from when the certificate will be revoked.
  type: time
  exposed: true
  stored: true

- name: serialNumber
  description: SerialNumber of the revoked certificate.
  type: string
  exposed: true
  stored: true
  creation_only: true
  format: free

- name: subject
  description: Subject of the certificate related to the revocation.
  type: string
  exposed: true
  stored: true
  creation_only: true
  format: free
