package resetPassword

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/smtp"
	"os"
	"regexp"
	"template_rest_api/api/app/user"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// signup user

type Database struct {
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
}

func SendMailSimple(password string, email string) {
	auth := smtp.PlainAuth(
		"",
		"ahmedbrahim0110@gmail.com",
		"euut fmcn qths tquq",
		"smtp.gmail.com",
	)

	// HTML content for the email
	htmlContent := `
    <html>
<head>
<style>
body {
font-family: Arial, sans-serif;
background-color: #f1f1f1;
}
.container {
max-width: 600px;
margin: 0 auto;
padding: 40px;
background-color: #ffffff;
box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1);
}
h1 {
color: #333333;
font-size: 24px;
margin-top: 0;
}
p {
color: #666666;
font-size: 16px;
line-height: 1.5;
margin-bottom: 20px;
}
.button {
display: inline-block;
padding: 10px 20px;
background-color: #4CAF50;
color: #ffffff;
text-decoration: none;
border-radius: 4px;
}
.button:hover {
background-color: #45a049;
}
</style>
</head>
<body>
<div class="container">
<h1>TeamFlow: Reset Password</h1>
<p>Your email  is: <strong>` + email + `</strong></p>
<p>Your new password is: <strong>` + password + `</strong></p>
<p>Make sure to keep your password secure and do not share it with anyone.</p>
<p>If you didn't request a password reset, please contact our support team immediately.</p>
</div>
</body>
</html>
    `

	msg := []byte(fmt.Sprintf(
		"Subject: My Special Subject\n"+
			"From: %s\n"+
			"To: %s\n"+
			"MIME-version: 1.0;\n"+
			"Content-Type: text/html; charset=\"UTF-8\";\n\n"+
			"%s",
		"fafa86404@gmail.com",
		"fafa86404@gmail.com,",
		htmlContent,
	))

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"ahmedbrahim0110@gmail.com",
		[]string{email},
		msg,
	)
	if err != nil {
		fmt.Print(err)
	}
}
func generateStrongPassword(length int) (string, error) {
	if length < 1 {
		return "", fmt.Errorf("length should be greater than 0")
	}

	// Determine the number of bytes needed to generate the password
	numBytes := (length * 3) / 4

	// Generate a random byte slice
	randomBytes := make([]byte, numBytes)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Encode the random bytes to base64
	password := base64.RawURLEncoding.EncodeToString(randomBytes)

	// Truncate the password to the desired length
	if len(password) > length {
		password = password[:length]
	}

	return password, nil
}

func (db Database) SignUpUser(ctx *gin.Context) {

	// init vars
	var account resetPassword
	empty_reg, _ := regexp.Compile(os.Getenv("EMPTY_REGEX"))
	email_reg, _ := regexp.Compile(os.Getenv("EMAIL_REGEX"))

	// upmarshal sent json
	if err := ctx.ShouldBindJSON(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check field validity
	if empty_reg.MatchString(account.Email) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "please complete all fields"})
		return
	}

	// check email validity
	if !email_reg.MatchString(account.Email) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid email"})
		return
	}

	if exists := CheckMailExists(db.DB, account.Email); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "email does not exist "})
		return
	}

	// hash password
	password, err := generateStrongPassword(12)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Print(password)

	users, err := user.GetUserByEmail(db.DB, account.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	SendMailSimple(password, account.Email)

	// update new password

	user.HashPassword(&password)
	err = UpdatePassword(db.DB, users.ID, password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "updated password & mail has been send"})
}
