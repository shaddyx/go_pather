package pather

import (
	"fmt"
	"strconv"
	"strings"
)

type PatherError struct {
	ParentErr error
	Path      string
	FullPath  string
}

func GetPath(obj any, path string) (any, *PatherError) {
	path = strings.TrimSpace(path)
	split := strings.Split(path, ".")
	fmt.Printf("Splitted: %v\n", split)
	if len(path) == 0 || len(split) == 0 {
		fmt.Println("Returning result")
		return obj, nil
	}
	fmt.Printf("len: %d\n", len(split))
	objMap, ok := obj.(map[string]any)

	if ok {
		val, ok := objMap[split[0]]
		if !ok {
			return nil, &PatherError{
				ParentErr: fmt.Errorf("no such index in path: %s, fullPath: %s", split[0], path),
				Path:      split[0],
				FullPath:  path,
			}
		}
		fmt.Printf("Obtaining path: %s from %v\n", split[0], obj)
		return GetPath(val, strings.Join(split[1:], "."))
	}

	objList, ok := obj.([]any)

	if ok {
		fmt.Printf("Obtaining path: %s from list %v\n", split[0], obj)
		index, err := strconv.Atoi(split[0])
		if err != nil {
			return nil, &PatherError{
				ParentErr: err,
				Path:      split[0],
				FullPath:  path,
			}
		}
		if index >= len(objList) {
			return nil, &PatherError{
				ParentErr: fmt.Errorf("no such index in path: %d, fullPath: %s", index, path),
				Path:      split[0],
				FullPath:  path,
			}
		}

		return GetPath(objList[index], strings.Join(split[1:], "."))
	}

	if len(path) == 0 {
		return obj, nil
	}

	return nil, &PatherError{
		Path:     split[0],
		FullPath: path,
	}
}
