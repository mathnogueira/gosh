#!/bin/sh

getFileContent_results=()
getFileContent () {
	local error
	local content
	x=`cat $1 2>&1`
	if [[ "$?" == "1" ]]; then
		error="error: $x"
		content=""
	else
		content="$x"
		error=""
	fi

	getFileContent_results=($content $error)
}

main () {
	local filename=$1
	getFileContent "$filename"
	local content=${getFileContent_results[1]}
	local err=${getFileContent_results[2]}

	if [[ "$err" != "" ]]; then
		echo "failed: $err"
		exit 1
	else
		echo $content
		exit 0
	fi
}

main $1