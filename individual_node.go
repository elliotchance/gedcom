package gedcom

import (
	"fmt"
	"strings"
	"time"
)

// IndividualNode represents a person.
type IndividualNode struct {
	*simpleDocumentNode
	cachedFamilies, cachedSpouses bool
	families                      FamilyNodes
	spouses                       []*IndividualNode
	cachedUniqueIDs               *StringSet
}

// SpouseChildren connects a single spouse to a set of children. The children
// may appear under multiple spouses. This is only useful when in combination
// with an individual (that would be the other spouse).
//
// The spouse can be nil, indicating that the spouse it not known for the
// assigned children. You should not assume that you can also recover the other
// spouse from one of the keys in this map as the map is valid to be empty or to
// only contain a nil key.
type SpouseChildren map[*IndividualNode]ChildNodes

func newIndividualNode(document *Document, pointer string, children ...Node) *IndividualNode {
	return &IndividualNode{
		newSimpleDocumentNode(document, TagIndividual, "", pointer, children...),
		false, false, nil, nil, nil,
	}
}

// TODO: Needs tests
//
// If the node is nil the result will also be nil.
func (node *IndividualNode) Name() *NameNode {
	if node == nil {
		return nil
	}

	nameTag := First(NodesWithTag(node, TagName))
	if nameTag != nil {
		return nameTag.(*NameNode)
	}

	return nil
}

// If the node is nil the result will also be nil.
func (node *IndividualNode) Names() []*NameNode {
	if node == nil {
		return nil
	}

	nameTags := NodesWithTag(node, TagName)
	names := make([]*NameNode, len(nameTags))

	for i, name := range nameTags {
		names[i] = name.(*NameNode)
	}

	return names
}

// If the node is nil the result will be SexUnknown.
func (node *IndividualNode) Sex() *SexNode {
	n := First(NodesWithTag(node, TagSex))
	if IsNil(n) {
		return nil
	}

	return n.(*SexNode)
}

// TODO: needs tests
//
// If the node is nil the result will also be nil.
func (node *IndividualNode) Spouses() (spouses IndividualNodes) {
	if node == nil {
		return nil
	}

	if node.cachedSpouses {
		return node.spouses
	}

	defer func() {
		node.spouses = spouses
		node.cachedSpouses = true
	}()

	spouses = IndividualNodes{}

	for _, family := range node.document.Families() {
		husband := family.Husband()
		wife := family.Wife()

		// We only care about families that have both parties (otherwise there
		// is no spouse to add).
		if husband == nil || wife == nil {
			continue
		}

		if husband.IsIndividual(node) {
			spouses = append(spouses, wife.Individual())
		}

		if wife.IsIndividual(node) {
			spouses = append(spouses, husband.Individual())
		}
	}

	return spouses
}

// TODO: needs tests
//
// If the node is nil the result will also be nil.
func (node *IndividualNode) Families() (families FamilyNodes) {
	if node == nil {
		return nil
	}

	if node.cachedFamilies {
		return node.families
	}

	defer func() {
		node.families = families
		node.cachedFamilies = true
	}()

	families = FamilyNodes{}

	for _, family := range node.document.Families() {
		hasChild := family.HasChild(node)
		isHusband := family.Husband().IsIndividual(node)
		isWife := family.Wife().IsIndividual(node)

		if hasChild || isHusband || isWife {
			families = append(families, family)
		}
	}

	return families
}

// TODO: needs tests
func (node *IndividualNode) Is(individual *IndividualNode) bool {
	if node == nil {
		return false
	}

	if individual == nil {
		return false
	}

	leftPointer := node.Pointer()
	rightPointer := individual.Pointer()

	return leftPointer == rightPointer
}

// TODO: needs tests
//
// If the node is nil the result will also be nil.
func (node *IndividualNode) FamilyWithSpouse(spouse *IndividualNode) *FamilyNode {
	if node == nil {
		return nil
	}

	for _, family := range node.document.Families() {
		a := family.Husband().IsIndividual(node) && family.Wife().IsIndividual(spouse)
		b := family.Wife().IsIndividual(node) && family.Husband().IsIndividual(spouse)

		if a || b {
			return family
		}
	}

	return nil
}

