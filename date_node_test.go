package gedcom_test

import (
	"strings"
	"testing"
	"time"

	"errors"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

func parseTime(s string) time.Time {
	if strings.HasSuffix(s, "23") {
		s += ":59:59.999999999"
	} else {
		s += ":00:00.000000000"
	}

	d, err := time.Parse("_2 Jan 2006 15:04:05.000000000", s)
	if err != nil {
		panic(err)
	}

	return d
}

var dateTests = map[string]struct {
	startDate gedcom.Date
	startTime time.Time
	endDate   gedcom.Date
	endTime   time.Time
	str       string
}{
	// Valid dates, testing each 3 digit month name. The days are a mix of DD
	// and D.
	"01 Jan 1980": {
		gedcom.Date{1, time.January, 1980, false, gedcom.DateConstraintExact, nil}, parseTime("1 Jan 1980 00"),
		gedcom.Date{1, time.January, 1980, true, gedcom.DateConstraintExact, nil}, parseTime("1 Jan 1980 23"),
		"1 Jan 1980",
	},
	"15 Feb 1880": {
		gedcom.Date{15, time.February, 1880, false, gedcom.DateConstraintExact, nil}, parseTime("15 Feb 1880 00"),
		gedcom.Date{15, time.February, 1880, true, gedcom.DateConstraintExact, nil}, parseTime("15 Feb 1880 23"),
		"15 Feb 1880",
	},
	"03 Mar 1870": {
		gedcom.Date{3, time.March, 1870, false, gedcom.DateConstraintExact, nil}, parseTime("3 Mar 1870 00"),
		gedcom.Date{3, time.March, 1870, true, gedcom.DateConstraintExact, nil}, parseTime("3 Mar 1870 23"),
		"3 Mar 1870",
	},
	"7 Apr 2020": {
		gedcom.Date{7, time.April, 2020, false, gedcom.DateConstraintExact, nil}, parseTime("7 Apr 2020 00"),
		gedcom.Date{7, time.April, 2020, true, gedcom.DateConstraintExact, nil}, parseTime("7 Apr 2020 23"),
		"7 Apr 2020",
	},
	"6 May 1989": {
		gedcom.Date{6, time.May, 1989, false, gedcom.DateConstraintExact, nil}, parseTime("6 May 1989 00"),
		gedcom.Date{6, time.May, 1989, true, gedcom.DateConstraintExact, nil}, parseTime("6 May 1989 23"),
		"6 May 1989",
	},
	"8 Jun 2001": {
		gedcom.Date{8, time.June, 2001, false, gedcom.DateConstraintExact, nil}, parseTime("8 Jun 2001 00"),
		gedcom.Date{8, time.June, 2001, true, gedcom.DateConstraintExact, nil}, parseTime("8 Jun 2001 23"),
		"8 Jun 2001",
	},
	"19 Jul 2003": {
		gedcom.Date{19, time.July, 2003, false, gedcom.DateConstraintExact, nil}, parseTime("19 Jul 2003 00"),
		gedcom.Date{19, time.July, 2003, true, gedcom.DateConstraintExact, nil}, parseTime("19 Jul 2003 23"),
		"19 Jul 2003",
	},
	"29 Aug 1640": {
		gedcom.Date{29, time.August, 1640, false, gedcom.DateConstraintExact, nil}, parseTime("29 Aug 1640 00"),
		gedcom.Date{29, time.August, 1640, true, gedcom.DateConstraintExact, nil}, parseTime("29 Aug 1640 23"),
		"29 Aug 1640",
	},
	"13 Sep 1733": {
		gedcom.Date{13, time.September, 1733, false, gedcom.DateConstraintExact, nil}, parseTime("13 Sep 1733 00"),
		gedcom.Date{13, time.September, 1733, true, gedcom.DateConstraintExact, nil}, parseTime("13 Sep 1733 23"),
		"13 Sep 1733",
	},
	"6 Oct 1848": {
		gedcom.Date{6, time.October, 1848, false, gedcom.DateConstraintExact, nil}, parseTime("6 Oct 1848 00"),
		gedcom.Date{6, time.October, 1848, true, gedcom.DateConstraintExact, nil}, parseTime("6 Oct 1848 23"),
		"6 Oct 1848",
	},
	"18 Nov 1992": {
		gedcom.Date{18, time.November, 1992, false, gedcom.DateConstraintExact, nil}, parseTime("18 Nov 1992 00"),
		gedcom.Date{18, time.November, 1992, true, gedcom.DateConstraintExact, nil}, parseTime("18 Nov 1992 23"),
		"18 Nov 1992",
	},
	"25 Dec 1901": {
		gedcom.Date{25, time.December, 1901, false, gedcom.DateConstraintExact, nil}, parseTime("25 Dec 1901 00"),
		gedcom.Date{25, time.December, 1901, true, gedcom.DateConstraintExact, nil}, parseTime("25 Dec 1901 23"),
		"25 Dec 1901",
	},

	// Valid dates, testing each full month name. The days are a mix of dd
	// and d.
	"01 January 1980": {
		gedcom.Date{1, time.January, 1980, false, gedcom.DateConstraintExact, nil}, parseTime("1 Jan 1980 00"),
		gedcom.Date{1, time.January, 1980, true, gedcom.DateConstraintExact, nil}, parseTime("1 Jan 1980 23"),
		"1 Jan 1980",
	},
	"15 February 1880": {
		gedcom.Date{15, time.February, 1880, false, gedcom.DateConstraintExact, nil}, parseTime("15 Feb 1880 00"),
		gedcom.Date{15, time.February, 1880, true, gedcom.DateConstraintExact, nil}, parseTime("15 Feb 1880 23"),
		"15 Feb 1880",
	},
	"03 March 1870": {
		gedcom.Date{3, time.March, 1870, false, gedcom.DateConstraintExact, nil}, parseTime("3 Mar 1870 00"),
		gedcom.Date{3, time.March, 1870, true, gedcom.DateConstraintExact, nil}, parseTime("3 Mar 1870 23"),
		"3 Mar 1870",
	},
	"7 April 2020": {
		gedcom.Date{7, time.April, 2020, false, gedcom.DateConstraintExact, nil}, parseTime("7 Apr 2020 00"),
		gedcom.Date{7, time.April, 2020, true, gedcom.DateConstraintExact, nil}, parseTime("7 Apr 2020 23"),
		"7 Apr 2020",
	},
	"8 June 2001": {
		gedcom.Date{8, time.June, 2001, false, gedcom.DateConstraintExact, nil}, parseTime("8 Jun 2001 00"),
		gedcom.Date{8, time.June, 2001, true, gedcom.DateConstraintExact, nil}, parseTime("8 Jun 2001 23"),
		"8 Jun 2001",
	},
	"19 July 2003": {
		gedcom.Date{19, time.July, 2003, false, gedcom.DateConstraintExact, nil}, parseTime("19 Jul 2003 00"),
		gedcom.Date{19, time.July, 2003, true, gedcom.DateConstraintExact, nil}, parseTime("19 Jul 2003 23"),
		"19 Jul 2003",
	},
	"29 August 1640": {
		gedcom.Date{29, time.August, 1640, false, gedcom.DateConstraintExact, nil}, parseTime("29 Aug 1640 00"),
		gedcom.Date{29, time.August, 1640, true, gedcom.DateConstraintExact, nil}, parseTime("29 Aug 1640 23"),
		"29 Aug 1640",
	},
	"13 September 1733": {
		gedcom.Date{13, time.September, 1733, false, gedcom.DateConstraintExact, nil}, parseTime("13 Sep 1733 00"),
		gedcom.Date{13, time.September, 1733, true, gedcom.DateConstraintExact, nil}, parseTime("13 Sep 1733 23"),
		"13 Sep 1733",
	},
	"6 October 1848": {
		gedcom.Date{6, time.October, 1848, false, gedcom.DateConstraintExact, nil}, parseTime("6 Oct 1848 00"),
		gedcom.Date{6, time.October, 1848, true, gedcom.DateConstraintExact, nil}, parseTime("6 Oct 1848 23"),
		"6 Oct 1848",
	},
	"18 November 1992": {
		gedcom.Date{18, time.November, 1992, false, gedcom.DateConstraintExact, nil}, parseTime("18 Nov 1992 00"),
		gedcom.Date{18, time.November, 1992, true, gedcom.DateConstraintExact, nil}, parseTime("18 Nov 1992 23"),
		"18 Nov 1992",
	},
	"25 December 1901": {
		gedcom.Date{25, time.December, 1901, false, gedcom.DateConstraintExact, nil}, parseTime("25 Dec 1901 00"),
		gedcom.Date{25, time.December, 1901, true, gedcom.DateConstraintExact, nil}, parseTime("25 Dec 1901 23"),
		"25 Dec 1901",
	},

	// Only month and year combinations.
	"Jan 1980": {
		gedcom.Date{0, time.January, 1980, false, gedcom.DateConstraintExact, nil}, parseTime("1 Jan 1980 00"),
		gedcom.Date{0, time.January, 1980, true, gedcom.DateConstraintExact, nil}, parseTime("31 Jan 1980 23"),
		"Jan 1980",
	},
	"Feb 1880": {
		gedcom.Date{0, time.February, 1880, false, gedcom.DateConstraintExact, nil}, parseTime("1 Feb 1880 00"),
		gedcom.Date{0, time.February, 1880, true, gedcom.DateConstraintExact, nil}, parseTime("29 Feb 1880 23"),
		"Feb 1880",
	},
	"Mar 1870": {
		gedcom.Date{0, time.March, 1870, false, gedcom.DateConstraintExact, nil}, parseTime("1 Mar 1870 00"),
		gedcom.Date{0, time.March, 1870, true, gedcom.DateConstraintExact, nil}, parseTime("31 Mar 1870 23"),
		"Mar 1870",
	},
	"Apr 2020": {
		gedcom.Date{0, time.April, 2020, false, gedcom.DateConstraintExact, nil}, parseTime("1 Apr 2020 00"),
		gedcom.Date{0, time.April, 2020, true, gedcom.DateConstraintExact, nil}, parseTime("30 Apr 2020 23"),
		"Apr 2020",
	},
	"May 1989": {
		gedcom.Date{0, time.May, 1989, false, gedcom.DateConstraintExact, nil}, parseTime("1 May 1989 00"),
		gedcom.Date{0, time.May, 1989, true, gedcom.DateConstraintExact, nil}, parseTime("31 May 1989 23"),
		"May 1989",
	},
	"Jun 2001": {
		gedcom.Date{0, time.June, 2001, false, gedcom.DateConstraintExact, nil}, parseTime("1 Jun 2001 00"),
		gedcom.Date{0, time.June, 2001, true, gedcom.DateConstraintExact, nil}, parseTime("30 Jun 2001 23"),
		"Jun 2001",
	},
	"Jul 2003": {
		gedcom.Date{0, time.July, 2003, false, gedcom.DateConstraintExact, nil}, parseTime("1 Jul 2003 00"),
		gedcom.Date{0, time.July, 2003, true, gedcom.DateConstraintExact, nil}, parseTime("31 Jul 2003 23"),
		"Jul 2003",
	},
	"Aug 1640": {
		gedcom.Date{0, time.August, 1640, false, gedcom.DateConstraintExact, nil}, parseTime("1 Aug 1640 00"),
		gedcom.Date{0, time.August, 1640, true, gedcom.DateConstraintExact, nil}, parseTime("31 Aug 1640 23"),
		"Aug 1640",
	},
	"Sep 1733": {
		gedcom.Date{0, time.September, 1733, false, gedcom.DateConstraintExact, nil}, parseTime("1 Sep 1733 00"),
		gedcom.Date{0, time.September, 1733, true, gedcom.DateConstraintExact, nil}, parseTime("30 Sep 1733 23"),
		"Sep 1733",
	},
	"Oct 1848": {
		gedcom.Date{0, time.October, 1848, false, gedcom.DateConstraintExact, nil}, parseTime("1 Oct 1848 00"),
		gedcom.Date{0, time.October, 1848, true, gedcom.DateConstraintExact, nil}, parseTime("31 Oct 1848 23"),
		"Oct 1848",
	},
	"Nov 1992": {
		gedcom.Date{0, time.November, 1992, false, gedcom.DateConstraintExact, nil}, parseTime("1 Nov 1992 00"),
		gedcom.Date{0, time.November, 1992, true, gedcom.DateConstraintExact, nil}, parseTime("30 Nov 1992 23"),
		"Nov 1992",
	},
	"Dec 1901": {
		gedcom.Date{0, time.December, 1901, false, gedcom.DateConstraintExact, nil}, parseTime("1 Dec 1901 00"),
		gedcom.Date{0, time.December, 1901, true, gedcom.DateConstraintExact, nil}, parseTime("31 Dec 1901 23"),
		"Dec 1901",
	},
	"January 1980": {
		gedcom.Date{0, time.January, 1980, false, gedcom.DateConstraintExact, nil}, parseTime("1 Jan 1980 00"),
		gedcom.Date{0, time.January, 1980, true, gedcom.DateConstraintExact, nil}, parseTime("31 Jan 1980 23"),
		"Jan 1980",
	},
	"February 1880": {
		gedcom.Date{0, time.February, 1880, false, gedcom.DateConstraintExact, nil}, parseTime("1 Feb 1880 00"),
		gedcom.Date{0, time.February, 1880, true, gedcom.DateConstraintExact, nil}, parseTime("29 Feb 1880 23"),
		"Feb 1880",
	},
	"March 1870": {
		gedcom.Date{0, time.March, 1870, false, gedcom.DateConstraintExact, nil}, parseTime("1 Mar 1870 00"),
		gedcom.Date{0, time.March, 1870, true, gedcom.DateConstraintExact, nil}, parseTime("31 Mar 1870 23"),
		"Mar 1870",
	},
	"April 2020": {
		gedcom.Date{0, time.April, 2020, false, gedcom.DateConstraintExact, nil}, parseTime("1 Apr 2020 00"),
		gedcom.Date{0, time.April, 2020, true, gedcom.DateConstraintExact, nil}, parseTime("30 Apr 2020 23"),
		"Apr 2020",
	},
	"June 2001": {
		gedcom.Date{0, time.June, 2001, false, gedcom.DateConstraintExact, nil}, parseTime("1 Jun 2001 00"),
		gedcom.Date{0, time.June, 2001, true, gedcom.DateConstraintExact, nil}, parseTime("30 Jun 2001 23"),
		"Jun 2001",
	},
	"July 2003": {
		gedcom.Date{0, time.July, 2003, false, gedcom.DateConstraintExact, nil}, parseTime("1 Jul 2003 00"),
		gedcom.Date{0, time.July, 2003, true, gedcom.DateConstraintExact, nil}, parseTime("31 Jul 2003 23"),
		"Jul 2003",
	},
	"August 1640": {
		gedcom.Date{0, time.August, 1640, false, gedcom.DateConstraintExact, nil}, parseTime("1 Aug 1640 00"),
		gedcom.Date{0, time.August, 1640, true, gedcom.DateConstraintExact, nil}, parseTime("31 Aug 1640 23"),
		"Aug 1640",
	},
	"September 1733": {
		gedcom.Date{0, time.September, 1733, false, gedcom.DateConstraintExact, nil}, parseTime("1 Sep 1733 00"),
		gedcom.Date{0, time.September, 1733, true, gedcom.DateConstraintExact, nil}, parseTime("30 Sep 1733 23"),
		"Sep 1733",
	},
	"October 1848": {
		gedcom.Date{0, time.October, 1848, false, gedcom.DateConstraintExact, nil}, parseTime("1 Oct 1848 00"),
		gedcom.Date{0, time.October, 1848, true, gedcom.DateConstraintExact, nil}, parseTime("31 Oct 1848 23"),
		"Oct 1848",
	},
	"November 1992": {
		gedcom.Date{0, time.November, 1992, false, gedcom.DateConstraintExact, nil}, parseTime("1 Nov 1992 00"),
		gedcom.Date{0, time.November, 1992, true, gedcom.DateConstraintExact, nil}, parseTime("30 Nov 1992 23"),
		"Nov 1992",
	},
	"December 1901": {
		gedcom.Date{0, time.December, 1901, false, gedcom.DateConstraintExact, nil}, parseTime("1 Dec 1901 00"),
		gedcom.Date{0, time.December, 1901, true, gedcom.DateConstraintExact, nil}, parseTime("31 Dec 1901 23"),
		"Dec 1901",
	},

	// Months with different capitalization.
	"DECEMBER 1901": {
		gedcom.Date{0, time.December, 1901, false, gedcom.DateConstraintExact, nil}, parseTime("1 Dec 1901 00"),
		gedcom.Date{0, time.December, 1901, true, gedcom.DateConstraintExact, nil}, parseTime("31 Dec 1901 23"),
		"Dec 1901",
	},
	"13 SEP 1733": {
		gedcom.Date{13, time.September, 1733, false, gedcom.DateConstraintExact, nil}, parseTime("13 Sep 1733 00"),
		gedcom.Date{13, time.September, 1733, true, gedcom.DateConstraintExact, nil}, parseTime("13 Sep 1733 23"),
		"13 Sep 1733",
	},

	// Only year.
	"834": {
		gedcom.Date{0, 0, 834, false, gedcom.DateConstraintExact, nil}, parseTime("1 Jan 0834 00"),
		gedcom.Date{0, 0, 834, true, gedcom.DateConstraintExact, nil}, parseTime("31 Dec 0834 23"),
		"834",
	},
	"0834": {
		gedcom.Date{0, 0, 834, false, gedcom.DateConstraintExact, nil}, parseTime("1 Jan 0834 00"),
		gedcom.Date{0, 0, 834, true, gedcom.DateConstraintExact, nil}, parseTime("31 Dec 0834 23"),
		"834",
	},
	"1901": {
		gedcom.Date{0, 0, 1901, false, gedcom.DateConstraintExact, nil}, parseTime("1 Jan 1901 00"),
		gedcom.Date{0, 0, 1901, true, gedcom.DateConstraintExact, nil}, parseTime("31 Dec 1901 23"),
		"1901",
	},
	"2020": {
		gedcom.Date{0, 0, 2020, false, gedcom.DateConstraintExact, nil}, parseTime("1 Jan 2020 00"),
		gedcom.Date{0, 0, 2020, true, gedcom.DateConstraintExact, nil}, parseTime("31 Dec 2020 23"),
		"2020",
	},
	"0066": {
		gedcom.Date{0, 0, 66, false, gedcom.DateConstraintExact, nil}, parseTime("1 Jan 0066 00"),
		gedcom.Date{0, 0, 66, true, gedcom.DateConstraintExact, nil}, parseTime("31 Dec 0066 23"),
		"66",
	},

	// Extra whitespace. The GEDCOM file should not allow values to contain new
	// lines or carriage returns in the node value so we do not need to test
	// those cases.
	"  18 November 1992": {
		gedcom.Date{18, time.November, 1992, false, gedcom.DateConstraintExact, nil}, parseTime("18 Nov 1992 00"),
		gedcom.Date{18, time.November, 1992, true, gedcom.DateConstraintExact, nil}, parseTime("18 Nov 1992 23"),
		"18 Nov 1992",
	},
	"15 Feb   1880": {
		gedcom.Date{15, time.February, 1880, false, gedcom.DateConstraintExact, nil}, parseTime("15 Feb 1880 00"),
		gedcom.Date{15, time.February, 1880, true, gedcom.DateConstraintExact, nil}, parseTime("15 Feb 1880 23"),
		"15 Feb 1880",
	},
	"Feb   1880": {
		gedcom.Date{0, time.February, 1880, false, gedcom.DateConstraintExact, nil}, parseTime("1 Feb 1880 00"),
		gedcom.Date{0, time.February, 1880, true, gedcom.DateConstraintExact, nil}, parseTime("29 Feb 1880 23"),
		"Feb 1880",
	},
	"25 December 1901  ": {
		gedcom.Date{25, time.December, 1901, false, gedcom.DateConstraintExact, nil}, parseTime("25 Dec 1901 00"),
		gedcom.Date{25, time.December, 1901, true, gedcom.DateConstraintExact, nil}, parseTime("25 Dec 1901 23"),
		"25 Dec 1901",
	},
	" 1901  ": {
		gedcom.Date{0, 0, 1901, false, gedcom.DateConstraintExact, nil}, parseTime("1 Jan 1901 00"),
		gedcom.Date{0, 0, 1901, true, gedcom.DateConstraintExact, nil}, parseTime("31 Dec 1901 23"),
		"1901",
	},

	// Before dates.
	"Before Feb 1907": {
		gedcom.Date{0, time.February, 1907, false, gedcom.DateConstraintBefore, nil}, parseTime("1 Feb 1907 00"),
		gedcom.Date{0, time.February, 1907, true, gedcom.DateConstraintBefore, nil}, parseTime("28 Feb 1907 23"),
		"Bef. Feb 1907",
	},
	"bef. 21 Dec 1884": {
		gedcom.Date{21, time.December, 1884, false, gedcom.DateConstraintBefore, nil}, parseTime("21 Dec 1884 00"),
		gedcom.Date{21, time.December, 1884, true, gedcom.DateConstraintBefore, nil}, parseTime("21 Dec 1884 23"),
		"Bef. 21 Dec 1884",
	},

	// After dates.
	"after Feb 1907": {
		gedcom.Date{0, time.February, 1907, false, gedcom.DateConstraintAfter, nil}, parseTime("1 Feb 1907 00"),
		gedcom.Date{0, time.February, 1907, true, gedcom.DateConstraintAfter, nil}, parseTime("28 Feb 1907 23"),
		"Aft. Feb 1907",
	},
	"Aft. 21 Dec 1884": {
		gedcom.Date{21, time.December, 1884, false, gedcom.DateConstraintAfter, nil}, parseTime("21 Dec 1884 00"),
		gedcom.Date{21, time.December, 1884, true, gedcom.DateConstraintAfter, nil}, parseTime("21 Dec 1884 23"),
		"Aft. 21 Dec 1884",
	},

	// Approximate dates.
	"Abt. 1945": {
		gedcom.Date{0, 0, 1945, false, gedcom.DateConstraintAbout, nil}, parseTime("1 Jan 1945 00"),
		gedcom.Date{0, 0, 1945, true, gedcom.DateConstraintAbout, nil}, parseTime("31 Dec 1945 23"),
		"Abt. 1945",
	},
	"about Feb 1907": {
		gedcom.Date{0, time.February, 1907, false, gedcom.DateConstraintAbout, nil}, parseTime("1 Feb 1907 00"),
		gedcom.Date{0, time.February, 1907, true, gedcom.DateConstraintAbout, nil}, parseTime("28 Feb 1907 23"),
		"Abt. Feb 1907",
	},
	"c. 8 Mar 1505": {
		gedcom.Date{8, time.March, 1505, false, gedcom.DateConstraintAbout, nil}, parseTime("8 Mar 1505 00"),
		gedcom.Date{8, time.March, 1505, true, gedcom.DateConstraintAbout, nil}, parseTime("8 Mar 1505 23"),
		"Abt. 8 Mar 1505",
	},
	"ca. 8 Mar 1505": {
		gedcom.Date{8, time.March, 1505, false, gedcom.DateConstraintAbout, nil}, parseTime("8 Mar 1505 00"),
		gedcom.Date{8, time.March, 1505, true, gedcom.DateConstraintAbout, nil}, parseTime("8 Mar 1505 23"),
		"Abt. 8 Mar 1505",
	},
	"CA 8 Mar 1505": {
		gedcom.Date{8, time.March, 1505, false, gedcom.DateConstraintAbout, nil}, parseTime("8 Mar 1505 00"),
		gedcom.Date{8, time.March, 1505, true, gedcom.DateConstraintAbout, nil}, parseTime("8 Mar 1505 23"),
		"Abt. 8 Mar 1505",
	},
	"cca. 8 Mar 1505": {
		gedcom.Date{8, time.March, 1505, false, gedcom.DateConstraintAbout, nil}, parseTime("8 Mar 1505 00"),
		gedcom.Date{8, time.March, 1505, true, gedcom.DateConstraintAbout, nil}, parseTime("8 Mar 1505 23"),
		"Abt. 8 Mar 1505",
	},
	"Cca 8 Mar 1505": {
		gedcom.Date{8, time.March, 1505, false, gedcom.DateConstraintAbout, nil}, parseTime("8 Mar 1505 00"),
		gedcom.Date{8, time.March, 1505, true, gedcom.DateConstraintAbout, nil}, parseTime("8 Mar 1505 23"),
		"Abt. 8 Mar 1505",
	},
	"circa 21 Dec 1884": {
		gedcom.Date{21, time.December, 1884, false, gedcom.DateConstraintAbout, nil}, parseTime("21 Dec 1884 00"),
		gedcom.Date{21, time.December, 1884, true, gedcom.DateConstraintAbout, nil}, parseTime("21 Dec 1884 23"),
		"Abt. 21 Dec 1884",
	},

	// Invalid dates.
	"25 D 1901": {
		gedcom.Date{0, 0, 0, false, gedcom.DateConstraintExact, errors.New(`parsing time "25 0 1901": month out of range`)}, time.Time{},
		gedcom.Date{0, 0, 0, true, gedcom.DateConstraintExact, errors.New(`parsing time "25 0 1901": month out of range`)}, time.Time{},
		"",
	},
	"5 Decmbr 1901": {
		gedcom.Date{0, 0, 0, false, gedcom.DateConstraintExact, errors.New(`parsing time "5 0 1901": month out of range`)}, time.Time{},
		gedcom.Date{0, 0, 0, true, gedcom.DateConstraintExact, errors.New(`parsing time "5 0 1901": month out of range`)}, time.Time{},
		"",
	},
	"13 Jan": {
		gedcom.Date{0, 0, 0, false, gedcom.DateConstraintExact, errors.New("unable to parse date: 13 Jan")}, time.Time{},
		gedcom.Date{0, 0, 0, true, gedcom.DateConstraintExact, errors.New("unable to parse date: 13 Jan")}, time.Time{},
		"",
	},
	"73 November 1992": {
		gedcom.Date{0, 0, 0, false, gedcom.DateConstraintExact, errors.New(`parsing time "73 11 1992": day out of range`)}, time.Time{},
		gedcom.Date{0, 0, 0, true, gedcom.DateConstraintExact, errors.New(`parsing time "73 11 1992": day out of range`)}, time.Time{},
		"",
	},
	"31 Feb 1992": {
		gedcom.Date{0, 0, 0, false, gedcom.DateConstraintExact, errors.New(`parsing time "31 2 1992": day out of range`)}, time.Time{},
		gedcom.Date{0, 0, 0, true, gedcom.DateConstraintExact, errors.New(`parsing time "31 2 1992": day out of range`)}, time.Time{},
		"",
	},
	"3 Febuary 1992": {
		gedcom.Date{0, 0, 0, false, gedcom.DateConstraintExact, errors.New(`parsing time "3 0 1992": month out of range`)}, time.Time{},
		gedcom.Date{0, 0, 0, true, gedcom.DateConstraintExact, errors.New(`parsing time "3 0 1992": month out of range`)}, time.Time{},
		"",
	},

	// Date ranges.
	"Bet 29 August 1640 and 19 Feb 1992": {
		gedcom.Date{29, time.August, 1640, false, gedcom.DateConstraintExact, nil}, parseTime("29 Aug 1640 00"),
		gedcom.Date{19, time.February, 1992, true, gedcom.DateConstraintExact, nil}, parseTime("19 Feb 1992 23"),
		"Bet. 29 Aug 1640 and 19 Feb 1992",
	},
	"Between July 2003 and 7 Dec 2020": {
		gedcom.Date{0, time.July, 2003, false, gedcom.DateConstraintExact, nil}, parseTime("1 Jul 2003 00"),
		gedcom.Date{7, time.December, 2020, true, gedcom.DateConstraintExact, nil}, parseTime("7 Dec 2020 23"),
		"Bet. Jul 2003 and 7 Dec 2020",
	},
	"Bet. 29 August 1640 AND 19 Feb 1992": {
		gedcom.Date{29, time.August, 1640, false, gedcom.DateConstraintExact, nil}, parseTime("29 Aug 1640 00"),
		gedcom.Date{19, time.February, 1992, true, gedcom.DateConstraintExact, nil}, parseTime("19 Feb 1992 23"),
		"Bet. 29 Aug 1640 and 19 Feb 1992",
	},
	"from 29 August 1640 to 19 Feb 1992": {
		gedcom.Date{29, time.August, 1640, false, gedcom.DateConstraintExact, nil}, parseTime("29 Aug 1640 00"),
		gedcom.Date{19, time.February, 1992, true, gedcom.DateConstraintExact, nil}, parseTime("19 Feb 1992 23"),
		"Bet. 29 Aug 1640 and 19 Feb 1992",
	},
	"FROM 29 August 1640 - 19 Feb 1992": {
		gedcom.Date{29, time.August, 1640, false, gedcom.DateConstraintExact, nil}, parseTime("29 Aug 1640 00"),
		gedcom.Date{19, time.February, 1992, true, gedcom.DateConstraintExact, nil}, parseTime("19 Feb 1992 23"),
		"Bet. 29 Aug 1640 and 19 Feb 1992",
	},

	// Edge cases.
	"foo circa 21 Dec 1884": {
		gedcom.Date{0, 0, 0, false, gedcom.DateConstraintExact, errors.New("unable to parse date: foo circa 21 Dec 1884")}, time.Time{},
		gedcom.Date{0, 0, 0, true, gedcom.DateConstraintExact, errors.New("unable to parse date: foo circa 21 Dec 1884")}, time.Time{},
		"",
	},
	"About 21 Dec 1884 never": {
		gedcom.Date{0, 0, 0, false, gedcom.DateConstraintExact, errors.New("unable to parse date: About 21 Dec 1884 never")}, time.Time{},
		gedcom.Date{0, 0, 0, true, gedcom.DateConstraintExact, errors.New("unable to parse date: About 21 Dec 1884 never")}, time.Time{},
		"",
	},

	// Extreme dates. These are not supported because all dates have to be
	// compatible with Go's time package which has hard limits on the year
	// 0 - 9999.
	//
	//   "3 Mar -1500"
	//   "17 Feb 17329"
}

func TestDateNode_StartDate(t *testing.T) {
	for date, test := range dateTests {
		t.Run(date, func(t *testing.T) {
			node := gedcom.NewDateNode(date)

			assertEqual(t, node.StartDate(), test.startDate)
		})
	}

	StartDate := tf.Function(t, (*gedcom.DateNode).StartDate)

	StartDate((*gedcom.DateNode)(nil)).Returns(gedcom.Date{})
}

func TestDateNode_EndDate(t *testing.T) {
	for date, test := range dateTests {
		t.Run(date, func(t *testing.T) {
			node := gedcom.NewDateNode(date)

			assertEqual(t, node.EndDate(), test.endDate)
		})
	}

	EndDate := tf.Function(t, (*gedcom.DateNode).EndDate)

	EndDate((*gedcom.DateNode)(nil)).Returns(gedcom.Date{})
}

func TestDateNode_String(t *testing.T) {
	for date, test := range dateTests {
		t.Run(date, func(t *testing.T) {
			node := gedcom.NewDateNode(date)

			assert.Equalf(t, test.str, node.String(), "%#+v", date)
		})
	}

	String := tf.Function(t, (*gedcom.DateNode).String)

	String((*gedcom.DateNode)(nil)).Returns("")
}

func TestDateNode_Years(t *testing.T) {
	tests := []struct {
		date     *gedcom.DateNode
		expected float64
	}{
		// Nil
		{nil, 0.0},

		// Zero
		{gedcom.NewDateNode(""), 0.0},

		// Year
		{gedcom.NewDateNode("750"), 750.5},
		{gedcom.NewDateNode("1845"), 1845.5},

		// Months
		{gedcom.NewDateNode("Jan 1845"), 1845.0437158469945},
		{gedcom.NewDateNode("Mar 1999"), 1999.204918032787},
		{gedcom.NewDateNode("Dec 1832"), 1832.956403269755},

		// Days
		{gedcom.NewDateNode("1 Jan 1789"), 1789.0027322404371},
		{gedcom.NewDateNode("31 Jan 1435"), 1435.0846994535518},
		{gedcom.NewDateNode("1 Feb 1601"), 1601.0874316939892},
		{gedcom.NewDateNode("1 Mar 845"), 845.1639344262295},
		{gedcom.NewDateNode("31 Dec 2010"), 2010.9972677595629},

		// Ranges
		{
			gedcom.NewDateNode("Bet. 1 Jan 1789 and 1 Mar 1789"),
			1789.0833333333335,
		},
		{
			gedcom.NewDateNode("Bet. 1 Jan 1789 and 1 Jan 1789"),
			// Same as "1 Jan 1789"
			1789.0027322404371,
		},
		{
			gedcom.NewDateNode("Bet. 1430 and 1435"),
			// From the start of 1430 to the end of 1435 is actually 6 years.
			1433,
		},

		// Invalid
		{gedcom.NewDateNode("Foo"), 0},
	}

	for _, test := range tests {
		t.Run(gedcom.Value(test.date), func(t *testing.T) {
			assert.Equal(t, test.expected, test.date.Years())
		})
	}
}

func TestDateNode_Similarity(t *testing.T) {
	tests := []struct {
		date1    *gedcom.DateNode
		date2    *gedcom.DateNode
		expected float64
	}{
		// Two unknown dates will be equal to each other.
		{
			gedcom.NewDateNode(""),
			gedcom.NewDateNode(""),
			1,
		},

		// The difference will be same regardless of time line so the two next
		// tests must return the same similarity.
		{
			gedcom.NewDateNode("500"),
			gedcom.NewDateNode("502"),
			0.96,
		},
		{
			gedcom.NewDateNode("2000"),
			gedcom.NewDateNode("2002"),
			0.96,
		},

		// A higher score is awarded to values that are closer to each other.
		{
			gedcom.NewDateNode("1900"),
			gedcom.NewDateNode("1901"),
			0.99,
		},
		{
			gedcom.NewDateNode("1900"),
			gedcom.NewDateNode("1904"),
			0.84,
		},

		// Months
		{
			gedcom.NewDateNode("Feb 2000"),
			gedcom.NewDateNode("Mar 2000"),
			0.9999331793984663,
		},
		{
			gedcom.NewDateNode("Feb 2000"),
			gedcom.NewDateNode("Feb 2001"),
			0.9900204627124954,
		},

		// Days
		{
			gedcom.NewDateNode("13 Feb 2000"),
			gedcom.NewDateNode("14 Feb 2000"),
			0.9999999257548872,
		},
		{
			gedcom.NewDateNode("13 Feb 2000"),
			gedcom.NewDateNode("13 Apr 2000"),
			0.9997327175938642,
		},

		// Exact matches
		{
			gedcom.NewDateNode("2000"),
			gedcom.NewDateNode("2000"),
			1,
		},
		{
			gedcom.NewDateNode("Mar 2000"),
			gedcom.NewDateNode("Mar 2000"),
			1,
		},
		{
			gedcom.NewDateNode("13 Mar 2000"),
			gedcom.NewDateNode("13 Mar 2000"),
			1,
		},
		{
			gedcom.NewDateNode("Bet. 2000 and 2003"),
			gedcom.NewDateNode("Between 2000 and 2003"),
			1,
		},
		{
			gedcom.NewDateNode("Bet. Mar 2000 and Oct 2000"),
			gedcom.NewDateNode("Bet. Mar 2000 and Oct 2000"),
			1,
		},
		{
			gedcom.NewDateNode("bet. 13 Mar 2000 and 17 March 2000"),
			gedcom.NewDateNode("Between 13 Mar 2000 and 17 March 2000"),
			1,
		},

		// These ranges are inverse so they have the same difference.
		{
			gedcom.NewDateNode("Bet. 2000 and 2003"),
			gedcom.NewDateNode("Between 2001 and 2003"),
			0.9975,
		},
		{
			gedcom.NewDateNode("Bet. 2001 and 2003"),
			gedcom.NewDateNode("Between 2000 and 2003"),
			0.9975,
		},

		// Range has the same difference but over different time periods.
		{
			gedcom.NewDateNode("Bet. 2000 and 2003"),
			gedcom.NewDateNode("Between 1997 and 2000"),
			0.91,
		},

		// Other ranges.
		{
			gedcom.NewDateNode("Bet. Mar 2000 and Oct 2000"),
			gedcom.NewDateNode("Bet. Feb 2000 and Oct 2000"),
			0.9999832948496166,
		},
		{
			gedcom.NewDateNode("bet. 15 Mar 2000 and 23 March 2000"),
			gedcom.NewDateNode("Between 15 Mar 2000 and 25 March 2000"),
			0.9999999257548872,
		},

		// Invalid
		{
			gedcom.NewDateNode("Foo"),
			gedcom.NewDateNode("13 Mar 2000"),
			0,
		},
		{
			gedcom.NewDateNode("13 Mar 2000"),
			gedcom.NewDateNode("Bar"),
			0,
		},

		// Nil cases
		{
			nil,
			gedcom.NewDateNode("Jan 1845"),
			0.5,
		},
		{
			gedcom.NewDateNode("Jan 1845"),
			nil,
			0.5,
		},
		{
			nil,
			nil,
			0.5,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			similarity := test.date1.Similarity(test.date2, 10)
			assert.Equal(t, test.expected, similarity)
		})
	}
}

