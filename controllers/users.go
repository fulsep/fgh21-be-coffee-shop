package controllers

import (
	"RGT/konis/dtos"
	"RGT/konis/lib"
	"RGT/konis/models"
	"RGT/konis/repository"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllUsers(c *gin.Context) {
	data, err := repository.FindAllUsers()

	if err != nil {
		lib.HandlerBadReq(c, "Invalid")
		return
	}
	lib.HandlerOK(c, "List all users", data, nil)
}

func GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		lib.HandlerBadReq(c, "Invalid")
		return
	}

	data, err := repository.FindUserById(id)

	if err != nil {
		lib.HandlerBadReq(c, "Data not found")
		return
	}
	lib.HandlerOK(c, "Get user by id", data, nil)
}

func CreateUser(c *gin.Context) {
	formUser := dtos.FormUser{}
	err := c.Bind(&formUser)

	if err != nil {
		lib.HandlerBadReq(c, "data not verified")
		return
	}

	roleId := 1

	data, err := repository.CreateUser(models.Users{
		Email:    formUser.Email,
		Password: formUser.Password,
		RoleId:   roleId,
	})

	if err != nil {
		lib.HandlerBadReq(c, "data not verified")
		return
	}

	lib.HandlerOK(c, "Create user success", data, nil)
}

func UpdateUserById(c *gin.Context) {
	formUser := dtos.FormUser{}
	err := c.Bind(&formUser)
	if err != nil {
		lib.HandlerBadReq(c, "data not verified")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		lib.HandlerBadReq(c, "data not verified")
		return
	}

	roleId := 1

	data, err := repository.UpdateUserById(models.Users{
		Email:    formUser.Email,
		Password: formUser.Password,
		RoleId:   roleId,
	}, id)

	if data.Id == 0 {
		lib.HandlerNotfound(c, "Data not found")
		return
	}

	if err != nil {
		lib.HandlerBadReq(c, "data not verified")
		return
	}

	lib.HandlerOK(c, "Update user success", data, nil)
}

func DeleteUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		lib.HandlerBadReq(c, "Invalid")
		return
	}

	data, err := repository.DeleteUserById(id)

	if data.Id == 0 {
		lib.HandlerNotfound(c, "Data not found")
		return
	}

	if err != nil {
		lib.HandlerBadReq(c, "Data not found")
		return
	}
	lib.HandlerOK(c, "Delete user success", data, nil)
}

func CreateUserWithProfile(c *gin.Context) {
	var input dtos.CreateUserProfileInput
	if err := c.Bind(&input); err != nil {
		lib.HandlerBadReq(c, "Invalid input data")
		return
	}

	// Validate image upload
	file, err := c.FormFile("profileImage")
	if err != nil {
		lib.HandlerBadReq(c, "Image upload failed")
		return
	}

	// Validate file extension
	allowedExtensions := map[string]bool{".jpg": true, ".jpeg": true, ".png": true}
	fileExt := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedExtensions[fileExt] {
		lib.HandlerBadReq(c, "Invalid file extension")
		return
	}

	// Generate new filename
	newFileName := uuid.New().String() + fileExt
	uploadDir := "./img/profile/"
	fullFilePath := uploadDir + newFileName

	// Save uploaded file
	if err := c.SaveUploadedFile(file, fullFilePath); err != nil {
		lib.HandlerBadReq(c, "Failed to upload image")
		return
	}

	// Create User
	user, err := repository.CreateinsertUser(models.InsertUsers{
		Email:    input.Email,
		Password: input.Password,
		RoleId:   input.RoleId,
	})
	if err != nil {
		// Print the actual error for debugging
		fmt.Println("Error creating user:", err)
		lib.HandlerBadReq(c, "Failed to create user: "+err.Error())
		return
	}

	// Create Profile
	profile, err := repository.CreateinsertProfile(models.InsertProfile{
		FullName:    input.FullName,
		PhoneNumber: &input.PhoneNumber,
		Address:     &input.Address,
		Image:       &fullFilePath,
		UserId:      user.Id, // Ensure the UserId is correctly passed
	})
	if err != nil {
		// Print the actual error for debugging
		fmt.Println("Error creating profile:", err)
		lib.HandlerBadReq(c, "Failed to create profile: "+err.Error())
		return
	}

	lib.HandlerOK(c, "User and Profile created successfully", profile, nil)
}
