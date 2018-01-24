FROM alpine
MAINTAINER eiffelqiu@qq.com

WORKDIR .
ENV PATH="./:${PATH}"
EXPOSE 80

COPY ./server ./
ENTRYPOINT "server"
