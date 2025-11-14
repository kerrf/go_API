A simple exercise project that sets up an API.
There is a MySQL database with three tables:
1. users
2. events
3. registrations

Anyone can create a user and login with their credentials. Afterwards, they receive an
auth. token.
Users can create events, update, and delete events. Updates and deletion can
only be done by the user that created the event.
Anyone can fetch the details of any event or retrieve all events.
Password are stored as hashes with bcrypt.
Users can register for an event and can deregister themselves.

The following API requests are possible:
GET /events
GET /events/"id"
POST /events
PUT /events/"id"
DELETE /events/"id"
PUT

## Example

`PUT http://localhost:8080/events/4
Content-type: application/json
authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im1hZ3V5IiwiZXhwIjoxNzYzMDc3NjYzLCJ1c2VySUQiOjIxfQ.zVBFZYhJol8bl8Mw8298-pca6SMN9Dk-DIAELUI1IZQ

{
    "name" : "UPDATED",
    "description" : "updated",
    "location" : "am sack",
    "dateTime": "2025-10-12T15:30:12.00Z"
}`
