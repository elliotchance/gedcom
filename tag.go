package gedcom

import "fmt"

const (
	// A constant that means none of the below flags are enabled.
	tagOptionNone = 0

	// tagOptionEvent describes is this tag is a type of event, such as a birth
	// (BIRT) for an individual or a marriage for a family (MARR). Generally
	// speaking all events have an optional date (DATE) and place (PLAC) as
	// child nodes.
	tagOptionEvent = 1 << iota
)

// knownTags is a private cache used to quickly and easily convert strings to
// tags. It is initialised by newTag.
var knownTags map[string]Tag

// Tag is the type of node. You should not initialise a Tag instance manually,
// but rather use TagFromString.
type Tag struct {
	// isKnown will be true when using one of the Tag constants. You should use
	// this if you have loaded a Tag from a string and you to check if all of
	// the other fields are loaded.
	isKnown bool

	// tag is the raw tag as it is represented in the GEDCOM file, like "INDI"
	// for an individual. Generally speaking non-standard (unofficial) tags
	// should be prefixed with an underscore. However, this is not guaranteed in
	// all cases so you should use the IsOfficial method.
	//
	// Tags are considered equal when their respective Tag properties are the
	// same, regardless of another other properties.
	tag string

	// name is the pretty name that can be used when printing tags to the user.
	// For example the "BARM" tag has a Name of "Bar Mitzvah".
	//
	// When using TagFromString on a tag that is not known the Name will be the
	// same as Tag.
	name string

	// A combination of flags from the tagOption constants.
	options int
}

