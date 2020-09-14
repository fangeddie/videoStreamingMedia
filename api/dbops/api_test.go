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

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
	t.Run("RegetVideo", testReGetVideoInfo)
}

var tempVid string

func testAddVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(1, "爆裂鼓手")
	if err != nil {
		t.Errorf("Error of AddVideoInfo: %v", err)
	}

	tempVid = vi.Id
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempVid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempVid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}

func testReGetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempVid)
	if err != nil || vi != nil {
		t.Errorf("Error of RegetVideoInfo: %v", err)
	}
}