// TODO: needs tests
//
// If the node is nil the result will also be nil.
func (node *IndividualNode) FamilyWithUnknownSpouse() *FamilyNode {
	if node == nil {
		return nil
	}

	for _, family := range node.document.Families() {
		a := family.Husband().IsIndividual(node) && family.Wife() == nil
		b := family.Wife().IsIndividual(node) && family.Husband() == nil

		if a || b {
			return family
		}
	}

	return nil
}

// IsLiving determines if the individual is living.
//
// If no death node exists, but the estimated birth date is found to be older
// than Document.MaxLivingAge then the individual will not be considered living.
//
// The default value for MaxLivingAge is 100 and can be modified on the document
// attached to this node. A MaxLivingAge of 0 means that it will only consider
// the individual to be not living if there is an explicit Death event.
//
// If there is no document attached DefaultMaxLivingAge will be used.
//
// If the node is nil the result will always be false.
func (node *IndividualNode) IsLiving() bool {
	if node == nil {
		return false
	}

	deaths := node.Deaths()
	if len(deaths) > 0 {
		return false
	}

	maxLivingAge := node.Document().MaxLivingAge
	if maxLivingAge == 0 {
		return true
	}

	nowYear := float64(time.Now().Year())
	birthDate, _ := node.EstimatedBirthDate()
	birthYear := Years(birthDate)
	age := nowYear - birthYear

	return birthYear == 0 || age <= maxLivingAge
}

// Births returns zero or more birth events for the individual.
//
// If the node is nil the result will also be nil.
func (node *IndividualNode) Births() (nodes []*BirthNode) {
	if node == nil {
		return nil
	}

	for _, n := range NodesWithTag(node, TagBirth) {
		nodes = append(nodes, n.(*BirthNode))
	}

	return
}

// Baptisms returns zero or more baptism events for the individual. The baptisms
// do not include LDS baptisms.
//
// If the node is nil the result will also be nil.
func (node *IndividualNode) Baptisms() []*BaptismNode {
	nodes := NodesWithTag(node, TagBaptism)

	return nodes.CastTo((*BaptismNode)(nil)).([]*BaptismNode)
}

// Deaths returns zero or more death events for the individual. It is common for
// individuals to not have a death event if the death date is not known. If you
// need to check if an individual is living you should use IsLiving().
//
// If the node is nil the result will also be nil.
func (node *IndividualNode) Deaths() []*DeathNode {
	nodes := NodesWithTag(node, TagDeath)

	return nodes.CastTo((*DeathNode)(nil)).([]*DeathNode)
}

// Burials returns zero or more burial events for the individual.
//
// If the node is nil the result will also be nil.
func (node *IndividualNode) Burials() []*BurialNode {
	nodes := NodesWithTag(node, TagBurial)

	return nodes.CastTo((*BurialNode)(nil)).([]*BurialNode)
}

// Parents returns the families for which this individual is a child. There may
// be zero or more parents for an individual. The families returned will all
// reference this individual as child. However the father, mother or both may
// not exist.
//
// It is also possible to have duplicate families. That is, families that have
// the same husband and wife combinations if these families are defined in the
// GEDCOM file.
//
// If the node is nil the result will also be nil.
func (node *IndividualNode) Parents() FamilyNodes {
	if node == nil {
		return nil
	}

	parents := FamilyNodes{}

	for _, family := range node.Families() {
		if family.HasChild(node) {
			parents = append(parents, family)
		}
	}

	return parents
}

// SpouseChildren maps the known spouses to their children. The spouse will be
// nil if the other parent is not known for some or all of the children.
// Children can appear under multiple spouses.
//
// If the node is nil the result will also be nil.
func (node *IndividualNode) SpouseChildren() SpouseChildren {
	spouseChildren := SpouseChildren{}

	for _, family := range node.Families() {
		if !family.HasChild(node) {
			var spouse *IndividualNode

			switch {
			case family.Husband().IsIndividual(node):
				if wife := family.Wife(); wife != nil {
					spouse = family.Wife().Individual()
				}

			case family.Wife().IsIndividual(node):
				if husband := family.Husband(); husband != nil {
					spouse = husband.Individual()
				}
			}

			familyWithSpouse := node.FamilyWithSpouse(spouse)
			var children ChildNodes
			if familyWithSpouse != nil {
				children = familyWithSpouse.Children()
			}
			spouseChildren[spouse] = children

			// Find children with unknown spouse.
			unknownSpouseFamily := node.FamilyWithUnknownSpouse()
			if unknownSpouseFamily != nil {
				spouseChildren[nil] = unknownSpouseFamily.Children()
			}
		}
	}

	return spouseChildren
}

