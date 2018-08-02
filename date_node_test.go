package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
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
		gedcom.Date{1, time.January, 1980, false, gedcom.DateConstraintExact}, parseTime("1 Jan 1980 00"),
		gedcom.Date{1, time.January, 1980, true, gedcom.DateConstraintExact}, parseTime("1 Jan 1980 23"),
		"1 Jan 1980",
	},
	"15 Feb 1880": {
		gedcom.Date{15, time.February, 1880, false, gedcom.DateConstraintExact}, parseTime("15 Feb 1880 00"),
		gedcom.Date{15, time.February, 1880, true, gedcom.DateConstraintExact}, parseTime("15 Feb 1880 23"),
		"15 Feb 1880",
	},
	"03 Mar 1870": {
		gedcom.Date{3, time.March, 1870, false, gedcom.DateConstraintExact}, parseTime("3 Mar 1870 00"),
		gedcom.Date{3, time.March, 1870, true, gedcom.DateConstraintExact}, parseTime("3 Mar 1870 23"),
		"3 Mar 1870",
	},
	"7 Apr 2020": {
		gedcom.Date{7, time.April, 2020, false, gedcom.DateConstraintExact}, parseTime("7 Apr 2020 00"),
		gedcom.Date{7, time.April, 2020, true, gedcom.DateConstraintExact}, parseTime("7 Apr 2020 23"),
		"7 Apr 2020",
	},
	"6 May 1989": {
		gedcom.Date{6, time.May, 1989, false, gedcom.DateConstraintExact}, parseTime("6 May 1989 00"),
		gedcom.Date{6, time.May, 1989, true, gedcom.DateConstraintExact}, parseTime("6 May 1989 23"),
		"6 May 1989",
	},
	"8 Jun 2001": {
		gedcom.Date{8, time.June, 2001, false, gedcom.DateConstraintExact}, parseTime("8 Jun 2001 00"),
		gedcom.Date{8, time.June, 2001, true, gedcom.DateConstraintExact}, parseTime("8 Jun 2001 23"),
		"8 Jun 2001",
	},
	"19 Jul 2003": {
		gedcom.Date{19, time.July, 2003, false, gedcom.DateConstraintExact}, parseTime("19 Jul 2003 00"),
		gedcom.Date{19, time.July, 2003, true, gedcom.DateConstraintExact}, parseTime("19 Jul 2003 23"),
		"19 Jul 2003",
	},
	"29 Aug 1640": {
		gedcom.Date{29, time.August, 1640, false, gedcom.DateConstraintExact}, parseTime("29 Aug 1640 00"),
		gedcom.Date{29, time.August, 1640, true, gedcom.DateConstraintExact}, parseTime("29 Aug 1640 23"),
		"29 Aug 1640",
	},
	"13 Sep 1733": {
		gedcom.Date{13, time.September, 1733, false, gedcom.DateConstraintExact}, parseTime("13 Sep 1733 00"),
		gedcom.Date{13, time.September, 1733, true, gedcom.DateConstraintExact}, parseTime("13 Sep 1733 23"),
		"13 Sep 1733",
	},
	"6 Oct 1848": {
		gedcom.Date{6, time.October, 1848, false, gedcom.DateConstraintExact}, parseTime("6 Oct 1848 00"),
		gedcom.Date{6, time.October, 1848, true, gedcom.DateConstraintExact}, parseTime("6 Oct 1848 23"),
		"6 Oct 1848",
	},
	"18 Nov 1992": {
		gedcom.Date{18, time.November, 1992, false, gedcom.DateConstraintExact}, parseTime("18 Nov 1992 00"),
		gedcom.Date{18, time.November, 1992, true, gedcom.DateConstraintExact}, parseTime("18 Nov 1992 23"),
		"18 Nov 1992",
	},
	"25 Dec 1901": {
		gedcom.Date{25, time.December, 1901, false, gedcom.DateConstraintExact}, parseTime("25 Dec 1901 00"),
		gedcom.Date{25, time.December, 1901, true, gedcom.DateConstraintExact}, parseTime("25 Dec 1901 23"),
		"25 Dec 1901",
	},

	// Valid dates, testing each full month name. The days are a mix of dd
	// and d.
	"01 January 1980": {
		gedcom.Date{1, time.January, 1980, false, gedcom.DateConstraintExact}, parseTime("1 Jan 1980 00"),
		gedcom.Date{1, time.January, 1980, true, gedcom.DateConstraintExact}, parseTime("1 Jan 1980 23"),
		"1 Jan 1980",
	},
	"15 February 1880": {
		gedcom.Date{15, time.February, 1880, false, gedcom.DateConstraintExact}, parseTime("15 Feb 1880 00"),
		gedcom.Date{15, time.February, 1880, true, gedcom.DateConstraintExact}, parseTime("15 Feb 1880 23"),
		"15 Feb 1880",
	},
	"03 March 1870": {
		gedcom.Date{3, time.March, 1870, false, gedcom.DateConstraintExact}, parseTime("3 Mar 1870 00"),
		gedcom.Date{3, time.March, 1870, true, gedcom.DateConstraintExact}, parseTime("3 Mar 1870 23"),
		"3 Mar 1870",
	},
	"7 April 2020": {
		gedcom.Date{7, time.April, 2020, false, gedcom.DateConstraintExact}, parseTime("7 Apr 2020 00"),
		gedcom.Date{7, time.April, 2020, true, gedcom.DateConstraintExact}, parseTime("7 Apr 2020 23"),
		"7 Apr 2020",
	},
	"8 June 2001": {
		gedcom.Date{8, time.June, 2001, false, gedcom.DateConstraintExact}, parseTime("8 Jun 2001 00"),
		gedcom.Date{8, time.June, 2001, true, gedcom.DateConstraintExact}, parseTime("8 Jun 2001 23"),
		"8 Jun 2001",
	},
	"19 July 2003": {
		gedcom.Date{19, time.July, 2003, false, gedcom.DateConstraintExact}, parseTime("19 Jul 2003 00"),
		gedcom.Date{19, time.July, 2003, true, gedcom.DateConstraintExact}, parseTime("19 Jul 2003 23"),
		"19 Jul 2003",
	},
	"29 August 1640": {
		gedcom.Date{29, time.August, 1640, false, gedcom.DateConstraintExact}, parseTime("29 Aug 1640 00"),
		gedcom.Date{29, time.August, 1640, true, gedcom.DateConstraintExact}, parseTime("29 Aug 1640 23"),
		"29 Aug 1640",
	},
	"13 September 1733": {
		gedcom.Date{13, time.September, 1733, false, gedcom.DateConstraintExact}, parseTime("13 Sep 1733 00"),
		gedcom.Date{13, time.September, 1733, true, gedcom.DateConstraintExact}, parseTime("13 Sep 1733 23"),
		"13 Sep 1733",
	},
	"6 October 1848": {
		gedcom.Date{6, time.October, 1848, false, gedcom.DateConstraintExact}, parseTime("6 Oct 1848 00"),
		gedcom.Date{6, time.October, 1848, true, gedcom.DateConstraintExact}, parseTime("6 Oct 1848 23"),
		"6 Oct 1848",
	},
	"18 November 1992": {
		gedcom.Date{18, time.November, 1992, false, gedcom.DateConstraintExact}, parseTime("18 Nov 1992 00"),
		gedcom.Date{18, time.November, 1992, true, gedcom.DateConstraintExact}, parseTime("18 Nov 1992 23"),
		"18 Nov 1992",
	},
	"25 December 1901": {
		gedcom.Date{25, time.December, 1901, false, gedcom.DateConstraintExact}, parseTime("25 Dec 1901 00"),
		gedcom.Date{25, time.December, 1901, true, gedcom.DateConstraintExact}, parseTime("25 Dec 1901 23"),
		"25 Dec 1901",
	},

	// Only month and year combinations.
	"Jan 1980": {
		gedcom.Date{0, time.January, 1980, false, gedcom.DateConstraintExact}, parseTime("1 Jan 1980 00"),
		gedcom.Date{0, time.January, 1980, true, gedcom.DateConstraintExact}, parseTime("31 Jan 1980 23"),
		"Jan 1980",
	},
	"Feb 1880": {
		gedcom.Date{0, time.February, 1880, false, gedcom.DateConstraintExact}, parseTime("1 Feb 1880 00"),
		gedcom.Date{0, time.February, 1880, true, gedcom.DateConstraintExact}, parseTime("29 Feb 1880 23"),
		"Feb 1880",
	},
	"Mar 1870": {
		gedcom.Date{0, time.March, 1870, false, gedcom.DateConstraintExact}, parseTime("1 Mar 1870 00"),
		gedcom.Date{0, time.March, 1870, true, gedcom.DateConstraintExact}, parseTime("31 Mar 1870 23"),
		"Mar 1870",
	},
	"Apr 2020": {
		gedcom.Date{0, time.April, 2020, false, gedcom.DateConstraintExact}, parseTime("1 Apr 2020 00"),
		gedcom.Date{0, time.April, 2020, true, gedcom.DateConstraintExact}, parseTime("30 Apr 2020 23"),
		"Apr 2020",
	},
	"May 1989": {
		gedcom.Date{0, time.May, 1989, false, gedcom.DateConstraintExact}, parseTime("1 May 1989 00"),
		gedcom.Date{0, time.May, 1989, true, gedcom.DateConstraintExact}, parseTime("31 May 1989 23"),
		"May 1989",
	},
	"Jun 2001": {
		gedcom.Date{0, time.June, 2001, false, gedcom.DateConstraintExact}, parseTime("1 Jun 2001 00"),
		gedcom.Date{0, time.June, 2001, true, gedcom.DateConstraintExact}, parseTime("30 Jun 2001 23"),
		"Jun 2001",
	},
	"Jul 2003": {
		gedcom.Date{0, time.July, 2003, false, gedcom.DateConstraintExact}, parseTime("1 Jul 2003 00"),
		gedcom.Date{0, time.July, 2003, true, gedcom.DateConstraintExact}, parseTime("31 Jul 2003 23"),
		"Jul 2003",
	},
	"Aug 1640": {
		gedcom.Date{0, time.August, 1640, false, gedcom.DateConstraintExact}, parseTime("1 Aug 1640 00"),
		gedcom.Date{0, time.August, 1640, true, gedcom.DateConstraintExact}, parseTime("31 Aug 1640 23"),
		"Aug 1640",
	},
	"Sep 1733": {
		gedcom.Date{0, time.September, 1733, false, gedcom.DateConstraintExact}, parseTime("1 Sep 1733 00"),
		gedcom.Date{0, time.September, 1733, true, gedcom.DateConstraintExact}, parseTime("30 Sep 1733 23"),
		"Sep 1733",
	},
	"Oct 1848": {
		gedcom.Date{0, time.October, 1848, false, gedcom.DateConstraintExact}, parseTime("1 Oct 1848 00"),
		gedcom.Date{0, time.October, 1848, true, gedcom.DateConstraintExact}, parseTime("31 Oct 1848 23"),
		"Oct 1848",
	},
	"Nov 1992": {
		gedcom.Date{0, time.November, 1992, false, gedcom.DateConstraintExact}, parseTime("1 Nov 1992 00"),
		gedcom.Date{0, time.November, 1992, true, gedcom.DateConstraintExact}, parseTime("30 Nov 1992 23"),
		"Nov 1992",
	},
	"Dec 1901": {
		gedcom.Date{0, time.December, 1901, false, gedcom.DateConstraintExact}, parseTime("1 Dec 1901 00"),
		gedcom.Date{0, time.December, 1901, true, gedcom.DateConstraintExact}, parseTime("31 Dec 1901 23"),
		"Dec 1901",
	},
	"January 1980": {
		gedcom.Date{0, time.January, 1980, false, gedcom.DateConstraintExact}, parseTime("1 Jan 1980 00"),
		gedcom.Date{0, time.January, 1980, true, gedcom.DateConstraintExact}, parseTime("31 Jan 1980 23"),
		"Jan 1980",
	},
	"February 1880": {
		gedcom.Date{0, time.February, 1880, false, gedcom.DateConstraintExact}, parseTime("1 Feb 1880 00"),
		gedcom.Date{0, time.February, 1880, true, gedcom.DateConstraintExact}, parseTime("29 Feb 1880 23"),
		"Feb 1880",
	},
	"March 1870": {
		gedcom.Date{0, time.March, 1870, false, gedcom.DateConstraintExact}, parseTime("1 Mar 1870 00"),
		gedcom.Date{0, time.March, 1870, true, gedcom.DateConstraintExact}, parseTime("31 Mar 1870 23"),
		"Mar 1870",
	},
	"April 2020": {
		gedcom.Date{0, time.April, 2020, false, gedcom.DateConstraintExact}, parseTime("1 Apr 2020 00"),
		gedcom.Date{0, time.April, 2020, true, gedcom.DateConstraintExact}, parseTime("30 Apr 2020 23"),
		"Apr 2020",
	},
	"June 2001": {
		gedcom.Date{0, time.June, 2001, false, gedcom.DateConstraintExact}, parseTime("1 Jun 2001 00"),
		gedcom.Date{0, time.June, 2001, true, gedcom.DateConstraintExact}, parseTime("30 Jun 2001 23"),
		"Jun 2001",
	},
	"July 2003": {
		gedcom.Date{0, time.July, 2003, false, gedcom.DateConstraintExact}, parseTime("1 Jul 2003 00"),
		gedcom.Date{0, time.July, 2003, true, gedcom.DateConstraintExact}, parseTime("31 Jul 2003 23"),
		"Jul 2003",
	},
	"August 1640": {
		gedcom.Date{0, time.August, 1640, false, gedcom.DateConstraintExact}, parseTime("1 Aug 1640 00"),
		gedcom.Date{0, time.August, 1640, true, gedcom.DateConstraintExact}, parseTime("31 Aug 1640 23"),
		"Aug 1640",
	},
	"September 1733": {
		gedcom.Date{0, time.September, 1733, false, gedcom.DateConstraintExact}, parseTime("1 Sep 1733 00"),
		gedcom.Date{0, time.September, 1733, true, gedcom.DateConstraintExact}, parseTime("30 Sep 1733 23"),
		"Sep 1733",
	},
	"October 1848": {
		gedcom.Date{0, time.October, 1848, false, gedcom.DateConstraintExact}, parseTime("1 Oct 1848 00"),
		gedcom.Date{0, time.October, 1848, true, gedcom.DateConstraintExact}, parseTime("31 Oct 1848 23"),
		"Oct 1848",
	},
	"November 1992": {
		gedcom.Date{0, time.November, 1992, false, gedcom.DateConstraintExact}, parseTime("1 Nov 1992 00"),
		gedcom.Date{0, time.November, 1992, true, gedcom.DateConstraintExact}, parseTime("30 Nov 1992 23"),
		"Nov 1992",
	},
	"December 1901": {
		gedcom.Date{0, time.December, 1901, false, gedcom.DateConstraintExact}, parseTime("1 Dec 1901 00"),
		gedcom.Date{0, time.December, 1901, true, gedcom.DateConstraintExact}, parseTime("31 Dec 1901 23"),
		"Dec 1901",
	},

	// Months with different capitalization.
	"DECEMBER 1901": {
		gedcom.Date{0, time.December, 1901, false, gedcom.DateConstraintExact}, parseTime("1 Dec 1901 00"),
		gedcom.Date{0, time.December, 1901, true, gedcom.DateConstraintExact}, parseTime("31 Dec 1901 23"),
		"Dec 1901",
	},
	"13 SEP 1733": {
		gedcom.Date{13, time.September, 1733, false, gedcom.DateConstraintExact}, parseTime("13 Sep 1733 00"),
		gedcom.Date{13, time.September, 1733, true, gedcom.DateConstraintExact}, parseTime("13 Sep 1733 23"),
		"13 Sep 1733",
	},

	// Only year.
	"834": {
		gedcom.Date{0, 0, 834, false, gedcom.DateConstraintExact}, parseTime("1 Jan 0834 00"),
		gedcom.Date{0, 0, 834, true, gedcom.DateConstraintExact}, parseTime("31 Dec 0834 23"),
		"834",
	},
	"0834": {
		gedcom.Date{0, 0, 834, false, gedcom.DateConstraintExact}, parseTime("1 Jan 0834 00"),
		gedcom.Date{0, 0, 834, true, gedcom.DateConstraintExact}, parseTime("31 Dec 0834 23"),
		"834",
	},
	"1901": {
		gedcom.Date{0, 0, 1901, false, gedcom.DateConstraintExact}, parseTime("1 Jan 1901 00"),
		gedcom.Date{0, 0, 1901, true, gedcom.DateConstraintExact}, parseTime("31 Dec 1901 23"),
		"1901",
	},
	"2020": {
		gedcom.Date{0, 0, 2020, false, gedcom.DateConstraintExact}, parseTime("1 Jan 2020 00"),
		gedcom.Date{0, 0, 2020, true, gedcom.DateConstraintExact}, parseTime("31 Dec 2020 23"),
		"2020",
	},
	"0066": {
		gedcom.Date{0, 0, 66, false, gedcom.DateConstraintExact}, parseTime("1 Jan 0066 00"),
		gedcom.Date{0, 0, 66, true, gedcom.DateConstraintExact}, parseTime("31 Dec 0066 23"),
		"66",
	},

	// Extra whitespace. The GEDCOM file should not allow values to contain new
	// lines or carriage returns in the node value so we do not need to test
	// those cases.
	"  18 November 1992": {
		gedcom.Date{18, time.November, 1992, false, gedcom.DateConstraintExact}, parseTime("18 Nov 1992 00"),
		gedcom.Date{18, time.November, 1992, true, gedcom.DateConstraintExact}, parseTime("18 Nov 1992 23"),
		"18 Nov 1992",
	},
	"15 Feb   1880": {
		gedcom.Date{15, time.February, 1880, false, gedcom.DateConstraintExact}, parseTime("15 Feb 1880 00"),
		gedcom.Date{15, time.February, 1880, true, gedcom.DateConstraintExact}, parseTime("15 Feb 1880 23"),
		"15 Feb 1880",
	},
	"Feb   1880": {
		gedcom.Date{0, time.February, 1880, false, gedcom.DateConstraintExact}, parseTime("1 Feb 1880 00"),
		gedcom.Date{0, time.February, 1880, true, gedcom.DateConstraintExact}, parseTime("29 Feb 1880 23"),
		"Feb 1880",
	},
	"25 December 1901  ": {
		gedcom.Date{25, time.December, 1901, false, gedcom.DateConstraintExact}, parseTime("25 Dec 1901 00"),
		gedcom.Date{25, time.December, 1901, true, gedcom.DateConstraintExact}, parseTime("25 Dec 1901 23"),
		"25 Dec 1901",
	},
	" 1901  ": {
		gedcom.Date{0, 0, 1901, false, gedcom.DateConstraintExact}, parseTime("1 Jan 1901 00"),
		gedcom.Date{0, 0, 1901, true, gedcom.DateConstraintExact}, parseTime("31 Dec 1901 23"),
		"1901",
	},

	// Before dates.
	"Before Feb 1907": {
		gedcom.Date{0, time.February, 1907, false, gedcom.DateConstraintBefore}, parseTime("1 Feb 1907 00"),
		gedcom.Date{0, time.February, 1907, true, gedcom.DateConstraintBefore}, parseTime("28 Feb 1907 23"),
		"Bef. Feb 1907",
	},
	"bef. 21 Dec 1884": {
		gedcom.Date{21, time.December, 1884, false, gedcom.DateConstraintBefore}, parseTime("21 Dec 1884 00"),
		gedcom.Date{21, time.December, 1884, true, gedcom.DateConstraintBefore}, parseTime("21 Dec 1884 23"),
		"Bef. 21 Dec 1884",
	},

	// After dates.
	"after Feb 1907": {
		gedcom.Date{0, time.February, 1907, false, gedcom.DateConstraintAfter}, parseTime("1 Feb 1907 00"),
		gedcom.Date{0, time.February, 1907, true, gedcom.DateConstraintAfter}, parseTime("28 Feb 1907 23"),
		"Aft. Feb 1907",
	},
	"Aft. 21 Dec 1884": {
		gedcom.Date{21, time.December, 1884, false, gedcom.DateConstraintAfter}, parseTime("21 Dec 1884 00"),
		gedcom.Date{21, time.December, 1884, true, gedcom.DateConstraintAfter}, parseTime("21 Dec 1884 23"),
		"Aft. 21 Dec 1884",
	},

	// Approximate dates.
	"Abt. 1945": {
		gedcom.Date{0, 0, 1945, false, gedcom.DateConstraintAbout}, parseTime("1 Jan 1945 00"),
		gedcom.Date{0, 0, 1945, true, gedcom.DateConstraintAbout}, parseTime("31 Dec 1945 23"),
		"Abt. 1945",
	},
	"about Feb 1907": {
		gedcom.Date{0, time.February, 1907, false, gedcom.DateConstraintAbout}, parseTime("1 Feb 1907 00"),
		gedcom.Date{0, time.February, 1907, true, gedcom.DateConstraintAbout}, parseTime("28 Feb 1907 23"),
		"Abt. Feb 1907",
	},
	"c. 8 Mar 1505": {
		gedcom.Date{8, time.March, 1505, false, gedcom.DateConstraintAbout}, parseTime("8 Mar 1505 00"),
		gedcom.Date{8, time.March, 1505, true, gedcom.DateConstraintAbout}, parseTime("8 Mar 1505 23"),
		"Abt. 8 Mar 1505",
	},
	"circa 21 Dec 1884": {
		gedcom.Date{21, time.December, 1884, false, gedcom.DateConstraintAbout}, parseTime("21 Dec 1884 00"),
		gedcom.Date{21, time.December, 1884, true, gedcom.DateConstraintAbout}, parseTime("21 Dec 1884 23"),
		"Abt. 21 Dec 1884",
	},

	// Invalid dates.
	"25 D 1901": {
		gedcom.Date{0, 0, 0, false, gedcom.DateConstraintExact}, time.Time{},
		gedcom.Date{0, 0, 0, true, gedcom.DateConstraintExact}, time.Time{},
		"",
	},
	"5 Decmbr 1901": {
		gedcom.Date{0, 0, 0, false, gedcom.DateConstraintExact}, time.Time{},
		gedcom.Date{0, 0, 0, true, gedcom.DateConstraintExact}, time.Time{},
		"",
	},
	"13 Jan": {
		gedcom.Date{0, 0, 0, false, gedcom.DateConstraintExact}, time.Time{},
		gedcom.Date{0, 0, 0, true, gedcom.DateConstraintExact}, time.Time{},
		"",
	},
	"73 November 1992": {
		gedcom.Date{0, 0, 0, false, gedcom.DateConstraintExact}, time.Time{},
		gedcom.Date{0, 0, 0, true, gedcom.DateConstraintExact}, time.Time{},
		"",
	},
	"31 Feb 1992": {
		gedcom.Date{0, 0, 0, false, gedcom.DateConstraintExact}, time.Time{},
		gedcom.Date{0, 0, 0, true, gedcom.DateConstraintExact}, time.Time{},
		"",
	},
	"3 Febuary 1992": {
		gedcom.Date{0, 0, 0, false, gedcom.DateConstraintExact}, time.Time{},
		gedcom.Date{0, 0, 0, true, gedcom.DateConstraintExact}, time.Time{},
		"",
	},

	// Date ranges.
	"Bet 29 August 1640 and 19 Feb 1992": {
		gedcom.Date{29, time.August, 1640, false, gedcom.DateConstraintExact}, parseTime("29 Aug 1640 00"),
		gedcom.Date{19, time.February, 1992, true, gedcom.DateConstraintExact}, parseTime("19 Feb 1992 23"),
		"Bet. 29 Aug 1640 and 19 Feb 1992",
	},
	"Between July 2003 and 7 Dec 2020": {
		gedcom.Date{0, time.July, 2003, false, gedcom.DateConstraintExact}, parseTime("1 Jul 2003 00"),
		gedcom.Date{7, time.December, 2020, true, gedcom.DateConstraintExact}, parseTime("7 Dec 2020 23"),
		"Bet. Jul 2003 and 7 Dec 2020",
	},
	"Bet. 29 August 1640 AND 19 Feb 1992": {
		gedcom.Date{29, time.August, 1640, false, gedcom.DateConstraintExact}, parseTime("29 Aug 1640 00"),
		gedcom.Date{19, time.February, 1992, true, gedcom.DateConstraintExact}, parseTime("19 Feb 1992 23"),
		"Bet. 29 Aug 1640 and 19 Feb 1992",
	},
	"from 29 August 1640 to 19 Feb 1992": {
		gedcom.Date{29, time.August, 1640, false, gedcom.DateConstraintExact}, parseTime("29 Aug 1640 00"),
		gedcom.Date{19, time.February, 1992, true, gedcom.DateConstraintExact}, parseTime("19 Feb 1992 23"),
		"Bet. 29 Aug 1640 and 19 Feb 1992",
	},
	"FROM 29 August 1640 - 19 Feb 1992": {
		gedcom.Date{29, time.August, 1640, false, gedcom.DateConstraintExact}, parseTime("29 Aug 1640 00"),
		gedcom.Date{19, time.February, 1992, true, gedcom.DateConstraintExact}, parseTime("19 Feb 1992 23"),
		"Bet. 29 Aug 1640 and 19 Feb 1992",
	},

	// Edge cases.
	"foo circa 21 Dec 1884": {
		gedcom.Date{0, 0, 0, false, gedcom.DateConstraintExact}, time.Time{},
		gedcom.Date{0, 0, 0, true, gedcom.DateConstraintExact}, time.Time{},
		"",
	},
	"About 21 Dec 1884 never": {
		gedcom.Date{0, 0, 0, false, gedcom.DateConstraintExact}, time.Time{},
		gedcom.Date{0, 0, 0, true, gedcom.DateConstraintExact}, time.Time{},
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
			node := gedcom.NewDateNode(date, "", nil)

			assert.Equal(t, node.StartDate(), test.startDate)
		})
	}
}

func TestDateNode_EndDate(t *testing.T) {
	for date, test := range dateTests {
		t.Run(date, func(t *testing.T) {
			node := gedcom.NewDateNode(date, "", nil)

			assert.Equal(t, node.EndDate(), test.endDate)
		})
	}
}

func TestDateNode_String(t *testing.T) {
	for date, test := range dateTests {
		t.Run(date, func(t *testing.T) {
			node := gedcom.NewDateNode(date, "", nil)

			assert.Equalf(t, test.str, node.String(), "%#+v", date)
		})
	}
}

func TestDateNode_Years(t *testing.T) {
	tests := []struct {
		date     *gedcom.DateNode
		expected float64
	}{
		// Zero
		{gedcom.NewDateNode("", "", nil), 0.0},

		// Year
		{gedcom.NewDateNode("750", "", nil), 750.5},
		{gedcom.NewDateNode("1845", "", nil), 1845.5},

		// Months
		{gedcom.NewDateNode("Jan 1845", "", nil), 1845.0437158469945},
		{gedcom.NewDateNode("Mar 1999", "", nil), 1999.204918032787},
		{gedcom.NewDateNode("Dec 1832", "", nil), 1832.956403269755},

		// Days
		{gedcom.NewDateNode("1 Jan 1789", "", nil), 1789.0027322404371},
		{gedcom.NewDateNode("31 Jan 1435", "", nil), 1435.0846994535518},
		{gedcom.NewDateNode("1 Feb 1601", "", nil), 1601.0874316939892},
		{gedcom.NewDateNode("1 Mar 845", "", nil), 845.1639344262295},
		{gedcom.NewDateNode("31 Dec 2010", "", nil), 2010.9972677595629},

		// Ranges
		{
			gedcom.NewDateNode("Bet. 1 Jan 1789 and 1 Mar 1789", "", nil),
			1789.0833333333335,
		},
		{
			gedcom.NewDateNode("Bet. 1 Jan 1789 and 1 Jan 1789", "", nil),
			// Same as "1 Jan 1789"
			1789.0027322404371,
		},
		{
			gedcom.NewDateNode("Bet. 1430 and 1435", "", nil),
			// From the start of 1430 to the end of 1435 is actually 6 years.
			1433,
		},

		// Invalid
		{gedcom.NewDateNode("Foo", "", nil), 0},
	}

	for _, test := range tests {
		t.Run(test.date.Value(), func(t *testing.T) {
			assert.Equal(t, test.expected, test.date.Years())
		})
	}
}
