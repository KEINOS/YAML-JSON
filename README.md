# YAML-JSON Converter

Dockerfile to convert YAML to JSON and vice versa.

## Usage

```bash
yaml-json [-json/-yaml] [data]
```

### Sample Usage

```shellsession
$ # Convert YAML to JSON
$ cat ./sample.yaml | docker run --rm -i yaml-json -json
$ # Convert JSON to YAML
$ cat ./sample.json | docker run --rm -i yaml-json -yaml
```

## Build

```shellsession
$ ls
Dockerfile    README.md    src
$
$ tree
.
├── Dockerfile
├── README.md
└── src
    ├── sample.yml
    ├── sample.json
    └── yaml-json.go
$ 
$ docker build -t yaml-json .
```


