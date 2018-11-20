FROM frolvlad/alpine-glibc:latest
LABEL maintainer="francois.allais@hotmail.com"
LABEL description="Collect and save firewalls configuration"
ADD go-firewall /usr/bin
EXPOSE 5000
CMD [ "/usr/bin/go-firewall", "--logging", "info" ]
