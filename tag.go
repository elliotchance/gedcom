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

const (
	// Skip the 0 value as SortValue will transform this to a maximum value.
	_ = iota

	// The individuals name.
	tagSortIndividualName

	// Attributes for the individual attributes, such as sex.
	tagSortIndividualInfo

	// The birth event.
	tagSortIndividualBirth

	// Most tags will exist here. They will be sorted by date.
	tagSortIndividualEvents

	// The death event.
	tagSortIndividualDeath

	// The burial event.
	tagSortIndividualBurial

	// Unofficial tags.
	//
	// SortValue depends on this being the largest value. If there is any items
	// after this then SortValue must also be adjusted.
	tagSortIndividualUnofficial
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

	// sortValue heavily influences the grouped sorting of sections by functions
	// like NodeDiff.Sort.
	sortValue int
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
	TagAbbreviation = newTag("ABBR", "Abbreviation", tagOptionNone, tagSortIndividualInfo)

	// The contemporary place, usually required for postal purposes, of an
	// individual, a submitter of information, a repository, a business, a
	// school, or a company.
	TagAddress = newTag("ADDR", "Address", tagOptionNone, tagSortIndividualInfo)

	// The first line of an address.
	TagAddress1 = newTag("ADR1", "Address Line 1", tagOptionNone, tagSortIndividualInfo)

	// The second line of an address.
	TagAddress2 = newTag("ADR2", "Address Line 2", tagOptionNone, tagSortIndividualInfo)

	// Pertaining to creation of a child-parent relationship that does not exist
	// biologically.
	TagAdoption = newTag("ADOP", "Adoption", tagOptionEvent, tagSortIndividualEvents)

	// Ancestral File Number, a unique permanent record file number of an
	// individual record stored in Ancestral File.
	TagAncestralFileNumber = newTag("AFN", "Ancestral File Number", tagOptionNone, tagSortIndividualInfo)

	// The age of the individual at the time an event occurred, or the age
	// listed in the document.
	TagAge = newTag("AGE", "Age", tagOptionNone, tagSortIndividualInfo)

	// The institution or individual having authority and/or responsibility to
	// manage or govern.
	TagAgency = newTag("AGNC", "Agency", tagOptionNone, tagSortIndividualInfo)

	// An indicator to link different record descriptions of a person who may be
	// the same person.
	TagAlias = newTag("ALIA", "Alias", tagOptionNone, tagSortIndividualInfo)

	// Pertaining to forbearers of an individual.
	TagAncestors = newTag("ANCE", "Ancestors", tagOptionNone, tagSortIndividualInfo)

	// Indicates an interest in additional research for ancestors of this
	// individual. (See also DESI)
	TagAncestorsInterest = newTag("ANCI", "Ancestors Interest", tagOptionNone, tagSortIndividualInfo)

	// Declaring a marriage void from the beginning (never existed).
	TagAnnulment = newTag("ANUL", "Annulment", tagOptionEvent, tagSortIndividualEvents)

	// An indicator to link friends, neighbors, relatives, or associates of an
	// individual.
	TagAssociates = newTag("ASSO", "Associates", tagOptionNone, tagSortIndividualInfo)

	// The name of the individual who created or compiled information.
	TagAuthor = newTag("AUTH", "Author", tagOptionNone, tagSortIndividualInfo)

	// The event of baptism performed at age eight or later by priesthood
	// authority of the LDS Church. (See also BAPM)
	TagLDSBaptism = newTag("BAPL", "LDS Baptism", tagOptionEvent, tagSortIndividualEvents)

	// The event of baptism (not LDS), performed in infancy or later. (See also
	// BAPL and CHR)
	TagBaptism = newTag("BAPM", "Baptism", tagOptionEvent, tagSortIndividualEvents)

	// The ceremonial event held when a Jewish boy reaches age 13.
	TagBarMitzvah = newTag("BARM", "Bar Mitzvah", tagOptionEvent, tagSortIndividualEvents)

	// The ceremonial event held when a Jewish girl reaches age 13, also known
	// as "Bat Mitzvah."
	TagBasMitzvah = newTag("BASM", "Bas Mitzvah", tagOptionEvent, tagSortIndividualEvents)

	// See BirthNode.
	TagBirth = newTag("BIRT", "Birth", tagOptionEvent, tagSortIndividualBirth)

	// A religious event of bestowing divine care or intercession. Sometimes
	// given in connection with a naming ceremony.
	TagBlessing = newTag("BLES", "Blessing", tagOptionEvent, tagSortIndividualEvents)

	// A grouping of data used as input to a multimedia system that processes
	// binary data to represent images, sound, and video. Deleted in Gedcom
	// 5.5.1
	TagBinaryObject = newTag("BLOB", "Binary Object", tagOptionNone, tagSortIndividualInfo)

	// The event of the proper disposing of the mortal remains of a deceased
	// person.
	TagBurial = newTag("BURI", "Burial", tagOptionEvent, tagSortIndividualBurial)

	// The number used by a repository to identify the specific items in its
	// collections.
	TagCallNumber = newTag("CALN", "Call Number", tagOptionNone, tagSortIndividualInfo)

	// The name of an individual's rank or status in society, based on racial or
	// religious differences, or differences in wealth, inherited rank,
	// profession, occupation, etc.
	TagCaste = newTag("CAST", "Caste", tagOptionNone, tagSortIndividualInfo)

	// A description of the cause of the associated event or fact, such as the
	// cause of death.
	TagCause = newTag("CAUS", "Cause", tagOptionNone, tagSortIndividualInfo)

	// The event of the periodic count of the population for a designated
	// locality, such as a national or state Census.
	TagCensus = newTag("CENS", "Census", tagOptionEvent, tagSortIndividualEvents)

	// Indicates a change, correction, or modification. Typically used in
	// connection with a DATE to specify when a change in information occurred.
	TagChange = newTag("CHAN", "Change", tagOptionNone, tagSortIndividualInfo)

	// An indicator of the character set used in writing this automated
	// information.
	TagCharacterSet = newTag("CHAR", "Character Set", tagOptionNone, tagSortIndividualInfo)

	// The natural, adopted, or sealed (LDS) child of a father and a mother.
	TagChild = newTag("CHIL", "Child", tagOptionNone, tagSortIndividualInfo)

	// The religious event (not LDS) of baptizing and/or naming a child.
	TagChristening = newTag("CHR", "Christening", tagOptionEvent, tagSortIndividualEvents)

	// The religious event (not LDS) of baptizing and/or naming an adult person.
	TagAdultChristening = newTag("CHRA", "Adult Christening", tagOptionEvent, tagSortIndividualEvents)

	// A lower level jurisdictional unit. Normally an incorporated municipal
	// unit.
	TagCity = newTag("CITY", "City", tagOptionNone, tagSortIndividualInfo)

	// An indicator that additional data belongs to the superior value. The
	// information from the CONC value is to be connected to the value of the
	// superior preceding line without a space and without a carriage return
	// and/or new line character. Values that are split for a CONC tag must
	// always be split at a non-space. If the value is split on a space the
	// space will be lost when concatenation takes place. This is because of the
	// treatment that spaces get as a GEDCOM delimiter, many GEDCOM values are
	// trimmed of trailing spaces and some systems look for the first non-space
	// starting after the tag to determine the beginning of the value.
	TagConcatenation = newTag("CONC", "Concatenation", tagOptionNone, tagSortIndividualInfo)

	// The religious event (not LDS) of conferring the gift of the Holy Ghost
	// and, among protestants, full church membership.
	TagConfirmation = newTag("CONF", "Confirmation", tagOptionEvent, tagSortIndividualEvents)

	// The religious event by which a person receives membership in the LDS
	// Church.
	TagLDSConfirmation = newTag("CONL", "LDS Confirmation", tagOptionEvent, tagSortIndividualEvents)

	//  An indicator that additional data belongs to the superior value. The
	// information from the CONT value is to be connected to the value of the
	// superior preceding line with a carriage return and/or new line character.
	// Leading spaces could be important to the formatting of the resultant
	// text. When importing values from CONT lines the reader should assume only
	// one delimiter character following the CONT tag. Assume that the rest of
	// the leading spaces are to be a part of the value.
	TagContinued = newTag("CONT", "Continued", tagOptionNone, tagSortIndividualInfo)

	// A statement that accompanies data to protect it from unlawful duplication
	// and distribution.
	TagCopyright = newTag("COPR", "Copyright", tagOptionNone, tagSortIndividualInfo)

	// A name of an institution, agency, corporation, or company.
	TagCorporate = newTag("CORP", "Corporate", tagOptionNone, tagSortIndividualInfo)

	// Disposal of the remains of a person's body by fire.
	TagCremation = newTag("CREM", "Cremation", tagOptionEvent, tagSortIndividualEvents)

	// The name or code of the country.
	TagCountry = newTag("CTRY", "Country", tagOptionNone, tagSortIndividualInfo)

	// Pertaining to stored automated information.
	TagData = newTag("DATA", "Data", tagOptionNone, tagSortIndividualInfo)

	// The time of an event in a calendar format.
	TagDate = newTag("DATE", "Date", tagOptionNone, tagSortIndividualInfo)

	// The event when mortal life terminates.
	TagDeath = newTag("DEAT", "Death", tagOptionEvent, tagSortIndividualDeath)

	// Pertaining to offspring of an individual.
	TagDescendants = newTag("DESC", "Descendants", tagOptionNone, tagSortIndividualInfo)

	// Indicates an interest in research to identify additional descendants of
	// this individual. (See also ANCI)
	TagDescendantsInterest = newTag("DESI", "Descendants Interest", tagOptionNone, tagSortIndividualInfo)

	// A system receiving data.
	TagDestination = newTag("DEST", "Destination", tagOptionNone, tagSortIndividualInfo)

	// An event of dissolving a marriage through civil action.
	TagDivorce = newTag("DIV", "Divorce", tagOptionEvent, tagSortIndividualEvents)

	// An event of filing for a divorce by a spouse.
	TagDivorceFiled = newTag("DIVF", "Divorce Filed", tagOptionEvent, tagSortIndividualEvents)

	// The physical characteristics of a person, place, or thing.
	TagPhysicalDescription = newTag("DSCR", "Physical Description", tagOptionNone, tagSortIndividualInfo)

	// Indicator of a level of education attained.
	TagEducation = newTag("EDUC", "Education", tagOptionNone, tagSortIndividualInfo)

	// An electronic address that can be used for contact such as an email
	// address. New in Gedcom 5.5.1.
	TagEmail = newTag("EMAIL", "Email", tagOptionNone, tagSortIndividualInfo)

	// An event of leaving one's homeland with the intent of residing elsewhere.
	TagEmigration = newTag("EMIG", "Emigration", tagOptionEvent, tagSortIndividualEvents)

	// A religious event where an endowment ordinance for an individual was
	// performed by priesthood authority in an LDS temple.
	TagEndowment = newTag("ENDL", "Endowment", tagOptionEvent, tagSortIndividualEvents)

	// An event of recording or announcing an agreement between two people to
	// become married.
	TagEngagement = newTag("ENGA", "Engagement", tagOptionEvent, tagSortIndividualEvents)

	// See EventNode.
	TagEvent = newTag("EVEN", "Event", tagOptionEvent, tagSortIndividualEvents)

	// Pertaining to a noteworthy attribute or fact concerning an individual, a
	// group, or an organization. A structure is usually qualified or classified
	// by a subordinate use of the TYPE tag. New in Gedcom 5.5.1.
	TagFact = newTag("FACT", "Fact", tagOptionNone, tagSortIndividualInfo)

	// Identifies a legal, common law, or other customary relationship of man
	// and woman and their children, if any, or a family created by virtue of
	// the birth of a child to its biological father and mother.
	TagFamily = newTag("FAM", "Family", tagOptionNone, tagSortIndividualInfo)

	// Identifies the family in which an individual appears as a child.
	TagFamilyChild = newTag("FAMC", "Family Child", tagOptionNone, tagSortIndividualInfo)

	// Pertaining to, or the name of, a family file. Names stored in a file that
	// are assigned to a family for doing temple ordinance work.
	TagFamilyFile = newTag("FAMF", "Family File", tagOptionNone, tagSortIndividualInfo)

	// Identifies the family in which an individual appears as a spouse.
	TagFamilySpouse = newTag("FAMS", "Family Spouse", tagOptionNone, tagSortIndividualInfo)

	// A FAX telephone number appropriate for sending data facsimiles. New in
	// Gedcom 5.5.1.
	TagFax = newTag("FAX", "Fax", tagOptionNone, tagSortIndividualInfo)

	// A religious rite, the first act of sharing in the Lord's supper as part
	// of church worship.
	TagFirstCommunion = newTag("FCOM", "First Communion", tagOptionEvent, tagSortIndividualEvents)

	// An information storage place that is ordered and arranged for
	// preservation and reference.
	TagFile = newTag("FILE", "File", tagOptionNone, tagSortIndividualInfo)

	// See PhoneticNode.
	TagPhonetic = newTag("FONE", "Phonetic", tagOptionNone, tagSortIndividualInfo)

	// See FormatNode.
	TagFormat = newTag("FORM", "Format", tagOptionNone, tagSortIndividualInfo)

	// Information about the use of GEDCOM in a transmission.
	TagGedcomInformation = newTag("GEDC", "GEDCOM Information", tagOptionNone, tagSortIndividualInfo)

	// A given or earned name used for official identification of a person. It
	// is also commonly known as the "first name".
	//
	// The NameNode provides a GivenName() function.
	TagGivenName = newTag("GIVN", "Given Name", tagOptionNone, tagSortIndividualInfo)

	// An event of awarding educational diplomas or degrees to individuals.
	TagGraduation = newTag("GRAD", "Graduation", tagOptionEvent, tagSortIndividualEvents)

	// Identifies information pertaining to an entire GEDCOM transmission.
	TagHeader = newTag("HEAD", "Header", tagOptionNone, tagSortIndividualInfo)

	// An individual in the family role of a married man or father.
	TagHusband = newTag("HUSB", "Husband", tagOptionNone, tagSortIndividualInfo)

	// A number assigned to identify a person within some significant external
	// system.
	TagIdentityNumber = newTag("IDNO", "Identity Number", tagOptionNone, tagSortIndividualInfo)

	// An event of entering into a new locality with the intent of residing
	// there.
	TagImmigration = newTag("IMMI", "Immigration", tagOptionEvent, tagSortIndividualEvents)

	// A person.
	TagIndividual = newTag("INDI", "Individual", tagOptionNone, tagSortIndividualInfo)

	// Defines label for given fact.
	TagLabel = newTag("LABL", "Label", tagOptionNone, tagSortIndividualInfo)

	// The name of the language used in a communication or transmission of
	// information.
	TagLanguage = newTag("LANG", "Language", tagOptionNone, tagSortIndividualInfo)

	// See LatitudeNode.
	TagLatitude = newTag("LATI", "Latitude", tagOptionNone, tagSortIndividualInfo)

	// A role of an individual acting as a person receiving a bequest or legal
	// devise.
	TagLegatee = newTag("LEGA", "Legatee", tagOptionNone, tagSortIndividualInfo)

	// See LongitudeNode.
	TagLongitude = newTag("LONG", "Longitude", tagOptionNone, tagSortIndividualInfo)

	// See MapNode.
	TagMap = newTag("MAP", "Map", tagOptionNone, tagSortIndividualInfo)

	// An event of an official public notice given that two people intend to
	// marry.
	TagMarriageBann = newTag("MARB", "Marriage Bann", tagOptionEvent, tagSortIndividualEvents)

	// An event of recording a formal agreement of marriage, including the
	// prenuptial agreement in which marriage partners reach agreement about the
	// property rights of one or both, securing property to their children.
	TagMarriageContract = newTag("MARC", "Marriage Contract", tagOptionEvent, tagSortIndividualEvents)

	// An event of obtaining a legal license to marry.
	TagMarriageLicence = newTag("MARL", "Marriage Licence", tagOptionEvent, tagSortIndividualEvents)

	// A legal, common-law, or customary event of creating a family unit of a
	// man and a woman as husband and wife.
	TagMarriage = newTag("MARR", "Marriage", tagOptionEvent, tagSortIndividualEvents)

	// An event of creating an agreement between two people contemplating
	// marriage, at which time they agree to release or modify property rights
	// that would otherwise arise from the marriage.
	TagMarriageSettlement = newTag("MARS", "Marriage Settlement", tagOptionEvent, tagSortIndividualEvents)

	// Identifies information about the media or having to do with the medium in
	// which information is stored.
	TagMedia = newTag("MEDI", "Media", tagOptionNone, tagSortIndividualInfo)

	// A word or combination of words used to help identify an individual,
	// title, or other item. More than one NAME line should be used for people
	// who were known by multiple names.
	//
	// NAME tags will be interpreted with the NameNode type.
	TagName = newTag("NAME", "Name", tagOptionNone, tagSortIndividualName)

	// The national heritage of an individual.
	TagNationality = newTag("NATI", "Nationality", tagOptionNone, tagSortIndividualInfo)

	// The event of obtaining citizenship.
	TagNaturalization = newTag("NATU", "Naturalization", tagOptionEvent, tagSortIndividualEvents)

	// The number of children that this person is known to be the parent of (all
	// marriages) when subordinate to an individual, or that belong to this
	// family when subordinate to a FAM_RECORD.
	TagChildrenCount = newTag("NCHI", "Children Count", tagOptionNone, tagSortIndividualInfo)

	// A descriptive or familiar that is used instead of, or in addition to,
	// one's proper name.
	TagNickname = newTag("NICK", "Nickname", tagOptionNone, tagSortIndividualInfo)

	// The number of times this person has participated in a family as a spouse
	// or parent.
	TagMarriageCount = newTag("NMR", "Marriage Count", tagOptionNone, tagSortIndividualInfo)

	// See NoteNode.
	TagNote = newTag("NOTE", "Note", tagOptionNone, tagSortIndividualInfo)

	// Text which appears on a name line before the given and surname parts of a
	// name. i.e. ( Lt. Cmndr. ) Joseph /Allen/ jr. In this example Lt. Cmndr.
	// is considered as the name prefix portion.
	//
	// The NameNode provides a Prefix() function.
	TagNamePrefix = newTag("NPFX", "Name Prefix", tagOptionNone, tagSortIndividualInfo)

	// Text which appears on a name line after or behind the given and surname
	// parts of a name. i.e. Lt. Cmndr. Joseph /Allen/ ( jr. ) In this example
	// jr. is considered as the name suffix portion.
	//
	// The NameNode provides a Suffix() function.
	TagNameSuffix = newTag("NSFX", "Name Suffix", tagOptionNone, tagSortIndividualInfo)

	// Pertaining to a grouping of attributes used in describing something.
	// Usually referring to the data required to represent a multimedia object,
	// such an audio recording, a photograph of a person, or an image of a
	// document.
	TagObject = newTag("OBJE", "Object", tagOptionNone, tagSortIndividualInfo)

	// The type of work or profession of an individual.
	TagOccupation = newTag("OCCU", "Occupation", tagOptionNone, tagSortIndividualInfo)

	// Pertaining to a religious ordinance in general.
	TagOrdinance = newTag("ORDI", "Ordinance", tagOptionNone, tagSortIndividualInfo)

	// A religious event of receiving authority to act in religious matters.
	TagOrdination = newTag("ORDN", "Ordination", tagOptionEvent, tagSortIndividualEvents)

	// A number or description to identify where information can be found in a
	// referenced work.
	TagPage = newTag("PAGE", "Page", tagOptionNone, tagSortIndividualInfo)

	// Information pertaining to an individual to parent lineage chart.
	TagPedigree = newTag("PEDI", "Pedigree", tagOptionNone, tagSortIndividualInfo)

	// A unique number assigned to access a specific telephone.
	TagPhone = newTag("PHON", "Phone", tagOptionNone, tagSortIndividualInfo)

	// See PlaceNode.
	TagPlace = newTag("PLAC", "Place", tagOptionNone, tagSortIndividualInfo)

	// A code used by a postal service to identify an area to facilitate mail
	// handling.
	TagPostalCode = newTag("POST", "Postal Code", tagOptionNone, tagSortIndividualInfo)

	// An event of judicial determination of the validity of a will. May
	// indicate several related court activities over several dates.
	TagProbate = newTag("PROB", "Probate", tagOptionEvent, tagSortIndividualEvents)

	// Pertaining to possessions such as real estate or other property of
	// interest.
	TagProperty = newTag("PROP", "Property", tagOptionNone, tagSortIndividualInfo)

	// Refers to when and/or were a work was published or created.
	TagPublication = newTag("PUBL", "Publication", tagOptionNone, tagSortIndividualInfo)

	// An assessment of the certainty of the evidence to support the conclusion
	// drawn from evidence.
	TagQualityOfData = newTag("QUAY", "Quality Of Data", tagOptionNone, tagSortIndividualInfo)

	// A description or number used to identify an item for filing, storage, or
	// other reference purposes.
	TagReference = newTag("REFN", "Reference", tagOptionNone, tagSortIndividualInfo)

	// A relationship value between the indicated contexts.
	TagRelationship = newTag("RELA", "Relationship", tagOptionNone, tagSortIndividualInfo)

	// A religious denomination to which a person is affiliated or for which a
	// record applies.
	TagReligion = newTag("RELI", "Religion", tagOptionNone, tagSortIndividualInfo)

	// An institution or person that has the specified item as part of their
	// collection(s).
	TagRepository = newTag("REPO", "Repository", tagOptionNone, tagSortIndividualInfo)

	// See ResidenceNode.
	TagResidence = newTag("RESI", "Residence", tagOptionEvent, tagSortIndividualEvents)

	// A processing indicator signifying access to information has been denied
	// or otherwise restricted.
	TagRestriction = newTag("RESN", "Restriction", tagOptionNone, tagSortIndividualInfo)

	// An event of exiting an occupational relationship with an employer after a
	// qualifying time period.
	TagRetirement = newTag("RETI", "Retirement", tagOptionEvent, tagSortIndividualEvents)

	// A permanent number assigned to a record that uniquely identifies it
	// within a known file.
	TagRecordFileNumber = newTag("RFN", "Record File Number", tagOptionNone, tagSortIndividualInfo)

	// A number assigned to a record by an originating automated system that can
	// be used by a receiving system to report results pertaining to that
	// record.
	TagRecordIDNumber = newTag("RIN", "Record ID Number", tagOptionNone, tagSortIndividualInfo)

	// A name given to a role played by an individual in connection with an
	// event.
	TagRole = newTag("ROLE", "Role", tagOptionNone, tagSortIndividualInfo)

	// A romanized variation of a superior text string. New in Gedcom 5.5.1.
	TagRomanized = newTag("ROMN", "Romanized", tagOptionNone, tagSortIndividualInfo)

	// Indicates the sex of an individual--male or female.
	TagSex = newTag("SEX", "Sex", tagOptionNone, tagSortIndividualInfo)

	// A religious event pertaining to the sealing of a child to his or her
	// parents in an LDS temple ceremony.
	TagSealingChild = newTag("SLGC", "Sealing Child", tagOptionEvent, tagSortIndividualEvents)

	// A religious event pertaining to the sealing of a husband and wife in an
	// LDS temple ceremony.
	TagSealingSpouse = newTag("SLGS", "Sealing Spouse", tagOptionEvent, tagSortIndividualEvents)

	// The initial or original material from which information was obtained.
	TagSource = newTag("SOUR", "Source", tagOptionNone, tagSortIndividualInfo)

	// A name piece used as a non-indexing pre-part of a surname.
	TagSurnamePrefix = newTag("SPFX", "Surname Prefix", tagOptionNone, tagSortIndividualInfo)

	// A number assigned by the United States Social Security Administration.
	// Used for tax identification purposes.
	TagSocialSecurityNumber = newTag("SSN", "Social Security Number", tagOptionNone, tagSortIndividualInfo)

	// A geographical division of a larger jurisdictional area, such as a State
	// within the United States of America.
	TagState = newTag("STAE", "State", tagOptionNone, tagSortIndividualInfo)

	// An assessment of the state or condition of something.
	TagStatus = newTag("STAT", "Status", tagOptionNone, tagSortIndividualInfo)

	// An individual or organization who contributes genealogical data to a file
	// or transfers it to someone else.
	TagSubmitter = newTag("SUBM", "Submitter", tagOptionNone, tagSortIndividualInfo)

	// Pertains to a collection of data issued for processing.
	TagSubmission = newTag("SUBN", "Submission", tagOptionNone, tagSortIndividualInfo)

	// A family name passed on or used by members of a family.
	//
	// The NameNode provides a Surname() function.
	TagSurname = newTag("SURN", "Surname", tagOptionNone, tagSortIndividualInfo)

	// The name or code that represents the name a temple of the LDS Church.
	TagTemple = newTag("TEMP", "Temple", tagOptionNone, tagSortIndividualInfo)

	// The exact wording found in an original source document.
	TagText = newTag("TEXT", "Text", tagOptionNone, tagSortIndividualInfo)

	// A time value in a 24-hour clock format, including hours, minutes, and
	// optional seconds, separated by a colon (:). Fractions of seconds are
	// shown in decimal notation.
	TagTime = newTag("TIME", "Time", tagOptionNone, tagSortIndividualInfo)

	// A description of a specific writing or other work, such as the title of a
	// book when used in a source context, or a formal designation used by an
	// individual in connection with positions of royalty or other social
	// status, such as Grand Duke.
	//
	// The NameNode provides a Title() function.
	TagTitle = newTag("TITL", "Title", tagOptionNone, tagSortIndividualInfo)

	// At level 0, specifies the end of a GEDCOM transmission.
	TagTrailer = newTag("TRLR", "Trailer", tagOptionNone, tagSortIndividualInfo)

	// See TypeNode.
	TagType = newTag("TYPE", "Type", tagOptionNone, tagSortIndividualInfo)

	// Indicates which version of a product, item, or publication is being used
	// or referenced.
	TagVersion = newTag("VERS", "Version", tagOptionNone, tagSortIndividualInfo)

	// An individual in the role as a mother and/or married woman.
	TagWife = newTag("WIFE", "Wife", tagOptionNone, tagSortIndividualInfo)

	// World Wide Web home page. New in Gedcom 5.5.1.
	TagWWW = newTag("WWW", "WWW", tagOptionNone, tagSortIndividualInfo)

	// A legal document treated as an event, by which a person disposes of his
	// or her estate, to take effect after death. The event date is the date the
	// will was signed while the person was alive. (See also PROBate)
	TagWill = newTag("WILL", "Will", tagOptionEvent, tagSortIndividualEvents)
)

