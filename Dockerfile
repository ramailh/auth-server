FROM scratch

COPY api/oauth2-test /api/oauth2-test

COPY .env .env

COPY template/index.html /template/index.html 

ENTRYPOINT ["/api/oauth2-test"]