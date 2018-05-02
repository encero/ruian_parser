package main

import (
	"bytes"
	is_ "github.com/matryer/is"
	"io"
	"net/http"
	"sort"
	"testing"
	"time"
)

type httpDoer func(req *http.Request) (*http.Response, error)

func (h httpDoer) Do(req *http.Request) (*http.Response, error) {
	return h(req)
}

func TestFetchFullLoadList(t *testing.T) {
	is := is_.New(t)

	api := RuianAPI{
		Doer: httpDoer(func(req *http.Request) (*http.Response, error) {
			is.Equal(req.Method, "GET")
			is.Equal(req.URL.String(), fullDataLinkListUrl)

			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewBufferString("1\n2\n")),
			}, nil
		}),
	}
	list, err := api.FullDataLinkList(testCtx())
	is.NoErr(err)

	is.Equal(list, []string{"1", "2"})

}

func TestIncrementalFullLoadList(t *testing.T) {
	is := is_.New(t)

	api := RuianAPI{
		Doer: httpDoer(func(req *http.Request) (*http.Response, error) {
			is.Equal(req.Method, "GET")
			is.Equal(req.URL.String(), "https://vdp.cuzk.cz/vdp/ruian/vymennyformat/seznamlinku?vf.pu=S&_vf.pu=on&_vf.pu=on&vf.cr=Z&vf.pd=15.11.2021&vf.ds=K&_vf.vu=on&_vf.vu=on&vf.vu=H&_vf.vu=on&_vf.vu=on&search=Vyhledat")

			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewBufferString("1\n2\n")),
			}, nil
		}),
	}
	fromDate, _ := time.Parse("2006-01-02", "2021-11-15")
	list, err := api.IncrementalDataLinkList(testCtx(), fromDate)
	is.NoErr(err)

	is.Equal(list, []string{"1", "2"})
}

func TestParseFileName(t *testing.T) {
	dateX := func(date string) time.Time {
		parsed, err := time.Parse("2006-01-02", date)
		if err != nil {
			panic(err)
		}
		return parsed
	}

	tests := []struct {
		filename string
		expected RuianFileDescriptor
	}{
		{
			filename: "20211031_OB_500071_UKSH.xml.zip",
			expected: RuianFileDescriptor{
				Raw:             "20211031_OB_500071_UKSH.xml.zip",
				Date:            dateX("2021-10-31"),
				CityCode:        "500071",
				Extent:          "OB",
				BaseDataPackage: false,
				Increments:      false,
				Actual:          true,
				Contents:        ContentCompleteWithDescriptionAndOriginalBorders,
			},
		},
		{
			filename: "20211006_ST_ZKSH.xml.zip",
			expected: RuianFileDescriptor{
				Raw:             "20211006_ST_ZKSH.xml.zip",
				Date:            dateX("2021-10-06"),
				CityCode:        "",
				Extent:          "ST",
				BaseDataPackage: false,
				Increments:      true,
				Actual:          true,
				Contents:        ContentCompleteWithDescriptionAndOriginalBorders,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.filename, func(t *testing.T) {
			is := is_.New(t)

			desc, err := ParseFileName(tc.filename)
			is.NoErr(err)

			is.Equal(desc, tc.expected)
		})
	}
}

func TestFilterNewestFromList(t *testing.T) {
	is := is_.New(t)

	input := []string{
		"https://vdp.cuzk.cz/vymenny_format/soucasna/20210930_OB_599808_UKSH.xml.zip",
		"https://vdp.cuzk.cz/vymenny_format/soucasna/20210930_OB_599832_UKSH.xml.zip",
		"https://vdp.cuzk.cz/vymenny_format/soucasna/20210930_OB_599867_UKSH.xml.zip",
		"https://vdp.cuzk.cz/vymenny_format/soucasna/20210930_OB_599905_UKSH.xml.zip",
		"https://vdp.cuzk.cz/vymenny_format/soucasna/20210930_OB_599921_UKSH.xml.zip",

		"https://vdp.cuzk.cz/vymenny_format/soucasna/20211031_OB_599808_UKSH.xml.zip",
		"https://vdp.cuzk.cz/vymenny_format/soucasna/20211031_OB_599832_UKSH.xml.zip",
		"https://vdp.cuzk.cz/vymenny_format/soucasna/20211031_OB_599867_UKSH.xml.zip",
		"https://vdp.cuzk.cz/vymenny_format/soucasna/20211031_OB_599905_UKSH.xml.zip",
		"https://vdp.cuzk.cz/vymenny_format/soucasna/20211031_OB_599921_UKSH.xml.zip",

		"https://vdp.cuzk.cz/vymenny_format/soucasna/20210830_OB_599808_UKSH.xml.zip",
		"https://vdp.cuzk.cz/vymenny_format/soucasna/20210830_OB_599832_UKSH.xml.zip",
		"https://vdp.cuzk.cz/vymenny_format/soucasna/20210830_OB_599867_UKSH.xml.zip",
		"https://vdp.cuzk.cz/vymenny_format/soucasna/20210830_OB_599921_UKSH.xml.zip",
		"https://vdp.cuzk.cz/vymenny_format/soucasna/20210830_OB_599905_UKSH.xml.zip",

		// duplicates
		"https://vdp.cuzk.cz/vymenny_format/soucasna/20211031_OB_599905_UKSH.xml.zip",
		"https://vdp.cuzk.cz/vymenny_format/soucasna/20211031_OB_599921_UKSH.xml.zip",
	}

	expected := []string{
		"https://vdp.cuzk.cz/vymenny_format/soucasna/20211031_OB_599808_UKSH.xml.zip",
		"https://vdp.cuzk.cz/vymenny_format/soucasna/20211031_OB_599832_UKSH.xml.zip",
		"https://vdp.cuzk.cz/vymenny_format/soucasna/20211031_OB_599867_UKSH.xml.zip",
		"https://vdp.cuzk.cz/vymenny_format/soucasna/20211031_OB_599905_UKSH.xml.zip",
		"https://vdp.cuzk.cz/vymenny_format/soucasna/20211031_OB_599921_UKSH.xml.zip",
	}



	got, err := FilterNewestFromList(input)
	is.NoErr(err)

	sort.Strings(expected)
	sort.Strings(got)


	is.Equal(expected, got)
}
