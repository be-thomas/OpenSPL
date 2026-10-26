// Package commands holds the command-name tables each SPL front-end uses to
// classify a stage start.
//
// Classification happens before parsing so a recognized-but-unimplemented
// command produces a precise diagnostic instead of a confusing parse failure
// that treats the command as a search term. The tables are per dialect: SPL1
// and the SPL2 subset know different command sets, and neither front-end
// imports the other.
package commands

// spl1Implemented are the classic SPL commands the SPL1 front-end parses.
var spl1Implemented = map[string]bool{
	"search":    true,
	"where":     true,
	"stats":     true,
	"timechart": true,
	"eval":      true,
	"rex":       true,
	"table":     true,
	"fields":    true,
	"sort":      true,
	"head":      true,
	"bin":       true,
}

// spl1Known are real classic SPL commands that SPL1 does not implement yet.
var spl1Known = map[string]bool{
	"append": true, "appendcols": true, "addtotals": true, "analyzefields": true,
	"anomalydetection": true, "apply": true, "associate": true, "bucket": true,
	"chart": true, "cluster": true, "collect": true, "contingency": true,
	"correlate": true, "datamodel": true, "dbinspect": true, "dedup": true,
	"delete": true, "diff": true, "entitlement": true, "eventcount": true,
	"eventstats": true, "extract": true, "fieldformat": true, "fieldsummary": true,
	"filldown": true, "fillnull": true, "findtypes": true, "fit": true,
	"folderize": true, "foreach": true, "format": true, "from": true,
	"geostats": true, "head": true, "highlight": true, "inputcsv": true,
	"inputlookup": true, "iplocation": true, "join": true, "kmeans": true,
	"kvform": true, "loadjob": true, "localize": true, "lookup": true,
	"makecontinuous": true, "makeresults": true, "map": true, "mcollect": true,
	"metadata": true, "metasearch": true, "ml": true, "mstats": true,
	"multikmeans": true, "multisearch": true, "mvcombine": true, "mvexpand": true,
	"nomv": true, "outlier": true, "outputcsv": true, "outputlookup": true,
	"outputtext": true, "pivot": true, "predict": true, "rangemap": true,
	"rare": true, "regex": true, "reltime": true, "rename": true, "replace": true,
	"rest": true, "return": true, "reverse": true, "rex": true, "script": true,
	"scrub": true, "search": true, "select": true, "selfjoin": true,
	"sendemail": true, "set": true, "setfields": true, "spath": true,
	"stats": true, "strcat": true, "streamstats": true, "table": true,
	"tags": true, "tail": true, "timechart": true, "top": true,
	"transaction": true, "transpose": true, "trendline": true, "tscollect": true,
	"tstats": true, "typeahead": true, "typer": true, "union": true,
	"uniq": true, "untable": true, "where": true, "x11": true, "xmlkv": true,
	"xmlunescape": true, "xpath": true,
}

// SPL1Implemented reports whether the SPL1 front-end implements a command.
func SPL1Implemented(name string) bool { return spl1Implemented[name] }

// SPL1Known reports whether name is a real classic SPL command.
func SPL1Known(name string) bool { return spl1Known[name] }

// spl2Implemented are the SPL2 commands in the supported subset. The pipeline
// itself always starts with `from`, which the front-end handles before it
// classifies a piped stage.
var spl2Implemented = map[string]bool{
	"from":  true,
	"where": true,
	"stats": true,
}

// spl2Known are SPL2 commands outside the supported subset. They are real
// commands, so they get `unsupported-command` instead of `unknown-command`.
var spl2Known = map[string]bool{
	"append": true, "dedup": true, "eval": true, "eventstats": true,
	"expand": true, "fields": true, "head": true, "join": true,
	"lookup": true, "multisearch": true, "rare": true, "rename": true,
	"rex": true, "search": true, "sort": true, "streamstats": true,
	"table": true, "timechart": true, "top": true,
}

// SPL2Implemented reports whether the SPL2 front-end implements a command.
func SPL2Implemented(name string) bool { return spl2Implemented[name] }

// SPL2Known reports whether name is a real SPL2 command.
func SPL2Known(name string) bool { return spl2Known[name] }