// LDSBaptisms returns zero or more LDS baptism events for the individual. These
// are not to be confused with Baptisms().
//
// If the node is nil the result will also be nil.
func (node *IndividualNode) LDSBaptisms() Nodes {
	return NodesWithTag(node, TagLDSBaptism)
}

// EstimatedBirthDate attempts to find the exact or approximate birth date of an
// individual. It does this by looking at the births, baptisms and LDS baptisms.
// If any of them contain a date then the lowest date value is returned based on
// the Years() value which takes in account the full date range.
//
// This logic is loosely based off the idea that if the birth date is not known
// that a baptism usually happens when the individual is quite young (and
// therefore close to the their birth date).
//
// It is worth noting that since EstimatedBirthDate returns the lowest possible
// date that an exact birth date will be ignored if another event happens in a
// range before that. For example, if an individual has a birth date of
// "9 Feb 1983" but the Baptism was "9 Jan 1983" then the Baptism is returned.
// This data must be wrong in either case but EstimatedBirthDate cannot make a
// sensible decision in this case so it always returned the earliest date.
//
// EstimatedBirthDate is useful when comparing individuals where the exact dates
// are less important that attempting to serve approximate information for
// comparison. You almost certainly do not want to use the EstimatedBirthDate
// value for anything meaningful aside from comparisons.
//
// If the node is nil the result will also be nil.
func (node *IndividualNode) EstimatedBirthDate() (date *DateNode, isBirthEvent bool) {
	births := Dates(NewNodes(node.Births())...)
	if len(births) > 0 {
		return births.Minimum(), true
	}

	baptisms := Dates(NewNodes(Compound(node.Baptisms(), node.LDSBaptisms()))...)
	if len(baptisms) > 0 {
		return baptisms.Minimum(), false
	}

	return nil, false
}

// EstimatedDeathDate attempts to find the exact or approximate death date of an
// individual. It does this by returning the earliest death date based on the
// value of Years(). If there are no death dates then it will attempt to return
// the minimum burial date.
//
// This logic is loosely based off the idea that if the death date is not known
// that a burial usually happens a short time after the death of the individual.
//
// It is worth noting that EstimatedDeathDate will always return a death date if
// one is present before falling back to a possibly more specific burial date.
// One example of this might be a death date that has a large range such as
// "1983 - 1993". The burial may be a much more specific date like "Apr 1985".
// This almost certainly indicates that the death date was around early 1985,
// however the larger death date range will still be returned.
//
// EstimatedDeathDate is useful when comparing individuals where the exact dates
// are less important that attempting to serve approximate information for
// comparison. You almost certainly do not want to use the EstimatedDeathDate
// value for anything meaningful aside from comparisons.
//
// If the node is nil the result will also be nil.
func (node *IndividualNode) EstimatedDeathDate() (date *DateNode, isDeathEvent bool) {
	deaths := Dates(NewNodes(node.Deaths())...)
	if len(deaths) > 0 {
		return deaths.Minimum(), true
	}
	burials := Dates(NewNodes(node.Burials())...)
	if len(burials) > 0 {
		return burials.Minimum(), false
	}

	// TODO: It might be good to include the probates as well?

	return nil, false
}

