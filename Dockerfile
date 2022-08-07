FROM gcr.io/distroless/static-debian11

COPY dist/testkube-flux2-receiver_linux_amd64_v1/testkube-flux2-receiver /

ENTRYPOINT ["/testkube-flux2-receiver"]
