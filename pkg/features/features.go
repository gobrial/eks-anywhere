package features

const (
	TinkerbellProviderEnvVar = "TINKERBELL_PROVIDER"
	CloudStackProviderEnvVar = "CLOUDSTACK_PROVIDER"
	SnowProviderEnvVar       = "SNOW_PROVIDER"
	FullLifecycleAPIEnvVar   = "FULL_LIFECYCLE_API"
	FullLifecycleGate        = "FullLifecycleAPI"
	CuratedPackagesEnvVar    = "CURATED_PACKAGES_SUPPORT"
)

func FeedGates(featureGates []string) {
	globalFeatures.feedGates(featureGates)
}

type Feature struct {
	Name     string
	IsActive func() bool
}

func IsActive(feature Feature) bool {
	return feature.IsActive()
}

func FullLifecycleAPI() Feature {
	return Feature{
		Name:     "Full lifecycle API support through the EKS-A controller",
		IsActive: globalFeatures.isActiveForEnvVarOrGate(FullLifecycleAPIEnvVar, FullLifecycleGate),
	}
}

func TinkerbellProvider() Feature {
	return Feature{
		Name:     "Tinkerbell provider support",
		IsActive: globalFeatures.isActiveForEnvVar(TinkerbellProviderEnvVar),
	}
}

func CloudStackProvider() Feature {
	return Feature{
		Name:     "CloudStack provider support",
		IsActive: globalFeatures.isActiveForEnvVar(CloudStackProviderEnvVar),
	}
}

func SnowProvider() Feature {
	return Feature{
		Name:     "Snow provider support",
		IsActive: globalFeatures.isActiveForEnvVar(SnowProviderEnvVar),
	}
}

func CuratedPackagesSupport() Feature {
	return Feature{
		Name:     "Curated Packages Support",
		IsActive: globalFeatures.isActiveForEnvVar(CuratedPackagesEnvVar),
	}
}
