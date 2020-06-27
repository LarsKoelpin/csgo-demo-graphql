package domain

var DefaultPositionTemplate = map[string]interface{}{
	"x": true,
	"y": true,
}

var DefaultPlayerTemplate = map[string]interface{}{
	"name":          true,
	"entityId":      true,
	"position":      DefaultPositionTemplate,
	"defusing":      true,
	"angleX":        true,
	"angleY":        true,
	"hp":            true,
	"team":          true,
	"armor":         true,
	"flashDuration": true,
	"npc":           true,
	"hasHelmet":     true,
	"hasDefuseKit":  true,
	"equipment": map[string]interface{}{
		"type":           true,
		"ammoInMagazine": true,
		"ammoReserve":    true,
		"ammoType":       true,
	},
	"planting":  true,
	"inBuyzone": true,
	"money":     true,
	"kills":     true,
	"deaths":    true,
	"firing":    true,
}

var DefaultTickTemplate = map[string]interface{}{
	"tick":              true,
	"players":           DefaultPlayerTemplate,
	"totalRoundsPlayed": true,
	"smokes": map[string]interface{}{
		"id": true,
		"position": map[string]interface{}{
			"x": true,
			"y": true,
		},
	},
}

var DefaultDemoTemplate = map[string]interface{}{
	"header": map[string]interface{}{
		"mapName":  true,
		"tickrate": true,
		"fps":      true,
	},
	"ticks":  DefaultTickTemplate,
	"events": DefaultEventsTemplate,
}

var DefaultEventsTemplate = map[string]interface{}{
	"SMOKE_STARTED": DefaultGrenadeEventMapping,
	"SMOKE_EXPIRED": DefaultGrenadeEventMapping,
	"FIRE_STARTED":  DefaultGrenadeEventMapping,
	"FIRE_EXPIRED":  DefaultGrenadeEventMapping,
	"ROUND_STARTED": map[string]interface{}{
		"name":      true,
		"tick":      true,
		"timelimit": true,
	},
	"ROUND_ENDED": map[string]interface{}{
		"name": true,
		"tick": true,
	},
	"MATCH_STARTED": map[string]interface{}{
		"name": true,
		"tick": true,
	},
	"FLASH_EXPLOSION": map[string]interface{}{
		"name":     true,
		"tick":     true,
		"position": DefaultPositionTemplate,
	},
	"HE_EXPLOSION": map[string]interface{}{
		"name":     true,
		"tick":     true,
		"position": DefaultPositionTemplate,
	},
	"WEAPON_FIRED": map[string]interface{}{
		"name": true,
		"tick": true,
		"weapon": map[string]interface{}{
			"type":           true,
			"ammoInMagazine": true,
			"ammoReserve":    true,
			"ammoType":       true,
		},
	},
	"BOMB_PLANTED": map[string]interface{}{
		"name":     true,
		"tick":     true,
		"bombsite": true,
	},
}

var DefaultGrenadeEventMapping = map[string]interface{}{
	"id":       true,
	"tick":     true,
	"name":     true,
	"position": DefaultPositionTemplate,
}
