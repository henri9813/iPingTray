#!/bin/bash
set -x

for icon in icons/*.ico; do
	icon_name=$(basename "${icon%.*}")
	capitalized_name=$(echo "${icon_name:0:1}" | tr '[:lower:]' '[:upper:]')${icon_name:1}
	"${GOPATH}/bin/2goarray" "${capitalized_name}" icons <"${icon}" >"icons/${icon_name}.go"
done
