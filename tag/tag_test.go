package tag_test

import (
	"github.com/elliotchance/gedcom/tag"
	"github.com/stretchr/testify/assert"
	"sort"
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
	tag        *tag.Tag
	str        string
}{
	// Official
	//        isKnown
	//        |  isOfficial
	//        |  |  isEvent
	"ABBR":  {y, y, n, &tag.TagAbbreviation, "Abbreviation"},
	"ADDR":  {y, y, n, &tag.TagAddress, "Address"},
	"ADOP":  {y, y, y, &tag.TagAdoption, "Adoption"},
	"ADR1":  {y, y, n, &tag.TagAddress1, "Address Line 1"},
	"ADR2":  {y, y, n, &tag.TagAddress2, "Address Line 2"},
	"AFN":   {y, y, n, &tag.TagAncestralFileNumber, "Ancestral File Number"},
	"AGE":   {y, y, n, &tag.TagAge, "Age"},
	"AGNC":  {y, y, n, &tag.TagAgency, "Agency"},
	"ALIA":  {y, y, n, &tag.TagAlias, "Alias"},
	"ANCE":  {y, y, n, &tag.TagAncestors, "Ancestors"},
	"ANCI":  {y, y, n, &tag.TagAncestorsInterest, "Ancestors Interest"},
	"ANUL":  {y, y, y, &tag.TagAnnulment, "Annulment"},
	"ASSO":  {y, y, n, &tag.TagAssociates, "Associates"},
	"AUTH":  {y, y, n, &tag.TagAuthor, "Author"},
	"BAPL":  {y, y, y, &tag.TagLDSBaptism, "LDS Baptism"},
	"BAPM":  {y, y, y, &tag.TagBaptism, "Baptism"},
	"BARM":  {y, y, y, &tag.TagBarMitzvah, "Bar Mitzvah"},
	"BASM":  {y, y, y, &tag.TagBasMitzvah, "Bas Mitzvah"},
	"BIRT":  {y, y, y, &tag.TagBirth, "Birth"},
	"BLES":  {y, y, y, &tag.TagBlessing, "Blessing"},
	"BLOB":  {y, y, n, &tag.TagBinaryObject, "Binary Object"},
	"BURI":  {y, y, y, &tag.TagBurial, "Burial"},
	"CALN":  {y, y, n, &tag.TagCallNumber, "Call Number"},
	"CAST":  {y, y, n, &tag.TagCaste, "Caste"},
	"CAUS":  {y, y, n, &tag.TagCause, "Cause"},
	"CENS":  {y, y, y, &tag.TagCensus, "Census"},
	"CHAN":  {y, y, n, &tag.TagChange, "Change"},
	"CHAR":  {y, y, n, &tag.TagCharacterSet, "Character Set"},
	"CHIL":  {y, y, n, &tag.TagChild, "Child"},
	"CHR":   {y, y, y, &tag.TagChristening, "Christening"},
	"CHRA":  {y, y, y, &tag.TagAdultChristening, "Adult Christening"},
	"CITY":  {y, y, n, &tag.TagCity, "City"},
	"CONC":  {y, y, n, &tag.TagConcatenation, "Concatenation"},
	"CONF":  {y, y, y, &tag.TagConfirmation, "Confirmation"},
	"CONL":  {y, y, y, &tag.TagLDSConfirmation, "LDS Confirmation"},
	"CONT":  {y, y, n, &tag.TagContinued, "Continued"},
	"COPR":  {y, y, n, &tag.TagCopyright, "Copyright"},
	"CORP":  {y, y, n, &tag.TagCorporate, "Corporate"},
	"CREM":  {y, y, y, &tag.TagCremation, "Cremation"},
	"CTRY":  {y, y, n, &tag.TagCountry, "Country"},
	"DATA":  {y, y, n, &tag.TagData, "Data"},
	"DATE":  {y, y, n, &tag.TagDate, "Date"},
	"DEAT":  {y, y, y, &tag.TagDeath, "Death"},
	"DESC":  {y, y, n, &tag.TagDescendants, "Descendants"},
	"DESI":  {y, y, n, &tag.TagDescendantsInterest, "Descendants Interest"},
	"DEST":  {y, y, n, &tag.TagDestination, "Destination"},
	"DIV":   {y, y, y, &tag.TagDivorce, "Divorce"},
	"DIVF":  {y, y, y, &tag.TagDivorceFiled, "Divorce Filed"},
	"DSCR":  {y, y, n, &tag.TagPhysicalDescription, "Physical Description"},
	"EDUC":  {y, y, n, &tag.TagEducation, "Education"},
	"EMAIL": {y, y, n, &tag.TagEmail, "Email"},
	"EMIG":  {y, y, y, &tag.TagEmigration, "Emigration"},
	"ENDL":  {y, y, y, &tag.TagEndowment, "Endowment"},
	"ENGA":  {y, y, y, &tag.TagEngagement, "Engagement"},
	"EVEN":  {y, y, y, &tag.TagEvent, "Event"},
	"FACT":  {y, y, n, &tag.TagFact, "Fact"},
	"FAM":   {y, y, n, &tag.TagFamily, "Family"},
	"FAMC":  {y, y, n, &tag.TagFamilyChild, "Family Child"},
	"FAMF":  {y, y, n, &tag.TagFamilyFile, "Family File"},
	"FAMS":  {y, y, n, &tag.TagFamilySpouse, "Family Spouse"},
	"FAX":   {y, y, n, &tag.TagFax, "Fax"},
	"FCOM":  {y, y, y, &tag.TagFirstCommunion, "First Communion"},
	"FILE":  {y, y, n, &tag.TagFile, "File"},
	"FONE":  {y, y, n, &tag.TagPhonetic, "Phonetic"},
	"FORM":  {y, y, n, &tag.TagFormat, "Format"},
	"GEDC":  {y, y, n, &tag.TagGedcomInformation, "GEDCOM Information"},
	"GIVN":  {y, y, n, &tag.TagGivenName, "Given Name"},
	"GRAD":  {y, y, y, &tag.TagGraduation, "Graduation"},
	"HEAD":  {y, y, n, &tag.TagHeader, "Header"},
	"HUSB":  {y, y, n, &tag.TagHusband, "Husband"},
	"IDNO":  {y, y, n, &tag.TagIdentityNumber, "Identity Number"},
	"IMMI":  {y, y, y, &tag.TagImmigration, "Immigration"},
	"INDI":  {y, y, n, &tag.TagIndividual, "Individual"},
	"LANG":  {y, y, n, &tag.TagLanguage, "Language"},
	"LABL":  {y, y, n, &tag.TagLabel, "Label"},
	"LATI":  {y, y, n, &tag.TagLatitude, "Latitude"},
	"LEGA":  {y, y, n, &tag.TagLegatee, "Legatee"},
	"LONG":  {y, y, n, &tag.TagLongitude, "Longitude"},
	"MAP":   {y, y, n, &tag.TagMap, "Map"},
	"MARB":  {y, y, y, &tag.TagMarriageBann, "Marriage Bann"},
	"MARC":  {y, y, y, &tag.TagMarriageContract, "Marriage Contract"},
	"MARL":  {y, y, y, &tag.TagMarriageLicence, "Marriage Licence"},
	"MARR":  {y, y, y, &tag.TagMarriage, "Marriage"},
	"MARS":  {y, y, y, &tag.TagMarriageSettlement, "Marriage Settlement"},
	"MEDI":  {y, y, n, &tag.TagMedia, "Media"},
	"NAME":  {y, y, n, &tag.TagName, "Name"},
	"NATI":  {y, y, n, &tag.TagNationality, "Nationality"},
	"NATU":  {y, y, y, &tag.TagNaturalization, "Naturalization"},
	"NCHI":  {y, y, n, &tag.TagChildrenCount, "Children Count"},
	"NICK":  {y, y, n, &tag.TagNickname, "Nickname"},
	"NMR":   {y, y, n, &tag.TagMarriageCount, "Marriage Count"},
	"NOTE":  {y, y, n, &tag.TagNote, "Note"},
	"NPFX":  {y, y, n, &tag.TagNamePrefix, "Name Prefix"},
	"NSFX":  {y, y, n, &tag.TagNameSuffix, "Name Suffix"},
	"OBJE":  {y, y, n, &tag.TagObject, "Object"},
	"OCCU":  {y, y, n, &tag.TagOccupation, "Occupation"},
	"ORDI":  {y, y, n, &tag.TagOrdinance, "Ordinance"},
	"ORDN":  {y, y, y, &tag.TagOrdination, "Ordination"},
	"PAGE":  {y, y, n, &tag.TagPage, "Page"},
	"PEDI":  {y, y, n, &tag.TagPedigree, "Pedigree"},
	"PHON":  {y, y, n, &tag.TagPhone, "Phone"},
	"PLAC":  {y, y, n, &tag.TagPlace, "Place"},
	"POST":  {y, y, n, &tag.TagPostalCode, "Postal Code"},
	"PROB":  {y, y, y, &tag.TagProbate, "Probate"},
	"PROP":  {y, y, n, &tag.TagProperty, "Property"},
	"PUBL":  {y, y, n, &tag.TagPublication, "Publication"},
	"QUAY":  {y, y, n, &tag.TagQualityOfData, "Quality Of Data"},
	"REFN":  {y, y, n, &tag.TagReference, "Reference"},
	"RELA":  {y, y, n, &tag.TagRelationship, "Relationship"},
	"RELI":  {y, y, n, &tag.TagReligion, "Religion"},
	"REPO":  {y, y, n, &tag.TagRepository, "Repository"},
	"RESI":  {y, y, y, &tag.TagResidence, "Residence"},
	"RESN":  {y, y, n, &tag.TagRestriction, "Restriction"},
	"RETI":  {y, y, y, &tag.TagRetirement, "Retirement"},
	"RFN":   {y, y, n, &tag.TagRecordFileNumber, "Record File Number"},
	"RIN":   {y, y, n, &tag.TagRecordIDNumber, "Record ID Number"},
	"ROLE":  {y, y, n, &tag.TagRole, "Role"},
	"ROMN":  {y, y, n, &tag.TagRomanized, "Romanized"},
	"SEX":   {y, y, n, &tag.TagSex, "Sex"},
	"SLGC":  {y, y, y, &tag.TagSealingChild, "Sealing Child"},
	"SLGS":  {y, y, y, &tag.TagSealingSpouse, "Sealing Spouse"},
	"SOUR":  {y, y, n, &tag.TagSource, "Source"},
	"SPFX":  {y, y, n, &tag.TagSurnamePrefix, "Surname Prefix"},
	"SSN":   {y, y, n, &tag.TagSocialSecurityNumber, "Social Security Number"},
	"STAE":  {y, y, n, &tag.TagState, "State"},
	"STAT":  {y, y, n, &tag.TagStatus, "Status"},
	"SUBM":  {y, y, n, &tag.TagSubmitter, "Submitter"},
	"SUBN":  {y, y, n, &tag.TagSubmission, "Submission"},
	"SURN":  {y, y, n, &tag.TagSurname, "Surname"},
	"TEMP":  {y, y, n, &tag.TagTemple, "Temple"},
	"TEXT":  {y, y, n, &tag.TagText, "Text"},
	"TIME":  {y, y, n, &tag.TagTime, "Time"},
	"TITL":  {y, y, n, &tag.TagTitle, "Title"},
	"TRLR":  {y, y, n, &tag.TagTrailer, "Trailer"},
	"TYPE":  {y, y, n, &tag.TagType, "Type"},
	"VERS":  {y, y, n, &tag.TagVersion, "Version"},
	"WIFE":  {y, y, n, &tag.TagWife, "Wife"},
	"WILL":  {y, y, y, &tag.TagWill, "Will"},
	"WWW":   {y, y, n, &tag.TagWWW, "WWW"},

	// Unofficial
	//          isKnown
	//          |  isOfficial
	//          |  |  isEvent
	"_COR":    {y, n, n, &tag.UnofficialTagCoordinates, "Coordinates"},
	"_CRE":    {y, n, n, &tag.UnofficialTagCreated, "Created"},
	"_FID":    {y, n, n, &tag.UnofficialTagFamilySearchID1, "FamilySearch ID"},
	"_FSFTID": {y, n, n, &tag.UnofficialTagFamilySearchID2, "FamilySearch ID"},
	"_LAD":    {y, n, n, &tag.UnofficialTagLatitudeDegrees, "Latitude Degrees"},
	"_LAM":    {y, n, n, &tag.UnofficialTagLatitudeMinutes, "Latitude Minutes"},
	"_LAS":    {y, n, n, &tag.UnofficialTagLatitudeSeconds, "Latitude Seconds"},
	"_LOD":    {y, n, n, &tag.UnofficialTagLongitudeDegress, "Longitude Degress"},
	"_LOM":    {y, n, n, &tag.UnofficialTagLongitudeMinutes, "Longitude Minutes"},
	"_LON":    {y, n, n, &tag.UnofficialTagLongitudeNorth, "Longitude North"},
	"_LOS":    {y, n, n, &tag.UnofficialTagLongitudeSeconds, "Longitude Seconds"},
	"_UID":    {y, n, n, &tag.UnofficialTagUniqueID, "Unique ID"},

	// Unknown
	//       isKnown
	//       |  isOfficial
	//       |  |  isEvent
	"FOBR": {n, n, n, nil, "FOBR"},
	"BRBZ": {n, n, n, nil, "BRBZ"},
	"":     {n, n, n, nil, ""},
}

