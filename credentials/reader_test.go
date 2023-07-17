package credentials

//use the convey method of testing
import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/mrmxf/opentsg-core/credentials/auth"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/afero"
)

func TestDecodeHTTP(t *testing.T) {

	//test getting http images with known urls
	addresses := []string{"https://mrmxf.com/r/project/msg-tpg/ramp-2022-02-28/multiramp-12b-pc-4k-hswp.png",
		"https://mrmxf.io/smpte-ra-demo/register/lmt/schema"}
	expec := []string{"9ecf3411b3ad6b252a3bcf45a45291def84dbcbaa489be5a61845d27b3f4c484",
		"f0537bd22e23f61dac9c1abcd8cdfa8dcf1d7b907bedb62fe55ae8b653ec525c"}

	//open the url and check it is good

	for i, ad := range addresses {
		htest := sha256.New()
		ftest, err := Decode(ad)

		htest.Write(ftest)
		//generate a sha of the file
		//fmt.Println(string(ftest))
		Convey("Checking that json and image files are extracted with http", t, func() {
			Convey(fmt.Sprintf("using a website of %v", ad), func() {
				Convey("A matching hash of a file is returned", func() {
					So(err, ShouldBeNil)
					So(fmt.Sprintf("%x", htest.Sum(nil)), ShouldResemble, expec[i]) //[]byte(expec[i]))
				})
			})
		})
	}
}

func TestDecodeGitHub(t *testing.T) {
	tokenB, err := os.ReadFile("./testdata/ghkey.txt")
	token := string(tokenB)
	//test getting http images with known urls
	if err == nil || token != "" {
		addresses := []string{"https://api.github.com/repos/mmTristan/public/contents/nested%2Fnest.json",
			"https://api.github.com/repos/mmTristan/ascmhl/contents/schema%2Fascmhl.xsd",
			"https://github.com/mmTristan/ascmhl/schema/ascmhl.xsd"}
		tokens := []string{"", token, token}
		expec := []string{"bb9dd8180d70abce882ccdb69aab2bffa1a96c8f86cdbcd631948b3085465ab9",
			"0235e307f3930b9f8142c37f23f6e4d55fedb534dda19227069c2c95948f2bcb",
			"0235e307f3930b9f8142c37f23f6e4d55fedb534dda19227069c2c95948f2bcb"}

		//open the url and check it is good

		for i, ad := range addresses {
			htest := sha256.New()
			auth.AuthInit("", []string{tokens[i]})
			ftest, err := Decode(ad)
			//fmt.Println(string(ftest))
			htest.Write(ftest)
			//generate a sha of the file
			//fmt.Println(string(ftest))
			Convey("Checking that json and image files are extracted with github", t, func() {
				Convey(fmt.Sprintf("using a website of %v", ad), func() {
					Convey("A matching hash of a file is returned", func() {
						So(err, ShouldBeNil)
						So(fmt.Sprintf("%x", htest.Sum(nil)), ShouldResemble, expec[i]) //[]byte(expec[i]))
					})
				})
			})
		}
	} else {
		fmt.Printf("github tests skipped due to the following errors %v opening the token file\n", err)
	}
}

func TestDecodeGit(t *testing.T) {
	//access the token so it is not saved in the test suite
	tokenB, err := os.ReadFile("./testdata/apikey.txt")
	token := string(tokenB)
	if err == nil || token != "" {
		//test getting http images with known urls
		addresses := []string{"https://gitlab.com/api/v4/projects/35946043/repository/files/go.mod?ref=main",
			"https://gitlab.com/api/v4/projects/33185381/repository/files/test%2Ftestapi.json?ref=main",
			"https://gitlab.com/mmTristan/publicgo/go.mod",
			"https://gitlab.com/mmTristan/basic-12bit-clour/test/testapi.json"}
		tokens := []string{"", token, "", token}
		expec := []string{"07f2ba17f973dda06607381921ab8811c177b0d2910e2d252e02cc0946b4cc7d",
			"87bad426555c30caa55b4332304f55e1ae79d331297ba75ab244539018798fbd",
			"07f2ba17f973dda06607381921ab8811c177b0d2910e2d252e02cc0946b4cc7d",
			"87bad426555c30caa55b4332304f55e1ae79d331297ba75ab244539018798fbd"}

		//open the url and check it is good

		for i, ad := range addresses {
			htest := sha256.New()
			auth.AuthInit("", []string{tokens[i]})
			ftest, err := Decode(ad)

			htest.Write(ftest)

			Convey("Checking that json and image files are extracted with git", t, func() {
				Convey(fmt.Sprintf("using a website the gitlab api of %v", ad), func() {
					Convey("A matching hash of the extracted file is returned", func() {
						So(err, ShouldBeNil)
						So(fmt.Sprintf("%x", htest.Sum(nil)), ShouldResemble, expec[i])
					})
				})
			})
		}
	} else {
		fmt.Printf("gitlab tests skipped due to the following errors %v opening the token file\n", err)
	}
}

