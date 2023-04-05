FROM golang:1.18-alpine AS builder
RUN apk update && apk add --no-cache make git
RUN apk add openssh-client
WORKDIR /go/src/github.com/forbole/bdjuno
COPY . ./

ARG SSH_PRIVATE_KEY
RUN mkdir /root/.ssh/
RUN echo "${SSH_PRIVATE_KEY}" > /root/.ssh/id_rsa
RUN chmod 600 /root/.ssh/id_rsa
RUN ssh-keyscan github.com >> /root/.ssh/known_hosts
RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/

######################################################
## Enabe the lines below if chain supports cosmwasm ##
## module to properly build docker image            ##
######################################################
#RUN apk update && apk add --no-cache ca-certificates build-base git
#ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.1.1/libwasmvm_muslc.aarch64.a /lib/libwasmvm_muslc.aarch64.a
#ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.1.1/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.x86_64.a
#RUN sha256sum /lib/libwasmvm_muslc.aarch64.a | grep 9ecb037336bd56076573dc18c26631a9d2099a7f2b40dc04b6cae31ffb4c8f9a
#RUN sha256sum /lib/libwasmvm_muslc.x86_64.a | grep 6e4de7ba9bad4ae9679c7f9ecf7e283dd0160e71567c6a7be6ae47c81ebe7f32
## Copy the library you want to the final location that will be found by the linker flag `-lwasmvm_muslc`
#RUN cp /lib/libwasmvm_muslc.$(uname -m).a /lib/libwasmvm_muslc.a

RUN go mod download
RUN make build
RUN rm /root/.ssh/id_rsa

##################################################
## Enabe line below if chain supports cosmwasm  ##
## module to properly build docker image        ##
##################################################
#RUN LINK_STATICALLY=true BUILD_TAGS="muslc" make build


FROM alpine:latest
##################################################
## Enabe line below if chain supports cosmwasm  ##
## module to properly build docker image        ##
##################################################
#RUN apk update && apk add --no-cache ca-certificates build-base
WORKDIR /bdjuno
COPY --from=builder /go/src/github.com/forbole/bdjuno/build/bdjuno /usr/bin/bdjuno
CMD [ "bdjuno"]