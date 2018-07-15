package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tagTests = map[string]struct {
	tag        gedcom.Tag
	s          string
	isOfficial bool
}{
	"_COR":  {gedcom.UnofficialTagCoordinates, "Coordinates", false},
	"_CRE":  {gedcom.UnofficialTagCreated, "Created", false},
	"_FID":  {gedcom.UnofficialTagFamilySearchID, "FamilySearchID", false},
	"_LAD":  {gedcom.UnofficialTagLatitudeDegrees, "LatitudeDegrees", false},
	"_LAM":  {gedcom.UnofficialTagLatitudeMinutes, "LatitudeMinutes", false},
	"_LAS":  {gedcom.UnofficialTagLatitudeSeconds, "LatitudeSeconds", false},
	"_LOD":  {gedcom.UnofficialTagLongitudeDegress, "LongitudeDegress", false},
	"_LOM":  {gedcom.UnofficialTagLongitudeMinutes, "LongitudeMinutes", false},
	"_LON":  {gedcom.UnofficialTagLongitudeNorth, "LongitudeNorth", false},
	"_LOS":  {gedcom.UnofficialTagLongitudeSeconds, "LongitudeSeconds", false},
	"ABBR":  {gedcom.TagAbbreviation, "Abbreviation", true},
	"ADDR":  {gedcom.TagAddress, "Address", true},
	"ADOP":  {gedcom.TagAdoption, "Adoption", true},
	"ADR1":  {gedcom.TagAddress1, "Address1", true},
	"ADR2":  {gedcom.TagAddress2, "Address2", true},
	"AFN":   {gedcom.TagAncestralFileNumber, "AncestralFileNumber", true},
	"AGE":   {gedcom.TagAge, "Age", true},
	"AGNC":  {gedcom.TagAgency, "Agency", true},
	"ALIA":  {gedcom.TagAlias, "Alias", true},
	"ANCE":  {gedcom.TagAncestors, "Ancestors", true},
	"ANCI":  {gedcom.TagAncestorsInterest, "AncestorsInterest", true},
	"ANUL":  {gedcom.TagAnnulment, "Annulment", true},
	"ASSO":  {gedcom.TagAssociates, "Associates", true},
	"AUTH":  {gedcom.TagAuthor, "Author", true},
	"BAPL":  {gedcom.TagLDSBaptism, "LDSBaptism", true},
	"BAPM":  {gedcom.TagBaptism, "Baptism", true},
	"BARM":  {gedcom.TagBarMitzvah, "BarMitzvah", true},
	"BASM":  {gedcom.TagBasMitzvah, "BasMitzvah", true},
	"BIRT":  {gedcom.TagBirth, "Birth", true},
	"BLES":  {gedcom.TagBlessing, "Blessing", true},
	"BLOB":  {gedcom.TagBinaryObject, "BinaryObject", true},
	"BURI":  {gedcom.TagBurial, "Burial", true},
	"CALN":  {gedcom.TagCallNumber, "CallNumber", true},
	"CAST":  {gedcom.TagCaste, "Caste", true},
	"CAUS":  {gedcom.TagCause, "Cause", true},
	"CENS":  {gedcom.TagCensus, "Census", true},
	"CHAN":  {gedcom.TagChange, "Change", true},
	"CHAR":  {gedcom.TagCharacterSet, "CharacterSet", true},
	"CHIL":  {gedcom.TagChild, "Child", true},
	"CHR":   {gedcom.TagChristening, "Christening", true},
	"CHRA":  {gedcom.TagAdultChristening, "AdultChristening", true},
	"CITY":  {gedcom.TagCity, "City", true},
	"CONC":  {gedcom.TagConcatenation, "Concatenation", true},
	"CONF":  {gedcom.TagConfirmation, "Confirmation", true},
	"CONL":  {gedcom.TagLDSConfirmation, "LDSConfirmation", true},
	"CONT":  {gedcom.TagContinued, "Continued", true},
	"COPR":  {gedcom.TagCopyright, "Copyright", true},
	"CORP":  {gedcom.TagCorporate, "Corporate", true},
	"CREM":  {gedcom.TagCremation, "Cremation", true},
	"CTRY":  {gedcom.TagCountry, "Country", true},
	"DATA":  {gedcom.TagData, "Data", true},
	"DATE":  {gedcom.TagDate, "Date", true},
	"DEAT":  {gedcom.TagDeath, "Death", true},
	"DESC":  {gedcom.TagDescendants, "Descendants", true},
	"DESI":  {gedcom.TagDescendantsInterest, "DescendantsInterest", true},
	"DEST":  {gedcom.TagDestination, "Destination", true},
	"DIV":   {gedcom.TagDivorce, "Divorce", true},
	"DIVF":  {gedcom.TagDivorceFiled, "DivorceFiled", true},
	"DSCR":  {gedcom.TagPhysicalDescription, "PhysicalDescription", true},
	"EDUC":  {gedcom.TagEducation, "Education", true},
	"EMAIL": {gedcom.TagEmail, "Email", true},
	"EMIG":  {gedcom.TagEmigration, "Emigration", true},
	"ENDL":  {gedcom.TagEndowment, "Endowment", true},
	"ENGA":  {gedcom.TagEngagement, "Engagement", true},
	"EVEN":  {gedcom.TagEvent, "Event", true},
	"FACT":  {gedcom.TagFact, "Fact", true},
	"FAM":   {gedcom.TagFamily, "Family", true},
	"FAMC":  {gedcom.TagFamilyChild, "FamilyChild", true},
	"FAMF":  {gedcom.TagFamilyFile, "FamilyFile", true},
	"FAMS":  {gedcom.TagFamilySpouse, "FamilySpouse", true},
	"FAX":   {gedcom.TagFax, "Fax", true},
	"FCOM":  {gedcom.TagFirstCommunion, "FirstCommunion", true},
	"FILE":  {gedcom.TagFile, "File", true},
	"FONE":  {gedcom.TagPhonetic, "Phonetic", true},
	"FORM":  {gedcom.TagFormat, "Format", true},
	"GEDC":  {gedcom.TagGedcomInformation, "GedcomInformation", true},
	"GIVN":  {gedcom.TagGivenName, "GivenName", true},
	"GRAD":  {gedcom.TagGraduation, "Graduation", true},
	"HEAD":  {gedcom.TagHeader, "Header", true},
	"HUSB":  {gedcom.TagHusband, "Husband", true},
	"IDNO":  {gedcom.TagIdentityNumber, "IdentityNumber", true},
	"IMMI":  {gedcom.TagImmigration, "Immigration", true},
	"INDI":  {gedcom.TagIndividual, "Individual", true},
	"LANG":  {gedcom.TagLanguage, "Language", true},
	"LATI":  {gedcom.TagLatitude, "Latitude", true},
	"LEGA":  {gedcom.TagLegatee, "Legatee", true},
	"LONG":  {gedcom.TagLongitude, "Longitude", true},
	"MAP":   {gedcom.TagMap, "Map", true},
	"MARB":  {gedcom.TagMarriageBann, "MarriageBann", true},
	"MARC":  {gedcom.TagMarriageContract, "MarriageContract", true},
	"MARL":  {gedcom.TagMarriageLicence, "MarriageLicence", true},
	"MARR":  {gedcom.TagMarriage, "Marriage", true},
	"MARS":  {gedcom.TagMarriageSettlement, "MarriageSettlement", true},
	"MEDI":  {gedcom.TagMedia, "Media", true},
	"NAME":  {gedcom.TagName, "Name", true},
	"NATI":  {gedcom.TagNationality, "Nationality", true},
	"NATU":  {gedcom.TagNaturalization, "Naturalization", true},
	"NCHI":  {gedcom.TagChildrenCount, "ChildrenCount", true},
	"NICK":  {gedcom.TagNickname, "Nickname", true},
	"NMR":   {gedcom.TagMarriageCount, "MarriageCount", true},
	"NOTE":  {gedcom.TagNote, "Note", true},
	"NPFX":  {gedcom.TagNamePrefix, "NamePrefix", true},
	"NSFX":  {gedcom.TagNameSuffix, "NameSuffix", true},
	"OBJE":  {gedcom.TagObject, "Object", true},
	"OCCU":  {gedcom.TagOccupation, "Occupation", true},
	"ORDI":  {gedcom.TagOrdinance, "Ordinance", true},
	"ORDN":  {gedcom.TagOrdination, "Ordination", true},
	"PAGE":  {gedcom.TagPage, "Page", true},
	"PEDI":  {gedcom.TagPedigree, "Pedigree", true},
	"PHON":  {gedcom.TagPhone, "Phone", true},
	"PLAC":  {gedcom.TagPlace, "Place", true},
	"POST":  {gedcom.TagPostalCode, "PostalCode", true},
	"PROB":  {gedcom.TagProbate, "Probate", true},
	"PROP":  {gedcom.TagProperty, "Property", true},
	"PUBL":  {gedcom.TagPublication, "Publication", true},
	"QUAY":  {gedcom.TagQualityOfData, "QualityOfData", true},
	"REFN":  {gedcom.TagReference, "Reference", true},
	"RELA":  {gedcom.TagRelationship, "Relationship", true},
	"RELI":  {gedcom.TagReligion, "Religion", true},
	"REPO":  {gedcom.TagRepository, "Repository", true},
	"RESI":  {gedcom.TagResidence, "Residence", true},
	"RESN":  {gedcom.TagRestriction, "Restriction", true},
	"RETI":  {gedcom.TagRetirement, "Retirement", true},
	"RFN":   {gedcom.TagRecordFileNumber, "RecordFileNumber", true},
	"RIN":   {gedcom.TagRecordIDNumber, "RecordIDNumber", true},
	"ROLE":  {gedcom.TagRole, "Role", true},
	"ROMN":  {gedcom.TagRomanized, "Romanized", true},
	"SEX":   {gedcom.TagSex, "Sex", true},
	"SLGC":  {gedcom.TagSealingChild, "SealingChild", true},
	"SLGS":  {gedcom.TagSealingSpouse, "SealingSpouse", true},
	"SOUR":  {gedcom.TagSource, "Source", true},
	"SPFX":  {gedcom.TagSurnamePrefix, "SurnamePrefix", true},
	"SSN":   {gedcom.TagSocialSecurityNumber, "SocialSecurityNumber", true},
	"STAE":  {gedcom.TagState, "State", true},
	"STAT":  {gedcom.TagStatus, "Status", true},
	"SUBM":  {gedcom.TagSubmitter, "Submitter", true},
	"SUBN":  {gedcom.TagSubmission, "Submission", true},
	"SURN":  {gedcom.TagSurname, "Surname", true},
	"TEMP":  {gedcom.TagTemple, "Temple", true},
	"TEXT":  {gedcom.TagText, "Text", true},
	"TIME":  {gedcom.TagTime, "Time", true},
	"TITL":  {gedcom.TagTitle, "Title", true},
	"TRLR":  {gedcom.TagTrailer, "Trailer", true},
	"TYPE":  {gedcom.TagType, "Type", true},
	"VERS":  {gedcom.TagVersion, "Version", true},
	"WIFE":  {gedcom.TagWife, "Wife", true},
	"WILL":  {gedcom.TagWill, "Will", true},
	"WWW":   {gedcom.TagWWW, "WWW", true},
}

func TestTags(t *testing.T) {
	t.Run("AllUnique", func(t *testing.T) {
		unique := map[gedcom.Tag]bool{}
		for _, actual := range tagTests {
			_, exists := unique[actual.tag]
			assert.False(t, exists, string(actual.tag))
			unique[actual.tag] = true
		}
	})

	for expected, actual := range tagTests {
		t.Run(expected, func(t *testing.T) {
			assert.Equal(t, gedcom.Tag(expected), actual.tag)
		})
	}
}

func TestTag_String(t *testing.T) {
	t.Run("Unknown", func(t *testing.T) {
		assert.Equal(t, gedcom.Tag("FOOBAR").String(), "FOOBAR")
	})

	for expected, actual := range tagTests {
		t.Run(expected, func(t *testing.T) {
			assert.Equal(t, gedcom.Tag(expected).String(), actual.tag.String())
		})
	}
}

func TestTag_IsOfficial(t *testing.T) {
	for expected, actual := range tagTests {
		t.Run(expected, func(t *testing.T) {
			assert.Equal(t, gedcom.Tag(expected).IsOfficial(),
				actual.isOfficial)
		})
	}
}
