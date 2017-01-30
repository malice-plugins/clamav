FROM malice/alpine

LABEL maintainer "https://github.com/blacktop"

COPY . /go/src/github.com/maliceio/malice-clamav
RUN apk --update add --no-cache clamav clamav-libunrar ca-certificates
RUN apk --update add --no-cache -t .build-deps \
                    build-base \
                    mercurial \
                    musl-dev \
                    openssl \
                    bash \
                    wget \
                    git \
                    gcc \
                    go \
  && echo "Building avscan Go binary..." \
  && cd /go/src/github.com/maliceio/malice-clamav \
  && export GOPATH=/go \
  && go version \
  && go get \
  && go build -ldflags "-X main.Version=$(cat VERSION) -X main.BuildTime=$(date -u +%Y%m%d)" -o /bin/avscan \
  && rm -rf /go /usr/local/go /usr/lib/go /tmp/* \
  && apk del --purge .build-deps

# Update ClamAV Definitions
RUN freshclam

# Add EICAR Test Virus File to malware folder
ADD http://www.eicar.org/download/eicar.com.txt /malware/EICAR

RUN chown malice -R /malware

WORKDIR /malware

ENTRYPOINT ["su-exec","malice","/sbin/tini","--","avscan"]
CMD ["--help"]