// Similarity calculates how similar two individuals are. The returned value
// will be between 0.0 and 1.0 where 1.0 means an exact match.
//
// You should prefer SurroundingSimilarity, a more advanced checker that uses
// this function as part of it's ultimate analysis.
//
// The similarity is based off three equally weighted components, the
// individuals name, estimated birth and estimated death date and is calculated
// as follows:
//
//   similarity = (nameSimilarity + birthSimilarity + deathSimilarity) / 3.0
//
// Individual names are compared with the StringSimilarity function that does
// not consider the punctuation and extra spacing.
//
// An individual may have more than one name, if this is the case then each name
// is checked and the highest matching combination is used.
//
// The birth and death dates use the EstimatedBirthDate and EstimatedDeathDate
// functions respectively. These functions are allowed to make some estimates
// when critical information like the birth date does not exist so there is more
// data to include in the comparison.
//
// Both dates are compared with the DateNode.Similarity function, which also
// returns a value of 0.0 to 1.0. To put simply the dates must existing within
// an error margin (for example, 10 years in either direction). Higher scores
// are awarded to dates that are more relatively closer to each other on a
// parabola. See DateNode.Similarity for a full explanation of how it deals with
// approximate dates and date ranges.
//
// It is safe to use Similarity when one or both of the individuals are nil.
// This will always result in a 0.5. It is a 0.5 not because it is a partial
// match but that a positive or negative match cannot be determined. This is
// important when Similarity is used is more extensive similarity calculations
// as to not unnecessarily skew the results.
//
// The options.MaxYears allows the error margin on dates to be adjusted. See
// DefaultMaxYearsForSimilarity for more information.
func (node *IndividualNode) Similarity(other *IndividualNode, options SimilarityOptions) float64 {
	if node == nil || other == nil {
		return 0.5
	}

	// Compare the matrix of names.
	nameSimilarity := 0.0

	for _, name1 := range node.Names() {
		for _, name2 := range other.Names() {
			similarity := StringSimilarity(name1.String(), name2.String(),
				options.JaroBoostThreshold, options.JaroPrefixSize)

			if similarity > nameSimilarity {
				nameSimilarity = similarity
			}
		}
	}

	// Compare the dates.
	leftEstimatedBirthDate, _ := node.EstimatedBirthDate()
	rightEstimatedBirthDate, _ := other.EstimatedBirthDate()
	birthSimilarity := leftEstimatedBirthDate.
		Similarity(rightEstimatedBirthDate, options.MaxYears)

	leftEstimatedDeathDate, _ := node.EstimatedDeathDate()
	rightEstimatedDeathDate, _ := other.EstimatedDeathDate()
	deathSimilarity := leftEstimatedDeathDate.
		Similarity(rightEstimatedDeathDate, options.MaxYears)

	// Final calculation.
	nameSimilarityRatio := nameSimilarity * options.NameToDateRatio
	avgBirthDeathSimilarity := (birthSimilarity + deathSimilarity) / 2.0
	inverseRatio := 1.0 - options.NameToDateRatio

	return nameSimilarityRatio + avgBirthDeathSimilarity*inverseRatio
}

