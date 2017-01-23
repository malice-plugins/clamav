# To update the AV run the following:

```bash
$ docker run --name=clamav malice/clamav update
```

## Then to use the updated clamav container:

```bash
$ docker commit clamav malice/clamav:updated
$ docker rm clamav # clean up updated container
$ docker run --rm malice/clamav:updated EICAR
```
