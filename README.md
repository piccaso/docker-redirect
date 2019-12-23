## Simple 301/302 HTTP redirect
A small server (written in go) which only does redirects and runs in an read-only environment.


### Usage:
```yml
version: "2.2"
services:
  redirect:
    image: 0xff/redirect
    read_only: true
    ports:
      - "9000:9000"
    restart: always
    cap_drop:
      - ALL
    environment:
      TARGET: "https://my-new-site.com/"
      CODE: "301"

```