// SurroundingSimilarity is a more advanced version of Similarity.
// SurroundingSimilarity also takes into account the immediate surrounding
// family. That is the parents, spouses and children have separate metrics
// calculated so they can be interpreted differently or together.
//
// Checking for surrounding family is critical for calculating the similarity of
// individuals that would otherwise be considered the same because of similar
// names and dates in large family trees.
//
// SurroundingSimilarity returns a structure of the same name, but really it
// calculates four discreet similarities:
//
// 1. IndividualSimilarity: This is the same as Individual.Similarity().
//
// 2. ParentsSimilarity: The similarity of the fathers and mothers of the
// individual. Each missing parent will be given 0.5. If both parents are
// missing the parent similarity will also be 0.5.
//
// An individual can have zero or more pairs of parents, but only a single
// ParentsSimilarity is returned. The ParentsSimilarity is the highest value
// when each of the parents are compared with the other parents of the other
// individual.
//
// 3. SpousesSimilarity: The similarity of the spouses is compared with
// IndividualNodes.Similarity() which is designed to compare several individuals
// at once. It also handles comparing a different number of individuals on
// either side.
//
// 4. ChildrenSimilarity: Children are also compared with
// IndividualNodes.Similarity() but without respect to their parents (which in
// this case would be the current individual and likely one of their spouses).
// It is done this way as to not skew the results if any particular parent is
// unknown or the child is connected to a different spouse.
//
// The options.MaxYears allows the error margin on dates to be adjusted. See
// DefaultMaxYearsForSimilarity for more information.
//
// The options.MinimumSimilarity is used when comparing slices of individuals.
// In this case that means for the spouses and children. A higher value makes
// the matching more strict. See DefaultMinimumSimilarity for more information.
func (node *IndividualNode) SurroundingSimilarity(other *IndividualNode, options SimilarityOptions, forceFullCalculation bool) (s *SurroundingSimilarity) {
	// Individual, spouse and children similarity only needs to be calculated
	// once. The parents similarity will be calculated from the matrix below.
	individualSimilarity := node.Similarity(other, options)

	// Comparing individuals is extremely expensive because of the matrix of
	// comparisons and the individual comparisons themselves need to utilise a
	// lot of surrounding data.
	//
	// The individual similarity weight is by far the greatest, at 80% by
	// default. This means we can in most cases avoid calculating the parents,
	// spouses and children if we know that even in the perfect case the result
	// similarity would not be above the minimum threshold.
	//
	// Consider the formula:
	//
	//   (IndividualSimilarity * IndividualWeight) +
	//   (ParentsSimilarity * ParentsWeight) +
	//   (SpousesSimilarity * SpousesWeight) +
	//   (ChildrenSimilarity * ChildrenWeight) > MinimumWeight
	//
	// Isolating the individual similarity (which is the only thing we have
	// calculated):
	//
	//   (IndividualSimilarity * IndividualWeight) >
	//   MinimumWeight -
	//   (ParentsSimilarity * ParentsWeight) -
	//   (SpousesSimilarity * SpousesWeight) -
	//   (ChildrenSimilarity * ChildrenWeight)
	//
	// Now, lets assume that we have perfect matches (1.0) for all other
	// elements:
	//
	//   (IndividualSimilarity * IndividualWeight) >
	//   MinimumWeight - ParentsWeight - SpousesWeight - ChildrenWeight
	//
	// If that statement is not true, there is no need to proceed because the
	// individuals will never be considered a match given the MinimumWeight.
	//
	// A higher MinimumWeight means that less work will actually need to be done
	// because there will be less possible candidates.
	//
	// ghost:ignore
	if !forceFullCalculation && options.canSkipExtraProcessing(individualSimilarity) {
		return NewSurroundingSimilarity(0, 0, 0, 0)
	}

	spousesSimilarity := node.Spouses().Similarity(other.Spouses(), options)
	childrenSimilarity := node.Children().Individuals().
		Similarity(other.Children().Individuals(), options)

	s = NewSurroundingSimilarity(
		0.0, // Parents. Filled in later.
		individualSimilarity,
		spousesSimilarity,
		childrenSimilarity,
	)
	s.Options = options

	didFindParents := false
	for _, parents1 := range node.Parents() {
		for _, parents2 := range other.Parents() {
			didFindParents = true

			// depth of 0 means only the wife/husband is compared.
			similarity := parents1.Similarity(parents2, 0, options)

			if similarity > s.ParentsSimilarity {
				s.ParentsSimilarity = similarity
			}
		}
	}

	if !didFindParents {
		s.ParentsSimilarity = 0.5
	}

	return
}

// TODO: Needs tests
//
// If the node is nil the result will also be nil.
func (node *IndividualNode) Children() ChildNodes {
	children := ChildNodes{}

	for _, family := range node.Families() {
		if !family.HasChild(node) {
			children = append(children, family.Children()...)
		}
	}

	return children
}

// AllEvents returns zero or more events of any kind for the individual.
//
// This is not to be confused with the EventNode.
func (node *IndividualNode) AllEvents() (nodes Nodes) {
	for _, n := range node.Nodes() {
		if n.Tag().IsEvent() {
			nodes = append(nodes, n)
		}
	}

	return
}

// Birth returns the first values for the date and place of the birth events.
func (node *IndividualNode) Birth() (*DateNode, *PlaceNode) {
	birthNodes := Compound(node.Births())

	return DateAndPlace(birthNodes...)
}

