# Copyright (C) 2022, Vi Grey
# All rights reserved.
#
# Redistribution and use in source and binary forms, with or without
# modification, are permitted provided that the following conditions
# are met:
#
# 1. Redistributions of source code must retain the above copyright
#    notice, this list of conditions and the following disclaimer.
# 2. Redistributions in binary form must reproduce the above copyright
#    notice, this list of conditions and the following disclaimer in the
#    documentation and/or other materials provided with the distribution.
#
# THIS SOFTWARE IS PROVIDED BY AUTHOR AND CONTRIBUTORS ``AS IS'' AND
# ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
# IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
# ARE DISCLAIMED. IN NO EVENT SHALL AUTHOR OR CONTRIBUTORS BE LIABLE
# FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
# DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
# OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
# HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
# LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
# OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
# SUCH DAMAGE.

PKG_NAME := sattrak
CURRENTDIR := $(dir $(realpath $(firstword $(MAKEFILE_LIST))))
DOC_NAME := draft-sattrak-v1-00

all:
	mkdir -p $(CURRENTDIR)build/bin; \
	mkdir -p $(CURRENTDIR)build/docs; \
	mkdir -p $(CURRENTDIR)build/nes; \
	go build -ldflags="-s -w" -o $(CURRENTDIR)build/bin/$(PKG_NAME) $(CURRENTDIR)src/$(PKG_NAME)/; \
	xml2rfc $(CURRENTDIR)src/docs/$(DOC_NAME).xml -p $(CURRENTDIR)build/docs --text --html --pdf --no-external-js --no-external-css --v3 --id-is-work-in-progress --no-pagination; \
	cd $(CURRENTDIR)src/nes/; \
	asm6 $(PKG_NAME).asm ../../build/nes/$(PKG_NAME).nes; \
	cd $(CURRENTDIR); \

clean:
	rm -rf -- $(CURRENTDIR)build; \
