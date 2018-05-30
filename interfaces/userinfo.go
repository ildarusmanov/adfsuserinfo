package interfaces

type Userinfo interface {
    GetId() string
    GetUsername() string
    GetEmail() string
    GetFullName() string
    GetPhotoUrl() string
}