package responses

import "github.com/gin-gonic/gin"

type UserResponse struct {
    Status  int    `json:"status"`
    Message string `json:"message"`
    Data    *gin.H `json:"data"`
}

type OrderResponse struct {
    Status  int    `json:"status"`
    Message string `json:"message"`
    Data    *gin.H `json:"data"`
}

type ProductResponse struct {
    Status  int    `json:"status"`
    Message string `json:"message"`
    Data    *gin.H `json:"data"`
}