// Death returns the first values for the date and place of the death events.
func (node *IndividualNode) Death() (*DateNode, *PlaceNode) {
	deathNodes := Compound(node.Deaths())

	return DateAndPlace(deathNodes...)
}

// Baptism returns the first values for the date and place of the baptism
// events.
func (node *IndividualNode) Baptism() (*DateNode, *PlaceNode) {
	baptismNodes := Compound(node.Baptisms())

	return DateAndPlace(baptismNodes...)
}

// Burial returns the first values for the date and place of the burial events.
func (node *IndividualNode) Burial() (*DateNode, *PlaceNode) {
	burialNodes := Compound(node.Burials())

	return DateAndPlace(burialNodes...)
}

// Age returns the best estimate minimum and maximum age of the individual if
// they are still living or their approximate age range at the time of their
// death.
//
// Age will use EstimatedBirthDate which will allow it to be more resilient if a
// birth date is strictly missing but this means that the value returned from
// Age may not always be exact. See IsEstimate and other fields returned in
// either Age return values.
//
// If the birth or death date is a range, such as "Between 1945 and 1947" the
// minimum and maximum possible ages will be returned.
//
// If the birth or death date is not an exact value then the ages will be
// estimates. See Age.IsEstimate.
//
// If no estimated birth date can be determined then Age.IsKnown will be false.
//
// IsLiving will be used to determine if the individual is still living which
// includes a maximum possible age when no death elements exist. However, in the
// case that the individual is not living then EstimatedDeathDate will be used
// to try and estimate the age at the time of death instead of simply using the
// maximum possible age (which is 100 by default).
func (node *IndividualNode) Age() (Age, Age) {
	now := NewDateRangeWithNow()
	startAge, endAge := node.ageAt(now)

	// Unlike AgeAt, we always want to trim back to the death date.
	if startAge.Constraint == AgeConstraintAfterDeath {
		estimatedBirthDate, _ := node.EstimatedBirthDate()
		estimatedDeathDate, _ := node.EstimatedDeathDate()
		ageInYears := estimatedDeathDate.Years() - estimatedBirthDate.Years()

		startAge = NewAgeWithYears(ageInYears, false, AgeConstraintAfterDeath)
		endAge = NewAgeWithYears(ageInYears, false, AgeConstraintAfterDeath)
	}

	return startAge, endAge
}

// AgeAt follows the same logic as Age but uses an event as the comparison
// instead of the current time.
//
// If there is more than one date associated with the event or a date contains a
// range the minimum and maximum values will be used to return the full range.
func (node *IndividualNode) AgeAt(event Node) (Age, Age) {
	dates := Dates(event).StripZero()

	if len(dates) > 0 {
		dateRange := dates.Range()

		return node.ageAt(dateRange)
	}

	// We cannot determine the date of the event so we would not know what age
	// they were at the time.
	return NewUnknownAge(), NewUnknownAge()
}

func (node *IndividualNode) ageAt(at DateRange) (Age, Age) {
	estimatedBirthDate, isBirthEvent := node.EstimatedBirthDate()

	// If we have no idea when they are born we cannot proceed with any age
	// estimate.
	if !estimatedBirthDate.IsValid() {
		return NewUnknownAge(), NewUnknownAge()
	}

	estimatedDeathDate, isDeathEvent := node.EstimatedDeathDate()
	birthRange := estimatedBirthDate.DateRange()
	startAge, endAge := at.Sub(birthRange).Age()

	// There are some cases where the endAge can be before the startAge in some
	// combinations of constraints. If this is the case we swap them to make the
	// output more sensible.
	if startAge.IsAfter(endAge) {
		startAge, endAge = endAge, startAge
	}

	isEstimate := !isBirthEvent || !isDeathEvent ||
		!birthRange.IsExact() || !at.IsExact()
	startAge.IsEstimate = isEstimate
	endAge.IsEstimate = isEstimate

	switch {
	case at.IsBefore(estimatedBirthDate.DateRange()):
		startAge.Constraint = AgeConstraintBeforeBirth
		endAge.Constraint = AgeConstraintBeforeBirth

	case at.IsAfter(estimatedDeathDate.DateRange()) && estimatedDeathDate != nil:
		startAge.Constraint = AgeConstraintAfterDeath
		endAge.Constraint = AgeConstraintAfterDeath

	default:
		startAge.Constraint = AgeConstraintLiving
		endAge.Constraint = AgeConstraintLiving
	}

	return startAge, endAge
}