// http://wiki-en.genealogy.net/GEDCOM-Tags
//
// Definition of all standard GEDCOM-Tags
//
// List of all field names (tags), used in GEDCOM Specification 5.5, or (with a
// special comment) or added/deleted in GEDCOM 5.5.1. These field names are used
// in a hierarchical structure to describe e.g. single persons in connection
// with their families. Field names may have different meaning and content
// depending on their position in the data structure.
//
// GEDCOM-standard permits user defined field names (one of the major causes of
// "misunderstandings" of different GEDCOM-compliant programs!). They have to
// begin with an underscore _ .
var (
	// A short name of a title, description, or name.
	TagAbbreviation = newTag("ABBR", "Abbreviation", tagOptionNone)

	// The contemporary place, usually required for postal purposes, of an
	// individual, a submitter of information, a repository, a business, a
	// school, or a company.
	TagAddress = newTag("ADDR", "Address", tagOptionNone)

	// The first line of an address.
	TagAddress1 = newTag("ADR1", "Address Line 1", tagOptionNone)

	// The second line of an address.
	TagAddress2 = newTag("ADR2", "Address Line 2", tagOptionNone)

	// Pertaining to creation of a child-parent relationship that does not exist
	// biologically.
	TagAdoption = newTag("ADOP", "Adoption", tagOptionEvent)

	// Ancestral File Number, a unique permanent record file number of an
	// individual record stored in Ancestral File.
	TagAncestralFileNumber = newTag("AFN", "Ancestral File Number", tagOptionNone)

	// The age of the individual at the time an event occurred, or the age
	// listed in the document.
	TagAge = newTag("AGE", "Age", tagOptionNone)

	// The institution or individual having authority and/or responsibility to
	// manage or govern.
	TagAgency = newTag("AGNC", "Agency", tagOptionNone)

	// An indicator to link different record descriptions of a person who may be
	// the same person.
	TagAlias = newTag("ALIA", "Alias", tagOptionNone)

	// Pertaining to forbearers of an individual.
	TagAncestors = newTag("ANCE", "Ancestors", tagOptionNone)

	// Indicates an interest in additional research for ancestors of this
	// individual. (See also DESI)
	TagAncestorsInterest = newTag("ANCI", "Ancestors Interest", tagOptionNone)

	// Declaring a marriage void from the beginning (never existed).
	TagAnnulment = newTag("ANUL", "Annulment", tagOptionEvent)

	// An indicator to link friends, neighbors, relatives, or associates of an
	// individual.
	TagAssociates = newTag("ASSO", "Associates", tagOptionNone)

	// The name of the individual who created or compiled information.
	TagAuthor = newTag("AUTH", "Author", tagOptionNone)

	// The event of baptism performed at age eight or later by priesthood
	// authority of the LDS Church. (See also BAPM)
	TagLDSBaptism = newTag("BAPL", "LDS Baptism", tagOptionEvent)

	// The event of baptism (not LDS), performed in infancy or later. (See also
	// BAPL and CHR)
	TagBaptism = newTag("BAPM", "Baptism", tagOptionEvent)

	// The ceremonial event held when a Jewish boy reaches age 13.
	TagBarMitzvah = newTag("BARM", "Bar Mitzvah", tagOptionEvent)

	// The ceremonial event held when a Jewish girl reaches age 13, also known
	// as "Bat Mitzvah."
	TagBasMitzvah = newTag("BASM", "Bas Mitzvah", tagOptionEvent)

	// See BirthNode.
	TagBirth = newTag("BIRT", "Birth", tagOptionEvent)

	// A religious event of bestowing divine care or intercession. Sometimes
	// given in connection with a naming ceremony.
	TagBlessing = newTag("BLES", "Blessing", tagOptionEvent)

	// A grouping of data used as input to a multimedia system that processes
	// binary data to represent images, sound, and video. Deleted in Gedcom
	// 5.5.1
	TagBinaryObject = newTag("BLOB", "Binary Object", tagOptionNone)

	// The event of the proper disposing of the mortal remains of a deceased
	// person.
	TagBurial = newTag("BURI", "Burial", tagOptionEvent)

	// The number used by a repository to identify the specific items in its
	// collections.
	TagCallNumber = newTag("CALN", "Call Number", tagOptionNone)

	// The name of an individual's rank or status in society, based on racial or
	// religious differences, or differences in wealth, inherited rank,
	// profession, occupation, etc.
	TagCaste = newTag("CAST", "Caste", tagOptionNone)

	// A description of the cause of the associated event or fact, such as the
	// cause of death.
	TagCause = newTag("CAUS", "Cause", tagOptionNone)

	// The event of the periodic count of the population for a designated
	// locality, such as a national or state Census.
	TagCensus = newTag("CENS", "Census", tagOptionEvent)

	// Indicates a change, correction, or modification. Typically used in
	// connection with a DATE to specify when a change in information occurred.
	TagChange = newTag("CHAN", "Change", tagOptionNone)

	// An indicator of the character set used in writing this automated
	// information.
	TagCharacterSet = newTag("CHAR", "Character Set", tagOptionNone)

	// The natural, adopted, or sealed (LDS) child of a father and a mother.
	TagChild = newTag("CHIL", "Child", tagOptionNone)

	// The religious event (not LDS) of baptizing and/or naming a child.
	TagChristening = newTag("CHR", "Christening", tagOptionEvent)

	// The religious event (not LDS) of baptizing and/or naming an adult person.
	TagAdultChristening = newTag("CHRA", "Adult Christening", tagOptionEvent)

	// A lower level jurisdictional unit. Normally an incorporated municipal
	// unit.
	TagCity = newTag("CITY", "City", tagOptionNone)

	// An indicator that additional data belongs to the superior value. The
	// information from the CONC value is to be connected to the value of the
	// superior preceding line without a space and without a carriage return
	// and/or new line character. Values that are split for a CONC tag must
	// always be split at a non-space. If the value is split on a space the
	// space will be lost when concatenation takes place. This is because of the
	// treatment that spaces get as a GEDCOM delimiter, many GEDCOM values are
	// trimmed of trailing spaces and some systems look for the first non-space
	// starting after the tag to determine the beginning of the value.
	TagConcatenation = newTag("CONC", "Concatenation", tagOptionNone)

	// The religious event (not LDS) of conferring the gift of the Holy Ghost
	// and, among protestants, full church membership.
	TagConfirmation = newTag("CONF", "Confirmation", tagOptionEvent)

	// The religious event by which a person receives membership in the LDS
	// Church.
	TagLDSConfirmation = newTag("CONL", "LDS Confirmation", tagOptionEvent)

	//  An indicator that additional data belongs to the superior value. The
	// information from the CONT value is to be connected to the value of the
	// superior preceding line with a carriage return and/or new line character.
	// Leading spaces could be important to the formatting of the resultant
	// text. When importing values from CONT lines the reader should assume only
	// one delimiter character following the CONT tag. Assume that the rest of
	// the leading spaces are to be a part of the value.
	TagContinued = newTag("CONT", "Continued", tagOptionNone)

	// A statement that accompanies data to protect it from unlawful duplication
	// and distribution.
	TagCopyright = newTag("COPR", "Copyright", tagOptionNone)

	// A name of an institution, agency, corporation, or company.
	TagCorporate = newTag("CORP", "Corporate", tagOptionNone)

	// Disposal of the remains of a person's body by fire.
	TagCremation = newTag("CREM", "Cremation", tagOptionEvent)

	// The name or code of the country.
	TagCountry = newTag("CTRY", "Country", tagOptionNone)

	// Pertaining to stored automated information.
	TagData = newTag("DATA", "Data", tagOptionNone)

	// The time of an event in a calendar format.
	TagDate = newTag("DATE", "Date", tagOptionNone)

	// The event when mortal life terminates.
	TagDeath = newTag("DEAT", "Death", tagOptionEvent)

	// Pertaining to offspring of an individual.
	TagDescendants = newTag("DESC", "Descendants", tagOptionNone)

	// Indicates an interest in research to identify additional descendants of
	// this individual. (See also ANCI)
	TagDescendantsInterest = newTag("DESI", "Descendants Interest", tagOptionNone)

	// A system receiving data.
	TagDestination = newTag("DEST", "Destination", tagOptionNone)

	// An event of dissolving a marriage through civil action.
	TagDivorce = newTag("DIV", "Divorce", tagOptionEvent)

	// An event of filing for a divorce by a spouse.
	TagDivorceFiled = newTag("DIVF", "Divorce Filed", tagOptionEvent)

	// The physical characteristics of a person, place, or thing.
	TagPhysicalDescription = newTag("DSCR", "Physical Description", tagOptionNone)

	// Indicator of a level of education attained.
	TagEducation = newTag("EDUC", "Education", tagOptionNone)

	// An electronic address that can be used for contact such as an email
	// address. New in Gedcom 5.5.1.
	TagEmail = newTag("EMAIL", "Email", tagOptionNone)

	// An event of leaving one's homeland with the intent of residing elsewhere.
	TagEmigration = newTag("EMIG", "Emigration", tagOptionEvent)

	// A religious event where an endowment ordinance for an individual was
	// performed by priesthood authority in an LDS temple.
	TagEndowment = newTag("ENDL", "Endowment", tagOptionEvent)

	// An event of recording or announcing an agreement between two people to
	// become married.
	TagEngagement = newTag("ENGA", "Engagement", tagOptionEvent)

	// A noteworthy happening related to an individual, a group, or an
	// organization.
	TagEvent = newTag("EVEN", "Event", tagOptionEvent)

	// Pertaining to a noteworthy attribute or fact concerning an individual, a
	// group, or an organization. A structure is usually qualified or classified
	// by a subordinate use of the TYPE tag. New in Gedcom 5.5.1.
	TagFact = newTag("FACT", "Fact", tagOptionNone)

	// Identifies a legal, common law, or other customary relationship of man
	// and woman and their children, if any, or a family created by virtue of
	// the birth of a child to its biological father and mother.
	TagFamily = newTag("FAM", "Family", tagOptionNone)

	// Identifies the family in which an individual appears as a child.
	TagFamilyChild = newTag("FAMC", "Family Child", tagOptionNone)

	// Pertaining to, or the name of, a family file. Names stored in a file that
	// are assigned to a family for doing temple ordinance work.
	TagFamilyFile = newTag("FAMF", "Family File", tagOptionNone)

	// Identifies the family in which an individual appears as a spouse.
	TagFamilySpouse = newTag("FAMS", "Family Spouse", tagOptionNone)

	// A FAX telephone number appropriate for sending data facsimiles. New in
	// Gedcom 5.5.1.
	TagFax = newTag("FAX", "Fax", tagOptionNone)

	// A religious rite, the first act of sharing in the Lord's supper as part
	// of church worship.
	TagFirstCommunion = newTag("FCOM", "First Communion", tagOptionEvent)

	// An information storage place that is ordered and arranged for
	// preservation and reference.
	TagFile = newTag("FILE", "File", tagOptionNone)

	// A phonetic variation of a superior text string. New in Gedcom 5.5.1
	TagPhonetic = newTag("FONE", "Phonetic", tagOptionNone)

	// An assigned name given to a consistent format in which information can be
	// conveyed.
	TagFormat = newTag("FORM", "Format", tagOptionNone)

	// Information about the use of GEDCOM in a transmission.
	TagGedcomInformation = newTag("GEDC", "GEDCOM Information", tagOptionNone)

	// A given or earned name used for official identification of a person. It
	// is also commonly known as the "first name".
	//
	// The NameNode provides a GivenName() function.
	TagGivenName = newTag("GIVN", "Given Name", tagOptionNone)

	// An event of awarding educational diplomas or degrees to individuals.
	TagGraduation = newTag("GRAD", "Graduation", tagOptionEvent)

	// Identifies information pertaining to an entire GEDCOM transmission.
	TagHeader = newTag("HEAD", "Header", tagOptionNone)

	// An individual in the family role of a married man or father.
	TagHusband = newTag("HUSB", "Husband", tagOptionNone)

	// A number assigned to identify a person within some significant external
	// system.
	TagIdentityNumber = newTag("IDNO", "Identity Number", tagOptionNone)

	// An event of entering into a new locality with the intent of residing
	// there.
	TagImmigration = newTag("IMMI", "Immigration", tagOptionEvent)

	// A person.
	TagIndividual = newTag("INDI", "Individual", tagOptionNone)

	// The name of the language used in a communication or transmission of
	// information.
	TagLanguage = newTag("LANG", "Language", tagOptionNone)

	// A value indicating a coordinate position on a line, plane, or space. New
	// in Gedcom 5.5.1.
	TagLatitude = newTag("LATI", "Latitude", tagOptionNone)

	// A role of an individual acting as a person receiving a bequest or legal
	// devise.
	TagLegatee = newTag("LEGA", "Legatee", tagOptionNone)

	// A value indicating a coordinate position on a line, plane, or space. New
	// in Gedcom 5.5.1.
	TagLongitude = newTag("LONG", "Longitude", tagOptionNone)

	// Pertains to a representation of measurements usually presented in a
	// graphical form. New in Gedcom 5.5.1
	TagMap = newTag("MAP", "Map", tagOptionNone)

	// An event of an official public notice given that two people intend to
	// marry.
	TagMarriageBann = newTag("MARB", "Marriage Bann", tagOptionEvent)

	// An event of recording a formal agreement of marriage, including the
	// prenuptial agreement in which marriage partners reach agreement about the
	// property rights of one or both, securing property to their children.
	TagMarriageContract = newTag("MARC", "Marriage Contract", tagOptionEvent)

	// An event of obtaining a legal license to marry.
	TagMarriageLicence = newTag("MARL", "Marriage Licence", tagOptionEvent)

	// A legal, common-law, or customary event of creating a family unit of a
	// man and a woman as husband and wife.
	TagMarriage = newTag("MARR", "Marriage", tagOptionEvent)

	// An event of creating an agreement between two people contemplating
	// marriage, at which time they agree to release or modify property rights
	// that would otherwise arise from the marriage.
	TagMarriageSettlement = newTag("MARS", "Marriage Settlement", tagOptionEvent)

	// Identifies information about the media or having to do with the medium in
	// which information is stored.
	TagMedia = newTag("MEDI", "Media", tagOptionNone)

	// A word or combination of words used to help identify an individual,
	// title, or other item. More than one NAME line should be used for people
	// who were known by multiple names.
	//
	// NAME tags will be interpreted with the NameNode type.
	TagName = newTag("NAME", "Name", tagOptionNone)

	// The national heritage of an individual.
	TagNationality = newTag("NATI", "Nationality", tagOptionNone)

	// The event of obtaining citizenship.
	TagNaturalization = newTag("NATU", "Naturalization", tagOptionEvent)

	// The number of children that this person is known to be the parent of (all
	// marriages) when subordinate to an individual, or that belong to this
	// family when subordinate to a FAM_RECORD.
	TagChildrenCount = newTag("NCHI", "Children Count", tagOptionNone)

	// A descriptive or familiar that is used instead of, or in addition to,
	// one's proper name.
	TagNickname = newTag("NICK", "Nickname", tagOptionNone)

	// The number of times this person has participated in a family as a spouse
	// or parent.
	TagMarriageCount = newTag("NMR", "Marriage Count", tagOptionNone)

	// Additional information provided by the submitter for understanding the
	// enclosing data.
	TagNote = newTag("NOTE", "Note", tagOptionNone)

	// Text which appears on a name line before the given and surname parts of a
	// name. i.e. ( Lt. Cmndr. ) Joseph /Allen/ jr. In this example Lt. Cmndr.
	// is considered as the name prefix portion.
	//
	// The NameNode provides a Prefix() function.
	TagNamePrefix = newTag("NPFX", "Name Prefix", tagOptionNone)

	// Text which appears on a name line after or behind the given and surname
	// parts of a name. i.e. Lt. Cmndr. Joseph /Allen/ ( jr. ) In this example
	// jr. is considered as the name suffix portion.
	//
	// The NameNode provides a Suffix() function.
	TagNameSuffix = newTag("NSFX", "Name Suffix", tagOptionNone)

	// Pertaining to a grouping of attributes used in describing something.
	// Usually referring to the data required to represent a multimedia object,
	// such an audio recording, a photograph of a person, or an image of a
	// document.
	TagObject = newTag("OBJE", "Object", tagOptionNone)

	// The type of work or profession of an individual.
	TagOccupation = newTag("OCCU", "Occupation", tagOptionNone)

	// Pertaining to a religious ordinance in general.
	TagOrdinance = newTag("ORDI", "Ordinance", tagOptionNone)

	// A religious event of receiving authority to act in religious matters.
	TagOrdination = newTag("ORDN", "Ordination", tagOptionEvent)

	// A number or description to identify where information can be found in a
	// referenced work.
	TagPage = newTag("PAGE", "Page", tagOptionNone)

	// Information pertaining to an individual to parent lineage chart.
	TagPedigree = newTag("PEDI", "Pedigree", tagOptionNone)

	// A unique number assigned to access a specific telephone.
	TagPhone = newTag("PHON", "Phone", tagOptionNone)

	// A jurisdictional name to identify the place or location of an event.
	TagPlace = newTag("PLAC", "Place", tagOptionNone)

	// A code used by a postal service to identify an area to facilitate mail
	// handling.
	TagPostalCode = newTag("POST", "Postal Code", tagOptionNone)

	// An event of judicial determination of the validity of a will. May
	// indicate several related court activities over several dates.
	TagProbate = newTag("PROB", "Probate", tagOptionEvent)

	// Pertaining to possessions such as real estate or other property of
	// interest.
	TagProperty = newTag("PROP", "Property", tagOptionNone)

	// Refers to when and/or were a work was published or created.
	TagPublication = newTag("PUBL", "Publication", tagOptionNone)

	// An assessment of the certainty of the evidence to support the conclusion
	// drawn from evidence.
	TagQualityOfData = newTag("QUAY", "Quality Of Data", tagOptionNone)

	// A description or number used to identify an item for filing, storage, or
	// other reference purposes.
	TagReference = newTag("REFN", "Reference", tagOptionNone)

	// A relationship value between the indicated contexts.
	TagRelationship = newTag("RELA", "Relationship", tagOptionNone)

	// A religious denomination to which a person is affiliated or for which a
	// record applies.
	TagReligion = newTag("RELI", "Religion", tagOptionNone)

	// An institution or person that has the specified item as part of their
	// collection(s).
	TagRepository = newTag("REPO", "Repository", tagOptionNone)

	// See ResidenceNode.
	TagResidence = newTag("RESI", "Residence", tagOptionEvent)

	// A processing indicator signifying access to information has been denied
	// or otherwise restricted.
	TagRestriction = newTag("RESN", "Restriction", tagOptionNone)

	// An event of exiting an occupational relationship with an employer after a
	// qualifying time period.
	TagRetirement = newTag("RETI", "Retirement", tagOptionEvent)

	// A permanent number assigned to a record that uniquely identifies it
	// within a known file.
	TagRecordFileNumber = newTag("RFN", "Record File Number", tagOptionNone)

	// A number assigned to a record by an originating automated system that can
	// be used by a receiving system to report results pertaining to that
	// record.
	TagRecordIDNumber = newTag("RIN", "Record ID Number", tagOptionNone)

	// A name given to a role played by an individual in connection with an
	// event.
	TagRole = newTag("ROLE", "Role", tagOptionNone)

	// A romanized variation of a superior text string. New in Gedcom 5.5.1.
	TagRomanized = newTag("ROMN", "Romanized", tagOptionNone)

	// Indicates the sex of an individual--male or female.
	TagSex = newTag("SEX", "Sex", tagOptionNone)

	// A religious event pertaining to the sealing of a child to his or her
	// parents in an LDS temple ceremony.
	TagSealingChild = newTag("SLGC", "Sealing Child", tagOptionEvent)

	// A religious event pertaining to the sealing of a husband and wife in an
	// LDS temple ceremony.
	TagSealingSpouse = newTag("SLGS", "Sealing Spouse", tagOptionEvent)

	// The initial or original material from which information was obtained.
	TagSource = newTag("SOUR", "Source", tagOptionNone)

	// A name piece used as a non-indexing pre-part of a surname.
	TagSurnamePrefix = newTag("SPFX", "Surname Prefix", tagOptionNone)

	// A number assigned by the United States Social Security Administration.
	// Used for tax identification purposes.
	TagSocialSecurityNumber = newTag("SSN", "Social Security Number", tagOptionNone)

	// A geographical division of a larger jurisdictional area, such as a State
	// within the United States of America.
	TagState = newTag("STAE", "State", tagOptionNone)

	// An assessment of the state or condition of something.
	TagStatus = newTag("STAT", "Status", tagOptionNone)

	// An individual or organization who contributes genealogical data to a file
	// or transfers it to someone else.
	TagSubmitter = newTag("SUBM", "Submitter", tagOptionNone)

	// Pertains to a collection of data issued for processing.
	TagSubmission = newTag("SUBN", "Submission", tagOptionNone)

	// A family name passed on or used by members of a family.
	//
	// The NameNode provides a Surname() function.
	TagSurname = newTag("SURN", "Surname", tagOptionNone)

	// The name or code that represents the name a temple of the LDS Church.
	TagTemple = newTag("TEMP", "Temple", tagOptionNone)

	// The exact wording found in an original source document.
	TagText = newTag("TEXT", "Text", tagOptionNone)

	// A time value in a 24-hour clock format, including hours, minutes, and
	// optional seconds, separated by a colon (:). Fractions of seconds are
	// shown in decimal notation.
	TagTime = newTag("TIME", "Time", tagOptionNone)

	// A description of a specific writing or other work, such as the title of a
	// book when used in a source context, or a formal designation used by an
	// individual in connection with positions of royalty or other social
	// status, such as Grand Duke.
	//
	// The NameNode provides a Title() function.
	TagTitle = newTag("TITL", "Title", tagOptionNone)

	// At level 0, specifies the end of a GEDCOM transmission.
	TagTrailer = newTag("TRLR", "Trailer", tagOptionNone)

	// A further qualification to the meaning of the associated superior tag.
	// The value does not have any computer processing reliability. It is more
	// in the form of a short one or two word note that should be displayed any
	// time the associated data is displayed.
	TagType = newTag("TYPE", "Type", tagOptionNone)

	// Indicates which version of a product, item, or publication is being used
	// or referenced.
	TagVersion = newTag("VERS", "Version", tagOptionNone)

	// An individual in the role as a mother and/or married woman.
	TagWife = newTag("WIFE", "Wife", tagOptionNone)

	// World Wide Web home page. New in Gedcom 5.5.1.
	TagWWW = newTag("WWW", "WWW", tagOptionNone)

	// A legal document treated as an event, by which a person disposes of his
	// or her estate, to take effect after death. The event date is the date the
	// will was signed while the person was alive. (See also PROBate)
	TagWill = newTag("WILL", "Will", tagOptionEvent)
)

