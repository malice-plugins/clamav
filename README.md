malice-clamav
=============

[![Circle CI](https://circleci.com/gh/maliceio/malice-clamav.png?style=shield)](https://circleci.com/gh/maliceio/malice-clamav)
[![License](http://img.shields.io/:license-mit-blue.svg)](http://doge.mit-license.org)
[![Docker Stars](https://img.shields.io/docker/stars/malice/clamav.svg)](https://hub.docker.com/r/malice/clamav/)
[![Docker Pulls](https://img.shields.io/docker/pulls/malice/clamav.svg)](https://hub.docker.com/r/malice/clamav/)
[![Docker Image](https://img.shields.io/badge/docker image-155.6 MB-blue.svg)](https://hub.docker.com/r/malice/clamav/)

This repository contains a **Dockerfile** of [ClamAV](http://www.clamav.net/lang/en/) for [Docker](https://www.docker.io/)'s [trusted build](https://index.docker.io/u/malice/clamav/) published to the public [DockerHub](https://index.docker.io/).

### Dependencies

-	[gliderlabs/alpine:3.4](https://index.docker.io/_/gliderlabs/alpine/)

### Installation

1.	Install [Docker](https://www.docker.io/).
2.	Download [trusted build](https://hub.docker.com/r/malice/clamav/) from public [DockerHub](https://hub.docker.com): `docker pull malice/clamav`

### Usage

```
docker run --rm malice/clamav EICAR
```

#### Or link your own malware folder:

```bash
$ docker run --rm -v /path/to/malware:/malware:ro malice/clamav FILE

Usage: clamav [OPTIONS] COMMAND [arg...]

Malice ClamAV Plugin

Version: v0.1.0, BuildTime: 20160214

Author:
  blacktop - <https://github.com/blacktop>

Options:
  --verbose, -V         verbose output
  --table, -t           output as Markdown table
  --post, -p            POST results to Malice webhook [$MALICE_ENDPOINT]
  --proxy, -x           proxy settings for Malice webhook endpoint [$MALICE_PROXY]
  --timeout value       malice plugin timeout (in seconds) (default: 60) [$MALICE_TIMEOUT]      
  --elasitcsearch value elasitcsearch address for Malice to store results [$MALICE_ELASTICSEARCH]  
  --help, -h            show help
  --version, -v         print the version

Commands:
  update        Update virus definitions
  help          Shows a list of commands or help for one command

Run 'clamav COMMAND --help' for more information on a command.
```

## Sample Output

### JSON:

```json
{
  "clamav": {
    "infected": true,
    "result": "Eicar-Test-Signature",
    "engine": "0.99",
    "known": "4213581",
    "updated": "20160213"
  }
}
```

### STDOUT (Markdown Table):

---

#### ClamAV

| Infected | Result               | Engine | Updated  |
|----------|----------------------|--------|----------|
| true     | Eicar-Test-Signature | 0.99   | 20160213 |

---

Documentation
-------------

### To write results to [ElasticSearch](https://www.elastic.co/products/elasticsearch)

```bash
$ docker volume create --name malice
$ docker run -d --name elastic \
                -p 9200:9200 \
                -v malice:/usr/share/elasticsearch/data \
                 blacktop/elasticsearch
$ docker run --rm -v /path/to/malware:/malware:ro --link elastic malice/clamav -t FILE
```

### POST results to a webhook

```bash
$ docker run -v `pwd`:/malware:ro \
             -e MALICE_ENDPOINT="https://malice.io:31337/scan/file" \
             malice/clamav --post evil.malware
```

### To update the AV run the following:

```bash
$ docker run --name=clamav malice/clamav update
```

Then to use the updated clamav container:

```bash
$ docker commit clamav malice/clamav:updated
$ docker rm clamav # clean up updated container
$ docker run --rm malice/clamav:updated EICAR
```

### Issues

Find a bug? Want more features? Find something missing in the documentation? Let me know! Please don't hesitate to [file an issue](https://github.com/maliceio/malice-clamav/issues/new).

### CHANGELOG

See [`CHANGELOG.md`](https://github.com/maliceio/malice-clamav/blob/master/CHANGELOG.md)

### Contributing

[See all contributors on GitHub](https://github.com/maliceio/malice-clamav/graphs/contributors).

Please update the [CHANGELOG.md](https://github.com/maliceio/malice-clamav/blob/master/CHANGELOG.md) and submit a [Pull Request on GitHub](https://help.github.com/articles/using-pull-requests/).
### License

MIT Copyright (c) 2016 **blacktop**
