FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

ENV OPERATOR=/usr/local/bin/must-gather-operator \
    USER_UID=1001 \
    USER_NAME=must-gather-operator

# install operator binary
COPY build/_output/bin/must-gather-operator ${OPERATOR}

RUN microdnf install tar gzip && microdnf clean all

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
