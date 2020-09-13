package dbops

import (
	"testing"
)

// init(dblogin, truncate tables) -> run tests -> clear data(truncate tables)

func clearTables() {
	dbConn.Exec("truncate users;")
	dbConn.Exec("truncate video_info;")
	dbConn.Exec("truncate comments;")
	dbConn.Exec("truncate sessions;")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("ReGet", testReGetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("fang", "123")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("fang")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser: %v", err)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("fang", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v", err)
	}
}

func testReGetUser(t *testing.T) {
	pwd, err := GetUserCredential("fang")
	if err != nil {
		t.Errorf("Error of ReGetUser: %v", err)
	}
	if pwd != "" {
		t.Error("ReGetUser failed.")
	}
}
