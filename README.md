# API Documentation

## Overview

### Public API

| Endpoint | Methods | Description |
| -------- | ------- | ----------- |
| `/api/healthz` | GET | Returns the status of the server |
| `/api/users` | POST, PUT | Create and update users |
| `/api/chirps` | POST, GET | Create and return chirps |
| `/api/chirps/{id}` | GET, DELETE | Delete and return chirp by chirp id |
| `/api/login` | POST | Login with user credentials |
| `/api/refresh` | POST | Refresh user's JWT |
| `/api/revoke` | POST | Revoke user's refresh token |

## Resources

### `users` - Chirpy users

**Payload Example**

```json
{
  "id": "00000000-0000-0000-0000-000000000000",  
  "created_at": "2025-09-12T13:34:08.714941Z",    
  "updated_at": "2025-09-12T13:34:08.714941Z",    
  "email": "test@example.com",
  "is_chirpy_red": true                
}
```

**Fields**

| Field | Type | Description |
| ----  | ---- | ----------- |
| `id` | string | user id (UUID) |
| `created_at` | string | ISO 8601/RFC 3339 timestamp in UTC |
| `updated_at` | string |  ISO 8601/RFC 3339 timestamp in UTC |
| `email` | string | user's email address |
| `is_chirpy_red` | boolean | subscription status |

---

### `chirps` - Chirpy posts

**Payload Example**

```json
{
  "id": "00000000-0000-0000-0000-000000000000",  
  "created_at": "2025-09-12T13:34:08.714941Z",    
  "updated_at": "2025-09-12T13:34:08.714941Z",    
  "body": "example chirp body text",
  "user_id": "10000000-1000-1000-1000-100000000000"      
}
```

**Fields**

| Field | Type | Description |
| ----  | ---- | ----------- |
| `id` | string | chirp id (UUID) |
| `created_at` | string | ISO 8601/RFC 3339 timestamp in UTC |
| `updated_at` | string |  ISO 8601/RFC 3339 timestamp in UTC |
| `body` | string |  chirp's body text |
| `user_id` | string |  user id (UUID) |

## Endpoints

### **`POST /api/users`** - Create a new Chirpy user

#### __Parameters__

- Headers
```bash
Content-Type: application/json
```

- Body (JSON)
```json
{
  "password": "examplePassword",
  "email": "test@example.com"
}
```
| Field | Type | Description |
| ----  | ---- | ----------- |
| `password` | string | user's password |
| `email` | string | user's email |


#### Response

- Response Body (JSON)
```json
{
  "id": "00000000-0000-0000-0000-000000000000",  
  "created_at": "2025-09-12T13:34:08.714941Z",    
  "updated_at": "2025-09-12T13:34:08.714941Z",
  "email": "test@example.com",
  "is_chirpy_red": true                
}
```
| Field | Type | Description |
| ----  | ---- | ----------- |
| `id` | string | user id (UUID) |
| `created_at` | string | ISO 8601/RFC 3339 timestamp in UTC |
| `updated_at` | string |  ISO 8601/RFC 3339 timestamp in UTC |
| `email` | string | user's email address |
| `is_chirpy_red` | boolean | subscription status |

--- 

### **`PUT /api/users`** - Update an existing Chirpy user

#### Parameters

- Headers
```bash
Content-Type: application/json
Authorization: Bearer <JWT>
```

- Body (JSON)
```json
{
  "password": "examplePassword",
  "email": "test@example.com"
}
```
| Field | Type | Description |
| ----  | ---- | ----------- |
| `password` | string | user's password |
| `email` | string | user's email |

#### Response

- Response Body (JSON)
```json
{
  "id": "00000000-0000-0000-0000-000000000000",  
  "created_at": "2025-09-12T13:34:08.714941Z",    
  "updated_at": "2025-09-12T13:34:08.714941Z",    
  "email": "test@example.com",
  "is_chirpy_red": true,                
}
```
| Field | Type | Description |
| ----  | ---- | ----------- |
| `id` | string | user id (UUID) |
| `created_at` | string | ISO 8601/RFC 3339 timestamp in UTC |
| `updated_at` | string |  ISO 8601/RFC 3339 timestamp in UTC |
| `email` | string | user's email address |
| `is_chirpy_red` | boolean | subscription status |

