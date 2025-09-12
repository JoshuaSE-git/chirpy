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

`users` - Chirpy users

**Payload Example**

```json
{
  "id": "00000000-0000-0000-0000-000000000000"  
  "created_at": "2025-09-12T13:34:08.714941Z"    
  "updated_at": "2025-09-12T13:34:08.714941Z"    
  "email": "test@example.com"
  "is_chirpy_red": true                
}
```

**Fields**

| Name | Description |
| ---- | ----------- |
| `id` | UUID (string) |
| `created_at` | ISO 8601/RFC 3339 timestamp in UTC |
| `updated_at` | ISO 8601/RFC 3339 timestamp in UTC |
| `email` | user's email address |
| `is_chirpy_red` | boolean, subscription status |

---

`chirps` - Chirpy posts

**Payload Example**

```json
{
  "id": "00000000-0000-0000-0000-000000000000"  
  "created_at": "2025-09-12T13:34:08.714941Z"    
  "updated_at": "2025-09-12T13:34:08.714941Z"    
  "body": "example chirp body text"
  "user_id": "10000000-1000-1000-1000-100000000000"      
}
```

**Fields**

| Name | Description |
| ---- | ----------- |
| `id` | UUID (string) |
| `created_at` | ISO 8601/RFC 3339 timestamp in UTC |
| `updated_at` | ISO 8601/RFC 3339 timestamp in UTC |
| `body` | chirp's body text |
| `user_id` | UUID (string) |


