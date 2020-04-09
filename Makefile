PHONY: all

all:
	 docker build -t haad/sigtrapper -f Dockerfile .
	 docker push haad/sigtrapper:latest