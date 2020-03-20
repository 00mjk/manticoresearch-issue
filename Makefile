#---* Makefile *---#
.SILENT :

# export GO111MODULE=on

## manticore-start		:	start local manticore searchd with qorpress config.
manticore-start:
	@searchd --config ./config/manticore.conf
.PHONY: manticore-start

## manticore-stop			:	stop local manticore searchd.
manticore-stop:
	@searchd --stop --config ./config/manticore.conf
.PHONY: manticore-stop

## manticore-index			:	stop local manticore searchd.
manticore-index:
	@indexer --config ./config/manticore.conf --all
.PHONY: manticore-index

## help				:	Print commands help.
help : Makefile
	@sed -n 's/^##//p' $<
.PHONY: help

# https://stackoverflow.com/a/6273809/1826109
%:
	@:
