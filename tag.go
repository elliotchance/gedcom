package gedcom

// Tag is the type of node.
type Tag string

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
// GEDCOM-standard permits userdefined field names (one of the major causes of
// "misunderstandings" of different GECOM-compliant programs!). They have to
// begin with an underscore _ .
const (
	// A short name of a title, description, or name.
	Abbreviation = Tag("ABBR")

	// The contemporary place, usually required for postal purposes, of an
	// individual, a submitter of information, a repository, a business, a
	// school, or a company.
	Address = Tag("ADDR")

	// The first line of an address.
	Address1 = Tag("ADR1")

	// The second line of an address.
	Address2 = Tag("ADR2")

	// Pertaining to creation of a child-parent relationship that does not exist
	// biologically.
	Adoption = Tag("ADOP")

	// Ancestral File Number, a unique permanent record file number of an
	// individual record stored in Ancestral File.
	AncestralFileNumber = Tag("AFN")

	// The age of the individual at the time an event occurred, or the age
	// listed in the document.
	Age = Tag("AGE")

	// The institution or individual having authority and/or responsibility to
	// manage or govern.
	Agency = Tag("AGNC")

	// An indicator to link different record descriptions of a person who may be
	// the same person.
	Alias = Tag("ALIA")

	// Pertaining to forbearers of an individual.
	Ancestors = Tag("ANCE")

	// Indicates an interest in additional research for ancestors of this
	// individual. (See also DESI)
	AncestorsInterest = Tag("ANCI")

	// Declaring a marriage void from the beginning (never existed).
	Annulment = Tag("ANUL")

	// An indicator to link friends, neighbors, relatives, or associates of an
	// individual.
	Associates = Tag("ASSO")

	// The name of the individual who created or compiled information.
	Author = Tag("AUTH")

	// The event of baptism performed at age eight or later by priesthood
	// authority of the LDS Church. (See also BAPM)
	LDSBaptism = Tag("BAPL")

	// The event of baptism (not LDS), performed in infancy or later. (See also
	// BAPL and CHR)
	Baptism = Tag("BAPM")

	// The ceremonial event held when a Jewish boy reaches age 13.
	BarMitzvah = Tag("BARM")

	// The ceremonial event held when a Jewish girl reaches age 13, also known
	// as "Bat Mitzvah."
	BasMitzvah = Tag("BASM")

	// The event of entering into life.
	Birth = Tag("BIRT")

	// A religious event of bestowing divine care or intercession. Sometimes
	// given in connection with a naming ceremony.
	Blessing = Tag("BLES")

	// A grouping of data used as input to a multimedia system that processes
	// binary data to represent images, sound, and video. Deleted in Gedcom
	// 5.5.1
	BinaryObject = Tag("BLOB")

	// The event of the proper disposing of the mortal remains of a deceased
	// person.
	Burial = Tag("BURI")

	// The number used by a repository to identify the specific items in its
	// collections.
	CallNumber = Tag("CALN")

	// The name of an individual's rank or status in society, based on racial or
	// religious differences, or differences in wealth, inherited rank,
	// profession, occupation, etc.
	Caste = Tag("CAST")

	// A description of the cause of the associated event or fact, such as the
	// cause of death.
	Cause = Tag("CAUS")

	// The event of the periodic count of the population for a designated
	// locality, such as a national or state Census.
	Census = Tag("CENS")

	// Indicates a change, correction, or modification. Typically used in
	// connection with a DATE to specify when a change in information occurred.
	Change = Tag("CHAN")

	// An indicator of the character set used in writing this automated
	// information.
	CharacterSet = Tag("CHAR")

	// The natural, adopted, or sealed (LDS) child of a father and a mother.
	Child = Tag("CHIL")

	// The religious event (not LDS) of baptizing and/or naming a child.
	Christening = Tag("CHR")

	// The religious event (not LDS) of baptizing and/or naming an adult person.
	AdultChristening = Tag("CHRA")

	// A lower level jurisdictional unit. Normally an incorporated municipal
	// unit.
	City = Tag("CITY")

	// An indicator that additional data belongs to the superior value. The
	// information from the CONC value is to be connected to the value of the
	// superior preceding line without a space and without a carriage return
	// and/or new line character. Values that are split for a CONC tag must
	// always be split at a non-space. If the value is split on a space the
	// space will be lost when concatenation takes place. This is because of the
	// treatment that spaces get as a GEDCOM delimiter, many GEDCOM values are
	// trimmed of trailing spaces and some systems look for the first non-space
	// starting after the tag to determine the beginning of the value.
	Concatenation = Tag("CONC")

	// The religious event (not LDS) of conferring the gift of the Holy Ghost
	// and, among protestants, full church membership.
	Confirmation = Tag("CONF")

	// The religious event by which a person receives membership in the LDS
	// Church.
	LDSConfirmation = Tag("CONL")

	//  An indicator that additional data belongs to the superior value. The
	// information from the CONT value is to be connected to the value of the
	// superior preceding line with a carriage return and/or new line character.
	// Leading spaces could be important to the formatting of the resultant
	// text. When importing values from CONT lines the reader should assume only
	// one delimiter character following the CONT tag. Assume that the rest of
	// the leading spaces are to be a part of the value.
	Continued = Tag("CONT")

	// A statement that accompanies data to protect it from unlawful duplication
	// and distribution.
	Copyright = Tag("COPR")

	// A name of an institution, agency, corporation, or company.
	Corporate = Tag("CORP")

	// Disposal of the remains of a person's body by fire.
	Cremation = Tag("CREM")

	// The name or code of the country.
	Country = Tag("CTRY")

	// Pertaining to stored automated information.
	Data = Tag("DATA")

	// The time of an event in a calendar format.
	Date = Tag("DATE")

	// The event when mortal life terminates.
	Death = Tag("DEAT")

	// Pertaining to offspring of an individual.
	Descendants = Tag("DESC")

	// Indicates an interest in research to identify additional descendants of
	// this individual. (See also ANCI)
	DescendantsInterest = Tag("DESI")

	// A system receiving data.
	Destination = Tag("DEST")

	// An event of dissolving a marriage through civil action.
	Divorce = Tag("DIV")

	// An event of filing for a divorce by a spouse.
	DivorceFiled = Tag("DIVF")

	// The physical characteristics of a person, place, or thing.
	PhysicalDescription = Tag("DSCR")

	// Indicator of a level of education attained.
	Education = Tag("EDUC")

	// An electronic address that can be used for contact such as an email
	// address. New in Gedcom 5.5.1.
	Email = Tag("EMAIL")

	// An event of leaving one's homeland with the intent of residing elsewhere.
	Emigration = Tag("EMIG")

	// A religious event where an endowment ordinance for an individual was
	// performed by priesthood authority in an LDS temple.
	Endowment = Tag("ENDL")

	// An event of recording or announcing an agreement between two people to
	// become married.
	Engagement = Tag("ENGA")

	// A noteworthy happening related to an individual, a group, or an
	// organization.
	Event = Tag("EVEN")

	// Pertaining to a noteworthy attribute or fact concerning an individual, a
	// group, or an organization. A structure is usually qualified or classified
	// by a subordinate use of the TYPE tag. New in Gedcom 5.5.1.
	Fact = Tag("FACT")

	// Identifies a legal, common law, or other customary relationship of man
	// and woman and their children, if any, or a family created by virtue of
	// the birth of a child to its biological father and mother.
	Family = Tag("FAM")

	// Identifies the family in which an individual appears as a child.
	FamilyChild = Tag("FAMC")

	// Pertaining to, or the name of, a family file. Names stored in a file that
	// are assigned to a family for doing temple ordinance work.
	FamilyFile = Tag("FAMF")

	// Identifies the family in which an individual appears as a spouse.
	FamilySpouse = Tag("FAMS")

	// A FAX telephone number appropriate for sending data facsimiles. New in
	// Gedcom 5.5.1.
	Fax = Tag("FAX")

	// A religious rite, the first act of sharing in the Lord's supper as part
	// of church worship.
	FirstCommunion = Tag("FCOM")

	// An information storage place that is ordered and arranged for
	// preservation and reference.
	File = Tag("FILE")

	// A phonetic variation of a superior text string. New in Gedcom 5.5.1
	Phonetic = Tag("FONE")

	// An assigned name given to a consistent format in which information can be
	// conveyed.
	Format = Tag("FORM")

	// Information about the use of GEDCOM in a transmission.
	GedcomInformation = Tag("GEDC")

	// A given or earned name used for official identification of a person. It
	// is also commonly known as the "first name".
	//
	// The NameNode provides a GivenName() function.
	GivenName = Tag("GIVN")

	// An event of awarding educational diplomas or degrees to individuals.
	Graduation = Tag("GRAD")

	// Identifies information pertaining to an entire GEDCOM transmission.
	Header = Tag("HEAD")

	// An individual in the family role of a married man or father.
	Husband = Tag("HUSB")

	// A number assigned to identify a person within some significant external
	// system.
	IdentityNumber = Tag("IDNO")

	// An event of entering into a new locality with the intent of residing
	// there.
	Immigration = Tag("IMMI")

	// A person.
	Individual = Tag("INDI")

	// The name of the language used in a communication or transmission of
	// information.
	Language = Tag("LANG")

	// A value indicating a coordinate position on a line, plane, or space. New
	// in Gedcom 5.5.1.
	Latitude = Tag("LATI")

	// A role of an individual acting as a person receiving a bequest or legal
	// devise.
	Legatee = Tag("LEGA")

	// A value indicating a coordinate position on a line, plane, or space. New
	// in Gedcom 5.5.1.
	Longitude = Tag("LONG")

	// Pertains to a representation of measurements usually presented in a
	// graphical form. New in Gedcom 5.5.1
	Map = Tag("MAP")

	// An event of an official public notice given that two people intend to
	// marry.
	MarriageBann = Tag("MARB")

	// An event of recording a formal agreement of marriage, including the
	// prenuptial agreement in which marriage partners reach agreement about the
	// property rights of one or both, securing property to their children.
	MarriageContract = Tag("MARC")

	// An event of obtaining a legal license to marry.
	MarriageLicence = Tag("MARL")

	// A legal, common-law, or customary event of creating a family unit of a
	// man and a woman as husband and wife.
	Marriage = Tag("MARR")

	// An event of creating an agreement between two people contemplating
	// marriage, at which time they agree to release or modify property rights
	// that would otherwise arise from the marriage.
	MarriageSettlement = Tag("MARS")

	// Identifies information about the media or having to do with the medium in
	// which information is stored.
	Media = Tag("MEDI")

	// A word or combination of words used to help identify an individual,
	// title, or other item. More than one NAME line should be used for people
	// who were known by multiple names.
	//
	// NAME tags will be interpreted with the NameNode type.
	Name = Tag("NAME")

	// The national heritage of an individual.
	Nationality = Tag("NATI")

	// The event of obtaining citizenship.
	Naturalization = Tag("NATU")

	// The number of children that this person is known to be the parent of (all
	// marriages) when subordinate to an individual, or that belong to this
	// family when subordinate to a FAM_RECORD.
	ChildrenCount = Tag("NCHI")

	// A descriptive or familiar that is used instead of, or in addition to,
	// one's proper name.
	Nickname = Tag("NICK")

	// The number of times this person has participated in a family as a spouse
	// or parent.
	MarriageCount = Tag("NMR")

	// Additional information provided by the submitter for understanding the
	// enclosing data.
	Note = Tag("NOTE")

	// Text which appears on a name line before the given and surname parts of a
	// name. i.e. ( Lt. Cmndr. ) Joseph /Allen/ jr. In this example Lt. Cmndr.
	// is considered as the name prefix portion.
	//
	// The NameNode provides a Prefix() function.
	NamePrefix = Tag("NPFX")

	// Text which appears on a name line after or behind the given and surname
	// parts of a name. i.e. Lt. Cmndr. Joseph /Allen/ ( jr. ) In this example
	// jr. is considered as the name suffix portion.
	//
	// The NameNode provides a Suffix() function.
	NameSuffix = Tag("NSFX")

	// Pertaining to a grouping of attributes used in describing something.
	// Usually referring to the data required to represent a multimedia object,
	// such an audio recording, a photograph of a person, or an image of a
	// document.
	Object = Tag("OBJE")

	// The type of work or profession of an individual.
	Occupation = Tag("OCCU")

	// Pertaining to a religious ordinance in general.
	Ordinance = Tag("ORDI")

	// A religious event of receiving authority to act in religious matters.
	Ordination = Tag("ORDN")

	// A number or description to identify where information can be found in a
	// referenced work.
	Page = Tag("PAGE")

	// Information pertaining to an individual to parent lineage chart.
	Pedigree = Tag("PEDI")

	// A unique number assigned to access a specific telephone.
	Phone = Tag("PHON")

	// A jurisdictional name to identify the place or location of an event.
	Place = Tag("PLAC")

	// A code used by a postal service to identify an area to facilitate mail
	// handling.
	PostalCode = Tag("POST")

	// An event of judicial determination of the validity of a will. May
	// indicate several related court activities over several dates.
	Probate = Tag("PROB")

	// Pertaining to possessions such as real estate or other property of
	// interest.
	Property = Tag("PROP")

	// Refers to when and/or were a work was published or created.
	Publication = Tag("PUBL")

	// An assessment of the certainty of the evidence to support the conclusion
	// drawn from evidence.
	QualityOfData = Tag("QUAY")

	// A description or number used to identify an item for filing, storage, or
	// other reference purposes.
	Reference = Tag("REFN")

	// A relationship value between the indicated contexts.
	Relationship = Tag("RELA")

	// A religious denomination to which a person is affiliated or for which a
	// record applies.
	Religion = Tag("RELI")

	// An institution or person that has the specified item as part of their
	// collection(s).
	Repository = Tag("REPO")

	// The act of dwelling at an address for a period of time.
	Residence = Tag("RESI")

	// A processing indicator signifying access to information has been denied
	// or otherwise restricted.
	Restriction = Tag("RESN")

	// An event of exiting an occupational relationship with an employer after a
	// qualifying time period.
	Retirement = Tag("RETI")

	// A permanent number assigned to a record that uniquely identifies it
	// within a known file.
	RecordFileNumber = Tag("RFN")

	// A number assigned to a record by an originating automated system that can
	// be used by a receiving system to report results pertaining to that
	// record.
	RecordIDNumber = Tag("RIN")

	// A name given to a role played by an individual in connection with an
	// event.
	Role = Tag("ROLE")

	// A romanized variation of a superior text string. New in Gedcom 5.5.1.
	Romanized = Tag("ROMN")

	// Indicates the sex of an individual--male or female.
	Sex = Tag("SEX")

	// A religious event pertaining to the sealing of a child to his or her
	// parents in an LDS temple ceremony.
	SealingChild = Tag("SLGC")

	// A religious event pertaining to the sealing of a husband and wife in an
	// LDS temple ceremony.
	SealingSpouse = Tag("SLGS")

	// The initial or original material from which information was obtained.
	Source = Tag("SOUR")

	// A name piece used as a non-indexing pre-part of a surname.
	SurnamePrefix = Tag("SPFX")

	// A number assigned by the United States Social Security Administration.
	// Used for tax identification purposes.
	SocialSecurityNumber = Tag("SSN")

	// A geographical division of a larger jurisdictional area, such as a State
	// within the United States of America.
	State = Tag("STAE")

	// An assessment of the state or condition of something.
	Status = Tag("STAT")

	// An individual or organization who contributes genealogical data to a file
	// or transfers it to someone else.
	Submitter = Tag("SUBM")

	// Pertains to a collection of data issued for processing.
	Submission = Tag("SUBN")

	// A family name passed on or used by members of a family.
	//
	// The NameNode provides a Surname() function.
	Surname = Tag("SURN")

	// The name or code that represents the name a temple of the LDS Church.
	Temple = Tag("TEMP")

	// The exact wording found in an original source document.
	Text = Tag("TEXT")

	// A time value in a 24-hour clock format, including hours, minutes, and
	// optional seconds, separated by a colon (:). Fractions of seconds are
	// shown in decimal notation.
	Time = Tag("TIME")

	// A description of a specific writing or other work, such as the title of a
	// book when used in a source context, or a formal designation used by an
	// individual in connection with positions of royalty or other social
	// status, such as Grand Duke.
	//
	// The NameNode provides a Title() function.
	Title = Tag("TITL")

	// At level 0, specifies the end of a GEDCOM transmission.
	Trailer = Tag("TRLR")

	// A further qualification to the meaning of the associated superior tag.
	// The value does not have any computer processing reliability. It is more
	// in the form of a short one or two word note that should be displayed any
	// time the associated data is displayed.
	Type = Tag("TYPE")

	// Indicates which version of a product, item, or publication is being used
	// or referenced.
	Version = Tag("VERS")

	// An individual in the role as a mother and/or married woman.
	Wife = Tag("WIFE")

	// World Wide Web home page. New in Gedcom 5.5.1.
	WWW = Tag("WWW")

	// A legal document treated as an event, by which a person disposes of his
	// or her estate, to take effect after death. The event date is the date the
	// will was signed while the person was alive. (See also PROBate)
	Will = Tag("WILL")

	// Unofficial. The unique identifier for the person on FamilySearch.org.
	// This has been seen exported from MacFamilyFree.
	FamilySearchID = Tag("_FID")

	// Unofficial. Latitude degrees. This has been seen exported from
	// MacFamilyFree.
	LatitudeDegrees = Tag("_LAD")

	// Unofficial. Latitude minutes. This has been seen exported from
	// MacFamilyFree.
	LatitudeMinutes = Tag("_LAM")

	// Unofficial. Latitude seconds. This has been seen exported from
	// MacFamilyFree.
	LatitudeSeconds = Tag("_LAS")

	// Unofficial. Longitude degrees. This has been seen exported from
	// MacFamilyFree.
	LongitudeDegress = Tag("_LOD")

	// Unofficial. Longitude minutes. This has been seen exported from
	// MacFamilyFree.
	LongitudeMinutes = Tag("_LOM")

	// Unofficial. Longitude north? This has been seen exported from
	// MacFamilyFree.
	LongitudeNorth = Tag("_LON")

	// Unofficial. Longitude seconds. This has been seen exported from
	// MacFamilyFree.
	LongitudeSeconds = Tag("_LOS")

	// Unofficial. Used to group the _LA* and _LO* tags for latitude and
	// longitude. This has been seen exported from MacFamilyFree.
	Coordinates = Tag("_COR")

	// Unofficial. The created date and/or time. This has been seen exported
	// from Ancestry.com.
	Created = Tag("_CRE")
)

