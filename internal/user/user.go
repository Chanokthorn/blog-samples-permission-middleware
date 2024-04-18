package user

// User Represents information of current user
type User struct {
	ShopOwnerID int  `json:"shopOwnerID"`
	Role        Role `json:"role"`
}
