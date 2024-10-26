package util

import (
	"fmt"
	"io"
	"os"
)

/**
 * @author Ted Yoo
 * @breif File util
 * @date 2024.10.26
 */

func ExistsFile(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateFile(path string) error {
	// create file if not exists
	if !ExistsFile(path) {
		file, err := os.Create(path)
		if isError(err) {
			fmt.Println(err.Error())
			return err
		}
		defer func(file *os.File) {
			if err := file.Close(); err != nil {
				fmt.Println(err.Error())
				return
			}
		}(file)
		fmt.Println("File Created Successfully", path)
	}

	return nil
}

func DeleteFile(path string) error {
	// delete file
	var err = os.Remove(path)
	if isError(err) {
		return err
	}

	fmt.Println("File Deleted")
	return nil
}

func WriteFile(path string) error {
	// Open file using READ & WRITE permission.
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return err
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
			return
		}
	}(file)

	// Write some text line-by-line to file.
	_, err = file.WriteString("Hello \n")
	if isError(err) {
		return err
	}
	_, err = file.WriteString("World \n")
	if isError(err) {
		return err
	}

	// Save file changes.
	err = file.Sync()
	if isError(err) {
		return err
	}

	fmt.Println("File Updated Successfully.")

	return nil
}

func ReadFile(path string) (string, error) {
	// Open file for reading.
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return "", err
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
			return
		}
	}(file)

	// Read file, line by line
	text := make([]byte, 1024)
	for {
		_, err = file.Read(text)

		// Break if finally arrived at end of file
		if err == io.EOF {
			break
		}

		// Break if error occured
		if err != nil && err != io.EOF {
			isError(err)
			break
		}
	}

	fmt.Println("Reading from file.")

	return string(text), nil
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return err != nil
}