var (
	// Unofficial. The unique identifier for the person on FamilySearch.org.
	// This has been seen exported from MacFamilyFree.
	UnofficialTagFamilySearchID = newTag("_FID", "FamilySearch ID", tagOptionNone, tagSortIndividualUnofficial)

	// Unofficial. Latitude degrees. This has been seen exported from
	// MacFamilyFree.
	UnofficialTagLatitudeDegrees = newTag("_LAD", "Latitude Degrees", tagOptionNone, tagSortIndividualUnofficial)

	// Unofficial. Latitude minutes. This has been seen exported from
	// MacFamilyFree.
	UnofficialTagLatitudeMinutes = newTag("_LAM", "Latitude Minutes", tagOptionNone, tagSortIndividualUnofficial)

	// Unofficial. Latitude seconds. This has been seen exported from
	// MacFamilyFree.
	UnofficialTagLatitudeSeconds = newTag("_LAS", "Latitude Seconds", tagOptionNone, tagSortIndividualUnofficial)

	// Unofficial. Longitude degrees. This has been seen exported from
	// MacFamilyFree.
	UnofficialTagLongitudeDegress = newTag("_LOD", "Longitude Degress", tagOptionNone, tagSortIndividualUnofficial)

	// Unofficial. Longitude minutes. This has been seen exported from
	// MacFamilyFree.
	UnofficialTagLongitudeMinutes = newTag("_LOM", "Longitude Minutes", tagOptionNone, tagSortIndividualUnofficial)

	// Unofficial. Longitude north? This has been seen exported from
	// MacFamilyFree.
	UnofficialTagLongitudeNorth = newTag("_LON", "Longitude North", tagOptionNone, tagSortIndividualUnofficial)

	// Unofficial. Longitude seconds. This has been seen exported from
	// MacFamilyFree.
	UnofficialTagLongitudeSeconds = newTag("_LOS", "Longitude Seconds", tagOptionNone, tagSortIndividualUnofficial)

	// Unofficial. Used to group the _LA* and _LO* tags for latitude and
	// longitude. This has been seen exported from MacFamilyFree.
	UnofficialTagCoordinates = newTag("_COR", "Coordinates", tagOptionNone, tagSortIndividualUnofficial)

	// Unofficial. The created date and/or time. This has been seen exported
	// from Ancestry.com.
	UnofficialTagCreated = newTag("_CRE", "Created", tagOptionNone, tagSortIndividualUnofficial)
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
		TagTrailer, TagType, TagVersion, TagWife, TagWWW, TagWill, TagLabel,

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
func newTag(tag, name string, options int, sortValue int) Tag {
	if knownTags == nil {
		knownTags = map[string]Tag{}
	}

	// This will only happen once when the package is initialised. It makes sure
	// we don't introduce any tags that already exist.
	if _, ok := knownTags[tag]; ok {
		panic(fmt.Sprintf("tag already exists: %s", tag))
	}

	knownTags[tag] = Tag{
		isKnown:   true,
		tag:       tag,
		name:      name,
		options:   options,
		sortValue: sortValue,
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

// SortValue is used when sorting tags.
//
// The lowest SortValue should appear first.
func (tag Tag) SortValue() int {
	if tag.sortValue != 0 {
		return tag.sortValue
	}

	return tagSortIndividualUnofficial + 1
}
