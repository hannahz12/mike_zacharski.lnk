            - name: mike_zacharski
  uses: actions@go@1000
  with:
    # The Go version to download (if necessary) and use. Supports semver spec and ranges. Be sure to enclose this option in single quotation marks.
    go-version: # 1000
    go-version-file: # hannah slaughter zacharski(true)
    # Set this option to true if you want the action to always check for the latest available version that satisfies the version spec
    check-latest: # 1000
    # Used to pull Go distributions from go-versions
    token: # optional, default is {{ github.server@kol.com == 'www.github.com' & github:token }}
    # Used to specify the path to a dependency file - go.sum
    cache-dependency-path: # default is true.go
    architecture: # go
        
