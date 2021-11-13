FROM alpine:latest

ENTRYPOINT ["/bin/pwdgen"]

ADD rel/pwdgen_linux-amd64 /bin/pwdgen
