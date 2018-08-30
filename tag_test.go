package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	y = true
	n = false
)

var tagTests = map[string]struct {
	isKnown    bool
	isOfficial bool
	isEvent    bool
	tag        *gedcom.Tag
	str        string
}{
	// Official
	"ABBR":  {y, y, n, &gedcom.TagAbbreviation, "Abbreviation"},
	"ADDR":  {y, y, n, &gedcom.TagAddress, "Address"},
	"ADOP":  {y, y, y, &gedcom.TagAdoption, "Adoption"},
	"ADR1":  {y, y, n, &gedcom.TagAddress1, "Address Line 1"},
	"ADR2":  {y, y, n, &gedcom.TagAddress2, "Address Line 2"},
	"AFN":   {y, y, n, &gedcom.TagAncestralFileNumber, "Ancestral File Number"},
	"AGE":   {y, y, n, &gedcom.TagAge, "Age"},
	"AGNC":  {y, y, n, &gedcom.TagAgency, "Agency"},
	"ALIA":  {y, y, n, &gedcom.TagAlias, "Alias"},
	"ANCE":  {y, y, n, &gedcom.TagAncestors, "Ancestors"},
	"ANCI":  {y, y, n, &gedcom.TagAncestorsInterest, "Ancestors Interest"},
	"ANUL":  {y, y, y, &gedcom.TagAnnulment, "Annulment"},
	"ASSO":  {y, y, n, &gedcom.TagAssociates, "Associates"},
	"AUTH":  {y, y, n, &gedcom.TagAuthor, "Author"},
	"BAPL":  {y, y, y, &gedcom.TagLDSBaptism, "LDS Baptism"},
	"BAPM":  {y, y, y, &gedcom.TagBaptism, "Baptism"},
	"BARM":  {y, y, y, &gedcom.TagBarMitzvah, "Bar Mitzvah"},
	"BASM":  {y, y, y, &gedcom.TagBasMitzvah, "Bas Mitzvah"},
	"BIRT":  {y, y, y, &gedcom.TagBirth, "Birth"},
	"BLES":  {y, y, y, &gedcom.TagBlessing, "Blessing"},
	"BLOB":  {y, y, n, &gedcom.TagBinaryObject, "Binary Object"},
	"BURI":  {y, y, y, &gedcom.TagBurial, "Burial"},
	"CALN":  {y, y, n, &gedcom.TagCallNumber, "Call Number"},
	"CAST":  {y, y, n, &gedcom.TagCaste, "Caste"},
	"CAUS":  {y, y, n, &gedcom.TagCause, "Cause"},
	"CENS":  {y, y, y, &gedcom.TagCensus, "Census"},
	"CHAN":  {y, y, n, &gedcom.TagChange, "Change"},
	"CHAR":  {y, y, n, &gedcom.TagCharacterSet, "Character Set"},
	"CHIL":  {y, y, n, &gedcom.TagChild, "Child"},
	"CHR":   {y, y, y, &gedcom.TagChristening, "Christening"},
	"CHRA":  {y, y, y, &gedcom.TagAdultChristening, "Adult Christening"},
	"CITY":  {y, y, n, &gedcom.TagCity, "City"},
	"CONC":  {y, y, n, &gedcom.TagConcatenation, "Concatenation"},
	"CONF":  {y, y, y, &gedcom.TagConfirmation, "Confirmation"},
	"CONL":  {y, y, y, &gedcom.TagLDSConfirmation, "LDS Confirmation"},
	"CONT":  {y, y, n, &gedcom.TagContinued, "Continued"},
	"COPR":  {y, y, n, &gedcom.TagCopyright, "Copyright"},
	"CORP":  {y, y, n, &gedcom.TagCorporate, "Corporate"},
	"CREM":  {y, y, y, &gedcom.TagCremation, "Cremation"},
	"CTRY":  {y, y, n, &gedcom.TagCountry, "Country"},
	"DATA":  {y, y, n, &gedcom.TagData, "Data"},
	"DATE":  {y, y, n, &gedcom.TagDate, "Date"},
	"DEAT":  {y, y, y, &gedcom.TagDeath, "Death"},
	"DESC":  {y, y, n, &gedcom.TagDescendants, "Descendants"},
	"DESI":  {y, y, n, &gedcom.TagDescendantsInterest, "Descendants Interest"},
	"DEST":  {y, y, n, &gedcom.TagDestination, "Destination"},
	"DIV":   {y, y, y, &gedcom.TagDivorce, "Divorce"},
	"DIVF":  {y, y, y, &gedcom.TagDivorceFiled, "Divorce Filed"},
	"DSCR":  {y, y, n, &gedcom.TagPhysicalDescription, "Physical Description"},
	"EDUC":  {y, y, n, &gedcom.TagEducation, "Education"},
	"EMAIL": {y, y, n, &gedcom.TagEmail, "Email"},
	"EMIG":  {y, y, y, &gedcom.TagEmigration, "Emigration"},
	"ENDL":  {y, y, y, &gedcom.TagEndowment, "Endowment"},
	"ENGA":  {y, y, y, &gedcom.TagEngagement, "Engagement"},
	"EVEN":  {y, y, y, &gedcom.TagEvent, "Event"},
	"FACT":  {y, y, n, &gedcom.TagFact, "Fact"},
	"FAM":   {y, y, n, &gedcom.TagFamily, "Family"},
	"FAMC":  {y, y, n, &gedcom.TagFamilyChild, "Family Child"},
	"FAMF":  {y, y, n, &gedcom.TagFamilyFile, "Family File"},
	"FAMS":  {y, y, n, &gedcom.TagFamilySpouse, "Family Spouse"},
	"FAX":   {y, y, n, &gedcom.TagFax, "Fax"},
	"FCOM":  {y, y, y, &gedcom.TagFirstCommunion, "First Communion"},
	"FILE":  {y, y, n, &gedcom.TagFile, "File"},
	"FONE":  {y, y, n, &gedcom.TagPhonetic, "Phonetic"},
	"FORM":  {y, y, n, &gedcom.TagFormat, "Format"},
	"GEDC":  {y, y, n, &gedcom.TagGedcomInformation, "GEDCOM Information"},
	"GIVN":  {y, y, n, &gedcom.TagGivenName, "Given Name"},
	"GRAD":  {y, y, y, &gedcom.TagGraduation, "Graduation"},
	"HEAD":  {y, y, n, &gedcom.TagHeader, "Header"},
	"HUSB":  {y, y, n, &gedcom.TagHusband, "Husband"},
	"IDNO":  {y, y, n, &gedcom.TagIdentityNumber, "Identity Number"},
	"IMMI":  {y, y, y, &gedcom.TagImmigration, "Immigration"},
	"INDI":  {y, y, n, &gedcom.TagIndividual, "Individual"},
	"LANG":  {y, y, n, &gedcom.TagLanguage, "Language"},
	"LATI":  {y, y, n, &gedcom.TagLatitude, "Latitude"},
	"LEGA":  {y, y, n, &gedcom.TagLegatee, "Legatee"},
	"LONG":  {y, y, n, &gedcom.TagLongitude, "Longitude"},
	"MAP":   {y, y, n, &gedcom.TagMap, "Map"},
	"MARB":  {y, y, y, &gedcom.TagMarriageBann, "Marriage Bann"},
	"MARC":  {y, y, y, &gedcom.TagMarriageContract, "Marriage Contract"},
	"MARL":  {y, y, y, &gedcom.TagMarriageLicence, "Marriage Licence"},
	"MARR":  {y, y, y, &gedcom.TagMarriage, "Marriage"},
	"MARS":  {y, y, y, &gedcom.TagMarriageSettlement, "Marriage Settlement"},
	"MEDI":  {y, y, n, &gedcom.TagMedia, "Media"},
	"NAME":  {y, y, n, &gedcom.TagName, "Name"},
	"NATI":  {y, y, n, &gedcom.TagNationality, "Nationality"},
	"NATU":  {y, y, y, &gedcom.TagNaturalization, "Naturalization"},
	"NCHI":  {y, y, n, &gedcom.TagChildrenCount, "Children Count"},
	"NICK":  {y, y, n, &gedcom.TagNickname, "Nickname"},
	"NMR":   {y, y, n, &gedcom.TagMarriageCount, "Marriage Count"},
	"NOTE":  {y, y, n, &gedcom.TagNote, "Note"},
	"NPFX":  {y, y, n, &gedcom.TagNamePrefix, "Name Prefix"},
	"NSFX":  {y, y, n, &gedcom.TagNameSuffix, "Name Suffix"},
	"OBJE":  {y, y, n, &gedcom.TagObject, "Object"},
	"OCCU":  {y, y, n, &gedcom.TagOccupation, "Occupation"},
	"ORDI":  {y, y, n, &gedcom.TagOrdinance, "Ordinance"},
	"ORDN":  {y, y, y, &gedcom.TagOrdination, "Ordination"},
	"PAGE":  {y, y, n, &gedcom.TagPage, "Page"},
	"PEDI":  {y, y, n, &gedcom.TagPedigree, "Pedigree"},
	"PHON":  {y, y, n, &gedcom.TagPhone, "Phone"},
	"PLAC":  {y, y, n, &gedcom.TagPlace, "Place"},
	"POST":  {y, y, n, &gedcom.TagPostalCode, "Postal Code"},
	"PROB":  {y, y, y, &gedcom.TagProbate, "Probate"},
	"PROP":  {y, y, n, &gedcom.TagProperty, "Property"},
	"PUBL":  {y, y, n, &gedcom.TagPublication, "Publication"},
	"QUAY":  {y, y, n, &gedcom.TagQualityOfData, "Quality Of Data"},
	"REFN":  {y, y, n, &gedcom.TagReference, "Reference"},
	"RELA":  {y, y, n, &gedcom.TagRelationship, "Relationship"},
	"RELI":  {y, y, n, &gedcom.TagReligion, "Religion"},
	"REPO":  {y, y, n, &gedcom.TagRepository, "Repository"},
	"RESI":  {y, y, y, &gedcom.TagResidence, "Residence"},
	"RESN":  {y, y, n, &gedcom.TagRestriction, "Restriction"},
	"RETI":  {y, y, y, &gedcom.TagRetirement, "Retirement"},
	"RFN":   {y, y, n, &gedcom.TagRecordFileNumber, "Record File Number"},
	"RIN":   {y, y, n, &gedcom.TagRecordIDNumber, "Record ID Number"},
	"ROLE":  {y, y, n, &gedcom.TagRole, "Role"},
	"ROMN":  {y, y, n, &gedcom.TagRomanized, "Romanized"},
	"SEX":   {y, y, n, &gedcom.TagSex, "Sex"},
	"SLGC":  {y, y, y, &gedcom.TagSealingChild, "Sealing Child"},
	"SLGS":  {y, y, y, &gedcom.TagSealingSpouse, "Sealing Spouse"},
	"SOUR":  {y, y, n, &gedcom.TagSource, "Source"},
	"SPFX":  {y, y, n, &gedcom.TagSurnamePrefix, "Surname Prefix"},
	"SSN":   {y, y, n, &gedcom.TagSocialSecurityNumber, "Social Security Number"},
	"STAE":  {y, y, n, &gedcom.TagState, "State"},
	"STAT":  {y, y, n, &gedcom.TagStatus, "Status"},
	"SUBM":  {y, y, n, &gedcom.TagSubmitter, "Submitter"},
	"SUBN":  {y, y, n, &gedcom.TagSubmission, "Submission"},
	"SURN":  {y, y, n, &gedcom.TagSurname, "Surname"},
	"TEMP":  {y, y, n, &gedcom.TagTemple, "Temple"},
	"TEXT":  {y, y, n, &gedcom.TagText, "Text"},
	"TIME":  {y, y, n, &gedcom.TagTime, "Time"},
	"TITL":  {y, y, n, &gedcom.TagTitle, "Title"},
	"TRLR":  {y, y, n, &gedcom.TagTrailer, "Trailer"},
	"TYPE":  {y, y, n, &gedcom.TagType, "Type"},
	"VERS":  {y, y, n, &gedcom.TagVersion, "Version"},
	"WIFE":  {y, y, n, &gedcom.TagWife, "Wife"},
	"WILL":  {y, y, y, &gedcom.TagWill, "Will"},
	"WWW":   {y, y, n, &gedcom.TagWWW, "WWW"},

	// Unofficial
	"_COR": {y, n, n, &gedcom.UnofficialTagCoordinates, "Coordinates"},
	"_CRE": {y, n, n, &gedcom.UnofficialTagCreated, "Created"},
	"_FID": {y, n, n, &gedcom.UnofficialTagFamilySearchID, "FamilySearch ID"},
	"_LAD": {y, n, n, &gedcom.UnofficialTagLatitudeDegrees, "Latitude Degrees"},
	"_LAM": {y, n, n, &gedcom.UnofficialTagLatitudeMinutes, "Latitude Minutes"},
	"_LAS": {y, n, n, &gedcom.UnofficialTagLatitudeSeconds, "Latitude Seconds"},
	"_LOD": {y, n, n, &gedcom.UnofficialTagLongitudeDegress, "Longitude Degress"},
	"_LOM": {y, n, n, &gedcom.UnofficialTagLongitudeMinutes, "Longitude Minutes"},
	"_LON": {y, n, n, &gedcom.UnofficialTagLongitudeNorth, "Longitude North"},
	"_LOS": {y, n, n, &gedcom.UnofficialTagLongitudeSeconds, "Longitude Seconds"},

	// Unknown
	"FOBR": {n, n, n, nil, "FOBR"},
	"BRBZ": {n, n, n, nil, "BRBZ"},
	"":     {n, n, n, nil, ""},
}

