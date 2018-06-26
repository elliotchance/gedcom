package gedcom_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/elliotchance/gedcom"
)

var tagTests = map[string]struct {
	tag gedcom.Tag
	s   string
}{
	"ABBR":  {gedcom.Abbreviation, "Abbreviation"},
	"ADDR":  {gedcom.Address, "Address"},
	"ADR1":  {gedcom.Address1, "Address1"},
	"ADR2":  {gedcom.Address2, "Address2"},
	"ADOP":  {gedcom.Adoption, "Adoption"},
	"AFN":   {gedcom.AncestralFileNumber, "AncestralFileNumber"},
	"AGE":   {gedcom.Age, "Age"},
	"AGNC":  {gedcom.Agency, "Agency"},
	"ALIA":  {gedcom.Alias, "Alias"},
	"ANCE":  {gedcom.Ancestors, "Ancestors"},
	"ANCI":  {gedcom.AncestorsInterest, "AncestorsInterest"},
	"ANUL":  {gedcom.Annulment, "Annulment"},
	"ASSO":  {gedcom.Associates, "Associates"},
	"AUTH":  {gedcom.Author, "Author"},
	"BAPL":  {gedcom.LDSBaptism, "LDSBaptism"},
	"BAPM":  {gedcom.Baptism, "Baptism"},
	"BARM":  {gedcom.BarMitzvah, "BarMitzvah"},
	"BASM":  {gedcom.BasMitzvah, "BasMitzvah"},
	"BIRT":  {gedcom.Birth, "Birth"},
	"BLES":  {gedcom.Blessing, "Blessing"},
	"BLOB":  {gedcom.BinaryObject, "BinaryObject"},
	"BURI":  {gedcom.Burial, "Burial"},
	"CALN":  {gedcom.CallNumber, "CallNumber"},
	"CAST":  {gedcom.Caste, "Caste"},
	"CAUS":  {gedcom.Cause, "Cause"},
	"CENS":  {gedcom.Census, "Census"},
	"CHAN":  {gedcom.Change, "Change"},
	"CHAR":  {gedcom.CharacterSet, "CharacterSet"},
	"CHIL":  {gedcom.Child, "Child"},
	"CHR":   {gedcom.Christening, "Christening"},
	"CHRA":  {gedcom.AdultChristening, "AdultChristening"},
	"CITY":  {gedcom.City, "City"},
	"CONC":  {gedcom.Concatenation, "Concatenation"},
	"CONF":  {gedcom.Confirmation, "Confirmation"},
	"CONL":  {gedcom.LDSConfirmation, "LDSConfirmation"},
	"CONT":  {gedcom.Continued, "Continued"},
	"COPR":  {gedcom.Copyright, "Copyright"},
	"CORP":  {gedcom.Corporate, "Corporate"},
	"CREM":  {gedcom.Cremation, "Cremation"},
	"CTRY":  {gedcom.Country, "Country"},
	"DATA":  {gedcom.Data, "Data"},
	"DATE":  {gedcom.Date, "Date"},
	"DEAT":  {gedcom.Death, "Death"},
	"DESC":  {gedcom.Descendants, "Descendants"},
	"DESI":  {gedcom.DescendantsInterest, "DescendantsInterest"},
	"DEST":  {gedcom.Destination, "Destination"},
	"DIV":   {gedcom.Divorce, "Divorce"},
	"DIVF":  {gedcom.DivorceFiled, "DivorceFiled"},
	"DSCR":  {gedcom.PhysicalDescription, "PhysicalDescription"},
	"EDUC":  {gedcom.Education, "Education"},
	"EMAIL": {gedcom.Email, "Email"},
	"EMIG":  {gedcom.Emigration, "Emigration"},
	"ENDL":  {gedcom.Endowment, "Endowment"},
	"ENGA":  {gedcom.Engagement, "Engagement"},
	"EVEN":  {gedcom.Event, "Event"},
	"FACT":  {gedcom.Fact, "Fact"},
	"FAM":   {gedcom.Family, "Family"},
	"FAMC":  {gedcom.FamilyChild, "FamilyChild"},
	"FAMF":  {gedcom.FamilyFile, "FamilyFile"},
	"FAMS":  {gedcom.FamilySpouse, "FamilySpouse"},
	"FAX":   {gedcom.Fax, "Fax"},
	"FCOM":  {gedcom.FirstCommunion, "FirstCommunion"},
	"FILE":  {gedcom.File, "File"},
	"FONE":  {gedcom.Phonetic, "Phonetic"},
	"FORM":  {gedcom.Format, "Format"},
	"GEDC":  {gedcom.GedcomInformation, "GedcomInformation"},
	"GIVN":  {gedcom.GivenName, "GivenName"},
	"GRAD":  {gedcom.Graduation, "Graduation"},
	"HEAD":  {gedcom.Header, "Header"},
	"HUSB":  {gedcom.Husband, "Husband"},
	"IDNO":  {gedcom.IdentityNumber, "IdentityNumber"},
	"IMMI":  {gedcom.Immigration, "Immigration"},
	"INDI":  {gedcom.Individual, "Individual"},
	"LANG":  {gedcom.Language, "Language"},
	"LATI":  {gedcom.Latitude, "Latitude"},
	"LEGA":  {gedcom.Legatee, "Legatee"},
	"LONG":  {gedcom.Longitude, "Longitude"},
	"MAP":   {gedcom.Map, "Map"},
	"MARB":  {gedcom.MarriageBann, "MarriageBann"},
	"MARC":  {gedcom.MarriageContract, "MarriageContract"},
	"MARL":  {gedcom.MarriageLicence, "MarriageLicence"},
	"MARR":  {gedcom.Marriage, "Marriage"},
	"MARS":  {gedcom.MarriageSettlement, "MarriageSettlement"},
	"MEDI":  {gedcom.Media, "Media"},
	"NAME":  {gedcom.Name, "Name"},
	"NATI":  {gedcom.Nationality, "Nationality"},
	"NATU":  {gedcom.Naturalization, "Naturalization"},
	"NCHI":  {gedcom.ChildrenCount, "ChildrenCount"},
	"NICK":  {gedcom.Nickname, "Nickname"},
	"NMR":   {gedcom.MarriageCount, "MarriageCount"},
	"NOTE":  {gedcom.Note, "Note"},
	"NPFX":  {gedcom.NamePrefix, "NamePrefix"},
	"NSFX":  {gedcom.NameSuffix, "NameSuffix"},
	"OBJE":  {gedcom.Object, "Object"},
	"OCCU":  {gedcom.Occupation, "Occupation"},
	"ORDI":  {gedcom.Ordinance, "Ordinance"},
	"ORDN":  {gedcom.Ordination, "Ordination"},
	"PAGE":  {gedcom.Page, "Page"},
	"PEDI":  {gedcom.Pedigree, "Pedigree"},
	"PHON":  {gedcom.Phone, "Phone"},
	"PLAC":  {gedcom.Place, "Place"},
	"POST":  {gedcom.PostalCode, "PostalCode"},
	"PROB":  {gedcom.Probate, "Probate"},
	"PROP":  {gedcom.Property, "Property"},
	"PUBL":  {gedcom.Publication, "Publication"},
	"QUAY":  {gedcom.QualityOfData, "QualityOfData"},
	"REFN":  {gedcom.Reference, "Reference"},
	"RELA":  {gedcom.Relationship, "Relationship"},
	"RELI":  {gedcom.Religion, "Religion"},
	"REPO":  {gedcom.Repository, "Repository"},
	"RESI":  {gedcom.Residence, "Residence"},
	"RESN":  {gedcom.Restriction, "Restriction"},
	"RETI":  {gedcom.Retirment, "Retirment"},
	"RFN":   {gedcom.RecordFileNumber, "RecordFileNumber"},
	"RIN":   {gedcom.RecordIDNumber, "RecordIDNumber"},
	"ROLE":  {gedcom.Role, "Role"},
	"ROMN":  {gedcom.Romanized, "Romanized"},
	"SEX":   {gedcom.Sex, "Sex"},
	"SLGC":  {gedcom.SealingChild, "SealingChild"},
	"SLGS":  {gedcom.SealingSpouse, "SealingSpouse"},
	"SOUR":  {gedcom.Source, "Source"},
	"SPFX":  {gedcom.SurnamePrefix, "SurnamePrefix"},
	"SSN":   {gedcom.SocialSecurityNumber, "SocialSecurityNumber"},
	"STAE":  {gedcom.State, "State"},
	"STAT":  {gedcom.Status, "Status"},
	"SUBM":  {gedcom.Submitter, "Submitter"},
	"SUBN":  {gedcom.Submission, "Submission"},
	"SURN":  {gedcom.Surname, "Surname"},
	"TEMP":  {gedcom.Temple, "Temple"},
	"TEXT":  {gedcom.Text, "Text"},
	"TIME":  {gedcom.Time, "Time"},
	"TITL":  {gedcom.Title, "Title"},
	"TRLR":  {gedcom.Trailer, "Trailer"},
	"TYPE":  {gedcom.Type, "Type"},
	"VERS":  {gedcom.Version, "Version"},
	"WIFE":  {gedcom.Wife, "Wife"},
	"WWW":   {gedcom.WWW, "WWW"},
	"WILL":  {gedcom.Will, "Will"},
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
