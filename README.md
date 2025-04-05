# Auth Protobuf

This repository stores proto file with generated grpc-client and grpc-server on Golang for account service  

## gRPC API:

### Login
- **Request**: {
    - `string login` (required)
    - `string password` (required)
  }
- **Response**: {
    - `string token`
  }

### SignUp
- **Request**: {
    - `string login` (required)
    - `string email` (required)
    - `string password` (required)
  }
- **Reponse** {
    - `string token`
  }