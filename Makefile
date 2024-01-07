#!/usr/bin/make --no-print-directory --jobs=1 --environment-overrides -f

VERSION_TAGS += CHDIRS
CHDIRS_MK_SUMMARY := go-corelibs/chdirs
CHDIRS_MK_VERSION := v1.0.2

include CoreLibs.mk