func TestDateNode_Equals(t *testing.T) {
	// d1 and d2 are the same value.
	d1 := gedcom.NewDateNode("15 SEP 1985")
	d2 := gedcom.NewDateNode("15 September 1985")

	// d3 and d4 represent the same enclosed ranges.
	d3 := gedcom.NewDateNode("Bet. Oct 2000 and 3 Apr 2008")
	d4 := gedcom.NewDateNode("From OCT 2000 to Bef. Jun 2008")

	// d5 has a different Start from d3 and d4.
	d5 := gedcom.NewDateNode("From Jun 2000 to 3 Apr 2008")

	// d6 to d8 are phrases.
	d6 := gedcom.NewDateNode("(Foo)")
	d7 := gedcom.NewDateNode("(15 SEP 1985)")
	d8 := gedcom.NewDateNode("(Foo)")

	// d9 to d11 are invalid.
	d9 := gedcom.NewDateNode("@#DJULIAN@ 01 JAN 1323")
	d10 := gedcom.NewDateNode("!! 15 SEP 1985")
	d11 := gedcom.NewDateNode("@#DJULIAN@ 01 JAN 1323")

	Equals := tf.Function(t, (*gedcom.DateNode).Equals)

	// nil values
	Equals((*gedcom.DateNode)(nil), d1).Returns(false)
	Equals(d1, (*gedcom.DateNode)(nil)).Returns(false)
	Equals((*gedcom.DateNode)(nil), (*gedcom.DateNode)(nil)).Returns(false)
	Equals(d6, (*gedcom.DateNode)(nil)).Returns(false)

	// Bad input
	Equals(d1, gedcom.NewNameNode("15 SEP 1985")).Returns(false)
	Equals(d6, gedcom.NewNameNode("15 SEP 1985")).Returns(false)

	// General cases.
	Equals(d1, d2).Returns(true)
	Equals(d2, d1).Returns(true)
	Equals(d3, d4).Returns(true)
	Equals(d4, d5).Returns(false)
	Equals(d1, d10).Returns(false)

	// Nothing equals a phrase except anther phrase that is exactly the same.
	Equals(d1, d6).Returns(false)
	Equals(d6, d1).Returns(false)
	Equals(d1, d7).Returns(false)
	Equals(d7, d1).Returns(false)
	Equals(d6, d7).Returns(false)
	Equals(d7, d6).Returns(false)
	Equals(d6, d8).Returns(true)
	Equals(d8, d6).Returns(true)

	// Invalid dates are only equal if they have exactly the same value.
	Equals(d9, d10).Returns(false)
	Equals(d10, d9).Returns(false)
	Equals(d9, d11).Returns(true)
	Equals(d11, d9).Returns(true)
}

