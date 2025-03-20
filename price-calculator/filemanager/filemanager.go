package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type Filemanager struct {
	InputPath  string
	OutputPath string
}

func New(inputPath string, outputPath string) Filemanager {
	return Filemanager{
		InputPath:  inputPath,
		OutputPath: outputPath,
	}
}

func (fm Filemanager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputPath)

	if err != nil {
		return nil, errors.New("Error opening file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		return nil, errors.New("Error reading file")
	}

	return lines, nil

}

func (fm Filemanager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputPath)

	if err != nil {
		return errors.New("Error creating file")
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("JSON conversion failed")
	}

	return nil
}