func TestTagFromString(t *testing.T) {
	tests := []struct {
		tag      string
		expected gedcom.Tag
	}{
		// Official and unofficial known tags we would expect to find.
		{"HEAD", gedcom.TagHeader},
		{"DATE", gedcom.TagDate},
		{"_CRE", gedcom.UnofficialTagCreated},

		// Unknown tags we should not find.
		{"FOBR", gedcom.TagFromString("FOBR")},
		{"BRBZ", gedcom.TagFromString("BRBZ")},
		{"", gedcom.TagFromString("")},
	}

	for _, test := range tests {
		t.Run(test.tag, func(t *testing.T) {
			actualTag := gedcom.TagFromString(test.tag)

			assert.True(t, test.expected.Is(actualTag))
		})
	}
}

func TestTags(t *testing.T) {
	tests := []struct {
		tag      string
		expected gedcom.Tag
	}{
		// Official and unofficial known tags we would expect to find.
		{"HEAD", gedcom.TagHeader},
		{"DATE", gedcom.TagDate},
		{"_CRE", gedcom.UnofficialTagCreated},

		// Unknown tags we should not find.
		{"FOBR", gedcom.TagFromString("FOBR")},
		{"BRBZ", gedcom.TagFromString("BRBZ")},
		{"", gedcom.TagFromString("")},
	}

	for _, test := range tests {
		t.Run(test.tag, func(t *testing.T) {
			if test.expected.IsKnown() {
				assert.Contains(t, gedcom.Tags(), test.expected)
			} else {
				assert.NotContains(t, gedcom.Tags(), test.expected)
			}
		})
	}
}

