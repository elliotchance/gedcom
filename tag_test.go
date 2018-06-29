package gedcom_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/elliotchance/gedcom"
)

var tagTests = map[string]struct {
	tag        gedcom.Tag
	s          string
	isOfficial bool
}{
	"_COR":  {gedcom.Coordinates, "Coordinates", false},
	"_CRE":  {gedcom.Created, "Created", false},
	"_FID":  {gedcom.FamilySearchID, "FamilySearchID", false},
	"_LAD":  {gedcom.LatitudeDegrees, "LatitudeDegrees", false},
	"_LAM":  {gedcom.LatitudeMinutes, "LatitudeMinutes", false},
	"_LAS":  {gedcom.LatitudeSeconds, "LatitudeSeconds", false},
	"_LOD":  {gedcom.LongitudeDegress, "LongitudeDegress", false},
	"_LOM":  {gedcom.LongitudeMinutes, "LongitudeMinutes", false},
	"_LON":  {gedcom.LongitudeNorth, "LongitudeNorth", false},
	"_LOS":  {gedcom.LongitudeSeconds, "LongitudeSeconds", false},
	"ABBR":  {gedcom.Abbreviation, "Abbreviation", true},
	"ADDR":  {gedcom.Address, "Address", true},
	"ADOP":  {gedcom.Adoption, "Adoption", true},
	"ADR1":  {gedcom.Address1, "Address1", true},
	"ADR2":  {gedcom.Address2, "Address2", true},
	"AFN":   {gedcom.AncestralFileNumber, "AncestralFileNumber", true},
	"AGE":   {gedcom.Age, "Age", true},
	"AGNC":  {gedcom.Agency, "Agency", true},
	"ALIA":  {gedcom.Alias, "Alias", true},
	"ANCE":  {gedcom.Ancestors, "Ancestors", true},
	"ANCI":  {gedcom.AncestorsInterest, "AncestorsInterest", true},
	"ANUL":  {gedcom.Annulment, "Annulment", true},
	"ASSO":  {gedcom.Associates, "Associates", true},
	"AUTH":  {gedcom.Author, "Author", true},
	"BAPL":  {gedcom.LDSBaptism, "LDSBaptism", true},
	"BAPM":  {gedcom.Baptism, "Baptism", true},
	"BARM":  {gedcom.BarMitzvah, "BarMitzvah", true},
	"BASM":  {gedcom.BasMitzvah, "BasMitzvah", true},
	"BIRT":  {gedcom.Birth, "Birth", true},
	"BLES":  {gedcom.Blessing, "Blessing", true},
	"BLOB":  {gedcom.BinaryObject, "BinaryObject", true},
	"BURI":  {gedcom.Burial, "Burial", true},
	"CALN":  {gedcom.CallNumber, "CallNumber", true},
	"CAST":  {gedcom.Caste, "Caste", true},
	"CAUS":  {gedcom.Cause, "Cause", true},
	"CENS":  {gedcom.Census, "Census", true},
	"CHAN":  {gedcom.Change, "Change", true},
	"CHAR":  {gedcom.CharacterSet, "CharacterSet", true},
	"CHIL":  {gedcom.Child, "Child", true},
	"CHR":   {gedcom.Christening, "Christening", true},
	"CHRA":  {gedcom.AdultChristening, "AdultChristening", true},
	"CITY":  {gedcom.City, "City", true},
	"CONC":  {gedcom.Concatenation, "Concatenation", true},
	"CONF":  {gedcom.Confirmation, "Confirmation", true},
	"CONL":  {gedcom.LDSConfirmation, "LDSConfirmation", true},
	"CONT":  {gedcom.Continued, "Continued", true},
	"COPR":  {gedcom.Copyright, "Copyright", true},
	"CORP":  {gedcom.Corporate, "Corporate", true},
	"CREM":  {gedcom.Cremation, "Cremation", true},
	"CTRY":  {gedcom.Country, "Country", true},
	"DATA":  {gedcom.Data, "Data", true},
	"DATE":  {gedcom.Date, "Date", true},
	"DEAT":  {gedcom.Death, "Death", true},
	"DESC":  {gedcom.Descendants, "Descendants", true},
	"DESI":  {gedcom.DescendantsInterest, "DescendantsInterest", true},
	"DEST":  {gedcom.Destination, "Destination", true},
	"DIV":   {gedcom.Divorce, "Divorce", true},
	"DIVF":  {gedcom.DivorceFiled, "DivorceFiled", true},
	"DSCR":  {gedcom.PhysicalDescription, "PhysicalDescription", true},
	"EDUC":  {gedcom.Education, "Education", true},
	"EMAIL": {gedcom.Email, "Email", true},
	"EMIG":  {gedcom.Emigration, "Emigration", true},
	"ENDL":  {gedcom.Endowment, "Endowment", true},
	"ENGA":  {gedcom.Engagement, "Engagement", true},
	"EVEN":  {gedcom.Event, "Event", true},
	"FACT":  {gedcom.Fact, "Fact", true},
	"FAM":   {gedcom.Family, "Family", true},
	"FAMC":  {gedcom.FamilyChild, "FamilyChild", true},
	"FAMF":  {gedcom.FamilyFile, "FamilyFile", true},
	"FAMS":  {gedcom.FamilySpouse, "FamilySpouse", true},
	"FAX":   {gedcom.Fax, "Fax", true},
	"FCOM":  {gedcom.FirstCommunion, "FirstCommunion", true},
	"FILE":  {gedcom.File, "File", true},
	"FONE":  {gedcom.Phonetic, "Phonetic", true},
	"FORM":  {gedcom.Format, "Format", true},
	"GEDC":  {gedcom.GedcomInformation, "GedcomInformation", true},
	"GIVN":  {gedcom.GivenName, "GivenName", true},
	"GRAD":  {gedcom.Graduation, "Graduation", true},
	"HEAD":  {gedcom.Header, "Header", true},
	"HUSB":  {gedcom.Husband, "Husband", true},
	"IDNO":  {gedcom.IdentityNumber, "IdentityNumber", true},
	"IMMI":  {gedcom.Immigration, "Immigration", true},
	"INDI":  {gedcom.Individual, "Individual", true},
	"LANG":  {gedcom.Language, "Language", true},
	"LATI":  {gedcom.Latitude, "Latitude", true},
	"LEGA":  {gedcom.Legatee, "Legatee", true},
	"LONG":  {gedcom.Longitude, "Longitude", true},
	"MAP":   {gedcom.Map, "Map", true},
	"MARB":  {gedcom.MarriageBann, "MarriageBann", true},
	"MARC":  {gedcom.MarriageContract, "MarriageContract", true},
	"MARL":  {gedcom.MarriageLicence, "MarriageLicence", true},
	"MARR":  {gedcom.Marriage, "Marriage", true},
	"MARS":  {gedcom.MarriageSettlement, "MarriageSettlement", true},
	"MEDI":  {gedcom.Media, "Media", true},
	"NAME":  {gedcom.Name, "Name", true},
	"NATI":  {gedcom.Nationality, "Nationality", true},
	"NATU":  {gedcom.Naturalization, "Naturalization", true},
	"NCHI":  {gedcom.ChildrenCount, "ChildrenCount", true},
	"NICK":  {gedcom.Nickname, "Nickname", true},
	"NMR":   {gedcom.MarriageCount, "MarriageCount", true},
	"NOTE":  {gedcom.Note, "Note", true},
	"NPFX":  {gedcom.NamePrefix, "NamePrefix", true},
	"NSFX":  {gedcom.NameSuffix, "NameSuffix", true},
	"OBJE":  {gedcom.Object, "Object", true},
	"OCCU":  {gedcom.Occupation, "Occupation", true},
	"ORDI":  {gedcom.Ordinance, "Ordinance", true},
	"ORDN":  {gedcom.Ordination, "Ordination", true},
	"PAGE":  {gedcom.Page, "Page", true},
	"PEDI":  {gedcom.Pedigree, "Pedigree", true},
	"PHON":  {gedcom.Phone, "Phone", true},
	"PLAC":  {gedcom.Place, "Place", true},
	"POST":  {gedcom.PostalCode, "PostalCode", true},
	"PROB":  {gedcom.Probate, "Probate", true},
	"PROP":  {gedcom.Property, "Property", true},
	"PUBL":  {gedcom.Publication, "Publication", true},
	"QUAY":  {gedcom.QualityOfData, "QualityOfData", true},
	"REFN":  {gedcom.Reference, "Reference", true},
	"RELA":  {gedcom.Relationship, "Relationship", true},
	"RELI":  {gedcom.Religion, "Religion", true},
	"REPO":  {gedcom.Repository, "Repository", true},
	"RESI":  {gedcom.Residence, "Residence", true},
	"RESN":  {gedcom.Restriction, "Restriction", true},
	"RETI":  {gedcom.Retirement, "Retirement", true},
	"RFN":   {gedcom.RecordFileNumber, "RecordFileNumber", true},
	"RIN":   {gedcom.RecordIDNumber, "RecordIDNumber", true},
	"ROLE":  {gedcom.Role, "Role", true},
	"ROMN":  {gedcom.Romanized, "Romanized", true},
	"SEX":   {gedcom.Sex, "Sex", true},
	"SLGC":  {gedcom.SealingChild, "SealingChild", true},
	"SLGS":  {gedcom.SealingSpouse, "SealingSpouse", true},
	"SOUR":  {gedcom.Source, "Source", true},
	"SPFX":  {gedcom.SurnamePrefix, "SurnamePrefix", true},
	"SSN":   {gedcom.SocialSecurityNumber, "SocialSecurityNumber", true},
	"STAE":  {gedcom.State, "State", true},
	"STAT":  {gedcom.Status, "Status", true},
	"SUBM":  {gedcom.Submitter, "Submitter", true},
	"SUBN":  {gedcom.Submission, "Submission", true},
	"SURN":  {gedcom.Surname, "Surname", true},
	"TEMP":  {gedcom.Temple, "Temple", true},
	"TEXT":  {gedcom.Text, "Text", true},
	"TIME":  {gedcom.Time, "Time", true},
	"TITL":  {gedcom.Title, "Title", true},
	"TRLR":  {gedcom.Trailer, "Trailer", true},
	"TYPE":  {gedcom.Type, "Type", true},
	"VERS":  {gedcom.Version, "Version", true},
	"WIFE":  {gedcom.Wife, "Wife", true},
	"WILL":  {gedcom.Will, "Will", true},
	"WWW":   {gedcom.WWW, "WWW", true},
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
