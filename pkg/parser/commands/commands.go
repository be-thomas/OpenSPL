// Package commands holds the command-name tables each SPL front-end uses to
// classify a stage start.
//
// Classification happens before parsing so a recognized-but-unimplemented
// command produces a precise diagnostic instead of a confusing parse failure
// that treats the command as a search term. The tables are per dialect: SPL1
// and the SPL2 subset know different command sets, and neither front-end
// imports the other.
package commands

import "sort"

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
	"rename":    true,
}

// spl1Known are real classic SPL commands that SPL1 does not implement yet.
// Kept in alphabetical order so the list can be checked against the search
// reference; a command missing from here is parsed as free text, which emits a
// query instead of a refusal.
var spl1Known = map[string]bool{
	"addcoltotals": true, "addtotals": true, "analyzefields": true, "anomalies": true,
	"anomalousvalue": true, "anomalydetection": true, "append": true, "appendcols": true,
	"apply": true, "associate": true, "audit": true, "bucket": true,
	"chart": true, "cluster": true, "collect": true, "concurrency": true,
	"contingency": true, "convert": true, "correlate": true, "datamodel": true,
	"dbinspect": true, "dedup": true, "delete": true, "delta": true,
	"diff": true, "entitlement": true, "erex": true, "eventcount": true,
	"eventstats": true, "extract": true, "fieldformat": true, "fieldsummary": true,
	"filldown": true, "fillnull": true, "findtypes": true, "fit": true,
	"folderize": true, "foreach": true, "format": true, "from": true,
	"gauge": true, "gentimes": true, "geoip": true, "geom": true,
	"geostats": true, "head": true, "highlight": true, "history": true,
	"inputcsv": true, "inputlookup": true, "iplocation": true, "join": true,
	"kmeans": true, "kvform": true, "loadjob": true, "localize": true,
	"localop": true, "lookup": true, "makecontinuous": true, "makemv": true,
	"makeresults": true, "map": true, "mcollect": true, "metadata": true,
	"metasearch": true, "meventcollect": true, "ml": true, "mpreview": true,
	"mstats": true, "multikmeans": true, "multisearch": true, "mvcombine": true,
	"mvexpand": true, "nomv": true, "noop": true, "outlier": true,
	"outputcsv": true, "outputlookup": true, "outputtext": true, "overlap": true,
	"pivot": true, "predict": true, "rangemap": true, "rare": true,
	"regex": true, "reltime": true, "replace": true,
	"rest": true, "return": true, "reverse": true, "rex": true,
	"savedsearch": true, "script": true, "scrub": true, "search": true,
	"searchtxn": true, "select": true, "selfjoin": true, "sendemail": true,
	"set": true, "setfields": true, "sitop": true, "spath": true,
	"stats": true, "strcat": true, "streamstats": true, "table": true,
	"tags": true, "tail": true, "timechart": true, "tojson": true,
	"top": true, "transaction": true, "transpose": true, "trendline": true,
	"tscollect": true, "tstats": true, "typeahead": true, "typer": true,
	"union": true, "uniq": true, "untable": true, "walklex": true,
	"where": true, "x11": true, "xmlkv": true, "xmlunescape": true,
	"xpath": true,
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
// Every command the SPL2 search reference documents is listed, so a command
// missing here is a gap in this table rather than a gap in the reference.
var spl2Known = map[string]bool{
	"addinfo": true, "append": true, "appendcols": true, "appendpipe": true,
	"bin": true, "branch": true, "convert": true, "decrypt": true,
	"dedup": true, "eval": true, "eventstats": true, "expand": true,
	"fields": true, "fieldsummary": true, "fillnull": true, "flatten": true,
	"head": true, "if": true, "into": true, "iplocation": true,
	"join": true, "loadjob": true, "lookup": true, "makemv": true,
	"makeresults": true, "mstats": true, "multisearch": true, "mvcombine": true,
	"mvexpand": true, "nomv": true, "ocsf": true, "rare": true,
	"rename": true, "replace": true, "reverse": true, "rex": true,
	"route": true, "search": true, "sort": true, "spath": true,
	"spl1": true, "streamstats": true, "table": true, "tags": true,
	"thru": true, "timechart": true, "timewrap": true, "top": true,
	"tstats": true, "typer": true, "union": true, "untable": true,
}

// SPL2Implemented reports whether the SPL2 front-end implements a command.
func SPL2Implemented(name string) bool { return spl2Implemented[name] }

// SPL2Known reports whether name is a real SPL2 command.
func SPL2Known(name string) bool { return spl2Known[name] }

// SPL1ImplementedNames returns the implemented SPL1 command names, sorted.
func SPL1ImplementedNames() []string { return names(spl1Implemented) }

// SPL2ImplementedNames returns the implemented SPL2 command names, sorted.
func SPL2ImplementedNames() []string { return names(spl2Implemented) }

func names(set map[string]bool) []string {
	out := make([]string, 0, len(set))
	for name := range set {
		out = append(out, name)
	}
	sort.Strings(out)
	return out
}
