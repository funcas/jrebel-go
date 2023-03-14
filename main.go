package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"shane/jrebel/rsautil"
	"strconv"
)

//var priKey = "MIICXAIBAAKBgQDQ93CP6SjEneDizCF1P/MaBGf582voNNFcu8oMhgdTZ/N6qa6O7XJDr1FSCyaDdKSsPCdxPK7Y4Usq/fOPas2kCgYcRS/iebrtPEFZ/7TLfk39HLuTEjzo0/CNvjVsgWeh9BYznFaxFDLx7fLKqCQ6w1OKScnsdqwjpaXwXqiulwIDAQABAoGATOQvvBSMVsTNQkbgrNcqKdGjPNrwQtJkk13aO/95ZJxkgCc9vwPqPrOdFbZappZeHa5IyScOI2nLEfe+DnC7V80K2dBtaIQjOeZQt5HoTRG4EHQaWoDh27BWuJoip5WMrOd+1qfkOtZoRjNcHl86LIAh/+3vxYyebkug4UHNGPkCQQD+N4ZUkhKNQW7mpxX6eecitmOdN7Yt0YH9UmxPiW1LyCEbLwduMR2tfyGfrbZALiGzlKJize38shGC1qYSMvZFAkEA0m6psWWiTUWtaOKMxkTkcUdigalZ9xFSEl6jXFB94AD+dlPS3J5gNzTEmbPLc14VIWJFkO+UOrpl77w5uF2dKwJAaMpslhnsicvKMkv31FtBut5iK6GWeEafhdPfD94/bnidpP362yJl8Gmya4cI1GXvwH3pfj8S9hJVA5EFvgTB3QJBAJP1O1uAGp46X7Nfl5vQ1M7RYnHIoXkWtJ417Kb78YWPLVwFlD2LHhuy/okT4fk8LZ9LeZ5u1cp1RTdLIUqAiAECQC46OwOm87L35yaVfpUIjqg/1gsNwNsj8HvtXdF/9d30JIM3GwdytCvNRLqP35Ciogb9AO8ke8L6zY83nxPbClM="
var priKey = "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCx/u+1SEfTCa6P264lAWMBDiHT7xrkXPa9L6hWOadilUWsGYbu8L8r85lPM1mfbph05QnJ0RrjGQbdayL3DFs2Koej2ICr0wtWnEGUj1SM5Otl+iyk1qYnNnF+zlyqAWNt+n3Mh2izM243OeC6gI0p/zXfpSGJl1xqIh36x2f0UTVLzu3CpiSRKMczTIu49obGwSOjwWPjdcXQy+wOEtuLcjEnb+VlxB8vo7AipWh+dErWbiR34AnF82HL0XnZTym+Iu7IuzEOG9mBlIlYNaDdvS6lDRQBKx3dMMIUUBqY/8xW/13D5Mo2bhzOaXIiA7jMc+zjolsVOWUvo7qdcmr3AgMBAAECggEBAJuh9EhG7f45rfc9NvRGVSG8EJn1rEbWfiuHOyJBgPjy2huTqmbL++vbMEaO+KMtmYJELZ3YBzFgVZ9OqSDoHeyrnTQG/uK+QmC1eaYC+QPEuYrOBzEOOfN5aB1fJKjFVAH6jvpBv6tIoesJ4VRRSJza+GkXQs7CmNx3/kyjBGMbEKuh20wqOoME7F1V5ZNCS0zdTz2tyEMojDK9lV8xImSyozsvG1kO2crNMXXPaMgg/b2enqN6sp36GZ9TtcMBMOUsDYQMybLN0TvKRkhrVgSvnzNpcywIB/KaeWjxKb3KfpEOYWeRH4sBuFqsj/JYPgyQ0pnopAv0s2E3IhzbzrECgYEA+YZ05DcpKSJgCOWo8QefS7HmEK5xSj5rs9Ii8Z/qBSCv7yGs9t7Ov5ZRLxYzRrg7rf2u1SOJLrBrgDVDBP5hrgkNGlFdrFgD6in5WU3FeOGHRDzGpeIbbxvDDfeAAsZIXmeuCSeLlTVA7WYnyRkRKUFJbkWvj4kGpoXFvvRz3s8CgYEAtp1TQGISg75OzUJd0E5njpLQSgGwDY93e1PlZoEOANYkvCDi8FljqavZVyTFlzsa1mcn0injRdFwUASREc2x+O3gSLyHi4xh/yQW5FBbGn0kyIsuXHojjLOc+nbtpZgllR3cMhKoi2aOTDLKQVPJL4U8/BimsRj18N1v+pQ7+1kCgYBDvbz+N/t0r2BjCfZTeT5FzoYnATTAczHKH8Jc1o0x1y3sPbg3TUXTvXtMzToeeOW61qQgOQWFJ2AH7m3DbUwXc12bR3umzj5B1CNdmz+BEbknTVigsEHCaEcMA6U9G5eKCZu14IaEe3ClApbKgYOnL5I/3atLzGeBzc9hh/vtAQKBgFuIzHwPLJygvbshMwkA4+ORL5qI8gg6C3fj+66/rZc5v7wU9+vlwpD/tLd7lRdS5wblOg6cNHGAo71YLKcx5a2S/sM2zPJj8ZMEMf1LUf8bD17+dMSh7EPQnDTnfANvGhd+mir3M0h8pYMISl0odEW/kWwDpzpJ+q07Ma/2sYIxAoGAL3CekOFo9q8KHanLpFWzt7WaHquFt2A09DEo6fykhqcfWY5I3NLoMcvshKC1GDQ8XTIzdM48R5lt3wdtz25mqUX8no+Jzr6Pocnxlv89MmPDDNdMAqaDp3GrDFTYjCy43AEx22cAC3PaSNXa6yum+SZbBF/RLexBJoNirKARMJA="
var port = flag.String("p", ":8081", "port")
var mode = flag.String("m", "debug", "run mode")
var logPath = flag.String("l", "access.log", "")

