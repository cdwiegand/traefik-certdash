# Traefik Certifcate Dashboard Plugin

This plugin will display certificates registered in Traefik's certbot store.

## Configuration example:

```
http:
  middlewares:
    traefik-certdash:
      plugin:
        traefik-certdash:
          FIXME: FIXME
  routers:
    traefik-certdash:
      rule: "FIXME PathPrefix(`/certs`)"
      middlewares:
        - traefik-certdash
      service: api@internal
  
```