func TestDecodeS3(t *testing.T) {
	//extract the access details for the test
	region := "eu-west-2"
	secretB, errS := os.ReadFile("./testdata/secret.txt")
	accessB, errA := os.ReadFile("./testdata/access.txt")
	secret := string(secretB)
	access := string(accessB)
	//run the test if there are no errors getting the details
	if errS == nil && errA == nil {
		//access the token so it is not saved in the test suite
		addresses := []string{"s3://mmh-cache/bot-tlh/staging/publish/multiramp-12b-pc-4k-zp.dpx", "http://s3.amazonaws.com/mmh-cache/bot-tlh/dev/schema/addimageschema.json"}
		expec := []string{"7490fd92c6292b6850bd6fd568abbe8bebae1a4ef3382c284d252921d94a6d4d",
			"679505dad56fb9aad50089ebb9e5b217912c41a9b585217f467633749bcbddac"}

		for i, ad := range addresses {
			htest := sha256.New()
			auth.AuthInit("", []string{region, access, secret})
			ftest, err := Decode(ad)
			htest.Write(ftest)
			Convey("Checking that json and image files are extracted with s3", t, func() {
				Convey(fmt.Sprintf("using an s3 link of %v", ad), func() {
					Convey("A matching hash of the extracted file is returned", func() {
						So(err, ShouldBeNil)
						So(fmt.Sprintf("%x", htest.Sum(nil)), ShouldResemble, expec[i])
					})
				})
			})
		}
		address := "s3://mmh-cache/bot-tlh/staging/publish/multiramp-12b-pc-4k-zp.dpx"
		expecHash := "7490fd92c6292b6850bd6fd568abbe8bebae1a4ef3382c284d252921d94a6d4d"

		passed := func() (string, string, string) { return region, secret, access }
		envGet := func() (string, string, string) {
			os.Setenv("AWS_ACCESS_KEY_ID", access)
			os.Setenv("AWS_DEFAULT_REGION", region)
			os.Setenv("AWS_SECRET_ACCESS_KEY", secret)
			return "", "", ""
		}
		cred := func() (string, string, string) {
			//generate all of the files necessary
			os.MkdirAll(config.DefaultSharedCredentialsFilename()+".json", 0755)
			os.Remove(config.DefaultSharedCredentialsFilename() + ".json") //remove the uneeded json
			f, err := os.OpenFile(config.DefaultSharedCredentialsFilename(), os.O_RDWR|os.O_CREATE, 0755)
			fmt.Println(err)
			defer f.Close()
			file := "[default]\n" + "aws_access_key_id=" + access + "\n" +
				"aws_secret_access_key=" + secret
			_, err = f.Write([]byte(file))
			fmt.Println(err, "DEFAULT")

			os.Setenv("AWS_ACCESS_KEY_ID", "")
			os.Setenv("AWS_DEFAULT_REGION", "")
			os.Setenv("AWS_SECRET_ACCESS_KEY", "")
			return "", "", ""
		}

		snapShot := afero.NewOsFs()
		snapKey := os.Getenv("AWS_ACCESS_KEY_ID")
		snapReg := os.Getenv("AWS_DEFAULT_REGION")
		snapSecret := os.Getenv("AWS_SECRET_ACCESS_KEY")
		//after the test has run, reset all values back to the snap shot
		accessorMethods := []func() (string, string, string){passed, envGet, cred}
		for _, ad := range accessorMethods {
			htest := sha256.New()
			passed1, passed2, passed3 := ad()
			auth.AuthInit("", []string{passed1, passed2, passed3})
			ftest, err := Decode(address)
			htest.Write(ftest)
			Convey("Checking that different authentication methods for s3 work", t, func() {
				Convey(fmt.Sprintf("using an s3 link of %v", address), func() {
					Convey("A matching hash of the extracted file is returned", func() {
						So(err, ShouldBeNil)
						So(fmt.Sprintf("%x", htest.Sum(nil)), ShouldResemble, expecHash)
					})
				})
			})
		}
		os.Setenv("AWS_ACCESS_KEY_ID", snapKey)
		os.Setenv("AWS_DEFAULT_REGION", snapReg)
		os.Setenv("AWS_SECRET_ACCESS_KEY", snapSecret)
		f, _ := os.Create(config.DefaultSharedCredentialsFilename())
		defer f.Close()
		oldBody, _ := snapShot.Open(config.DefaultSharedCredentialsFilename())
		old, _ := io.ReadAll(oldBody)
		f.Write(old)
		//f.Write(afero.O

		//test authorisation errors make a array of functions that return strings

	} else {
		fmt.Printf("s3 tests skipped due to the following errors %v, %v opening the token files\n", errS, errA)
	}
}

func TestErrors(t *testing.T) {

	//test getting http images with known urls
	baddresses := []string{"not even a website", "https://a.really.fake.website/not/real",
		"https://gitlab.com/api/v4/projects/33185381/repository/files/test%2Ftestapi.json?ref=main",
		"https://mrmxf.com/supersecret", "https://mrmxf.com/user"}
	want := []string{`Get "not%20even%20a%20website": unsupported protocol scheme ""`,
		`Get "https://a.really.fake.website/not/real": dial tcp: lookup a.really.fake.website on 1.1.1.1:53: no such host`,
		"404 Not Found", "404 Not Found", "401 Unauthorized"}

	//open the url and check it is good

	for i, ad := range baddresses {
		//open the files to be tested

		_, err := Decode(ad)
		Convey("Checking that errors are returned for incorrect websites", t, func() {
			Convey(fmt.Sprintf("using a website of %v", ad), func() {
				Convey(fmt.Sprintf("An error of %v is returned", want[i]), func() {
					So(err.Error(), ShouldEqual, want[i])
				})
			})
		})
	}
}
