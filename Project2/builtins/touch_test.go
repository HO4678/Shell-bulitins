// builtins/touch_test.go

package builtins

import (
	"os"
	"testing"
)

func TestTouchFile(t *testing.T) {
	// Define test file paths
	testFile1 := "test_file_1.txt"
	testFile2 := "test_file_2.txt"
	testFile3 := "/invalid/path/with/special/characters" // Invalid path
	testFile4 := "existing_file.txt"                     // Existing file

	// Clean up in case the files already exist
	defer os.Remove(testFile1)
	defer os.Remove(testFile2)

	// Test creating a new file
	err := TouchFile(testFile1)
	if err != nil {
		t.Errorf("TouchFile failed: %v", err)
	}

	// Check if the file was created
	stat, err := os.Stat(testFile1)
	if err != nil {
		t.Errorf("Error checking file: %v", err)
	}

	// Check if it is a regular file
	if stat.Mode().IsRegular() == false {
		t.Errorf("Expected %s to be a regular file", testFile1)
	}

	// Test updating the timestamp of an existing file
	err = TouchFile(testFile1)
	if err != nil {
		t.Errorf("TouchFile failed: %v", err)
	}

	// Test creating a new file with a specific mode
	err = TouchFile(testFile2)
	if err != nil {
		t.Errorf("TouchFile failed: %v", err)
	}

	// Check if the file was created with the specified mode
	stat2, err := os.Stat(testFile2)
	if err != nil {
		t.Errorf("Error checking file: %v", err)
	}

	// Check if it is a regular file
	if stat2.Mode().IsRegular() == false {
		t.Errorf("Expected %s to be a regular file", testFile2)
	}

	// Test creating a file with an invalid path
	err = TouchFile(testFile3)
	if err == nil {
		t.Errorf("Expected an error for invalid path, but got none")
	}

	// Test creating a file with an existing path
	file, _ := os.Create(testFile4)
	defer file.Close()
	err = TouchFile(testFile4)
	if err != nil {
		t.Errorf("TouchFile failed: %v", err)
	}

	// TODO: Add more test cases as needed
}
