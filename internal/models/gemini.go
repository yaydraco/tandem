package models

const (
	ProviderGemini ModelProvider = "gemini"

	// Models
	Gemini25Flash     ModelID = "gemini-2.5-flash"
	Gemini25Pro       ModelID = "gemini-2.5-pro"
	Gemini25FlashLite ModelID = "gemini-2.5-flash-lite"

	Gemini20Flash     ModelID = "gemini-2.0-flash"
	Gemini20FlashLite ModelID = "gemini-2.0-flash-lite"
)

var GeminiModels = map[ModelID]Model{

	Gemini25Flash: {
		ID:                  Gemini25Flash,
		Name:                "Gemini 2.5 Flash",
		Provider:            ProviderGemini,
		APIModel:            "gemini-2.5-flash",
		CostPer1MIn:         0.15,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0,
		CostPer1MOut:        0.60,
		ContextWindow:       1000000,
		DefaultMaxTokens:    50000,
		SupportsAttachments: true,
	},
	Gemini25Pro: {
		ID:                  Gemini25Pro,
		Name:                "Gemini 2.5 Pro",
		Provider:            ProviderGemini,
		APIModel:            "gemini-2.5-pro",
		CostPer1MIn:         1.25,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0,
		CostPer1MOut:        10,
		ContextWindow:       1000000,
		DefaultMaxTokens:    50000,
		SupportsAttachments: true,
	},
	Gemini25FlashLite: {
		ID:                  Gemini25FlashLite,
		Name:                "Gemini 2.5 Flash Lite",
		Provider:            ProviderGemini,
		APIModel:            "gemini-2.5-flash-lite",
		CostPer1MIn:         0.10,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0.03,
		CostPer1MOut:        0.40,
		ContextWindow:       65536,
		DefaultMaxTokens:    65556,
		SupportsAttachments: true,
	},

	Gemini20Flash: {
		ID:                  Gemini20Flash,
		Name:                "Gemini 2.0 Flash",
		Provider:            ProviderGemini,
		APIModel:            "gemini-2.0-flash",
		CostPer1MIn:         0.10,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0,
		CostPer1MOut:        0.40,
		ContextWindow:       1000000,
		DefaultMaxTokens:    6000,
		SupportsAttachments: true,
	},
	Gemini20FlashLite: {
		ID:                  Gemini20FlashLite,
		Name:                "Gemini 2.0 Flash Lite",
		Provider:            ProviderGemini,
		APIModel:            "gemini-2.0-flash-lite",
		CostPer1MIn:         0.05,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0,
		CostPer1MOut:        0.30,
		ContextWindow:       1000000,
		DefaultMaxTokens:    6000,
		SupportsAttachments: true,
	},
}
