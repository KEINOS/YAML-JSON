[![](https://images.microbadger.com/badges/image/keinos/alpine.svg)](https://microbadger.com/images/keinos/alpine "View image info on microbadger.com") [![](https://img.shields.io/docker/cloud/automated/keinos/yaml-json.svg)](https://hub.docker.com/r/keinos/yaml-json/builds "View builds on Docker Hub") [![](https://img.shields.io/docker/cloud/build/keinos/yaml-json.svg)](https://hub.docker.com/r/keinos/yaml-json "View on Docker Hub")

# YAML-JSON Converter

Dockerfile to convert YAML to JSON and vice versa.

```bash
docker pull keinos/yaml-json
```

- Golang: `golang:alpine` @ Docker Hub
- Base Image: `busybox` @ Docker Hub
- Repositories:
  - Image: https://hub.docker.com/r/keinos/yaml-json @ Docker Hub
  - Source: https://github.com/KEINOS/YAML-JSON @ GitHub
- Issues: https://github.com/KEINOS/YAML-JSON/issues @ GitHub

## Usage

```bash
yaml-json [-json | -yaml] [data]
```

- `-json`: Converts YAML to JSON
- `-yaml`: Converts JSON to YAML
- `data`: String data of YAML or JSON

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
$ git clone https://github.com/KEINOS/YAML-JSON.git YAML-JSON
$ cd $_
$ docker build -t yaml-json .
$ cd tests
$ cat ./sample.json | docker run --rm -i yaml-json -yaml
```