--- 

### **`POST /api/chirps`** - Create a new Chirpy post

#### Parameters

- Headers
```bash
Content-Type: application/json
Authorization: Bearer <JWT>
```

- Body (JSON)
```json
{
  "body": "example chirpy post body",
}
```
| Field | Type | Description |
| ----  | ---- | ----------- |
| `body` | string | chirp body text (140 char limit) |

#### Response

- Response Body (JSON)
```json
{
  "id": "00000000-0000-0000-0000-000000000000",  
  "created_at": "2025-09-12T13:34:08.714941Z",    
  "updated_at": "2025-09-12T13:34:08.714941Z",    
  "body": "example chirp body text",
  "user_id": "10000000-1000-1000-1000-100000000000"      
}
```

| Field | Type | Description |
| ----  | ---- | ----------- |
| `id` | string | chirp id (UUID) |
| `created_at` | string | ISO 8601/RFC 3339 timestamp in UTC |
| `updated_at` | string |  ISO 8601/RFC 3339 timestamp in UTC |
| `body` | string |  chirp's body text |
| `user_id` | string |  user id (UUID) |

--- 

### **`GET /api/chirps`** - Get a list of Chirpy posts

#### Parameters

- Queries

`/api/chirps?author_id=uuid&sort=desc`

| Query | Description |
| ----  | ----------- |
| `author_id` | Get chirps from a target user `(UUID)` |
| `sort` | Sort chirps in `asc` or `desc` order |

#### Response

- Response Body (JSON)
```json
[
  {
    "id": "00000000-0000-0000-0000-000000000000",  
    "created_at": "2025-09-12T13:34:08.714941Z",    
    "updated_at": "2025-09-12T13:34:08.714941Z",    
    "body": "example chirp body text",
    "user_id": "10000000-1000-1000-1000-100000000000"      
  },
  {
    "id": "10001001-0101-0101-0100-010010110100",  
    "created_at": "2025-09-12T13:34:08.714941Z",    
    "updated_at": "2025-09-12T13:34:08.714941Z",    
    "body": "example chirp body text",
    "user_id": "20000000-2000-2000-2000-200000000000"      
  },
  {
    "...": "more chirp objects"
  }
]
```

| Field | Type | Description |
| ----  | ---- | ----------- |
| `id` | string | chirp id (UUID) |
| `created_at` | string | ISO 8601/RFC 3339 timestamp in UTC |
| `updated_at` | string |  ISO 8601/RFC 3339 timestamp in UTC |
| `body` | string |  chirp's body text |
| `user_id` | string |  user id (UUID) |

--- 

### **`GET /api/chirps/{id}`** - Get a Chirpy post by ID

#### Parameters

- Headers
```bash
Content-Type: application/json
Authorization: Bearer <JWT>
```

- Body (JSON)
```json
{
  "body": "example chirpy post body",
}
```
| Field | Type | Description |
| ----  | ---- | ----------- |
| `body` | string | chirp body text (140 char limit) |

#### Response

- Response Body (JSON)
```json
{
  "id": "00000000-0000-0000-0000-000000000000",  
  "created_at": "2025-09-12T13:34:08.714941Z",    
  "updated_at": "2025-09-12T13:34:08.714941Z",    
  "body": "example chirp body text",
  "user_id": "10000000-1000-1000-1000-100000000000"      
}
```

| Field | Type | Description |
| ----  | ---- | ----------- |
| `id` | string | chirp id (UUID) |
| `created_at` | string | ISO 8601/RFC 3339 timestamp in UTC |
| `updated_at` | string |  ISO 8601/RFC 3339 timestamp in UTC |
| `body` | string |  chirp's body text |
| `user_id` | string |  user id (UUID) |
