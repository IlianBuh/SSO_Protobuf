# SSO Protobuf

This repository stores proto files with generated grpc-client and grpc-server on Golang for sso service(auth and userinfo)

## Auth gRPC API:

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

  
## UserInfo gRPC API:

### Users
- **Request**: {
    - `repeated int32 uuids` (required)
  }
- **Response**: {
    - `repeated User users`
  }

### User
- **Request**: {
    - `int32 uuid` (required)
  }
- **Reponse** {
    - `User user`
  }

Object `User` has following structure:
`User {
  int32 uuid = 1;
  string login = 2;
  string email = 3; 
}`