package profile

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Handler to save user details
func Handler(ctx *gin.Context) {
	// Check if user is authenticated
	session := sessions.Default(ctx)
	profile := session.Get("profile")
	if profile == nil {
		ctx.String(http.StatusUnauthorized, "User is not authenticated. Please sign in first.")
		return
	}

	// Parse form data
	firstname := ctx.PostForm("firstname")
	middlename := ctx.PostForm("middlename")
	lastname := ctx.PostForm("lastname")
	phoneno := ctx.PostForm("phoneno")
	city := ctx.PostForm("city")
	state := ctx.PostForm("state")
	country := ctx.PostForm("country")

	// Connect to MySQL database
	db, err := sql.Open("mysql", os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		log.Printf("Database connection error: %v", err)
		ctx.String(http.StatusInternalServerError, "Database connection failed")
		return
	}
	defer db.Close()

	// Insert user data into the database
	query := `
    INSERT INTO users (firstname, middlename, lastname, phoneno, city, state, country)
    VALUES (?, ?, ?, ?, ?, ?, ?)
`
	_, err = db.Exec(query, firstname, middlename, lastname, phoneno, city, state, country)
	if err != nil {
		log.Printf("Error inserting user data: %v", err)
		ctx.String(http.StatusInternalServerError, "Failed to save data")
		return
	}

	log.Println("User details saved successfully!")
	ctx.String(http.StatusOK, "User details saved successfully!")
}