var (
	// Unofficial. The unique identifier for the person on FamilySearch.org.
	// This has been seen exported from MacFamilyFree.
	UnofficialTagFamilySearchID = newTag("_FID", "FamilySearch ID", tagOptionNone)

	// Unofficial. Latitude degrees. This has been seen exported from
	// MacFamilyFree.
	UnofficialTagLatitudeDegrees = newTag("_LAD", "Latitude Degrees", tagOptionNone)

	// Unofficial. Latitude minutes. This has been seen exported from
	// MacFamilyFree.
	UnofficialTagLatitudeMinutes = newTag("_LAM", "Latitude Minutes", tagOptionNone)

	// Unofficial. Latitude seconds. This has been seen exported from
	// MacFamilyFree.
	UnofficialTagLatitudeSeconds = newTag("_LAS", "Latitude Seconds", tagOptionNone)

	// Unofficial. Longitude degrees. This has been seen exported from
	// MacFamilyFree.
	UnofficialTagLongitudeDegress = newTag("_LOD", "Longitude Degress", tagOptionNone)

	// Unofficial. Longitude minutes. This has been seen exported from
	// MacFamilyFree.
	UnofficialTagLongitudeMinutes = newTag("_LOM", "Longitude Minutes", tagOptionNone)

	// Unofficial. Longitude north? This has been seen exported from
	// MacFamilyFree.
	UnofficialTagLongitudeNorth = newTag("_LON", "Longitude North", tagOptionNone)

	// Unofficial. Longitude seconds. This has been seen exported from
	// MacFamilyFree.
	UnofficialTagLongitudeSeconds = newTag("_LOS", "Longitude Seconds", tagOptionNone)

	// Unofficial. Used to group the _LA* and _LO* tags for latitude and
	// longitude. This has been seen exported from MacFamilyFree.
	UnofficialTagCoordinates = newTag("_COR", "Coordinates", tagOptionNone)

	// Unofficial. The created date and/or time. This has been seen exported
	// from Ancestry.com.
	UnofficialTagCreated = newTag("_CRE", "Created", tagOptionNone)
)

// TagFromString returns known tag constant like TagHeader from it's raw string
// representation, like "HEAD". It will also return unofficial tags as well.
//
// If the tag is not found a Tag will still be returned (and is find to use) but
// it will not include extra information like if it is a type of event
// (IsEvent).
//
// You can check the Valid property to test if the returned Tag is known.
func TagFromString(tag string) Tag {
	if foundTag, ok := knownTags[tag]; ok {
		return foundTag
	}

	return Tag{
		tag:  tag,
		name: tag,
	}
}

// Tags returns all of the known GEDCOM tags. This includes official and
// unofficial tags.
//
// If a tag exists as a constant but is not registered here then it will not
// covered by tests and the conversion from a string to a tag with TagFromString
// will not work correctly.
func Tags() []Tag {
	return []Tag{
		// Official
		TagAbbreviation, TagAddress, TagAddress1, TagAddress2, TagAdoption,
		TagAncestralFileNumber, TagAge, TagAgency, TagAlias, TagAncestors,
		TagAncestorsInterest, TagAnnulment, TagAssociates, TagAuthor,
		TagLDSBaptism, TagBaptism, TagBarMitzvah, TagBasMitzvah, TagBirth,
		TagBlessing, TagBinaryObject, TagBurial, TagCallNumber, TagCaste,
		TagCause, TagCensus, TagChange, TagCharacterSet, TagChild,
		TagChristening, TagAdultChristening, TagCity, TagConcatenation,
		TagConfirmation, TagLDSConfirmation, TagContinued, TagCopyright,
		TagCorporate, TagCremation, TagCountry, TagData, TagDate, TagDeath,
		TagDescendants, TagDescendantsInterest, TagDestination, TagDivorce,
		TagDivorceFiled, TagPhysicalDescription, TagEducation, TagEmail,
		TagEmigration, TagEndowment, TagEngagement, TagEvent, TagFact,
		TagFamily, TagFamilyChild, TagFamilyFile, TagFamilySpouse, TagFax,
		TagFirstCommunion, TagFile, TagPhonetic, TagFormat,
		TagGedcomInformation, TagGivenName, TagGraduation, TagHeader,
		TagHusband, TagIdentityNumber, TagImmigration, TagIndividual,
		TagLanguage, TagLatitude, TagLegatee, TagLongitude, TagMap,
		TagMarriageBann, TagMarriageContract, TagMarriageLicence, TagMarriage,
		TagMarriageSettlement, TagMedia, TagName, TagNationality,
		TagNaturalization, TagChildrenCount, TagNickname, TagMarriageCount,
		TagNote, TagNamePrefix, TagNameSuffix, TagObject, TagOccupation,
		TagOrdinance, TagOrdination, TagPage, TagPedigree, TagPhone, TagPlace,
		TagPostalCode, TagProbate, TagProperty, TagPublication,
		TagQualityOfData, TagReference, TagRelationship, TagReligion,
		TagRepository, TagResidence, TagRestriction, TagRetirement,
		TagRecordFileNumber, TagRecordIDNumber, TagRole, TagRomanized, TagSex,
		TagSealingChild, TagSealingSpouse, TagSource, TagSurnamePrefix,
		TagSocialSecurityNumber, TagState, TagStatus, TagSubmitter,
		TagSubmission, TagSurname, TagTemple, TagText, TagTime, TagTitle,
		TagTrailer, TagType, TagVersion, TagWife, TagWWW, TagWill,

		// Unofficial
		UnofficialTagFamilySearchID, UnofficialTagLatitudeDegrees,
		UnofficialTagLatitudeMinutes, UnofficialTagLatitudeSeconds,
		UnofficialTagLongitudeDegress, UnofficialTagLongitudeMinutes,
		UnofficialTagLongitudeNorth, UnofficialTagLongitudeSeconds,
		UnofficialTagCoordinates, UnofficialTagCreated,
	}
}

// String returned the descriptive name of the tag, like "Bar Mitzvah".
func (tag Tag) String() string {
	return tag.name
}

// IsOfficial returns true if the tag is part of the GEDCOM 5.5 standard.
func (tag Tag) IsOfficial() bool {
	if len(tag.tag) == 0 {
		return false
	}

	return tag.tag[0] != '_'
}

// IsEvent return true if the tag is a type of event. Events can be attached to
// individuals or families and generally include a date and/or place in the
// child nodes.
//
// It is important to note that tags thar are not seen as events can still have
// dates and place attached to them.
func (tag Tag) IsEvent() bool {
	return tag.options&tagOptionEvent != 0
}

// Tag returns the raw GEDCOM name for the tag, like "MARR".
func (tag Tag) Tag() string {
	return tag.tag
}

// IsKnown returns true if the tag can provide extra information like IsEvent,
// IsOfficial, etc. If you are using the Tag or UnofficialTag variables provided
// by this will always return true.
func (tag Tag) IsKnown() bool {
	return tag.isKnown
}

// newTag is used to create and register a tag. It is for internal use only and
// should only be used for tags that are known, otherwise initialise the tag
// manually with Tag{}.
func newTag(tag, name string, options int) Tag {
	if knownTags == nil {
		knownTags = map[string]Tag{}
	}

	// This will only happen once when the package is initialised. It makes sure
	// we don't introduce any tags that already exist.
	if _, ok := knownTags[tag]; ok {
		panic(fmt.Sprintf("tag already exists: %s", tag))
	}

	knownTags[tag] = Tag{
		isKnown: true,
		tag:     tag,
		name:    name,
		options: options,
	}

	return knownTags[tag]
}

// Is test if two tags are the same. Tags are deemed to be the same solely on
// their Tag attribute (the GEDCOM representation). All other properties will be
// ignored. This makes it safe to compare real/registered tags with ones
// manually created or provided from TagFromString.
func (tag Tag) Is(tag2 Tag) bool {
	return tag.tag == tag2.tag
}
