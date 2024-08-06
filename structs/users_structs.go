package Structs

type User struct {
    ID       string `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

type LoginUser struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type RegisterUser struct {
    ID              string `json:"id"`
    Name            string `json:"name"`
    Email           string `json:"email"`
    Password        string `json:"password"`
    ConfirmPassword string `json:"confirm_password"`
}