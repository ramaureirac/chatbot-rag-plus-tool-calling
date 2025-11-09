# Consuming API

When serving the WebApplicatio you can go to http://localhost:8080 to use the Vue.js frontend. If needed you can also interact directly with the API of the applications.

## Autenticate

**URL:** `http://localhost:8080/login`  
**Method:** `POST`  
**Description:** `Authenticate the user and returns a UUID in case of success.`  
**Headers:**  `Content-Type: application/json`  
**Response: 200**  

    {
        "message": "123e4567-e89b-12d3-a456-426614174000"
    }

## Ask

**URL:** `http://localhost:8080/ask`  
**Method:** `POST`  
**Description:** `Send a query to the ChatBot.`  
**Headers:**  
- `Content-Type: application/json`
- `X-Anon-ID: 123e4567-e89b-12d3-a456-426614174000`  

**Body:**  

    {
        "message": "hello there!" 
    }

**Response: 200**  

    {
        "message": "hi"
    }