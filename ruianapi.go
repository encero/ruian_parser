package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"path"
	"regexp"
	"strings"
	"time"
)

type RuianFileContents int

const (
	ContentBasic RuianFileContents = iota
	ContentCompleteWithDescriptionAndGeneralizedBorders
	ContentCompleteWithDescriptionAndOriginalBorders
	ContentImages

	ExtentCity  = "OB"
	ExtentState = "ST"
)

type RuianFileDescriptor struct {
	Raw             string
	Date            time.Time
	CityCode        string
	Extent          string
	BaseDataPackage bool
	Increments      bool
	Actual          bool
	Contents        RuianFileContents
}

type RuianAPI struct {
	Doer interface {
		Do(req *http.Request) (*http.Response, error)
	}
}

const (
	userAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36"

	fullDataLinkListUrl                = "https://vdp.cuzk.cz/vdp/ruian/vymennyformat/seznamlinku?vf.pu=S&_vf.pu=on&_vf.pu=on&vf.cr=U&vf.up=OB&vf.ds=K&_vf.vu=on&_vf.vu=on&_vf.vu=on&_vf.vu=on&vf.uo=A&search=Vyhledat"
	incrementalDataLinkListUrlTemplate = "https://vdp.cuzk.cz/vdp/ruian/vymennyformat/seznamlinku?vf.pu=S&_vf.pu=on&_vf.pu=on&vf.cr=Z&vf.pd={_date_}&vf.ds=K&_vf.vu=on&_vf.vu=on&vf.vu=H&_vf.vu=on&_vf.vu=on&search=Vyhledat"

	defaultFullResultLinkCount = 20_000
)

func (api RuianAPI) FullDataLinkList(ctx context.Context) ([]string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fullDataLinkListUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)
	resp, err := api.Doer.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	result := readLines(resp.Body)

	return result, nil
}

func (api RuianAPI) IncrementalDataLinkList(ctx context.Context, fromDate time.Time) ([]string, error) {
	url := strings.Replace(incrementalDataLinkListUrlTemplate, "{_date_}", fromDate.Format("02.01.2006"), 1)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)
	resp, err := api.Doer.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	result := readLines(resp.Body)

	return result, nil
}

var filenameValidationEx = regexp.MustCompile("[0-9]{8}_((OB_[0-9]+)|ST)_[UZ][ZK][SH][ZGHO]")

// ParseFileName extract metadata from filename of ruian export file
func ParseFileName(filename string) (RuianFileDescriptor, error) {
	desc := RuianFileDescriptor{
		Raw: filename,
	}

	base := path.Base(filename)

	if !filenameValidationEx.MatchString(base) {
		return desc, fmt.Errorf("invalid filename: %s", filename)
	}

	parts := strings.Split(strings.SplitN(base, ".", 1)[0], "_")

	if len(parts) != 4 && len(parts) != 3 {
		return RuianFileDescriptor{}, fmt.Errorf("invalid ruian filename part count: %s", filename)
	}

	var index = 0

	var err error
	desc.Date, err = time.Parse("20060102", parts[index])
	if err != nil {
		return RuianFileDescriptor{}, fmt.Errorf("invalid ruian filename date: %s", filename)
	}
	index++

	switch parts[index] {
	case ExtentCity:
		desc.Extent = ExtentCity
	case ExtentState:
		desc.Extent = ExtentState
	default:
		return RuianFileDescriptor{}, fmt.Errorf("invalid ruian filename extent: %s", filename)
	}
	index++

	if len(parts) == 4 {
		desc.CityCode = parts[index]
		index++
	}

	details := strings.Split(parts[index], "")
	desc.Increments = details[0] == "Z"
	desc.BaseDataPackage = details[1] == "Z"
	desc.Actual = details[2] == "S"

	switch details[3] {
	case "H":
		desc.Contents = ContentCompleteWithDescriptionAndOriginalBorders
	case "O":
		desc.Contents = ContentImages
	case "Z":
		desc.Contents = ContentBasic
	case "G":
		desc.Contents = ContentCompleteWithDescriptionAndGeneralizedBorders
	}

	return desc, nil
}

func FilterNewestFromList(input []string) ([]string, error) {
	newest := make(map[string]RuianFileDescriptor, len(input)/3)

	for i, link := range input {
		desc, err := ParseFileName(link)
		if err != nil {
			return nil, fmt.Errorf("can't parse input [%s] at: %d: %w", link, i, err)
		}

		if desc.CityCode == "" {
			return nil, fmt.Errorf("Can't filter non City source files file: [%s] at: %d", link, i)
		}

		if _, ok := newest[desc.CityCode]; !ok {
			newest[desc.CityCode] = desc
		} else if newest[desc.CityCode].Date.Before(desc.Date) {
			newest[desc.CityCode] = desc
		}

	}

	output := make([]string, 0, len(newest))
	for _, desc := range newest {
		output = append(output, desc.Raw)
	}

	return output, nil
}

func readLines(r io.Reader) []string {
	scanner := bufio.NewScanner(r)

	result := make([]string, 0, defaultFullResultLinkCount)

	for scanner.Scan() {
		row := scanner.Text()

		result = append(result, strings.TrimSpace(row))
	}

	return result
}
