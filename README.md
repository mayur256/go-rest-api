# go-rest-api
Basic REST API example for golang. Will be using using GIN framework

## 
- This project is REST API implementation for a store to allowing storing and sharing albums with it's customers.
- Users can add, update and remove albums to/from a store.

## Endpoints
1. GET /albums - Get a list of all albums, returned as JSON.
2. POST /albums – Add a new album from request data sent as JSON.
3. GET /albums/:id - Get an album by its ID, returning the album data as JSON.
4. DELETE /albums/:id - Removes an album by its ID.
