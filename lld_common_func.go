package main

import (
	"fmt"
	"log"
)

// ==================== Interfaces ====================

// CommonActions defines operations available to all users
type CommonActions interface {
	A()
	B()
	C()
}

// AdminActions defines operations only available to admins
type AdminActions interface {
	D()
	E()
}

// SystemUser represents the base user type
type SystemUser interface {
	CommonActions
	ID() string
	Name() string
	IsAdmin() bool
}

// AdminUser extends SystemUser with admin capabilities
type AdminUser interface {
	SystemUser
	AdminActions
}

// ==================== Implementations ====================

// user implements the basic user type
type user struct {
	id   string
	name string
}

func (u *user) ID() string {
	return u.id
}

func (u *user) Name() string {
	return u.name
}

func (u *user) IsAdmin() bool {
	return false
}

func (u *user) A() {
	fmt.Printf("User %s executing A()\n", u.name)
}

func (u *user) B() {
	fmt.Printf("User %s executing B()\n", u.name)
}

func (u *user) C() {
	fmt.Printf("User %s executing C()\n", u.name)
}

// admin embeds user and adds admin capabilities
type admin struct {
	user
}

func (a *admin) IsAdmin() bool {
	return true
}

func (a *admin) D() {
	fmt.Printf("Admin %s executing privileged D()\n", a.name)
}

func (a *admin) E() {
	fmt.Printf("Admin %s executing privileged E()\n", a.name)
}

// ==================== Factory Functions ====================

// NewUser creates a regular user
func NewUser(id, name string) SystemUser {
	return &user{
		id:   id,
		name: name,
	}
}

// NewAdmin creates an admin user
func NewAdmin(id, name string) AdminUser {
	return &admin{
		user: user{
			id:   id,
			name: name,
		},
	}
}

// ==================== Service Implementation ====================

// SystemService handles operations that need to check permissions
type SystemService struct{}

func (s *SystemService) ProcessAction(u SystemUser) {
	// Common actions available to all users
	u.A()
	u.B()
	u.C()

	// Check for admin privileges
	if adminUser, ok := u.(AdminUser); ok {
		adminUser.D()
		adminUser.E()
	}
}

// ==================== Usage Example ====================

func main() {
	// Create users
	regularUser := NewUser("u1", "John Doe")
	adminUser := NewAdmin("a1", "Jane Admin")

	// Create service
	service := &SystemService{}

	// Process regular user
	fmt.Println("Processing regular user:")
	service.ProcessAction(regularUser)

	// Process admin user
	fmt.Println("\nProcessing admin user:")
	service.ProcessAction(adminUser)

	// Type assertion example
	if admin, ok := adminUser.(AdminActions); ok {
		fmt.Println("\nDirect admin actions:")
		admin.D()
		admin.E()
	}

	// Attempt to use admin functions as regular user (will panic)
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic:", r)
		}
	}()

	// This demonstrates the protection - don't actually do this in production
	// Instead always check IsAdmin() first
	// fakeAdmin := regularUser.(AdminUser)
	// fakeAdmin.D()
}

// ==================== Testing Considerations ====================

/*
// Mock implementations for testing
type mockUser struct {
	CommonActions
}

type mockAdmin struct {
	mockUser
	AdminActions
}
*/