func TestTag_String(t *testing.T) {
	for _, tag := range gedcom.Tags() {
		t.Run(tag.String(), func(t *testing.T) {
			expected := tagTests[tag.Tag()]
			assert.Equal(t, expected.str, tag.String())
		})
	}
}

func TestTag_IsOfficial(t *testing.T) {
	for _, tag := range gedcom.Tags() {
		t.Run(tag.String(), func(t *testing.T) {
			expected := tagTests[tag.Tag()]
			assert.Equal(t, expected.isOfficial, tag.IsOfficial())
		})
	}
}

func TestTag_IsEvent(t *testing.T) {
	for _, tag := range gedcom.Tags() {
		t.Run(tag.String(), func(t *testing.T) {
			expected := tagTests[tag.Tag()]
			assert.Equal(t, expected.isEvent, tag.IsEvent())
		})
	}
}

func TestTag_Tag(t *testing.T) {
	for _, tag := range gedcom.Tags() {
		t.Run(tag.String(), func(t *testing.T) {
			expected := tagTests[tag.Tag()]
			assert.Equal(t, expected.tag.Tag(), tag.Tag())
		})
	}
}

func TestTag_Known(t *testing.T) {
	for _, tag := range gedcom.Tags() {
		t.Run(tag.String(), func(t *testing.T) {
			expected := tagTests[tag.Tag()]
			assert.Equal(t, expected.isKnown, tag.IsKnown())
		})
	}
}

func TestTag_Is(t *testing.T) {
	tests := []struct {
		a, b  gedcom.Tag
		match bool
	}{
		{gedcom.TagHeader, gedcom.TagHeader, true},
		{gedcom.TagFromString("DATE"), gedcom.TagDate, true},
		{gedcom.TagBirth, gedcom.TagFromString("BIRT"), true},
		{gedcom.TagFromString("BRBZ"), gedcom.TagFromString("BRBZ"), true},
		{gedcom.TagFromString(""), gedcom.TagFromString(""), true},

		{gedcom.TagHeader, gedcom.TagVersion, false},
		{gedcom.TagFromString("DATE"), gedcom.TagHeader, false},
		{gedcom.TagHusband, gedcom.TagFromString("BIRT"), false},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.match, test.a.Is(test.b))
		})
	}
}