// String returns a human-readable representation of the individual like:
//
//   (no name) (b. Aft. 1983)
//   Bob Smith (b. 1943)
//   John Chance
//   Jane Doe (b. 3 Apr 1923, bur. Abt. 1943)
//
// Ideally it will use birth (b.) and death (d.) if available. However, it will
// fall back to the baptism (bap.) or burial (bur.) respectively.
func (node *IndividualNode) String() string {
	name := String(node.Name())
	if name == "" {
		name = "(no name)"
	}

	dateParts := []string{}

	if birth, _ := node.Birth(); birth != nil && birth.DateRange().IsValid() {
		dateParts = append(dateParts, fmt.Sprintf("b. %s", birth.String()))
	} else if baptism, _ := node.Baptism(); baptism != nil && baptism.DateRange().IsValid() {
		dateParts = append(dateParts, fmt.Sprintf("bap. %s", baptism.String()))
	}

	if death, _ := node.Death(); death != nil && death.DateRange().IsValid() {
		dateParts = append(dateParts, fmt.Sprintf("d. %s", death.String()))
	} else if burial, _ := node.Burial(); burial != nil && burial.DateRange().IsValid() {
		dateParts = append(dateParts, fmt.Sprintf("bur. %s", burial.String()))
	}

	if len(dateParts) == 0 {
		return name
	}

	return fmt.Sprintf("%s (%s)", name, strings.Join(dateParts, ", "))
}

func (node *IndividualNode) FamilySearchIDs() (nodes []*FamilySearchIDNode) {
	if node == nil {
		return nil
	}

	for _, tag := range FamilySearchIDNodeTags() {
		for _, n := range NodesWithTag(node, tag) {
			nodes = append(nodes, n.(*FamilySearchIDNode))
		}
	}

	return
}

func (node *IndividualNode) UniqueIDs() (nodes []*UniqueIDNode) {
	if node == nil {
		return nil
	}

	for _, n := range NodesWithTag(node, UnofficialTagUniqueID) {
		nodes = append(nodes, n.(*UniqueIDNode))
	}

	return
}

// UniqueIdentifiers returns any strings that can be used to uniquely this
// individual in this document, but also between documents. For this reason the
// individual pointer is not included.
//
// It's quite possible to receive an empty set, but it could also contain
// commonly unique identifiers such as the FamilySearch ID or UUID generated by
// some applications.
func (node *IndividualNode) UniqueIdentifiers() *StringSet {
	if node.cachedUniqueIDs == nil {
		node.cachedUniqueIDs = NewStringSet()

		for _, id := range node.UniqueIDs() {
			if uuid, err := id.UUID(); err == nil {
				node.cachedUniqueIDs.Add(uuid.String())
			}
		}

		for _, id := range node.FamilySearchIDs() {
			node.cachedUniqueIDs.Add(id.String())
		}
	}

	return node.cachedUniqueIDs
}

func (node *IndividualNode) resetCache() {
	node.cachedFamilies = false
	node.cachedSpouses = false
	node.families = nil
	node.spouses = nil
	node.cachedUniqueIDs = nil
}

func (node *IndividualNode) AddName(name string) *IndividualNode {
	node.AddNode(NewNameNode(name))

	return node
}

func (node *IndividualNode) AddBirthDate(birthDate string) *IndividualNode {
	existingBirth := First(node.Births())
	if existingBirth == nil {
		existingBirth = NewBirthNode("")
		node.AddNode(existingBirth)
	}

	existingBirth.AddNode(NewDateNode(birthDate))

	return node
}

// SetSex adds or replaces tge gender of an individual. You should use one of
// the SexMale, SexFemale or SexUnknown constants.
func (node *IndividualNode) SetSex(sex string) *IndividualNode {
	existingSex := First(NodesWithTag(node, TagSex))
	if existingSex == nil {
		node.AddNode(NewNode(TagSex, string(sex), ""))
	} else {
		existingSex.RawSimpleNode().value = string(sex)
	}

	return node
}

