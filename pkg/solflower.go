package pkg

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

// Node struct is needed for generating the SOLFLOWER json file
type Node struct {
	Name     string  `json:"name"`
	Children []*Node `json:"children,omitempty"`
	Language string  `json:"language,omitempty"`
	Size     int     `json:"size,omitempty"`
}

type File struct {
	Language string
	Code     int
}

// GenerateSolFlowerJSON converts the json provided by cloc to another json file.
// This json file is stored in a specified output path.
// This json file is needed for generating the code flower shape.
func GenerateSolFlowerJSON(jsonData []byte, outputPath string) ([]byte, error) {
	log.Println("Generating SOLFLOWER JSON")
	root := &Node{
		Name:     "root",
		Children: nil,
		Size:     0,
	}

	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}

	for key, value := range data {
		if key == "header" || key == "SUM" {
			continue
		}

		fields := strings.Split(key, string(os.PathSeparator))

		temp := root
		for _, field := range fields[1 : len(fields)-1] {
			flag := false
			for _, child := range temp.Children {
				if child.Name == field {
					temp = child
					flag = true
					break
				}
			}

			if flag {
				continue
			}
			child := Node{
				Name:     field,
				Children: nil,
				Size:     0,
			}

			temp.Children = append(temp.Children, &child)
			temp = &child
		}

		b, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}
		var file File
		err = json.Unmarshal(b, &file)
		if err != nil {
			return nil, err
		}

		leaf := Node{
			Name:     fields[len(fields)-1],
			Children: nil,
			Language: file.Language,
			Size:     file.Code,
		}
		temp.Children = append(temp.Children, &leaf)
	}

	b, err := json.Marshal(*root)
	if err != nil {
		return nil, err
	}

	return b, err
}
