# Gatling Docker Image

Run with:

```
docker build -t mefellows/gatling .
docker run -v <conf>:/opt/gatling/conf \
           -v <results>:/opt/gatling/results \
           -v <user-files>:/opt/gatling/user-files -it mefellows/gatling
```