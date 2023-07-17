// package credentials returns the bytes from a http location
package credentials

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/mrmxf/opentsg-core/credentials/auth"
)

var regGitAPI = regexp.MustCompile("https://gitlab\\.com/api/v4/")
var regGitL = regexp.MustCompilePOSIX("https://gitlab\\.com/")

var regGitHbAPI = regexp.MustCompile("https://api\\.github\\.com/")
var regGitH = regexp.MustCompilePOSIX("https://github\\.com/")

var regS3 = regexp.MustCompile("^s3://[\\w\\-\\.]{3,63}/")
var regS3AWS = regexp.MustCompile("^http://s3\\.amazonaws\\.com/[\\w\\-\\.]{3,63}/")

// Decode returns the body of a url and an error if the information could not be extracted.
func Decode(url string) ([]byte, error) {

	//insert a credentials manager
	tokenGen := auth.AuthGet()
	switch {
	case regGitL.MatchString(url), regGitAPI.MatchString(url):
		return gitDecode(url, tokenGen["git_auth"])
	case regGitH.MatchString(url), regGitHbAPI.MatchString(url):
		return gitHubDecode(url, tokenGen["github_auth"])
		//develop functions for each regex string
	case regS3.MatchString(url), regS3AWS.MatchString(url):
		return s3Decode(url, tokenGen["s3_profile"])
		//develop functions for each regex string
	default: //make this for any other http decode
		return httpDecode(url)
	}

}

func httpDecode(url string) ([]byte, error) {
	//look at implementing these
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	} else if resperr := repsonseHelper(resp); resperr != nil {
		return nil, resperr
	}
	return io.ReadAll(resp.Body)
}

type access struct {
	Content string `json:"content,omitempty"`
	Name    string `json:"name,omitempty"`
}

type jsonID struct {
	Id      int    `json:"id,omitempty"`
	Default string `json:"default_branch,omitempty"`
}

func gitDecode(url string, a auth.Auth) ([]byte, error) {

	//get the body of the gitlab api
	token := "Bearer " + a.Token.(string)

	//convert gitlab links to gitlab api calls
	if regGitL.MatchString(url) && !regGitAPI.MatchString(url) {
		//get the owner and the repo
		owner, wantrepo := bucketToString(url, 19)
		//split the repo into repo and file
		repo, file := bucketToString(wantrepo, 0)

		//extract the api json
		idGetURL := "https://gitlab.com/api/v4/projects/" + owner + "%2f" + repo
		idGetJson, err := getRequest(idGetURL, token)
		if err != nil {
			return nil, err
		}
		var i jsonID
		json.Unmarshal(idGetJson, &i)
		if i.Id == 0 {
			return nil, fmt.Errorf("Error no valid id found for %v", url)
		}
		newfile := strings.ReplaceAll(file, "/", "%2F")
		url = "https://gitlab.com/api/v4/projects/" + fmt.Sprintf("%v", i.Id) + "/repository/files/" + newfile + "?ref=" + i.Default
	}

	body, err := getRequest(url, token)
	if err != nil {
		return nil, err
	}

	//decode from a json with base 64 to the actual contents of the file
	var c access
	json.Unmarshal(body, &c)

	dst := make([]byte, base64.StdEncoding.DecodedLen(len(c.Content)))
	_, err = base64.StdEncoding.Decode(dst, []byte(c.Content))

	return dst, err
}

func getRequest(url, token string) ([]byte, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	// set a token if one is provided
	if token != "" {
		//req.Header.Set("PRIVATE-TOKEN", token)
		req.Header.Set("Authorization", token)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	} else if resperr := repsonseHelper(resp); resperr != nil {
		return nil, resperr
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func gitHubDecode(url string, a auth.Auth) ([]byte, error) {
	var token string
	if a.Token.(string) == "" {
		token = a.Token.(string)
	} else {
		token = "token " + a.Token.(string)
	}

	//if not api then we'll ammend it to that
	//https://github.com/mmTristan/ascmhl/schema/ascmhl.xsd
	//https://api.github.com/repos/mmTristan/ascmhl/contents/schema%2Fascmhl.xsd
	//convert gitlab links to gitlab api calls
	if regGitH.MatchString(url) && !regGitHbAPI.MatchString(url) {
		//get the owner and the repo
		owner, wantrepo := bucketToString(url, 19)
		//split the repo into repo and file
		repo, file := bucketToString(wantrepo, 0)

		newfile := strings.ReplaceAll(file, "/", "%2F")
		url = "https://api.github.com/repos/" + owner + "/" + repo + "/contents/" + newfile
	}

	body, err := getRequest(url, token)
	if err != nil {
		return nil, err
	}

	var c access
	json.Unmarshal(body, &c)

	dst := make([]byte, base64.StdEncoding.DecodedLen(len(c.Content)))
	_, err = base64.StdEncoding.Decode(dst, []byte(c.Content))

	return dst, err
}

func s3Decode(url string, a auth.Auth) ([]byte, error) {
	opt := (a.Token).(*auth.S3AuthDetail)

	//https://s3.console.aws.amazon.com/s3/object/mmh-cache?region=eu-west-2&prefix=bot-tlh/dev/schema/addimageschema.json
	//http://s3.amazonaws.com/[bucket_name]/object/mhl.jdon
	//split the string to the bucket and file for use with s3 sdk
	var bucket, file string
	if regS3.MatchString(url) {
		bucket, file = bucketToString(url, 5)
	} else {
		bucket, file = bucketToString(url, 24)
	}
	region := opt.Region
	if region == "" {
		region = "eu-west-2"
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: opt.Creds,
	})
	if err != nil {
		return nil, err
	}

	downloader := s3manager.NewDownloader(sess)

	//Download the item from the bucket and check for errors before returning
	buf := aws.NewWriteAtBuffer([]byte{})
	_, err = downloader.Download(buf,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(file),
		})
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// bucketToString splits the s3 url into a bucket and file for use with the aws sdk
// split up the other string type as well
func bucketToString(url string, start int) (string, string) {
	var bucket int
	if start > len(url) {
		//if the start is out of range then nip it in the bud
		return url, ""
	}
	for i, let := range url[start:] {
		//search for the end of the bucket name
		if let == rune('/') {
			bucket = i + start
			break
		}
	}
	//there's some error checking to be done here
	return url[start:bucket], url[bucket+1:]
}

func repsonseHelper(resp *http.Response) error {
	stat := resp.Status
	valid := regexp.MustCompile("OK")
	if !valid.MatchString(stat) {
		return fmt.Errorf(stat)
	} else {
		return nil
	}
}
