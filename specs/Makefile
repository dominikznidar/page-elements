SOURCE_FILES = $(wildcard *.proto)
BUILT_FILES = $(subst .proto,.pb.go,$(SOURCE_FILES))

all: $(BUILT_FILES)

$(BUILT_FILES): $(SOURCE_FILES)
	docker run -it --rm -v $$PWD:/tmp trifs/protoc:3.1.0 -I=/tmp --go_out=plugins=grpc:/tmp $(addprefix /tmp/,$(SOURCE_FILES))
