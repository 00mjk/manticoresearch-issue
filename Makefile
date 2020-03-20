#---* Makefile *---#
.SILENT :

export GO111MODULE=on

## manticore-start		:	start local manticore searchd with qorpress config.
manticore-start:
	@searchd --config ./shared/config/manticore.conf
.PHONY: manticore-start

## manticore-stop			:	stop local manticore searchd.
manticore-stop:
	@searchd --stop --config ./shared/config/manticore.conf
.PHONY: manticore-stop

## manticore-index			:	stop local manticore searchd.
manticore-index:
	@indexer --config ./shared/config/manticore.conf --all
.PHONY: manticore-index

## help				:	Print commands help.
help : Makefile
	@sed -n 's/^##//p' $<
.PHONY: help

# https://stackoverflow.com/a/6273809/1826109
%:
	@:
