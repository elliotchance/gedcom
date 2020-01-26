// "gedcom tune" is used to calculate the ideal weights and similarities for the
// main gedcom package.
//
// It works by comparing two GEDCOM files that are mostly the same, but must
// have the same pointers for individuals. It uses tries to calculate the best
// values that would lead the Similarity functions to the highest number of
// matches (which are confirmed by the individual pointers).
//
// The process works like this:
//
// 1. Load the two GEDCOM files.
//
// 2. Set predefined or random values for weightings.
//
// 3. Match the two files. One point is awarded to a successful match and one
// point is removed for each unsuccessful match.
//
// 4. Steps 2 and 3 are repeated many more times with different weightings.
//
// 5. The weighting values that scored the highest points are returned.
package main

import (
	"flag"
	"fmt"
	"github.com/elliotchance/gedcom"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
)

type TuneFlags struct {
	// Input files.
	optionGedcomFile1 string
	optionGedcomFile2 string
	optionRandom      bool

	// Profiling.
	optionCPUProfileOutput string

	// minimumSimilarity
	optionSimilarityMin  float64
	optionSimilarityMax  float64
	optionSimilarityStep float64

	// maxYears
	optionYearsMin  float64
	optionYearsMax  float64
	optionYearsStep float64

	// Weights
	optionsWeightIndividualMin  float64
	optionsWeightIndividualMax  float64
	optionsWeightIndividualStep float64

	// Name to date ratio
	optionsNameToDateRatioMin  float64
	optionsNameToDateRatioMax  float64
	optionsNameToDateRatioStep float64

	// Jaro boost threshold
	optionsJaroBoostMin  float64
	optionsJaroBoostMax  float64
	optionsJaroBoostStep float64

	// Jaro prefix size
	optionsJaroPrefixSizeMin  int
	optionsJaroPrefixSizeMax  int
	optionsJaroPrefixSizeStep int
}

func runTuneCommand() {
	tuneFlags := &TuneFlags{}

	// Input files. Must be provided.
	flag.StringVar(&tuneFlags.optionGedcomFile1, "gedcom1", "", "First GEDCOM file.")
	flag.StringVar(&tuneFlags.optionGedcomFile2, "gedcom2", "", "Second GEDCOM file.")
	flag.BoolVar(&tuneFlags.optionRandom, "random", false, "Run forever with random values.")

	// Profiling.
	flag.StringVar(&tuneFlags.optionCPUProfileOutput, "cpu-profile", "", "If enabled "+
		"the CPU profile file will be created or replaced. This is needed to "+
		"optimise the comparison process.")

	// minimumSimilarity
	flag.Float64Var(&tuneFlags.optionSimilarityMin, "similarity-min",
		gedcom.DefaultMinimumSimilarity, "Lower bound for minimumSimilarity.")
	flag.Float64Var(&tuneFlags.optionSimilarityMax, "similarity-max",
		gedcom.DefaultMinimumSimilarity, "Upper bound for minimumSimilarity.")
	flag.Float64Var(&tuneFlags.optionSimilarityStep, "similarity-step", 0.1,
		"Step size for minimumSimilarity.")

	// maxYears
	flag.Float64Var(&tuneFlags.optionYearsMin, "years-min",
		gedcom.DefaultMaxYearsForSimilarity, "Lower bound for maxYears.")
	flag.Float64Var(&tuneFlags.optionYearsMax, "years-max",
		gedcom.DefaultMaxYearsForSimilarity, "Upper bound for maxYears.")
	flag.Float64Var(&tuneFlags.optionYearsStep, "years-step", 1,
		"Step size for maxYears.")

	// Weights
	flag.Float64Var(&tuneFlags.optionsWeightIndividualMin, "weight-individual-min", 0.8,
		"Lower bound for individual weight.")
	flag.Float64Var(&tuneFlags.optionsWeightIndividualMax, "weight-individual-max", 0.8,
		"Upper bound for individual weight.")
	flag.Float64Var(&tuneFlags.optionsWeightIndividualStep, "weight-individual-step", 0.05,
		"Step size for individual weight.")

	// Name ratio
	flag.Float64Var(&tuneFlags.optionsNameToDateRatioMin, "name-ratio-min", 0.5,
		"Lower bound for name to date ratio.")
	flag.Float64Var(&tuneFlags.optionsNameToDateRatioMax, "name-ratio-max", 0.5,
		"Upper bound for name to date ratio.")
	flag.Float64Var(&tuneFlags.optionsNameToDateRatioStep, "name-ratio-step", 0.1,
		"Step size for name to date ratio.")

	// Jaro boost threshold
	flag.Float64Var(&tuneFlags.optionsJaroBoostMin, "jaro-boost-min", 0.0,
		"Lower bound for jaro boost threshold.")
	flag.Float64Var(&tuneFlags.optionsJaroBoostMax, "jaro-boost-max", 0.0,
		"Upper bound for jaro boost threshold.")
	flag.Float64Var(&tuneFlags.optionsJaroBoostStep, "jaro-boost-step", 0.1,
		"Step size for jaro boost threshold.")

	// Jaro prefix size
	flag.IntVar(&tuneFlags.optionsJaroPrefixSizeMin, "jaro-prefix-min", 8,
		"Lower bound for jaro prefix size.")
	flag.IntVar(&tuneFlags.optionsJaroPrefixSizeMax, "jaro-prefix-max", 8,
		"Upper bound for jaro prefix size.")
	flag.IntVar(&tuneFlags.optionsJaroPrefixSizeStep, "jaro-prefix-step", 1,
		"Step size for jaro prefix size.")

	err := flag.CommandLine.Parse(os.Args[2:])
	if err != nil {
		fatalln(err)
	}

	if tuneFlags.optionGedcomFile1 == "" {
		fatalln("-gedcom1 is required")
	}

	if tuneFlags.optionGedcomFile2 == "" {
		fatalln("-gedcom2 is required")
	}

	// CPU profiler
	if tuneFlags.optionCPUProfileOutput != "" {
		log.Printf("Starting CPU profiler.")
		startCPUProfiler(tuneFlags.optionCPUProfileOutput)

		defer pprof.StopCPUProfile()
	}

	gedcom1, err := gedcom.NewDocumentFromGEDCOMFile(tuneFlags.optionGedcomFile1)
	if err != nil {
		fatalln(err)
	}

	gedcom2, err := gedcom.NewDocumentFromGEDCOMFile(tuneFlags.optionGedcomFile2)
	if err != nil {
		fatalln(err)
	}

	// Calculate ideal score.
	idealScore := 0
	for _, i1 := range gedcom1.Individuals() {
		for _, i2 := range gedcom2.Individuals() {
			if i1.Pointer() == i2.Pointer() {
				idealScore += 1
			}
		}
	}

	// Run compare.
	options := gedcom.NewSimilarityOptions()
	if tuneFlags.optionRandom {
		for {
			options.MinimumSimilarity = random(tuneFlags.optionSimilarityMin, tuneFlags.optionSimilarityMax)
			options.MinimumWeightedSimilarity = random(tuneFlags.optionSimilarityMin, tuneFlags.optionSimilarityMax)
			options.IndividualWeight = random(tuneFlags.optionsWeightIndividualMin, tuneFlags.optionsWeightIndividualMax)
			options.SpousesWeight = (1.0 - options.IndividualWeight) / 3
			options.ParentsWeight = (1.0 - options.IndividualWeight) / 3
			options.ChildrenWeight = (1.0 - options.IndividualWeight) / 3
			options.MaxYears = random(tuneFlags.optionYearsMin, tuneFlags.optionYearsMax)
			options.NameToDateRatio = random(tuneFlags.optionsNameToDateRatioMin, tuneFlags.optionsNameToDateRatioMax)
			options.JaroBoostThreshold = random(tuneFlags.optionsJaroBoostMin, tuneFlags.optionsJaroBoostMax)
			options.JaroPrefixSize = int(random(float64(tuneFlags.optionsJaroPrefixSizeMin), float64(tuneFlags.optionsJaroPrefixSizeMax)))

			run(gedcom1, gedcom2, idealScore, options)
		}
	}

	runMinimumSimilarity(gedcom1, gedcom2, idealScore, options, tuneFlags)
}

func random(min, max float64) float64 {
	// ghost:ignore
	return min + rand.Float64()*(max-min)
}

