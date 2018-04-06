FROM ubuntu:latest
COPY bin/curl-anim /opt/curl-anim
COPY frames /opt/frames

CMD ["/opt/curl-anim", "--frames-path=/opt/frames", "--port=8081"]
