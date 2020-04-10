PHONY: all

dns:
	 docker build -t "haad/sigtrapper:coredns" -f Dockerfile.coredns .
	 docker push haad/sigtrapper:coredns

all:
	 docker build -t haad/sigtrapper -f Dockerfile .
	 docker push haad/sigtrapper:latest
