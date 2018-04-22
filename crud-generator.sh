#!/bin/sh
docker run --rm -w /data -v $(pwd):/data jerson/crud-generator $@

