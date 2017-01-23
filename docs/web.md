# Create a ClamAV scan micro-service :new: :construction:

```bash
$ docker run -d -p 3993:3993 malice/clamav web

INFO[0000] web service listening on port :3993
```

## Now you can perform scans like so

```bash
$ http -f localhost:3993/scan malware@/path/to/evil/malware
```

> **NOTE:** I am using **httpie** to POST to the malice micro-service

```bash
HTTP/1.1 200 OK
Content-Length: 124
Content-Type: application/json; charset=UTF-8
Date: Sat, 21 Jan 2017 05:39:29 GMT

{
  "clamav": {
    "infected": true,
    "result": "Eicar-Test-Signature",
    "engine": "0.99.2",
    "known": "5630857",
    "updated": "20170123",
    "error": ""
  }
}
```
