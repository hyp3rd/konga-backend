#!/bin/bash
set -Eeuo pipefail

traperr() {
  echo "ERROR: ${BASH_SOURCE[1]} at about line ${BASH_LINENO[0]} ${ERR}"
}

set -o errtrace
trap traperr ERR

report () {
	cat >&2 <<-'EOF'

I've successfully initialized all the modules.

	EOF
}

deps () {
	go list ./...
}

unique_repos () {
	cut -d '/' -f-3 | sort | uniq
}

no_konga () {
	grep -v '*konga*'
}

go_get_update () {
	while read d
	do
		echo $d
		export GO111MODULE=on
		go get -u $d/... || echo "failed, trying again with master" && cd $GOPATH/src/$d && git checkout master && go get -u -x $d
	done
}

ini_modules () {
    modules=('.' 'transport' 'middleware' 'transport/http' 'inmemory' 'cockroachdb' 'implementation' 'cmd/konga')

    for i in "${modules[@]}"; do
        cd $i ; rm -rf go.* ; go mod init ;  cd -  # ; go mod tidy ; GO111MODULE=on go build ; cd -
    done

	git add -A . ; git commit -m "modules update" || : ; git push || :

	report
}

# deps | unique_repos | go_get_update

ini_modules
