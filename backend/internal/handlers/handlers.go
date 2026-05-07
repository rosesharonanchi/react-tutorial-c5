package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"piggy.com/internal/models"
	"piggy.com/internal/piggyservice"
)

type Handler struct {
	service *piggyservice.Service
	authService *piggyservice.AuthService
}

func NewHandler(service *piggyservice.Service, as *piggyservice.AuthService) *Handler {
	return &Handler{service: service,
		authService: as,
	}

}

func (h *Handler) CreateTransaction(c *gin.Context) {
   
	//userID := c.MustGet("userID").(*int32) // Get user ID from context, set by auth middleware
   idStr := c.GetHeader("X-User-ID")
val, err := strconv.ParseInt(idStr, 10, 32)
if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
    return
}
userID := int32(val)
	var payload models.CreateTransactionPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := h.service.CreateTransaction(c, &userID,  payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transaction)
}


func (h *Handler) GetTransactions(c *gin.Context) {
	val, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Please log in again"})
        return
    }
	//important
   var userID int32
    if ptr, ok := val.(*int32); ok {
        userID = *ptr
    } else if v, ok := val.(int32); ok {
        userID = v
    } else {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid User ID type in context"})
        return
    }
	
	// Get the transaction type from query parameters, default to empty string if not provided
	transactionType := c.Query("type")
	transactions, err := h.service.GetTransactions(c.Request.Context() , userID,  transactionType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if transactions == nil {
        c.JSON(http.StatusOK, []interface{}{})
        return
    }

	c.JSON(http.StatusOK, transactions)
}


//register and login
func (h *Handler) SignUp(c *gin.Context) {
    var payload models.RegisterPayload
    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
        return
    }

    user, err := h.authService.Register(c.Request.Context(), payload.Name, 
        payload.Email, 
        payload.Password,)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "User registered successfully",
        "user_id": user.ID,
    })
}

func (h *Handler) Login(c *gin.Context) {
    var payload models.LoginPayload
    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password required"})
        return
    }

    user, err := h.authService.Login(c.Request.Context(), payload.Email, payload.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Login successful",
        "user_id": user.ID,
    })
}