func TestTagFromString(t *testing.T) {
	tests := []struct {
		tag      string
		expected tag.Tag
	}{
		// Official and unofficial known tags we would expect to find.
		{"HEAD", tag.TagHeader},
		{"DATE", tag.TagDate},
		{"_CRE", tag.UnofficialTagCreated},

		// Unknown tags we should not find.
		{"FOBR", tag.TagFromString("FOBR")},
		{"BRBZ", tag.TagFromString("BRBZ")},
		{"", tag.TagFromString("")},
	}

	for _, test := range tests {
		t.Run(test.tag, func(t *testing.T) {
			actualTag := tag.TagFromString(test.tag)

			assert.True(t, test.expected.Is(actualTag))
		})
	}
}

func TestTags(t *testing.T) {
	tests := []struct {
		tag      string
		expected tag.Tag
	}{
		// Official and unofficial known tags we would expect to find.
		{"HEAD", tag.TagHeader},
		{"DATE", tag.TagDate},
		{"_CRE", tag.UnofficialTagCreated},

		// Unknown tags we should not find.
		{"FOBR", tag.TagFromString("FOBR")},
		{"BRBZ", tag.TagFromString("BRBZ")},
		{"", tag.TagFromString("")},
	}

	for _, test := range tests {
		t.Run(test.tag, func(t *testing.T) {
			if test.expected.IsKnown() {
				assert.Contains(t, tag.Tags(), test.expected)
			} else {
				assert.NotContains(t, tag.Tags(), test.expected)
			}
		})
	}

	t.Run("Sorting", func(t *testing.T) {
		tags := tag.Tags()
		sort.SliceStable(tags, func(i, j int) bool {
			if tags[i].SortValue() != tags[j].SortValue() {
				return tags[i].SortValue() < tags[j].SortValue()
			}

			return tags[i].String() < tags[j].String()
		})

		assert.Equal(t, []tag.Tag{
			tag.TagName,

			tag.TagAbbreviation,
			tag.TagAddress,
			tag.TagAddress1,
			tag.TagAddress2,
			tag.TagAge,
			tag.TagAgency,
			tag.TagAlias,
			tag.TagAncestors,
			tag.TagAncestorsInterest,
			tag.TagAncestralFileNumber,
			tag.TagAssociates,
			tag.TagAuthor,
			tag.TagBinaryObject,
			tag.TagCallNumber,
			tag.TagCaste,
			tag.TagCause,
			tag.TagChange,
			tag.TagCharacterSet,
			tag.TagChild,
			tag.TagChildrenCount,
			tag.TagCity,
			tag.TagConcatenation,
			tag.TagContinued,
			tag.TagCopyright,
			tag.TagCorporate,
			tag.TagCountry,
			tag.TagData,
			tag.TagDate,
			tag.TagDescendants,
			tag.TagDescendantsInterest,
			tag.TagDestination,
			tag.TagEducation,
			tag.TagEmail,
			tag.TagFact,
			tag.TagFamily,
			tag.TagFamilyChild,
			tag.TagFamilyFile,
			tag.TagFamilySpouse,
			tag.TagFax,
			tag.TagFile,
			tag.TagFormat,
			tag.TagGedcomInformation,
			tag.TagGivenName,
			tag.TagHeader,
			tag.TagHusband,
			tag.TagIdentityNumber,
			tag.TagIndividual,
			tag.TagLabel,
			tag.TagLanguage,
			tag.TagLatitude,
			tag.TagLegatee,
			tag.TagLongitude,
			tag.TagMap,
			tag.TagMarriageCount,
			tag.TagMedia,
			tag.TagNamePrefix,
			tag.TagNameSuffix,
			tag.TagNationality,
			tag.TagNickname,
			tag.TagNote,
			tag.TagObject,
			tag.TagOccupation,
			tag.TagOrdinance,
			tag.TagPage,
			tag.TagPedigree,
			tag.TagPhone,
			tag.TagPhonetic,
			tag.TagPhysicalDescription,
			tag.TagPlace,
			tag.TagPostalCode,
			tag.TagProperty,
			tag.TagPublication,
			tag.TagQualityOfData,
			tag.TagRecordFileNumber,
			tag.TagRecordIDNumber,
			tag.TagReference,
			tag.TagRelationship,
			tag.TagReligion,
			tag.TagRepository,
			tag.TagRestriction,
			tag.TagRole,
			tag.TagRomanized,
			tag.TagSex,
			tag.TagSocialSecurityNumber,
			tag.TagSource,
			tag.TagState,
			tag.TagStatus,
			tag.TagSubmission,
			tag.TagSubmitter,
			tag.TagSurname,
			tag.TagSurnamePrefix,
			tag.TagTemple,
			tag.TagText,
			tag.TagTime,
			tag.TagTitle,
			tag.TagTrailer,
			tag.TagType,
			tag.TagVersion,
			tag.TagWWW,
			tag.TagWife,

			tag.TagBirth,

			tag.TagAdoption,
			tag.TagAdultChristening,
			tag.TagAnnulment,
			tag.TagBaptism,
			tag.TagBarMitzvah,
			tag.TagBasMitzvah,
			tag.TagBlessing,
			tag.TagCensus,
			tag.TagChristening,
			tag.TagConfirmation,
			tag.TagCremation,
			tag.TagDivorce,
			tag.TagDivorceFiled,
			tag.TagEmigration,
			tag.TagEndowment,
			tag.TagEngagement,
			tag.TagEvent,
			tag.TagFirstCommunion,
			tag.TagGraduation,
			tag.TagImmigration,
			tag.TagLDSBaptism,
			tag.TagLDSConfirmation,
			tag.TagMarriage,
			tag.TagMarriageBann,
			tag.TagMarriageContract,
			tag.TagMarriageLicence,
			tag.TagMarriageSettlement,
			tag.TagNaturalization,
			tag.TagOrdination,
			tag.TagProbate,
			tag.TagResidence,
			tag.TagRetirement,
			tag.TagSealingChild,
			tag.TagSealingSpouse,
			tag.TagWill,

			tag.TagDeath,

			tag.TagBurial,

			tag.UnofficialTagCoordinates,
			tag.UnofficialTagCreated,
			tag.UnofficialTagFamilySearchID1,
			tag.UnofficialTagFamilySearchID2,
			tag.UnofficialTagLatitudeDegrees,
			tag.UnofficialTagLatitudeMinutes,
			tag.UnofficialTagLatitudeSeconds,
			tag.UnofficialTagLongitudeDegress,
			tag.UnofficialTagLongitudeMinutes,
			tag.UnofficialTagLongitudeNorth,
			tag.UnofficialTagLongitudeSeconds,
			tag.UnofficialTagUniqueID,
		}, tags)
	})
}