func TestDateNode_StartAndEndDates(t *testing.T) {
	StartAndEndDates := tf.Function(t, (*gedcom.DateNode).StartAndEndDates)

	StartAndEndDates((*gedcom.DateNode)(nil)).Returns(gedcom.Date{}, gedcom.Date{})
}

func TestDateNode_IsPhrase(t *testing.T) {
	IsPhrase := tf.Function(t, (*gedcom.DateNode).IsPhrase)

	IsPhrase((*gedcom.DateNode)(nil)).Returns(false)
	IsPhrase(gedcom.NewDateNode("")).Returns(false)
	IsPhrase(gedcom.NewDateNode("(")).Returns(false)
	IsPhrase(gedcom.NewDateNode(")")).Returns(false)
	IsPhrase(gedcom.NewDateNode("3 Mar 1981")).Returns(false)

	IsPhrase(gedcom.NewDateNode("()")).Returns(true)
	IsPhrase(gedcom.NewDateNode("(Foo)")).Returns(true)
	IsPhrase(gedcom.NewDateNode("(Foo BAR)")).Returns(true)
}

var dateWarningTests = map[string]struct {
	date     *gedcom.DateNode
	expected []string
}{
	"Valid1": {
		gedcom.NewDateNode("3 Sep 1943"),
		nil,
	},
	"Valid2": {
		gedcom.NewDateNode("  3 Sep 1943 "),
		nil,
	},
	"Unparsable1": {
		gedcom.NewDateNode("foo bar"),
		[]string{
			`Unparsable date "foo bar"`,
		},
	},
	"Unparsable2": {
		gedcom.NewDateNode("abt. world war 2"),
		[]string{
			`Unparsable date "abt. world war 2"`,
		},
	},
}

func TestDateNode_Warnings(t *testing.T) {
	for testName, test := range dateWarningTests {
		t.Run(testName, func(t *testing.T) {
			assertEqual(t, test.date.Warnings().Strings(), test.expected)
		})
	}
}
