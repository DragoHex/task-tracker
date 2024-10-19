package lib

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func Test_initData(t *testing.T) {
	// test data file creation
	tasks.DataFile = filepath.Join("..", "..", "data", "test_data.json")
	initData()
	_, err := os.ReadFile(tasks.DataFile)
	if err != nil {
		t.Errorf("data file not created, getting this error when trying to read it: %s", err)
	}
	os.Remove(tasks.DataFile)

	// test imported data
	test_data := `{
  "current_id": 1,
  "count": 1,
  "tasks": [
    {
      "id": 1,
      "description": "new task 1",
      "status": 1,
      "create_at": "2024-10-19T15:09:57.495819+05:30",
      "updated_at": "2024-10-19T15:09:57.495819+05:30"
    }
  ]
}`
	err = os.WriteFile(DataFile, []byte(test_data), 0644)
	if err != nil {
		t.Errorf("error in writing to file: %s", err)
	}

	expected_task := Task{
		Id:          1,
		Description: "new task 1",
	}

	DataFile = filepath.Join("..", "..", "data", "test_data.json")
	initData()
	data, err := os.ReadFile(DataFile)
	if err != nil {
		t.Errorf("data file not created, getting this error when trying to read it: %s", err)
	}

	var taskData TaskData
	err = json.Unmarshal(data, &taskData)
	if err != nil {
		t.Errorf("error in unmarshalling: %s", err)
	}

	if expected_task.Id != taskData.Tasks[0].Id ||
		expected_task.Description != taskData.Tasks[0].Description {
		t.Error("imported data doesn't match the expected data")
	}

	os.Remove(DataFile)
}