func runMinimumSimilarity(gedcom1 *gedcom.Document, gedcom2 *gedcom.Document, idealScore int, options gedcom.SimilarityOptions, tuneFlags *TuneFlags) {
	for minimumSimilarity := tuneFlags.optionSimilarityMin; minimumSimilarity <= tuneFlags.optionSimilarityMax; minimumSimilarity += tuneFlags.optionSimilarityStep {
		options.MinimumWeightedSimilarity = minimumSimilarity
		options.MinimumSimilarity = minimumSimilarity

		runMaxYears(gedcom1, gedcom2, idealScore, options, tuneFlags)
	}
}

func runMaxYears(gedcom1 *gedcom.Document, gedcom2 *gedcom.Document, idealScore int, options gedcom.SimilarityOptions, tuneFlags *TuneFlags) {
	for maxYears := tuneFlags.optionYearsMin; maxYears <= tuneFlags.optionYearsMax; maxYears += tuneFlags.optionYearsStep {
		options.MaxYears = maxYears

		runIndividualWeight(gedcom1, gedcom2, idealScore, options, tuneFlags)
	}
}

func runIndividualWeight(gedcom1 *gedcom.Document, gedcom2 *gedcom.Document, idealScore int, options gedcom.SimilarityOptions, tuneFlags *TuneFlags) {
	for x := tuneFlags.optionsWeightIndividualMin; x <= tuneFlags.optionsWeightIndividualMax; x += tuneFlags.optionsWeightIndividualStep {
		options.IndividualWeight = x
		options.SpousesWeight = (1.0 - x) / 3
		options.ParentsWeight = (1.0 - x) / 3
		options.ChildrenWeight = (1.0 - x) / 3

		runNameToDateRatio(gedcom1, gedcom2, idealScore, options, tuneFlags)
	}
}

func runNameToDateRatio(gedcom1 *gedcom.Document, gedcom2 *gedcom.Document, idealScore int, options gedcom.SimilarityOptions, tuneFlags *TuneFlags) {
	for x := tuneFlags.optionsNameToDateRatioMin; x <= tuneFlags.optionsNameToDateRatioMax; x += tuneFlags.optionsNameToDateRatioStep {
		options.NameToDateRatio = x

		runJaroBoost(gedcom1, gedcom2, idealScore, options, tuneFlags)
	}
}

func runJaroBoost(gedcom1 *gedcom.Document, gedcom2 *gedcom.Document, idealScore int, options gedcom.SimilarityOptions, tuneFlags *TuneFlags) {
	for x := tuneFlags.optionsJaroBoostMin; x <= tuneFlags.optionsJaroBoostMax; x += tuneFlags.optionsJaroBoostStep {
		options.JaroBoostThreshold = x

		runJaroPrefixSize(gedcom1, gedcom2, idealScore, options, tuneFlags)
	}
}

func runJaroPrefixSize(gedcom1 *gedcom.Document, gedcom2 *gedcom.Document, idealScore int, options gedcom.SimilarityOptions, tuneFlags *TuneFlags) {
	for x := tuneFlags.optionsJaroPrefixSizeMin; x <= tuneFlags.optionsJaroPrefixSizeMax; x += tuneFlags.optionsJaroPrefixSizeStep {
		options.JaroPrefixSize = x

		run(gedcom1, gedcom2, idealScore, options)
	}
}

func startCPUProfiler(optionCPUProfileOutput string) {
	f, err := os.Create(optionCPUProfileOutput)
	if err != nil {
		panic(err)
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}
}

func run(gedcom1, gedcom2 *gedcom.Document, idealScore int, options gedcom.SimilarityOptions) {
	compareOptions := gedcom.NewIndividualNodesCompareOptions()
	compareOptions.SimilarityOptions = options

	comparisons := gedcom1.Individuals().Compare(gedcom2.Individuals(), compareOptions)

	score := 0.0
	for _, comparison := range comparisons {
		if comparison.Left != nil && comparison.Right != nil {
			if comparison.Left.Pointer() == comparison.Right.Pointer() {
				score += 1
			} else {
				score -= 1
			}
		}
	}

	adjustedScore := score / float64(idealScore)
	fmt.Printf("%s, Score:%.6f\n", options, adjustedScore)
}