func (node *IndividualNode) AddBurialDate(burialDate string) *IndividualNode {
	existingBurial := First(node.Burials())
	if existingBurial == nil {
		existingBurial = NewBurialNode("")
		node.AddNode(existingBurial)
	}

	existingBurial.AddNode(NewDateNode(burialDate))

	return node
}

func (node *IndividualNode) AddBaptismDate(baptismDate string) *IndividualNode {
	existingBaptism := First(node.Baptisms())
	if existingBaptism == nil {
		existingBaptism = NewBaptismNode("")
		node.AddNode(existingBaptism)
	}

	existingBaptism.AddNode(NewDateNode(baptismDate))

	return node
}

type eventAndDate struct {
	Event Node
	Date  *DateNode
}

func (node *IndividualNode) incorrectEventOrderWarnings() (warnings Warnings) {
	// Event order describes the boundaries of groups of events. That is to say
	// that any baptism or LDS baptism events must be after a birth event but
	// also much be before the any death event.
	eventOrder := []*struct {
		Tags   []Tag
		Events []eventAndDate
	}{
		{
			Tags: []Tag{TagBirth},
		},
		{
			Tags: []Tag{TagBaptism, TagLDSBaptism},
		},
		{
			Tags: []Tag{TagDeath},
		},
		{
			Tags: []Tag{TagBurial},
		},
	}

	// Collect all of the dates. This just makes it easier to compare later.
	// Each of the dates found in events are partnered with the original event
	// as we will need both if it turns into a warning.
	for _, group := range eventOrder {
		for _, tag := range group.Tags {
			nodes := NodesWithTag(node, tag)
			for _, node := range nodes {
				for _, date := range NodesWithTag(node, TagDate) {
					group.Events = append(group.Events, eventAndDate{
						Event: node,
						Date:  date.(*DateNode),
					})
				}
			}
		}
	}

	// Check the order. This is an iterative approach where each group is
	// compared with all of the events that proceed it.
	for i, group := range eventOrder {
		for _, event := range group.Events {
			for _, futureGroup := range eventOrder[i+1:] {
				for _, futureEvent := range futureGroup.Events {
					if event.Date.IsValid() &&
						futureEvent.Date.IsValid() {
						comparison := futureEvent.Date.DateRange().
							Compare(event.Date.DateRange())

						if comparison == DateRangeComparisonEntirelyBefore {
							warning := NewIncorrectEventOrderWarning(
								futureEvent.Event, futureEvent.Date.DateRange(),
								event.Event, event.Date.DateRange(),
							)
							warning.SetContext(WarningContext{
								Individual: node,
							})
							warnings = append(warnings, warning)
						}
					}
				}
			}
		}
	}

	return
}

func (node *IndividualNode) tooOldWarnings() (warnings Warnings) {
	estimatedDeathDate, _ := node.EstimatedDeathDate()
	_, max := node.Age()
	if max.Years() > DefaultMaxLivingAge && estimatedDeathDate != nil {
		warnings = Warnings{
			NewIndividualTooOldWarning(node, max.Years()),
		}
	}

	return
}

func (node *IndividualNode) multipleSexesWarnings() Warnings {
	sexes := castNodesWithTag(node, TagSex, (*SexNode)(nil)).([]*SexNode)
	if len(sexes) > 1 {
		return Warnings{
			NewMultipleSexesWarning(node, sexes),
		}
	}

	return nil
}

func (node *IndividualNode) Warnings() (warnings Warnings) {
	warnings = append(warnings, node.incorrectEventOrderWarnings()...)
	warnings = append(warnings, node.tooOldWarnings()...)
	warnings = append(warnings, node.multipleSexesWarnings()...)

	return
}

func (node *IndividualNode) AddDeathDate(deathDate string) *IndividualNode {
	existingDeath := First(node.Deaths())
	if existingDeath == nil {
		existingDeath = NewDeathNode("")
		node.AddNode(existingDeath)
	}

	existingDeath.AddNode(NewDateNode(deathDate))

	return node
}
