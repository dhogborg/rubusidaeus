FROM alexellis2/raspistill

MAINTAINER David Högborg <d@hogborg.se>

ADD bin/rubusidaeus /usr/bin/

EXPOSE 8080

ENTRYPOINT ["rubusidaeus"]