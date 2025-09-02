package gomodhaikom

import (
	"os"
	"testing"
)

//func TestGetUser(t *testing.T) {
// Set up test data
// token := "test-token"
// client := "test-client"
// requestid := "test-requestid"
// userXML := `
// 	<response>
// 		<user>
// 			<customerid>12345</customerid>
// 			<username>test-user</username>
// 			<email>test@example.com</email>
// 			<firstname>Test</firstname>
// 			<lastname>User</lastname>
// 		</user>
// 	</response>
// `
//
//
// // Set up mock HTTP client
// mockHaikom := mockHaikom{}

//  func (h mockHaikom) GetUser(token, client, requestid string) (User, error) {
//    user := User{}
//    user.CustomerId = "12345"
//    user.Username = "test-user"
//    user.Email = "test@example.com"
//    user.Firstname = "Test"
//    user.Lastname = "User"
//    return user, nil
//
// 	// Return mock response
// }
//
// t.Run("TestGetUser", func(t *testing.T) {
//    // Call function to be tested
//    UserMocClient = NewHaikomUserClient(mockHaikom)
//    user, err := UserMocClient.User.GetUser(token, client, requestid)
//
//    // Check that the function returns the expected values
//    assert.Equal(t, "12345", user.CustomerId)
//    assert.Equal(t, "test-user", user.Username)
//    assert.Equal(t, "test@example.com", user.Email)
//    assert.Equal(t, "Test", user.Firstname)
//    assert.Equal(t, "User", user.Lastname)
//    assert.NoError(t, err)
//  })

//}

func TestGetUserReal(t *testing.T) {
	// Set up test data
	token := os.Getenv("HAIKOM_TOKEN")
	client := "billes"
	requestid := "test-requestid"

	haikom := HaikomUser{
		User:     os.Getenv("HAIKOM_USER"),
		Password: os.Getenv("HAIKOM_PASSWORD"),
		Project:  os.Getenv("HAIKOM_PROJECT"),
		Url:      os.Getenv("HAIKOM_URL"),
	}

	t.Run("TestGetUser", func(t *testing.T) {
		// Call function to be tested
		UserTestService := NewUserService(haikom)
		user, err := UserTestService.User.GetUser(token, client, requestid)
		if err != nil {
			t.Errorf("Error getting user: %v", err)
		}
		t.Log(user)
	})
}
