# Project Crawler

## Description
Project Crawler is a Go-based utility designed to recursively scan directories, capture their structure, and output the data in a structured JSON format. It also supports excluding specified directories and files via a configuration file.

## Features
- **Recursive Directory Scanning**: Traverse directories and subdirectories to capture the complete structure.
- **File Content Extraction**: Read and include the content of files within the structure.
- **Exclusion Configuration**: Exclude specified directories and files via a JSON configuration file (`crawlerIgnore.json`).
- **JSON Output**: Generate a JSON file representing the directory structure.

## Usage


### 1. Clone the Repository
```sh
git clone https://github.com/marcosfsalazar/project-crawler.git
cd project-crawler
```

### 2. Build the project
```go
go build
```

### 2.1 Exclude Specific Diretories and Files (optional)

Create a crawlerIgnore.json file in the following format to exclude specified directories and files:
```json
{
  "dirs": ["dir_to_exclude1", "dir_to_exclude2"],
  "files": ["file_to_exclude1", "file_to_exclude2"]
}
```

### 3 Run the executable
<root-directory> is the directory that the program is gonna crawl
```sh
./projectCrawler <root-directory>
```

### Example

Let's suppose that you have the following directory structure:
```sh
/example
  /subdir1
    file1.txt
  /subdir2
    file2.ts
  file3.go
```

running the tool will generate a JSON output like this:

```json
{
  "sub_dirs": {
    "subdir1": {
      "sub_dirs": {},
      "files": {
        "file1.txt": {
          "content": "..."
        }
      }
    },
    "subdir2": {
      "sub_dirs": {},
      "files": {
        "file2.ts": {
          "content": "..."
        }
      }
    }
  },
  "files": {
    "file3.go": {
      "content": "..."
    }
  }
}
```

### Directory structure

The project is structured as follows:
```bash
./
|-- main.go
|-- go.mod
|-- crawler/
|   |-- crawler.go
|-- utils/
|   |-- file.go
|-- models/
    |-- directory.go
    |-- ignore.go

```
- **models**: Contains data structures for directory and file representation.
- **utils**: Utility functions for directory creation, file reading, and exclusion logic.
- **crawler**: Core logic for recursively reading directories and files.
- **main**.go: Entry point of the application.

### License

This project is licensed under the MIT License.