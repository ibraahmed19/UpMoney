package common1

type CasbinRule struct {
	ID    uint   `json:"id"`
	V0    string `json:"role"`
	V1    string `json:"object"`
	V2    string `json:"action"`
	Roles []Role `gorm:"many2many:role_permission"`
}