func (tag Tag) String() string {
	m := map[Tag]string{
		Abbreviation:         "Abbreviation",
		Address1:             "Address1",
		Address2:             "Address2",
		Address:              "Address",
		Adoption:             "Adoption",
		AdultChristening:     "AdultChristening",
		Age:                  "Age",
		Agency:               "Agency",
		Alias:                "Alias",
		Ancestors:            "Ancestors",
		AncestorsInterest:    "AncestorsInterest",
		AncestralFileNumber:  "AncestralFileNumber",
		Annulment:            "Annulment",
		Associates:           "Associates",
		Author:               "Author",
		Baptism:              "Baptism",
		BarMitzvah:           "BarMitzvah",
		BasMitzvah:           "BasMitzvah",
		BinaryObject:         "BinaryObject",
		Birth:                "Birth",
		Blessing:             "Blessing",
		Burial:               "Burial",
		CallNumber:           "CallNumber",
		Caste:                "Caste",
		Cause:                "Cause",
		Census:               "Census",
		Change:               "Change",
		CharacterSet:         "CharacterSet",
		Child:                "Child",
		ChildrenCount:        "ChildrenCount",
		Christening:          "Christening",
		City:                 "City",
		Concatenation:        "Concatenation",
		Confirmation:         "Confirmation",
		Continued:            "Continued",
		Coordinates:          "Coordinates",
		Copyright:            "Copyright",
		Corporate:            "Corporate",
		Country:              "Country",
		Created:              "Created",
		Cremation:            "Cremation",
		Data:                 "Data",
		Date:                 "Date",
		Death:                "Death",
		Descendants:          "Descendants",
		DescendantsInterest:  "DescendantsInterest",
		Destination:          "Destination",
		Divorce:              "Divorce",
		DivorceFiled:         "DivorceFiled",
		Education:            "Education",
		Email:                "Email",
		Emigration:           "Emigration",
		Endowment:            "Endowment",
		Engagement:           "Engagement",
		Event:                "Event",
		Fact:                 "Fact",
		Family:               "Family",
		FamilyChild:          "FamilyChild",
		FamilyFile:           "FamilyFile",
		FamilySearchID:       "FamilySearchID",
		FamilySpouse:         "FamilySpouse",
		Fax:                  "Fax",
		File:                 "File",
		FirstCommunion:       "FirstCommunion",
		Format:               "Format",
		GedcomInformation:    "GedcomInformation",
		GivenName:            "GivenName",
		Graduation:           "Graduation",
		Header:               "Header",
		Husband:              "Husband",
		IdentityNumber:       "IdentityNumber",
		Immigration:          "Immigration",
		Individual:           "Individual",
		Language:             "Language",
		Latitude:             "Latitude",
		LatitudeDegrees:      "LatitudeDegrees",
		LatitudeMinutes:      "LatitudeMinutes",
		LatitudeSeconds:      "LatitudeSeconds",
		LDSBaptism:           "LDSBaptism",
		LDSConfirmation:      "LDSConfirmation",
		Legatee:              "Legatee",
		Longitude:            "Longitude",
		LongitudeDegress:     "LongitudeDegress",
		LongitudeMinutes:     "LongitudeMinutes",
		LongitudeNorth:       "LongitudeNorth",
		LongitudeSeconds:     "LongitudeSeconds",
		Map:                  "Map",
		Marriage:             "Marriage",
		MarriageBann:         "MarriageBann",
		MarriageContract:     "MarriageContract",
		MarriageCount:        "MarriageCount",
		MarriageLicence:      "MarriageLicence",
		MarriageSettlement:   "MarriageSettlement",
		Media:                "Media",
		Name:                 "Name",
		NamePrefix:           "NamePrefix",
		NameSuffix:           "NameSuffix",
		Nationality:          "Nationality",
		Naturalization:       "Naturalization",
		Nickname:             "Nickname",
		Note:                 "Note",
		Object:               "Object",
		Occupation:           "Occupation",
		Ordinance:            "Ordinance",
		Ordination:           "Ordination",
		Page:                 "Page",
		Pedigree:             "Pedigree",
		Phone:                "Phone",
		Phonetic:             "Phonetic",
		PhysicalDescription:  "PhysicalDescription",
		Place:                "Place",
		PostalCode:           "PostalCode",
		Probate:              "Probate",
		Property:             "Property",
		Publication:          "Publication",
		QualityOfData:        "QualityOfData",
		RecordFileNumber:     "RecordFileNumber",
		RecordIDNumber:       "RecordIDNumber",
		Reference:            "Reference",
		Relationship:         "Relationship",
		Religion:             "Religion",
		Repository:           "Repository",
		Residence:            "Residence",
		Restriction:          "Restriction",
		Retirement:           "Retirement",
		Role:                 "Role",
		Romanized:            "Romanized",
		SealingChild:         "SealingChild",
		SealingSpouse:        "SealingSpouse",
		Sex:                  "Sex",
		SocialSecurityNumber: "SocialSecurityNumber",
		Source:               "Source",
		State:                "State",
		Status:               "Status",
		Submission:           "Submission",
		Submitter:            "Submitter",
		Surname:              "Surname",
		SurnamePrefix:        "SurnamePrefix",
		Temple:               "Temple",
		Text:                 "Text",
		Time:                 "Time",
		Title:                "Title",
		Trailer:              "Trailer",
		Type:                 "Type",
		Version:              "Version",
		Wife:                 "Wife",
		Will:                 "Will",
		WWW:                  "WWW",
	}

	if s, ok := m[tag]; ok {
		return s
	}

	return string(tag)
}

// IsOfficial returns true if the tag is part of the GEDCOM 5.5 standard.
func (tag Tag) IsOfficial() bool {
	return tag[0] != '_'
}
