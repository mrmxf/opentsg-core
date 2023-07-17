# credentials

extract credentials from websites such as gitlab.

## set up the configuration body to be used in the program

```go
auth.AuthInit([]string{key1, key2 key3})
```

the keys for gitlab, github and aws do not need to be declared and are assigned to the relevant domain automatically.

## get your file

Go get your the bytes of the file from the website you want with the following function.

```go
fileBytes, errHttp := mmReader.Decode("example.com/pathto/important/file.json")
```

## List of acceptable domain styles

### Gitlab

- https://gitlab.com/api/v4/projects/{project-id-number}/repository/files/path%2Fto%2Ffile.json?ref={branch-name}
- https://gitlab.com/{usernmae}/{repo}/{pathtofile}

### Github

- https://api.github.com/repos/{username}/{repo}/contents/path%2Fto%2Ffile.json
- https://github.com/{username}/{repo}/{filepath}/{filepath}/{file.txt}

### S3

- s3://{bucketname}/{filepath}/{filepath}/{file.txt}
- http://s3.amazonaws.com/{bucketname}/{filepath}/{filepath}/{file.txt}

### http

Any style is acceptable for http(s) but it is expected to follow this layout.

- https://example.com/path/to/name/of/file.json
