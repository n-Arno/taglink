taglink
=======

Proof of Concept of linking Scaleway Loadbalancer to Instances via tag.

How it works:
- `taglink <lb-id>`
- Find LB via ID
- Read LB tags
- Find Instances via tags
- Get public IPs of Instances
- Find LB first Backend
- Update Backend Server IPs with Instances Public IPs
