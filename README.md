![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/keinos/yaml-json.svg)

# YAML-JSON Converter

Dockerfile to convert YAML to JSON and vice versa.

## Usage

```bash
yaml-json [-json | -yaml] [data]
```

### Sample Usage

```shellsession
$ # Pull image
$ docker pull keinos/yaml-json
$ # Convert YAML to JSON
$ cat ./sample.yaml | docker run --rm -i keinos/yaml-json -json
$ # Convert JSON to YAML
$ cat ./sample.json | docker run --rm -i keinos/yaml-json -yaml
```

## Build

```shellsession
$ mkdir yaml-json && cd $_
$ git clone https://github.com/KEINOS/YAML-JSON.git
$ docker build -t yaml-json .
$ cd tests
$ cat ./sample.json | docker run --rm -i yaml-json -yaml
```
