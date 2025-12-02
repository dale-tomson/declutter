package organizer

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"time"
)

type FileInfo struct {
	Path    string
	ModTime time.Time
	Name    string
}

type Organizer struct {
	sourceDir   string
	logCallback func(string)
}

func New(sourceDir string, logCallback func(string)) *Organizer {
	return &Organizer{
		sourceDir:   sourceDir,
		logCallback: logCallback,
	}
}

func (o *Organizer) SourceDir() string {
	return o.sourceDir
}

func (o *Organizer) log(message string) {
	if o.logCallback != nil {
		o.logCallback(message)
	}
}

func (o *Organizer) GetFiles() ([]FileInfo, error) {
	var files []FileInfo

	entries, err := os.ReadDir(o.sourceDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			o.log(fmt.Sprintf("Warning: Could not get info for %s: %v", entry.Name(), err))
			continue
		}

		files = append(files, FileInfo{
			Path:    filepath.Join(o.sourceDir, entry.Name()),
			ModTime: info.ModTime(),
			Name:    entry.Name(),
		})
	}

	return files, nil
}

func (o *Organizer) OrganizeFiles(files []FileInfo) (int, int, error) {
	movedCount := 0
	skippedCount := 0
	createdFolders := make(map[string]bool)

	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime.Before(files[j].ModTime)
	})

	for _, file := range files {
		year := file.ModTime.Year()
		month := file.ModTime.Month()

		yearFolder := filepath.Join(o.sourceDir, fmt.Sprintf("%d", year))
		monthFolder := filepath.Join(yearFolder, fmt.Sprintf("%02d-%s", int(month), month.String()))

		if !createdFolders[yearFolder] {
			if err := o.ensureDir(yearFolder); err != nil {
				o.log(fmt.Sprintf("Error creating year folder %s: %v", yearFolder, err))
				continue
			}
			createdFolders[yearFolder] = true
		}

		if !createdFolders[monthFolder] {
			if err := o.ensureDir(monthFolder); err != nil {
				o.log(fmt.Sprintf("Error creating month folder %s: %v", monthFolder, err))
				continue
			}
			createdFolders[monthFolder] = true
		}

		destPath := filepath.Join(monthFolder, file.Name)

		if _, err := os.Stat(destPath); err == nil {
			o.log(fmt.Sprintf("Skipped (already exists): %s", file.Name))
			skippedCount++
			continue
		}

		if err := o.moveFile(file.Path, destPath); err != nil {
			o.log(fmt.Sprintf("Error moving %s: %v", file.Name, err))
			continue
		}

		o.log(fmt.Sprintf("Moved: %s → %d/%02d-%s/", file.Name, year, int(month), month.String()))
		movedCount++
	}

	return movedCount, skippedCount, nil
}

func (o *Organizer) ensureDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		o.log(fmt.Sprintf("Creating folder: %s", filepath.Base(path)))
		return os.MkdirAll(path, 0755)
	}
	return nil
}

func (o *Organizer) moveFile(src, dst string) error {
	if err := os.Rename(src, dst); err == nil {
		return nil
	}

	if err := o.copyFile(src, dst); err != nil {
		return err
	}

	return os.Remove(src)
}

func (o *Organizer) copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, sourceFile); err != nil {
		return err
	}

	srcInfo, err := os.Stat(src)
	if err == nil {
		os.Chtimes(dst, srcInfo.ModTime(), srcInfo.ModTime())
	}

	return destFile.Sync()
}

func GetYearMonthPath(baseDir string, t time.Time) string {
	year := t.Year()
	month := t.Month()
	yearFolder := filepath.Join(baseDir, fmt.Sprintf("%d", year))
	return filepath.Join(yearFolder, fmt.Sprintf("%02d-%s", int(month), month.String()))
}
