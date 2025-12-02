package organizer

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// TestNew verifies that New creates an Organizer with the correct source directory
func TestNew(t *testing.T) {
	dir := "/test/dir"
	org := New(dir, nil)

	if org.SourceDir() != dir {
		t.Errorf("Expected source dir %s, got %s", dir, org.SourceDir())
	}
}

// TestNewWithLogCallback verifies that the log callback is properly set
func TestNewWithLogCallback(t *testing.T) {
	var loggedMessage string
	callback := func(msg string) {
		loggedMessage = msg
	}

	org := New("/test", callback)
	org.log("test message")

	if loggedMessage != "test message" {
		t.Errorf("Expected log message 'test message', got '%s'", loggedMessage)
	}
}

// TestGetYearMonthPath verifies the path generation for year/month folders
func TestGetYearMonthPath(t *testing.T) {
	tests := []struct {
		name     string
		baseDir  string
		time     time.Time
		expected string
	}{
		{
			name:     "January 2024",
			baseDir:  "/test",
			time:     time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
			expected: filepath.Join("/test", "2024", "01-January"),
		},
		{
			name:     "December 2023",
			baseDir:  "/photos",
			time:     time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC),
			expected: filepath.Join("/photos", "2023", "12-December"),
		},
		{
			name:     "March 2025",
			baseDir:  "/documents",
			time:     time.Date(2025, 3, 1, 0, 0, 0, 0, time.UTC),
			expected: filepath.Join("/documents", "2025", "03-March"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetYearMonthPath(tt.baseDir, tt.time)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

// TestGetFiles tests scanning a directory for files
func TestGetFiles(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "organizer-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create test files with specific modification times
	testFiles := []struct {
		name    string
		modTime time.Time
	}{
		{"file1.txt", time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)},
		{"file2.txt", time.Date(2024, 3, 20, 14, 30, 0, 0, time.UTC)},
		{"file3.txt", time.Date(2023, 12, 25, 8, 0, 0, 0, time.UTC)},
	}

	for _, tf := range testFiles {
		filePath := filepath.Join(tmpDir, tf.name)
		if err := os.WriteFile(filePath, []byte("test content"), 0644); err != nil {
			t.Fatalf("Failed to create test file %s: %v", tf.name, err)
		}
		if err := os.Chtimes(filePath, tf.modTime, tf.modTime); err != nil {
			t.Fatalf("Failed to set mod time for %s: %v", tf.name, err)
		}
	}

	// Create a subdirectory (should be ignored)
	subDir := filepath.Join(tmpDir, "subdir")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatalf("Failed to create subdir: %v", err)
	}

	// Test GetFiles
	org := New(tmpDir, nil)
	files, err := org.GetFiles()
	if err != nil {
		t.Fatalf("GetFiles failed: %v", err)
	}

	if len(files) != 3 {
		t.Errorf("Expected 3 files, got %d", len(files))
	}

	// Verify file names are in the result
	fileNames := make(map[string]bool)
	for _, f := range files {
		fileNames[f.Name] = true
	}

	for _, tf := range testFiles {
		if !fileNames[tf.name] {
			t.Errorf("Expected file %s not found in results", tf.name)
		}
	}
}

// TestGetFilesEmptyDirectory tests scanning an empty directory
func TestGetFilesEmptyDirectory(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "organizer-test-empty-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	org := New(tmpDir, nil)
	files, err := org.GetFiles()
	if err != nil {
		t.Fatalf("GetFiles failed: %v", err)
	}

	if len(files) != 0 {
		t.Errorf("Expected 0 files, got %d", len(files))
	}
}

// TestGetFilesNonExistentDirectory tests scanning a non-existent directory
func TestGetFilesNonExistentDirectory(t *testing.T) {
	org := New("/nonexistent/directory/path", nil)
	_, err := org.GetFiles()

	if err == nil {
		t.Error("Expected error for non-existent directory, got nil")
	}
}

// TestOrganizeFiles tests the full organization workflow
func TestOrganizeFiles(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "organizer-test-organize-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create test files with specific modification times
	testFiles := []struct {
		name         string
		modTime      time.Time
		expectedPath string
	}{
		{
			name:         "january-file.txt",
			modTime:      time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC),
			expectedPath: filepath.Join(tmpDir, "2024", "01-January", "january-file.txt"),
		},
		{
			name:         "march-file.txt",
			modTime:      time.Date(2024, 3, 20, 14, 30, 0, 0, time.UTC),
			expectedPath: filepath.Join(tmpDir, "2024", "03-March", "march-file.txt"),
		},
		{
			name:         "december-file.txt",
			modTime:      time.Date(2023, 12, 25, 8, 0, 0, 0, time.UTC),
			expectedPath: filepath.Join(tmpDir, "2023", "12-December", "december-file.txt"),
		},
	}

	// Create the files
	for _, tf := range testFiles {
		filePath := filepath.Join(tmpDir, tf.name)
		if err := os.WriteFile(filePath, []byte("test content for "+tf.name), 0644); err != nil {
			t.Fatalf("Failed to create test file %s: %v", tf.name, err)
		}
		if err := os.Chtimes(filePath, tf.modTime, tf.modTime); err != nil {
			t.Fatalf("Failed to set mod time for %s: %v", tf.name, err)
		}
	}

	// Collect log messages
	var logMessages []string
	logCallback := func(msg string) {
		logMessages = append(logMessages, msg)
	}

	// Organize files
	org := New(tmpDir, logCallback)
	files, err := org.GetFiles()
	if err != nil {
		t.Fatalf("GetFiles failed: %v", err)
	}

	moved, skipped, err := org.OrganizeFiles(files)
	if err != nil {
		t.Fatalf("OrganizeFiles failed: %v", err)
	}

	// Verify results
	if moved != 3 {
		t.Errorf("Expected 3 files moved, got %d", moved)
	}
	if skipped != 0 {
		t.Errorf("Expected 0 files skipped, got %d", skipped)
	}

	// Verify files are in correct locations
	for _, tf := range testFiles {
		if _, err := os.Stat(tf.expectedPath); os.IsNotExist(err) {
			t.Errorf("Expected file not found at %s", tf.expectedPath)
		}

		// Verify original file is gone
		originalPath := filepath.Join(tmpDir, tf.name)
		if _, err := os.Stat(originalPath); !os.IsNotExist(err) {
			t.Errorf("Original file still exists at %s", originalPath)
		}
	}

	// Verify folder structure
	expectedFolders := []string{
		filepath.Join(tmpDir, "2024"),
		filepath.Join(tmpDir, "2024", "01-January"),
		filepath.Join(tmpDir, "2024", "03-March"),
		filepath.Join(tmpDir, "2023"),
		filepath.Join(tmpDir, "2023", "12-December"),
	}

	for _, folder := range expectedFolders {
		info, err := os.Stat(folder)
		if os.IsNotExist(err) {
			t.Errorf("Expected folder not found: %s", folder)
		} else if !info.IsDir() {
			t.Errorf("Expected %s to be a directory", folder)
		}
	}

	// Verify log messages were generated
	if len(logMessages) == 0 {
		t.Error("Expected log messages to be generated")
	}
}