func TestTag_String(t *testing.T) {
	for _, tag := range tag.Tags() {
		t.Run(tag.String(), func(t *testing.T) {
			expected := tagTests[tag.Tag()]
			assert.Equal(t, expected.str, tag.String())
		})
	}
}

func TestTag_IsOfficial(t *testing.T) {
	for _, tag := range tag.Tags() {
		t.Run(tag.String(), func(t *testing.T) {
			expected := tagTests[tag.Tag()]
			assert.Equal(t, expected.isOfficial, tag.IsOfficial())
		})
	}
}

func TestTag_IsEvent(t *testing.T) {
	for _, tag := range tag.Tags() {
		t.Run(tag.String(), func(t *testing.T) {
			expected := tagTests[tag.Tag()]
			assert.Equal(t, expected.isEvent, tag.IsEvent())
		})
	}
}

func TestTag_Tag(t *testing.T) {
	for _, tag := range tag.Tags() {
		t.Run(tag.String(), func(t *testing.T) {
			expected := tagTests[tag.Tag()]
			assert.Equal(t, expected.tag.Tag(), tag.Tag())
		})
	}
}

func TestTag_Known(t *testing.T) {
	for _, tag := range tag.Tags() {
		t.Run(tag.String(), func(t *testing.T) {
			expected := tagTests[tag.Tag()]
			assert.Equal(t, expected.isKnown, tag.IsKnown())
		})
	}
}

func TestTag_Is(t *testing.T) {
	tests := []struct {
		a, b tag.Tag
		match bool
	}{
		{tag.TagHeader, tag.TagHeader, true},
		{tag.TagFromString("DATE"), tag.TagDate, true},
		{tag.TagBirth, tag.TagFromString("BIRT"), true},
		{tag.TagFromString("BRBZ"), tag.TagFromString("BRBZ"), true},
		{tag.TagFromString(""), tag.TagFromString(""), true},

		{tag.TagHeader, tag.TagVersion, false},
		{tag.TagFromString("DATE"), tag.TagHeader, false},
		{tag.TagHusband, tag.TagFromString("BIRT"), false},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.match, test.a.Is(test.b))
		})
	}
}