func main() {
	flag.Parse()
	file, _ := os.Create(*logPath)
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	gin.SetMode(*mode)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		m := make(map[string]interface{})
		m["a"] = 1
		m["b"] = "h"
		c.JSON(200, m)
	})
	r.POST("/jrebel/leases", func(c *gin.Context) {
		leasesHandler(c)
	})
	r.POST("/agent/leases", func(c *gin.Context) {
		leasesHandler(c)
	})
	r.DELETE("/jrebel/leases/1", func(c *gin.Context) {
		leases1Handler(c)
	})
	r.DELETE("/agent/leases/1", func(context *gin.Context) {
		leases1Handler(context)
	})
	r.POST("/jrebel/validate-connection", func(context *gin.Context) {
		jrebelValidateHandler(context)
	})
	r.GET("/jrebel/validate-connection", func(context *gin.Context) {
		jrebelValidateHandler(context)
	})
	r.POST("/rpc/obtainTicket.action", func(c *gin.Context) {
		obtainTicketHandler(c)
	})
	r.Run(*port)
}

func pingHandler(c *gin.Context) {

}

func jrebelValidateHandler(c *gin.Context) {
	jsonStr := "{\n" +
		"    \"serverVersion\": \"3.2.4\",\n" +
		"    \"serverProtocolVersion\": \"1.1\",\n" +
		"    \"serverGuid\": \"a1b4aea8-b031-4302-b602-670a990272cb\",\n" +
		"    \"groupType\": \"managed\",\n" +
		"    \"statusCode\": \"SUCCESS\",\n" +
		"    \"company\": \"Administrator\",\n" +
		"    \"canGetLease\": true,\n" +
		"    \"licenseType\": 1,\n" +
		"    \"evaluationLicense\": false,\n" +
		"    \"seatPoolType\": \"standalone\"\n" +
		"}\n"
	m := make(map[string]interface{}, 4)

	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(http.StatusOK, m)
}