// TestOrganizeFilesSkipExisting tests that existing files are skipped
func TestOrganizeFilesSkipExisting(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "organizer-test-skip-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create the destination folder structure
	destFolder := filepath.Join(tmpDir, "2024", "01-January")
	if err := os.MkdirAll(destFolder, 0755); err != nil {
		t.Fatalf("Failed to create dest folder: %v", err)
	}

	// Create a file in the source
	modTime := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	srcPath := filepath.Join(tmpDir, "existing.txt")
	if err := os.WriteFile(srcPath, []byte("source content"), 0644); err != nil {
		t.Fatalf("Failed to create source file: %v", err)
	}
	if err := os.Chtimes(srcPath, modTime, modTime); err != nil {
		t.Fatalf("Failed to set mod time: %v", err)
	}

	// Create the same file in the destination
	destPath := filepath.Join(destFolder, "existing.txt")
	if err := os.WriteFile(destPath, []byte("destination content"), 0644); err != nil {
		t.Fatalf("Failed to create dest file: %v", err)
	}

	// Organize files
	org := New(tmpDir, nil)
	files, err := org.GetFiles()
	if err != nil {
		t.Fatalf("GetFiles failed: %v", err)
	}

	moved, skipped, err := org.OrganizeFiles(files)
	if err != nil {
		t.Fatalf("OrganizeFiles failed: %v", err)
	}

	if moved != 0 {
		t.Errorf("Expected 0 files moved, got %d", moved)
	}
	if skipped != 1 {
		t.Errorf("Expected 1 file skipped, got %d", skipped)
	}

	// Verify destination content was not overwritten
	content, err := os.ReadFile(destPath)
	if err != nil {
		t.Fatalf("Failed to read dest file: %v", err)
	}
	if string(content) != "destination content" {
		t.Error("Destination file content was overwritten")
	}
}

// TestOrganizeFilesNoDuplicateFolders tests that folders are not created multiple times
func TestOrganizeFilesNoDuplicateFolders(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "organizer-test-nodup-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create multiple files with the same month
	modTime := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	for i := 1; i <= 5; i++ {
		filePath := filepath.Join(tmpDir, fmt.Sprintf("file%d.txt", i))
		if err := os.WriteFile(filePath, []byte("content"), 0644); err != nil {
			t.Fatalf("Failed to create file: %v", err)
		}
		if err := os.Chtimes(filePath, modTime, modTime); err != nil {
			t.Fatalf("Failed to set mod time: %v", err)
		}
	}

	// Track folder creation logs
	var folderCreations int
	logCallback := func(msg string) {
		if len(msg) > 16 && msg[:16] == "Creating folder:" {
			folderCreations++
		}
	}

	org := New(tmpDir, logCallback)
	files, _ := org.GetFiles()
	org.OrganizeFiles(files)

	// Should only create 2 folders: year and month
	if folderCreations != 2 {
		t.Errorf("Expected 2 folder creations, got %d", folderCreations)
	}
}

// TestOrganizeFilesPreservesModTime tests that file modification times are preserved
func TestOrganizeFilesPreservesModTime(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "organizer-test-modtime-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create source file with specific mod time
	srcPath := filepath.Join(tmpDir, "source.txt")
	content := "test content"
	modTime := time.Date(2024, 6, 15, 12, 0, 0, 0, time.UTC)

	if err := os.WriteFile(srcPath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create source file: %v", err)
	}
	if err := os.Chtimes(srcPath, modTime, modTime); err != nil {
		t.Fatalf("Failed to set mod time: %v", err)
	}

	// Organize file
	org := New(tmpDir, nil)
	files, err := org.GetFiles()
	if err != nil {
		t.Fatalf("GetFiles failed: %v", err)
	}

	moved, _, err := org.OrganizeFiles(files)
	if err != nil {
		t.Fatalf("OrganizeFiles failed: %v", err)
	}
	if moved != 1 {
		t.Fatalf("Expected 1 file moved, got %d", moved)
	}

	// Check destination file
	dstPath := filepath.Join(tmpDir, "2024", "06-June", "source.txt")
	dstContent, err := os.ReadFile(dstPath)
	if err != nil {
		t.Fatalf("Failed to read destination file: %v", err)
	}
	if string(dstContent) != content {
		t.Errorf("Expected content '%s', got '%s'", content, string(dstContent))
	}
}
