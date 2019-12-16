# ci_plugin_travis_go

## The CLI needs to accept the following args

## Arguments
1.  -p or --project is the project name in the format of orgname/project
1.  -c or --commit is the commit reference
1.  -t or --tag is the tag name

## Status executable and response format
1. The output is **tab delimited**
1. The **first line** is a **header**
1. The **second line** is data 
1. The **status** should be success, failure, or running
1. The **build_url** should be the url where the status was found

```
./ci_plugin_travis_go status -p "linkerd/linkerd2" -c f27d7b65
status  build_url
success https://travis-ci.org/crosscloudci/testproj/builds/572521581 
```