func leases1Handler(c *gin.Context) {
	username := c.Query("username")
	jsonStr := fmt.Sprintf("{\n"+
		"    \"serverVersion\": \"3.2.4\",\n"+
		"    \"serverProtocolVersion\": \"1.1\",\n"+
		"    \"serverGuid\": \"a1b4aea8-b031-4302-b602-670a990272cb\",\n"+
		"    \"groupType\": \"managed\",\n"+
		"    \"statusCode\": \"SUCCESS\",\n"+
		"    \"msg\": null,\n"+
		"    \"statusMessage\": null\n"+
		"    \"company\": \"%s\"\n"+
		"}\n", username)
	//m := make(map[string]interface{}, 4)
	//
	//err := json.Unmarshal([]byte(jsonStr), &m)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	c.String(http.StatusOK, "%s", jsonStr)
}

func leasesHandler(c *gin.Context) {
	cr := c.PostForm("randomness")
	username := c.PostForm("username")
	guid := c.PostForm("guid")
	offline, _ := strconv.ParseBool(c.PostForm("offline"))
	validFrom := "null"
	validTo := "null"
	if offline {
		ct, _ := strconv.ParseInt(c.PostForm("clientTime"), 10, 64)
		ut, _ := strconv.ParseInt(c.PostForm("offlineDays"), 10, 64)
		until := ct + ut*24*3600*1000
		validFrom = strconv.FormatInt(ct, 10)
		validTo = strconv.FormatInt(until, 10)
	}

	sign := rsautil.CreateSign(cr, guid, offline, validFrom, validTo)
	jsonStr := fmt.Sprintf("{\n"+
		"    \"serverVersion\": \"3.2.4\",\n"+
		"    \"serverProtocolVersion\": \"1.1\",\n"+
		"    \"serverGuid\": \"a1b4aea8-b031-4302-b602-670a990272cb\",\n"+
		"    \"groupType\": \"managed\",\n"+
		"    \"id\": 1,\n"+
		"    \"licenseType\": 1,\n"+
		"    \"evaluationLicense\": false,\n"+
		"    \"signature\": \"%s\",\n"+
		"    \"serverRandomness\": \"H2ulzLlh7E0=\",\n"+
		"    \"seatPoolType\": \"standalone\",\n"+
		"    \"statusCode\": \"SUCCESS\",\n"+
		"    \"offline\": "+strconv.FormatBool(offline)+",\n"+
		"    \"validFrom\": "+validFrom+",\n"+
		"    \"validUntil\": "+validTo+",\n"+
		"    \"company\": \"%s\",\n"+
		"    \"orderId\": \"\",\n"+
		"    \"zeroIds\": [\n"+
		"        \n"+
		"    ],\n"+
		"    \"licenseValidFrom\": 1490544001000,\n"+
		"    \"licenseValidUntil\": 1924905600000\n"+
		"}", sign, username)
	//fmt.Println(jsonStr)
	m := make(map[string]interface{}, 4)

	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(http.StatusOK, m)
}

func obtainTicketHandler(c *gin.Context) {
	salt := c.PostForm("salt")
	username := c.PostForm("userName")
	prolongationPeriod := "607875500"
	xmlContent := "<ObtainTicketResponse><message></message><prolongationPeriod>" + prolongationPeriod + "</prolongationPeriod><responseCode>OK</responseCode><salt>" + salt + "</salt><ticketId>1</ticketId><ticketProperties>licensee=" + username + "\tlicenseType=0\t</ticketProperties></ObtainTicketResponse>"
	xmlSignature, _ := rsautil.PrivateSignMd5([]byte(xmlContent))
	body := "<!-- " + xmlSignature + " -->\n" + xmlContent
	c.String(http.StatusOK, "%s", body)
}
